// Copyright 2020 OpenTelemetry Authors
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

// +build tools

package tools

// This file follows the recommendation at
// https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
// on how to pin tooling dependencies to a go.mod file.
// This ensures that all systems use the same version of tools in addition to regular dependencies.

import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/google/addlicense"
	_ "github.com/google/uuid"
	_ "github.com/onsi/ginkgo"
	_ "github.com/onsi/gomega"
	_ "github.com/securego/gosec/v2"
	_ "github.com/securego/gosec/v2/report/sarif"
	_ "github.com/securego/gosec/v2/report/text"
	_ "github.com/tcnksm/ghr"
	_ "github.com/vektra/mockery"
)
