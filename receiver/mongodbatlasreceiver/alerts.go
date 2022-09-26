// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mongodbatlasreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver"

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha1" // #nosec G505 -- SHA1 is the algorithm mongodbatlas uses, it must be used to calculate the HMAC signature
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/tidwall/buntdb"
	"go.mongodb.org/atlas/mongodbatlas"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configtls"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/multierr"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver/internal"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver/internal/model"
)

// maxContentLength is the maximum payload size we will accept from incoming requests.
// Requests are generally ~1000 bytes, so we overshoot that by an order of magnitude.
// This is to protect from overly large requests.
const (
	maxContentLength    int64  = 16384
	signatureHeaderName string = "X-MMS-Signature"

	alertModeListen    alertMode = "listen"
	alertModeRetrieval alertMode = "retrieve"

	defaultAlertsPollInterval = 30 * time.Second
)

type alertMode string

type alertsClient interface {
	GetProject(context.Context, string) (*mongodbatlas.Project, error)
	GetAlerts(context.Context, string) ([]mongodbatlas.Alert, error)
}

type alertsReceiver struct {
	addr        string
	secret      string
	server      *http.Server
	mode        alertMode
	tlsSettings *configtls.TLSServerSetting
	consumer    consumer.Logs
	wg          *sync.WaitGroup
	logger      *zap.Logger

	// only relevant in `retrieval` mode
	projects      []ProjectConfig
	client        alertsClient
	privateKey    string
	publicKey     string
	retrySettings exporterhelper.RetrySettings
	pollInterval  time.Duration
	cache         *buntdb.DB
	doneChan      chan bool
}

func newAlertsReceiver(logger *zap.Logger, baseConfig *Config, consumer consumer.Logs) (*alertsReceiver, error) {
	cfg := baseConfig.Alerts
	var tlsConfig *tls.Config

	if cfg.TLS != nil {
		var err error

		tlsConfig, err = cfg.TLS.LoadTLSConfig()
		if err != nil {
			return nil, err
		}
	}

	recv := &alertsReceiver{
		addr:          cfg.Endpoint,
		secret:        cfg.Secret,
		tlsSettings:   cfg.TLS,
		consumer:      consumer,
		mode:          alertMode(cfg.Mode),
		projects:      cfg.Projects,
		retrySettings: baseConfig.RetrySettings,
		publicKey:     baseConfig.PublicKey,
		privateKey:    baseConfig.PrivateKey,
		wg:            &sync.WaitGroup{},
		pollInterval:  baseConfig.Alerts.PollInterval,
		doneChan:      make(chan bool, 1),
		logger:        logger,
	}

	if recv.mode == alertModeRetrieval {
		client, err := internal.NewMongoDBAtlasClient(recv.publicKey, recv.privateKey, recv.retrySettings, recv.logger)
		if err != nil {
			return nil, err
		}
		recv.client = client
	} else {
		s := &http.Server{
			TLSConfig: tlsConfig,
			Handler:   http.HandlerFunc(recv.handleRequest),
		}

		recv.server = s
	}

	return recv, nil
}

func (a alertsReceiver) Start(ctx context.Context, host component.Host) error {
	switch a.mode {
	case alertModeListen:
		return a.startListening(ctx, host)
	case alertModeRetrieval:
		return a.startRetrieving(ctx, host)
	default:
		return fmt.Errorf("receiver attempted to start without a known mode: %s", a.mode)
	}
}

func (a alertsReceiver) startRetrieving(ctx context.Context, _ component.Host) error {
	// may want to consider using disk storage for tracking which alerts have been sent
	// or may want to revisit how we get alerts
	// this is a memory cache used to keep track of unupdated alerts
	// eventually it should be transferred over to the storage extension
	cache, err := buntdb.Open(":memory:")
	if err != nil {
		return fmt.Errorf("unable to initialize cache for retrieval client: %w", err)
	}
	a.cache = cache

	time := time.NewTicker(a.pollInterval)
	a.wg.Add(1)
	go func() {
		defer a.wg.Done()
		for {
			select {
			case <-time.C:
				if err := a.retrieveAndProcessAlerts(ctx); err != nil {
					a.logger.Error("unable to retrieve alerts", zap.Error(err))
				}
			case <-a.doneChan:
				return
			case <-ctx.Done():
				return
			}
		}
	}()
	return nil
}

