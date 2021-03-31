// Copyright 2019, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package translator

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net/url"
	"regexp"
	"strings"
	"time"

	awsP "github.com/aws/aws-sdk-go/aws"
	"go.opentelemetry.io/collector/consumer/pdata"
	semconventions "go.opentelemetry.io/collector/translator/conventions"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/awsxray"
)

// AWS X-Ray acceptable values for origin field.
const (
	OriginEC2        = "AWS::EC2::Instance"
	OriginECS        = "AWS::ECS::Container"
	OriginECSEC2     = "AWS::ECS::EC2"
	OriginECSFargate = "AWS::ECS::Fargate"
	OriginEB         = "AWS::ElasticBeanstalk::Environment"
	OriginEKS        = "AWS::EKS::Container"
)

var (
	// reInvalidSpanCharacters defines the invalid letters in a span name as per
	// https://docs.aws.amazon.com/xray/latest/devguide/xray-api-segmentdocuments.html
	reInvalidSpanCharacters = regexp.MustCompile(`[^ 0-9\p{L}N_.:/%&#=+,\-@]`)
)

const (
	// defaultSpanName will be used if there are no valid xray characters in the span name
	defaultSegmentName = "span"
	// maxSegmentNameLength the maximum length of a Segment name
	maxSegmentNameLength = 200
)

const (
	traceIDLength    = 35 // fixed length of aws trace id
	identifierOffset = 11 // offset of identifier within traceID
)

var (
	writers = newWriterPool(2048)
)

// MakeSegmentDocumentString converts an OpenTelemetry Span to an X-Ray Segment and then serialzies to JSON
func MakeSegmentDocumentString(span pdata.Span, resource pdata.Resource, indexedAttrs []string, indexAllAttrs bool) (string, error) {
	segment, err := MakeSegment(span, resource, indexedAttrs, indexAllAttrs)
	if err != nil {
		return "", err
	}
	w := writers.borrow()
	if err := w.Encode(*segment); err != nil {
		return "", err
	}
	jsonStr := w.String()
	writers.release(w)
	return jsonStr, nil
}

// MakeSegment converts an OpenTelemetry Span to an X-Ray Segment
func MakeSegment(span pdata.Span, resource pdata.Resource, indexedAttrs []string, indexAllAttrs bool) (*awsxray.Segment, error) {
	var segmentType string

	storeResource := true
	if span.Kind() != pdata.SpanKindSERVER &&
		!span.ParentSpanID().IsEmpty() {
		segmentType = "subsegment"
		// We only store the resource information for segments, the local root.
		storeResource = false
	}

	// convert trace id
	traceID, err := convertToAmazonTraceID(span.TraceID())
	if err != nil {
		return nil, err
	}

	var (
		startTime                              = timestampToFloatSeconds(span.StartTimestamp())
		endTime                                = timestampToFloatSeconds(span.EndTimestamp())
		httpfiltered, http                     = makeHTTP(span)
		isError, isFault, causefiltered, cause = makeCause(span, httpfiltered, resource)
		origin                                 = determineAwsOrigin(resource)
		awsfiltered, aws                       = makeAws(causefiltered, resource)
		service                                = makeService(resource)
		sqlfiltered, sql                       = makeSQL(awsfiltered)
		user, annotations, metadata            = makeXRayAttributes(sqlfiltered, resource, storeResource, indexedAttrs, indexAllAttrs)
		name                                   string
		namespace                              string
	)

	// X-Ray segment names are service names, unlike span names which are methods. Try to find a service name.

	attributes := span.Attributes()

	// peer.service should always be prioritized for segment names when set because it is what the user decided.
	if peerService, ok := attributes.Get(semconventions.AttributePeerService); ok {
		name = peerService.StringVal()
	}

	if name == "" {
		if awsService, ok := attributes.Get(awsxray.AWSServiceAttribute); ok {
			// Generally spans are named something like "Method" or "Service.Method" but for AWS spans, X-Ray expects spans
			// to be named "Service"
			name = awsService.StringVal()

			namespace = "aws"
		}
	}

	if name == "" {
		if dbInstance, ok := attributes.Get(semconventions.AttributeDBName); ok {
			// For database queries, the segment name convention is <db name>@<db host>
			name = dbInstance.StringVal()
			if dbURL, ok := attributes.Get(semconventions.AttributeDBConnectionString); ok {
				if parsed, _ := url.Parse(dbURL.StringVal()); parsed != nil {
					if parsed.Hostname() != "" {
						name += "@" + parsed.Hostname()
					}
				}
			}
		}
	}

	if name == "" && span.Kind() == pdata.SpanKindSERVER {
		// Only for a server span, we can use the resource.
		if service, ok := resource.Attributes().Get(semconventions.AttributeServiceName); ok {
			name = service.StringVal()
		}
	}

	if name == "" {
		if rpcservice, ok := attributes.Get(semconventions.AttributeRPCService); ok {
			name = rpcservice.StringVal()
		}
	}

	if name == "" {
		if host, ok := attributes.Get(semconventions.AttributeHTTPHost); ok {
			name = host.StringVal()
		}
	}

	if name == "" {
		if peer, ok := attributes.Get(semconventions.AttributeNetPeerName); ok {
			name = peer.StringVal()
		}
	}

	if name == "" {
		name = fixSegmentName(span.Name())
	}

	if namespace == "" && span.Kind() == pdata.SpanKindCLIENT {
		namespace = "remote"
	}

	return &awsxray.Segment{
		ID:          awsxray.String(span.SpanID().HexString()),
		TraceID:     awsxray.String(traceID),
		Name:        awsxray.String(name),
		StartTime:   awsP.Float64(startTime),
		EndTime:     awsP.Float64(endTime),
		ParentID:    awsxray.String(span.ParentSpanID().HexString()),
		Fault:       awsP.Bool(isFault),
		Error:       awsP.Bool(isError),
		Cause:       cause,
		Origin:      awsxray.String(origin),
		Namespace:   awsxray.String(namespace),
		User:        awsxray.String(user),
		HTTP:        http,
		AWS:         aws,
		Service:     service,
		SQL:         sql,
		Annotations: annotations,
		Metadata:    metadata,
		Type:        awsxray.String(segmentType),
	}, nil
}

