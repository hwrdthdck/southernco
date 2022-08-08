// Copyright The OpenTelemetry Authors
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

//go:build windows
// +build windows

package windows // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/input/windows"

import (
	"encoding/xml"
	"fmt"
	"time"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
)

// EventXML is the rendered xml of an event.
type EventXML struct {
	EventID          EventID     `xml:"System>EventID"`
	Provider         Provider    `xml:"System>Provider"`
	Computer         string      `xml:"System>Computer"`
	Channel          string      `xml:"System>Channel"`
	RecordID         uint64      `xml:"System>EventRecordID"`
	TimeCreated      TimeCreated `xml:"System>TimeCreated"`
	Message          string      `xml:"RenderingInfo>Message"`
	RenderedLevel    string      `xml:"RenderingInfo>Level"`
	Level            string      `xml:"System>Level"`
	RenderedTask     string      `xml:"RenderingInfo>Task"`
	Task             string      `xml:"System>Task"`
	RenderedOpcode   string      `xml:"RenderingInfo>Opcode"`
	Opcode           string      `xml:"System>Opcode"`
	RenderedKeywords []string    `xml:"RenderingInfo>Keywords>Keyword"`
	Keywords         []string    `xml:"System>Keywords"`
	EventData        []string    `xml:"EventData>Data"`
}

// parseTimestamp will parse the timestamp of the event.
func (e *EventXML) parseTimestamp() time.Time {
	if timestamp, err := time.Parse(time.RFC3339Nano, e.TimeCreated.SystemTime); err == nil {
		return timestamp
	}
	return time.Now()
}

// parseRenderedSeverity will parse the severity of the event.
func (e *EventXML) parseRenderedSeverity() entry.Severity {
	switch e.RenderedLevel {
	case "":
		return e.parseSeverity()
	case "Critical":
		return entry.Fatal
	case "Error":
		return entry.Error
	case "Warning":
		return entry.Warn
	case "Information":
		return entry.Info
	default:
		return entry.Default
	}
}

// parseSeverity will parse the severity of the event when RenderingInfo is not populated
func (e *EventXML) parseSeverity() entry.Severity {
	switch e.Level {
	case "1":
		return entry.Fatal
	case "2":
		return entry.Error
	case "3":
		return entry.Warn
	case "4":
		return entry.Info
	default:
		return entry.Default
	}
}

// parseBody will parse a body from the event.
func (e *EventXML) parseBody() map[string]interface{} {
	message, details := e.parseMessage()

	level := e.RenderedLevel
	if level == "" {
		level = e.Level
	}

	task := e.RenderedTask
	if task == "" {
		task = e.Task
	}

	opcode := e.RenderedOpcode
	if opcode == "" {
		opcode = e.Opcode
	}

	keywords := e.RenderedKeywords
	if keywords == nil {
		keywords = e.Keywords
	}

	body := map[string]interface{}{
		"event_id": map[string]interface{}{
			"qualifiers": e.EventID.Qualifiers,
			"id":         e.EventID.ID,
		},
		"provider": map[string]interface{}{
			"name":         e.Provider.Name,
			"guid":         e.Provider.GUID,
			"event_source": e.Provider.EventSourceName,
		},
		"system_time": e.TimeCreated.SystemTime,
		"computer":    e.Computer,
		"channel":     e.Channel,
		"record_id":   e.RecordID,
		"level":       level,
		"message":     message,
		"task":        task,
		"opcode":      opcode,
		"keywords":    keywords,
		"event_data":  e.EventData,
	}
	if len(details) > 0 {
		body["details"] = details
	}
	return body
}

// parseMessage will attempt to parse a message into a message and details
func (e *EventXML) parseMessage() (string, map[string]interface{}) {
	switch e.Channel {
	case "Security":
		return parseSecurity(e.Message)
	default:
		return e.Message, nil
	}
}

// unmarshalEventXML will unmarshal EventXML from xml bytes.
func unmarshalEventXML(bytes []byte) (EventXML, error) {
	var eventXML EventXML
	if err := xml.Unmarshal(bytes, &eventXML); err != nil {
		return EventXML{}, fmt.Errorf("failed to unmarshal xml bytes into event: %w (%s)", err, string(bytes))
	}
	return eventXML, nil
}

// EventID is the identifier of the event.
type EventID struct {
	Qualifiers uint16 `xml:"Qualifiers,attr"`
	ID         uint32 `xml:",chardata"`
}

// TimeCreated is the creation time of the event.
type TimeCreated struct {
	SystemTime string `xml:"SystemTime,attr"`
}

// Provider is the provider of the event.
type Provider struct {
	Name            string `xml:"Name,attr"`
	GUID            string `xml:"Guid,attr"`
	EventSourceName string `xml:"EventSourceName,attr"`
}