func (a alertsReceiver) retrieveAndProcessAlerts(ctx context.Context) error {
	var alerts []mongodbatlas.Alert
	for _, p := range a.projects {
		project, err := a.client.GetProject(ctx, p.Name)
		if err != nil {
			a.logger.Error("error retrieving project "+p.Name+":", zap.Error(err))
			continue
		}
		projectAlerts, err := a.client.GetAlerts(ctx, project.ID)
		if err != nil {
			a.logger.Error("unable to get alerts for project", zap.Error(err))
			continue
		}
		alerts = append(alerts, projectAlerts...)
	}
	now := pcommon.NewTimestampFromTime(time.Now())
	logs, err := a.convertAlerts(now, alerts)
	if err != nil {
		return err
	}
	if logs.LogRecordCount() > 0 {
		return a.consumer.ConsumeLogs(ctx, logs)
	}
	return nil
}

func (a alertsReceiver) startListening(ctx context.Context, host component.Host) error {
	// We use a.server.Serve* over a.server.ListenAndServe*
	// So that we can catch and return errors relating to binding to network interface on start.
	var lc net.ListenConfig

	l, err := lc.Listen(ctx, "tcp", a.addr)
	if err != nil {
		return err
	}

	a.wg.Add(1)
	if a.tlsSettings != nil {
		go func() {
			defer a.wg.Done()

			a.logger.Debug("Starting ServeTLS",
				zap.String("address", a.addr),
				zap.String("certfile", a.tlsSettings.CertFile),
				zap.String("keyfile", a.tlsSettings.KeyFile))

			err := a.server.ServeTLS(l, a.tlsSettings.CertFile, a.tlsSettings.KeyFile)

			a.logger.Debug("Serve TLS done")

			if err != http.ErrServerClosed {
				a.logger.Error("ServeTLS failed", zap.Error(err))
				host.ReportFatalError(err)
			}
		}()
	} else {
		go func() {
			defer a.wg.Done()

			a.logger.Debug("Starting Serve", zap.String("address", a.addr))

			err := a.server.Serve(l)

			a.logger.Debug("Serve done")

			if err != http.ErrServerClosed {
				a.logger.Error("Serve failed", zap.Error(err))
				host.ReportFatalError(err)
			}
		}()
	}
	return nil
}

func (a alertsReceiver) handleRequest(rw http.ResponseWriter, req *http.Request) {
	if req.ContentLength < 0 {
		rw.WriteHeader(http.StatusLengthRequired)
		a.logger.Debug("Got request with no Content-Length specified", zap.String("remote", req.RemoteAddr))
		return
	}

	if req.ContentLength > maxContentLength {
		rw.WriteHeader(http.StatusRequestEntityTooLarge)
		a.logger.Debug("Got request with large Content-Length specified",
			zap.String("remote", req.RemoteAddr),
			zap.Int64("content-length", req.ContentLength),
			zap.Int64("max-content-length", maxContentLength))
		return
	}

	payloadSigHeader := req.Header.Get(signatureHeaderName)
	if payloadSigHeader == "" {
		rw.WriteHeader(http.StatusBadRequest)
		a.logger.Debug("Got payload with no HMAC signature, dropping...")
		return
	}

	payload := make([]byte, req.ContentLength)
	_, err := io.ReadFull(req.Body, payload)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		a.logger.Debug("Failed to read alerts payload", zap.Error(err), zap.String("remote", req.RemoteAddr))
		return
	}

	if err = verifyHMACSignature(a.secret, payload, payloadSigHeader); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		a.logger.Debug("Got payload with invalid HMAC signature, dropping...", zap.Error(err), zap.String("remote", req.RemoteAddr))
		return
	}

	logs, err := payloadToLogs(time.Now(), payload)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		a.logger.Error("Failed to convert log payload to log record", zap.Error(err))
		return
	}

	if err := a.consumer.ConsumeLogs(req.Context(), logs); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		a.logger.Error("Failed to consumer alert as log", zap.Error(err))
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (a alertsReceiver) Shutdown(ctx context.Context) error {
	switch a.mode {
	case alertModeListen:
		return a.shutdownListener(ctx)
	case alertModeRetrieval:
		return a.shutdownRetriever(ctx)
	default:
		return errors.New("alert receiver mode not set, unable to shutdown")
	}
}

func (a alertsReceiver) shutdownListener(ctx context.Context) error {
	a.logger.Debug("Shutting down server")
	err := a.server.Shutdown(ctx)
	if err != nil {
		return err
	}

	a.logger.Debug("Waiting for shutdown to complete.")
	a.wg.Wait()
	return nil
}

func (a alertsReceiver) shutdownRetriever(_ context.Context) error {
	a.logger.Debug("Shutting down client")
	a.doneChan <- true
	a.wg.Wait()
	return nil
}

