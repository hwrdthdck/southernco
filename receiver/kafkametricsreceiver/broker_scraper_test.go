// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package kafkametricsreceiver

import (
	"context"
	"fmt"
	"testing"

	"github.com/Shopify/sarama"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/receiver/receivertest"
)

func TestBrokerShutdown(t *testing.T) {
	client := newMockClient()
	client.closed = false
	client.close = nil
	client.Mock.
		On("Close").Return(nil).
		On("Closed").Return(false)
	scraper := brokerScraper{
		client:   client,
		settings: receivertest.NewNopCreateSettings(),
		config:   Config{},
	}
	_ = scraper.shutdown(context.Background())
	client.AssertExpectations(t)
}

func TestBrokerShutdown_closed(t *testing.T) {
	client := newMockClient()
	client.closed = true
	client.Mock.
		On("Closed").Return(true)
	scraper := brokerScraper{
		client:   client,
		settings: receivertest.NewNopCreateSettings(),
		config:   Config{},
	}
	_ = scraper.shutdown(context.Background())
	client.AssertExpectations(t)
}

func TestBrokerScraper_Name(t *testing.T) {
	s := brokerScraper{}
	assert.Equal(t, s.Name(), brokersScraperName)
}

func TestBrokerScraper_createBrokerScraper(t *testing.T) {
	sc := sarama.NewConfig()
	newSaramaClient = mockNewSaramaClient
	bs, err := createBrokerScraper(context.Background(), Config{}, sc, receivertest.NewNopCreateSettings())
	assert.NoError(t, err)
	assert.NotNil(t, bs)
}

func TestBrokerScraperStart(t *testing.T) {
	newSaramaClient = mockNewSaramaClient
	sc := sarama.NewConfig()
	bs, err := createBrokerScraper(context.Background(), Config{}, sc, receivertest.NewNopCreateSettings())
	assert.NoError(t, err)
	assert.NotNil(t, bs)
	assert.NoError(t, bs.Start(context.Background(), nil))
}

func TestBrokerScraper_scrape_handles_client_error(t *testing.T) {
	newSaramaClient = func(addrs []string, conf *sarama.Config) (sarama.Client, error) {
		return nil, fmt.Errorf("new client failed")
	}
	sc := sarama.NewConfig()
	bs, err := createBrokerScraper(context.Background(), Config{}, sc, receivertest.NewNopCreateSettings())
	assert.NoError(t, err)
	assert.NotNil(t, bs)
	_, err = bs.Scrape(context.Background())
	assert.Error(t, err)
}

func TestBrokerScraper_shutdown_handles_nil_client(t *testing.T) {
	newSaramaClient = func(addrs []string, conf *sarama.Config) (sarama.Client, error) {
		return nil, fmt.Errorf("new client failed")
	}
	sc := sarama.NewConfig()
	bs, err := createBrokerScraper(context.Background(), Config{}, sc, receivertest.NewNopCreateSettings())
	assert.NoError(t, err)
	assert.NotNil(t, bs)
	err = bs.Shutdown(context.Background())
	assert.NoError(t, err)
}

func TestBrokerScraper_scrape(t *testing.T) {
	newSaramaClient = func(addrs []string, conf *sarama.Config) (sarama.Client, error) {
		return nil, fmt.Errorf("new client failed")
	}
	sc := sarama.NewConfig()
	bs, err := createBrokerScraper(context.Background(), Config{}, sc, receivertest.NewNopCreateSettings())
	assert.NoError(t, err)
	assert.NotNil(t, bs)
	err = bs.Shutdown(context.Background())
	assert.NoError(t, err)
}

func TestBrokersScraper_createBrokerScraper(t *testing.T) {
	sc := sarama.NewConfig()
	newSaramaClient = mockNewSaramaClient
	bs, err := createBrokerScraper(context.Background(), Config{}, sc, receivertest.NewNopCreateSettings())
	assert.NoError(t, err)
	assert.NotNil(t, bs)
}
