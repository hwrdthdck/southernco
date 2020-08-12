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
	"github.com/aws/aws-sdk-go/aws"
	semconventions "go.opentelemetry.io/collector/translator/conventions"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/common/awsxray"
)

func makeSQL(attributes map[string]string) (map[string]string, *awsxray.SQLData) {
	var (
		filtered    = make(map[string]string)
		sqlData     awsxray.SQLData
		dbURL       string
		dbType      string
		dbInstance  string
		dbStatement string
		dbUser      string
	)

	for key, value := range attributes {
		switch key {
		case semconventions.AttributeDBConnectionString:
			dbURL = value
		case semconventions.AttributeDBSystem:
			dbType = value
		case semconventions.AttributeDBName:
			dbInstance = value
		case semconventions.AttributeDBStatement:
			dbStatement = value
		case semconventions.AttributeDBUser:
			dbUser = value
		default:
			filtered[key] = value
		}
	}

	if dbType != "sql" {
		// Either no DB attributes or this is not an SQL DB.
		return attributes, nil
	}

	if dbURL == "" {
		dbURL = "localhost"
	}
	url := dbURL + "/" + dbInstance
	sqlData = awsxray.SQLData{
		URL:            aws.String(url),
		DatabaseType:   aws.String(dbType),
		User:           aws.String(dbUser),
		SanitizedQuery: aws.String(dbStatement),
	}
	return filtered, &sqlData
}