func (a alertsReceiver) convertAlerts(now pcommon.Timestamp, alerts []mongodbatlas.Alert) (plog.Logs, error) {
	logs := plog.NewLogs()

	var errs error
	for _, alert := range alerts {
		if a.hasProcessed(alert) {
			continue
		}

		resourceLogs := logs.ResourceLogs().AppendEmpty()
		logRecord := resourceLogs.ScopeLogs().AppendEmpty().LogRecords().AppendEmpty()
		logRecord.SetObservedTimestamp(now)

		ts, err := time.Parse(time.RFC3339, alert.Updated)
		if err != nil {
			a.logger.Warn("unable to interpret updated time for alert with timestamp. Expecting a RFC3339 timestamp", zap.String("timestamp", alert.Updated))
			continue
		}

		logRecord.SetTimestamp(pcommon.NewTimestampFromTime(ts))
		logRecord.SetSeverityNumber(severityFromAPIAlert(alert))
		// this could be fairly expensive to do, maybe should evaulate
		bodyBytes, err := json.Marshal(alert)
		if err != nil {
			a.logger.Warn("unable to marshal alert into a body string")
			continue
		}

		logRecord.Body().SetStringVal(string(bodyBytes))

		resourceAttrs := resourceLogs.Resource().Attributes()
		resourceAttrs.PutString("mongodbatlas.group.id", alert.GroupID)
		resourceAttrs.PutString("mongodbatlas.alert.config.id", alert.AlertConfigID)
		putStringToMapNotNil(resourceAttrs, "mongodbatlas.cluster.name", &alert.ClusterName)
		putStringToMapNotNil(resourceAttrs, "mongodbatlas.replica_set.name", &alert.ReplicaSetName)

		attrs := logRecord.Attributes()
		// These attributes are always present
		attrs.PutString("event.domain", "mongodbatlas")
		attrs.PutString("event.name", alert.EventTypeName)
		attrs.PutString("status", alert.Status)
		attrs.PutString("created", alert.Created)
		attrs.PutString("updated", alert.Updated)
		attrs.PutString("id", alert.ID)

		// These attributes are optional and may not be present, depending on the alert type.
		putStringToMapNotNil(attrs, "metric.name", &alert.MetricName)
		putStringToMapNotNil(attrs, "type_name", &alert.EventTypeName)
		putStringToMapNotNil(attrs, "last_notified", &alert.LastNotified)
		putStringToMapNotNil(attrs, "resolved", &alert.Resolved)
		putStringToMapNotNil(attrs, "acknowledgement.comment", &alert.AcknowledgementComment)
		putStringToMapNotNil(attrs, "acknowledgement.username", &alert.AcknowledgingUsername)
		putStringToMapNotNil(attrs, "acknowledgement.until", &alert.AcknowledgedUntil)

		if alert.CurrentValue != nil {
			attrs.PutDouble("metric.value", *alert.CurrentValue.Number)
			attrs.PutString("metric.units", alert.CurrentValue.Units)
		}

		host, portStr, err := net.SplitHostPort(alert.HostnameAndPort)
		if err != nil {
			errs = multierr.Append(errs, fmt.Errorf("failed to split host:port %s: %w", alert.HostnameAndPort, err))
			continue
		}

		port, err := strconv.ParseInt(portStr, 10, 64)
		if err != nil {
			errs = multierr.Append(errs, fmt.Errorf("failed to parse port %s: %w", portStr, err))
			continue
		}

		attrs.PutString("net.peer.name", host)
		attrs.PutInt("net.peer.port", port)

		errs = multierr.Append(errs, a.cache.Update(func(tx *buntdb.Tx) error {
			_, _, err := tx.Set(alertKey(alert), "processed", &buntdb.SetOptions{
				Expires: true,
				TTL:     30 * time.Minute,
			})
			return err
		}))
	}

	return logs, errs
}

func verifyHMACSignature(secret string, payload []byte, signatureHeader string) error {
	b64Decoder := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(signatureHeader))
	payloadSig, err := io.ReadAll(b64Decoder)
	if err != nil {
		return err
	}

	h := hmac.New(sha1.New, []byte(secret))
	h.Write(payload)
	calculatedSig := h.Sum(nil)

	if !hmac.Equal(calculatedSig, payloadSig) {
		return errors.New("calculated signature does not equal header signature")
	}

	return nil
}

