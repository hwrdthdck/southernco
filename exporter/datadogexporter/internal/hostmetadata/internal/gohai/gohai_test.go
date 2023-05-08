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
//
// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package gohai

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestGetPayload(t *testing.T) {
	logger := zap.NewNop()
	gohai := NewPayload(logger)
	assert.NotNil(t, gohai.Gohai.gohai.CPU)
	assert.NotNil(t, gohai.Gohai.gohai.FileSystem)
	assert.NotNil(t, gohai.Gohai.gohai.Memory)
	assert.NotNil(t, gohai.Gohai.gohai.Network)
	assert.NotNil(t, gohai.Gohai.gohai.Platform)
}
