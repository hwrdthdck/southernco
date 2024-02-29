// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package datasource // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/datasource"

import (
	"context"

	"cloud.google.com/go/spanner"
	"google.golang.org/api/option"
)

const fgacRole = "spanner_sys_reader"

type Database struct {
	client     *spanner.Client
	databaseID *DatabaseID
}

func (database *Database) Client() *spanner.Client {
	return database.client
}

func (database *Database) DatabaseID() *DatabaseID {
	return database.databaseID
}

func NewDatabase(ctx context.Context, databaseID *DatabaseID, credentialsFilePath string, useFGACDatabaseRole bool) (*Database, error) {
	var client *spanner.Client
	var err error

	if credentialsFilePath != "" {
		credentialsFileClientOption := option.WithCredentialsFile(credentialsFilePath)
		client, err = spanner.NewClient(ctx, databaseID.ID(), credentialsFileClientOption)
	} else if useFGACDatabaseRole {
		cfg := spanner.ClientConfig{
			DatabaseRole: fgacRole,
		}
		client, err = spanner.NewClientWithConfig(ctx, databaseID.ID(), cfg)
	} else {
		// Fallback to Application Default Credentials(https://google.aip.dev/auth/4110)
		client, err = spanner.NewClient(ctx, databaseID.ID())
	}

	if err != nil {
		return nil, err
	}

	return NewDatabaseFromClient(client, databaseID), nil
}

func NewDatabaseFromClient(client *spanner.Client, databaseID *DatabaseID) *Database {
	return &Database{
		client:     client,
		databaseID: databaseID,
	}
}