func payloadToLogs(now time.Time, payload []byte) (plog.Logs, error) {
	var alert model.Alert

	err := json.Unmarshal(payload, &alert)
	if err != nil {
		return plog.Logs{}, err
	}

	logs := plog.NewLogs()
	resourceLogs := logs.ResourceLogs().AppendEmpty()
	logRecord := resourceLogs.ScopeLogs().AppendEmpty().LogRecords().AppendEmpty()

	logRecord.SetObservedTimestamp(pcommon.NewTimestampFromTime(now))
	logRecord.SetTimestamp(timestampFromAlert(alert))
	logRecord.SetSeverityNumber(severityFromAlert(alert))
	logRecord.Body().SetStringVal(string(payload))

	resourceAttrs := resourceLogs.Resource().Attributes()
	resourceAttrs.PutString("mongodbatlas.group.id", alert.GroupID)
	resourceAttrs.PutString("mongodbatlas.alert.config.id", alert.AlertConfigID)
	putStringToMapNotNil(resourceAttrs, "mongodbatlas.cluster.name", alert.ClusterName)
	putStringToMapNotNil(resourceAttrs, "mongodbatlas.replica_set.name", alert.ReplicaSetName)

	attrs := logRecord.Attributes()
	// These attributes are always present
	attrs.PutString("event.domain", "mongodbatlas")
	attrs.PutString("event.name", alert.EventType)
	attrs.PutString("message", alert.HumanReadable)
	attrs.PutString("status", alert.Status)
	attrs.PutString("created", alert.Created)
	attrs.PutString("updated", alert.Updated)
	attrs.PutString("id", alert.ID)

	// These attributes are optional and may not be present, depending on the alert type.
	putStringToMapNotNil(attrs, "metric.name", alert.MetricName)
	putStringToMapNotNil(attrs, "type_name", alert.TypeName)
	putStringToMapNotNil(attrs, "user_alias", alert.UserAlias)
	putStringToMapNotNil(attrs, "last_notified", alert.LastNotified)
	putStringToMapNotNil(attrs, "resolved", alert.Resolved)
	putStringToMapNotNil(attrs, "acknowledgement.comment", alert.AcknowledgementComment)
	putStringToMapNotNil(attrs, "acknowledgement.username", alert.AcknowledgementUsername)
	putStringToMapNotNil(attrs, "acknowledgement.until", alert.AcknowledgedUntil)

	if alert.CurrentValue != nil {
		attrs.PutDouble("metric.value", alert.CurrentValue.Number)
		attrs.PutString("metric.units", alert.CurrentValue.Units)
	}

	if alert.HostNameAndPort != nil {
		host, portStr, err := net.SplitHostPort(*alert.HostNameAndPort)
		if err != nil {
			return plog.Logs{}, fmt.Errorf("failed to split host:port %s: %w", *alert.HostNameAndPort, err)
		}

		port, err := strconv.ParseInt(portStr, 10, 64)
		if err != nil {
			return plog.Logs{}, fmt.Errorf("failed to parse port %s: %w", portStr, err)
		}

		attrs.PutString("net.peer.name", host)
		attrs.PutInt("net.peer.port", port)

	}

	return logs, nil
}

func (a *alertsReceiver) hasProcessed(alert mongodbatlas.Alert) bool {
	err := a.cache.View(func(tx *buntdb.Tx) error {
		_, err := tx.Get(alertKey(alert))
		if err != nil {
			return err
		}
		return nil
	})
	return err == nil
}

// alertKey is a key to index the alert to see if it has already been processed.
func alertKey(alert mongodbatlas.Alert) string {
	return fmt.Sprintf("%s|%s", alert.ID, alert.Status)
}

func timestampFromAlert(a model.Alert) pcommon.Timestamp {
	if time, err := time.Parse(time.RFC3339, a.Updated); err == nil {
		return pcommon.NewTimestampFromTime(time)
	}

	return pcommon.Timestamp(0)
}

// severityFromAlert maps the alert to a severity number.
// Currently, it just maps "OPEN" alerts to WARN, and everything else to INFO.
func severityFromAlert(a model.Alert) plog.SeverityNumber {
	// Status is defined here: https://www.mongodb.com/docs/atlas/reference/api/alerts-get-alert/#response-elements
	// It may also be "INFORMATIONAL" for single-fire alerts (events)
	switch a.Status {
	case "OPEN":
		return plog.SeverityNumberWarn
	default:
		return plog.SeverityNumberInfo
	}
}

// severityFromAPIAlert is a workaround for shared types between the API and the model
func severityFromAPIAlert(a mongodbatlas.Alert) plog.SeverityNumber {
	switch a.Status {
	case "OPEN":
		return plog.SeverityNumberWarn
	default:
		return plog.SeverityNumberInfo
	}
}

func putStringToMapNotNil(m pcommon.Map, k string, v *string) {
	if v != nil {
		m.PutString(k, *v)
	}
}