// newSegmentID generates a new valid X-Ray SegmentID
func newSegmentID() pdata.SpanID {
	var r [8]byte
	_, err := rand.Read(r[:])
	if err != nil {
		panic(err)
	}
	return pdata.NewSpanID(r)
}

func determineAwsOrigin(resource pdata.Resource) string {
	if resource.Attributes().Len() == 0 {
		return ""
	}

	if provider, ok := resource.Attributes().Get(semconventions.AttributeCloudProvider); ok {
		if provider.StringVal() != semconventions.AttributeCloudProviderAWS {
			return ""
		}
	}

	// TODO(willarmiros): Only use platform for origin resolution once detectors for all AWS environments are
	// implemented for robustness
	if is, present := resource.Attributes().Get(semconventions.AttributeCloudPlatform); present {
		switch is.StringVal() {
		case "EKS":
			return OriginEKS
		case "ElasticBeanstalk":
			return OriginEB
		case "ECS":
			lt, present := resource.Attributes().Get("aws.ecs.launchtype")
			if !present {
				return OriginECS
			}
			switch lt.StringVal() {
			case "ec2":
				return OriginECSEC2
			case "fargate":
				return OriginECSFargate
			default:
				return OriginECS
			}
		case "EC2":
			return OriginEC2

		// If infrastructure_service is defined with a non-AWS value, we should not assign it an AWS origin
		default:
			return ""
		}
	}

	// EKS > EB > ECS > EC2
	_, eks := resource.Attributes().Get(semconventions.AttributeK8sCluster)
	if eks {
		return OriginEKS
	}
	_, eb := resource.Attributes().Get(semconventions.AttributeServiceInstance)
	if eb {
		return OriginEB
	}
	_, ecs := resource.Attributes().Get(semconventions.AttributeContainerName)
	if ecs {
		return OriginECS
	}
	_, ec2 := resource.Attributes().Get(semconventions.AttributeHostID)
	if ec2 {
		return OriginEC2
	}
	return ""
}

// convertToAmazonTraceID converts a trace ID to the Amazon format.
//
// A trace ID unique identifier that connects all segments and subsegments
// originating from a single client request.
//  * A trace_id consists of three numbers separated by hyphens. For example,
//    1-58406520-a006649127e371903a2de979. This includes:
//  * The version number, that is, 1.
//  * The time of the original request, in Unix epoch time, in 8 hexadecimal digits.
//  * For example, 10:00AM December 2nd, 2016 PST in epoch time is 1480615200 seconds,
//    or 58406520 in hexadecimal.
//  * A 96-bit identifier for the trace, globally unique, in 24 hexadecimal digits.
func convertToAmazonTraceID(traceID pdata.TraceID) (string, error) {
	const (
		// maxAge of 28 days.  AWS has a 30 day limit, let's be conservative rather than
		// hit the limit
		maxAge = 60 * 60 * 24 * 28

		// maxSkew allows for 5m of clock skew
		maxSkew = 60 * 5
	)

	var (
		content      = [traceIDLength]byte{}
		epochNow     = time.Now().Unix()
		traceIDBytes = traceID.Bytes()
		epoch        = int64(binary.BigEndian.Uint32(traceIDBytes[0:4]))
		b            = [4]byte{}
	)

	// If AWS traceID originally came from AWS, no problem.  However, if oc generated
	// the traceID, then the epoch may be outside the accepted AWS range of within the
	// past 30 days.
	//
	// In that case, we return invalid traceid error
	if delta := epochNow - epoch; delta > maxAge || delta < -maxSkew {
		return "", fmt.Errorf("invalid xray traceid: %s", traceID.HexString())
	}

	binary.BigEndian.PutUint32(b[0:4], uint32(epoch))

	content[0] = '1'
	content[1] = '-'
	hex.Encode(content[2:10], b[0:4])
	content[10] = '-'
	hex.Encode(content[identifierOffset:], traceIDBytes[4:16]) // overwrite with identifier

	return string(content[0:traceIDLength]), nil
}

