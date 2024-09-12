// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package translation

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/schemaprocessor/internal/fixture"
)

func TestRaceTranslationSpanChanges(t *testing.T) {
	t.Parallel()

	tn, err := newTranslatorFromReader(
		zap.NewNop(),
		"https://example.com/1.7.0",
		LoadTranslationVersion(t, "complex_changeset.yml"),
	)
	require.NoError(t, err, "Must not error when creating translator")

	ctx, done := context.WithCancel(context.Background())
	t.Cleanup(done)

	fixture.ParallelRaceCompute(t, 10, func() error {
		for i := 0; i < 10; i++ {
			v := &Version{1, 0, 0}
			spans := NewExampleSpans(t, *v)
			for i := 0; i < spans.ResourceSpans().Len(); i++ {
				rSpan := spans.ResourceSpans().At(i)
				tn.ApplyAllResourceChanges(ctx, rSpan, rSpan.SchemaUrl())
				for j := 0; j < rSpan.ScopeSpans().Len(); j++ {
					span := rSpan.ScopeSpans().At(j)
					tn.ApplyScopeSpanChanges(ctx, span, span.SchemaUrl())
				}
			}
		}
		return nil
	})
}

func TestRaceTranslationMetricChanges(t *testing.T) {
	t.Parallel()

	tn, err := newTranslatorFromReader(
		zap.NewNop(),
		"https://example.com/1.7.0",
		LoadTranslationVersion(t, "complex_changeset.yml"),
	)
	require.NoError(t, err, "Must not error when creating translator")

	ctx, done := context.WithCancel(context.Background())
	t.Cleanup(done)

	fixture.ParallelRaceCompute(t, 10, func() error {
		for i := 0; i < 10; i++ {
			spans := NewExampleSpans(t, Version{1, 0, 0})
			for i := 0; i < spans.ResourceSpans().Len(); i++ {
				rSpan := spans.ResourceSpans().At(i)
				tn.ApplyAllResourceChanges(ctx, rSpan, rSpan.SchemaUrl())
				for j := 0; j < rSpan.ScopeSpans().Len(); j++ {
					span := rSpan.ScopeSpans().At(j)
					tn.ApplyScopeSpanChanges(ctx, span, span.SchemaUrl())
				}
			}
		}
		return nil
	})
}

func TestRaceTranslationLogChanges(t *testing.T) {
	t.Parallel()

	tn, err := newTranslatorFromReader(
		zap.NewNop(),
		"https://example.com/1.7.0",
		LoadTranslationVersion(t, "complex_changeset.yml"),
	)
	require.NoError(t, err, "Must not error when creating translator")

	ctx, done := context.WithCancel(context.Background())
	t.Cleanup(done)

	fixture.ParallelRaceCompute(t, 10, func() error {
		for i := 0; i < 10; i++ {
			metrics := NewExampleMetrics(t, Version{1, 0, 0})
			for i := 0; i < metrics.ResourceMetrics().Len(); i++ {
				rMetrics := metrics.ResourceMetrics().At(i)
				tn.ApplyAllResourceChanges(ctx, rMetrics, rMetrics.SchemaUrl())
				for j := 0; j < rMetrics.ScopeMetrics().Len(); j++ {
					metric := rMetrics.ScopeMetrics().At(j)
					tn.ApplyScopeMetricChanges(ctx, metric, metric.SchemaUrl())
				}
			}
		}
		return nil
	})
}
