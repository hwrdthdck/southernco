// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build windows

package windows // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/input/windows"

import (
	"errors"
	"fmt"
	"unicode/utf16"
	"unsafe"
)

// errUnknownNextFrame is an error returned when a systemcall indicates the next frame is 0 bytes.
var errUnknownNextFrame = errors.New("the buffer size needed by the next frame of a render syscall was 0, unable to determine size of next frame")

// systemPropertiesRenderContext stores a custom rendering context to get only the event properties.
var systemPropertiesRenderContext = uintptr(0)
var systemPropertiesRenderContextErr error

func init() {
	// This is not expected to fail, however, collecting the error if a new failure mode appears.
	systemPropertiesRenderContext, systemPropertiesRenderContextErr = evtCreateRenderContext(0, nil, EvtRenderContextSystem)
}

// Event is an event stored in windows event log.
type Event struct {
	handle uintptr
}

// GetPublisherName will get the publisher name of the event.
func (e *Event) GetPublisherName(buffer Buffer) (string, error) {
	if e.handle == 0 {
		return "", fmt.Errorf("event handle does not exist")
	}

	if systemPropertiesRenderContextErr != nil {
		return "", fmt.Errorf("failed to create render context: %w", systemPropertiesRenderContextErr)
	}

	bufferUsed, err := evtRender(systemPropertiesRenderContext, e.handle, EvtRenderEventValues, buffer.SizeBytes(), buffer.FirstByte())
	if errors.Is(err, ErrorInsufficientBuffer) {
		buffer.UpdateSizeBytes(*bufferUsed)
		return e.GetPublisherName(buffer)
	}

	if err != nil {
		return "", fmt.Errorf("failed to get provider name: %w", err)
	}

	utf16Ptr := (**uint16)(unsafe.Pointer(buffer.FirstByte()))
	providerName := utf16PtrToString(*utf16Ptr)

	return providerName, nil
}

// utf16PtrToString converts Windows API LPWSTR (pointer to string) to go string
func utf16PtrToString(s *uint16) string {
	if s == nil {
		return ""
	}

	utf16Len := 0
	curPtr := unsafe.Pointer(s)
	for *(*uint16)(curPtr) != 0 {
		curPtr = unsafe.Pointer(uintptr(curPtr) + unsafe.Sizeof(*s))
		utf16Len++
	}

	slice := unsafe.Slice(s, utf16Len)
	return string(utf16.Decode(slice))
}

// RenderSimple will render the event as EventXML without formatted info.
func (e *Event) RenderSimple(buffer Buffer) (EventXML, error) {
	if e.handle == 0 {
		return EventXML{}, fmt.Errorf("event handle does not exist")
	}

	bufferUsed, err := evtRender(0, e.handle, EvtRenderEventXML, buffer.SizeBytes(), buffer.FirstByte())
	if errors.Is(err, ErrorInsufficientBuffer) {
		buffer.UpdateSizeBytes(*bufferUsed)
		return e.RenderSimple(buffer)
	}

	if err != nil {
		return EventXML{}, fmt.Errorf("syscall to 'EvtRender' failed: %w", err)
	}

	bytes, err := buffer.ReadBytes(*bufferUsed)
	if err != nil {
		return EventXML{}, fmt.Errorf("failed to read bytes from buffer: %w", err)
	}

	return unmarshalEventXML(bytes)
}

// RenderFormatted will render the event as EventXML with formatted info.
func (e *Event) RenderFormatted(buffer Buffer, publisher Publisher) (EventXML, error) {
	if e.handle == 0 {
		return EventXML{}, fmt.Errorf("event handle does not exist")
	}

	bufferUsed, err := evtFormatMessage(publisher.handle, e.handle, 0, 0, 0, EvtFormatMessageXML, buffer.SizeWide(), buffer.FirstByte())
	if errors.Is(err, ErrorInsufficientBuffer) {
		// If the bufferUsed is 0 return an error as we don't want to make a recursive call with no buffer
		if *bufferUsed == 0 {
			return EventXML{}, errUnknownNextFrame
		}

		buffer.UpdateSizeWide(*bufferUsed)
		return e.RenderFormatted(buffer, publisher)
	}

	if err != nil {
		return EventXML{}, fmt.Errorf("syscall to 'EvtFormatMessage' failed: %w", err)
	}

	bytes, err := buffer.ReadWideChars(*bufferUsed)
	if err != nil {
		return EventXML{}, fmt.Errorf("failed to read bytes from buffer: %w", err)
	}

	return unmarshalEventXML(bytes)
}

// Close will close the event handle.
func (e *Event) Close() error {
	if e.handle == 0 {
		return nil
	}

	if err := evtClose(e.handle); err != nil {
		return fmt.Errorf("failed to close event handle: %w", err)
	}

	e.handle = 0
	return nil
}

// NewEvent will create a new event from an event handle.
func NewEvent(handle uintptr) Event {
	return Event{
		handle: handle,
	}
}