func timestampToFloatSeconds(ts pdata.Timestamp) float64 {
	return float64(ts) / float64(time.Second)
}

func makeXRayAttributes(attributes map[string]string, resource pdata.Resource, storeResource bool, indexedAttrs []string, indexAllAttrs bool) (
	string, map[string]interface{}, map[string]map[string]interface{}) {
	var (
		annotations = map[string]interface{}{}
		metadata    = map[string]map[string]interface{}{}
		user        string
	)
	userid, ok := attributes[semconventions.AttributeEnduserID]
	if ok {
		user = userid
		delete(attributes, semconventions.AttributeEnduserID)
	}

	if len(attributes) == 0 && (!storeResource || resource.Attributes().Len() == 0) {
		return user, nil, nil
	}

	defaultMetadata := map[string]interface{}{}

	indexedKeys := map[string]bool{}
	if !indexAllAttrs {
		for _, name := range indexedAttrs {
			indexedKeys[name] = true
		}
	}

	if storeResource {
		resource.Attributes().ForEach(func(key string, value pdata.AttributeValue) {
			key = "otel.resource." + key
			annoVal := annotationValue(value)
			indexed := indexAllAttrs || indexedKeys[key]
			if annoVal != nil && indexed {
				key = fixAnnotationKey(key)
				annotations[key] = annoVal
			} else {
				metaVal := metadataValue(value)
				if metaVal != nil {
					defaultMetadata[key] = metaVal
				}
			}
		})
	}

	if indexAllAttrs {
		for key, value := range attributes {
			key = fixAnnotationKey(key)
			annotations[key] = value
		}
	} else {
		for key, value := range attributes {
			if indexedKeys[key] {
				key = fixAnnotationKey(key)
				annotations[key] = value
			} else {
				defaultMetadata[key] = value
			}
		}
	}

	if len(defaultMetadata) > 0 {
		metadata["default"] = defaultMetadata
	}

	return user, annotations, metadata
}

func annotationValue(value pdata.AttributeValue) interface{} {
	switch value.Type() {
	case pdata.AttributeValueSTRING:
		return value.StringVal()
	case pdata.AttributeValueINT:
		return value.IntVal()
	case pdata.AttributeValueDOUBLE:
		return value.DoubleVal()
	case pdata.AttributeValueBOOL:
		return value.BoolVal()
	}
	return nil
}

func metadataValue(value pdata.AttributeValue) interface{} {
	switch value.Type() {
	case pdata.AttributeValueSTRING:
		return value.StringVal()
	case pdata.AttributeValueINT:
		return value.IntVal()
	case pdata.AttributeValueDOUBLE:
		return value.DoubleVal()
	case pdata.AttributeValueBOOL:
		return value.BoolVal()
	case pdata.AttributeValueMAP:
		converted := map[string]interface{}{}
		value.MapVal().ForEach(func(key string, value pdata.AttributeValue) {
			converted[key] = metadataValue(value)
		})
		return converted
	case pdata.AttributeValueARRAY:
		arrVal := value.ArrayVal()
		converted := make([]interface{}, arrVal.Len())
		for i := 0; i < arrVal.Len(); i++ {
			converted[i] = metadataValue(arrVal.At(i))
		}
		return converted
	}
	return nil
}

// fixSegmentName removes any invalid characters from the span name.  AWS X-Ray defines
// the list of valid characters here:
// https://docs.aws.amazon.com/xray/latest/devguide/xray-api-segmentdocuments.html
func fixSegmentName(name string) string {
	if reInvalidSpanCharacters.MatchString(name) {
		// only allocate for ReplaceAllString if we need to
		name = reInvalidSpanCharacters.ReplaceAllString(name, "")
	}

	if length := len(name); length > maxSegmentNameLength {
		name = name[0:maxSegmentNameLength]
	} else if length == 0 {
		name = defaultSegmentName
	}

	return name
}

// fixAnnotationKey removes any invalid characters from the annotaiton key.  AWS X-Ray defines
// the list of valid characters here:
// https://docs.aws.amazon.com/xray/latest/devguide/xray-api-segmentdocuments.html
func fixAnnotationKey(key string) string {
	return strings.Map(func(r rune) rune {
		switch {
		case '0' <= r && r <= '9':
			fallthrough
		case 'A' <= r && r <= 'Z':
			fallthrough
		case 'a' <= r && r <= 'z':
			return r
		default:
			return '_'
		}
	}, key)
}
