// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package geoipprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/geoipprocessor"

import (
	"context"
	"errors"
	"fmt"
	"net"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/otel/attribute"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/geoipprocessor/internal/provider"
)

const (
	ResourceSource  = "resource_attribute"
	AttributeSource = "attribute"
)

var (
	errIPNotFound        = errors.New("no IP address found in the resource attributes")
	errParseIP           = errors.New("could not parse IP address")
	errUnspecifiedIP     = errors.New("unspecified address")
	errUnspecifiedSource = errors.New("no source attributes defined")
)

// newGeoIPProcessor creates a new instance of geoIPProcessor with the specified fields.
type geoIPProcessor struct {
	providers          []provider.GeoIPProvider
	resourceAttributes []attribute.Key

	sourceConfig SourceConfig
}

func newGeoIPProcessor(sourceConfig SourceConfig, resourceAttributes []attribute.Key, providers []provider.GeoIPProvider) *geoIPProcessor {
	return &geoIPProcessor{
		resourceAttributes: resourceAttributes,
		providers:          providers,
		sourceConfig:       sourceConfig,
	}
}

// parseIP parses a string to a net.IP type and returns an error if the IP is invalid or unspecified.
func parseIP(strIP string) (net.IP, error) {
	ip := net.ParseIP(strIP)
	if ip == nil {
		return nil, fmt.Errorf("%w address: %s", errParseIP, strIP)
	} else if ip.IsUnspecified() {
		return nil, fmt.Errorf("%w address: %s", errUnspecifiedIP, strIP)
	}
	return ip, nil
}

// ipFromAttributes extracts an IP address from the given attributes based on the specified fields.
// It returns the first IP address if found, or an error if no valid IP address is found.
func ipFromAttributes(attributes []attribute.Key, resource pcommon.Map) (net.IP, error) {
	for _, attr := range attributes {
		if ipField, found := resource.Get(string(attr)); found {
			// The attribute might contain a domain name. Skip any net.ParseIP error until we have a fine-grained error propagation strategy.
			// TODO: propagate an error once error_mode configuration option is available (e.g. transformprocessor)
			ipAttribute, err := parseIP(ipField.AsString())
			if err == nil && ipAttribute != nil {
				return ipAttribute, nil
			}
		}
	}

	return nil, errIPNotFound
}

// geoLocation fetches geolocation information for the given IP address using the configured providers.
// It returns a set of attributes containing the geolocation data, or an error if the location could not be determined.
func (g *geoIPProcessor) geoLocation(ctx context.Context, ip net.IP) (attribute.Set, error) {
	allAttributes := attribute.EmptySet()
	for _, provider := range g.providers {
		geoAttributes, err := provider.Location(ctx, ip)
		if err != nil {
			return attribute.Set{}, err
		}
		*allAttributes = attribute.NewSet(append(allAttributes.ToSlice(), geoAttributes.ToSlice()...)...)
	}

	return *allAttributes, nil
}

// processMetadata processes a pcommon.Map by adding geolocation attributes based on the found IP address.
func (g *geoIPProcessor) processAttributes(ctx context.Context, metadata pcommon.Map) error {
	ipAddr, err := ipFromAttributes(g.resourceAttributes, metadata)
	if err != nil {
		// TODO: log IP error not found
		if errors.Is(err, errIPNotFound) {
			return nil
		}
		return err
	}

	attributes, err := g.geoLocation(ctx, ipAddr)
	if err != nil {
		return err
	}

	for _, geoAttr := range attributes.ToSlice() {
		metadata.PutStr(string(geoAttr.Key), geoAttr.Value.AsString())
	}

	return nil
}
