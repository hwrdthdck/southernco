// Copyright  The OpenTelemetry Authors
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

package processscraper

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/processor/filterset"
)

func TestIncludeCommandNameFilter(t *testing.T) {
	var filterConfig FilterConfig

	filterConfig.IncludeCommands.Commands = []string{"stringMatch"}
	filterConfig.IncludeCommands.Config = filterset.Config{MatchType: filterset.Strict}
	filter, err := createFilter(filterConfig)
	assert.Nil(t, err)

	match := filter.includeCommand("stringMatch", "")
	assert.True(t, match)

	match = filter.includeCommand("noMatch", "args")
	assert.False(t, match)

	// Test Regular Expression
	filterConfig.IncludeCommands.Commands = []string{"^([a-zA-Z0-9_\\-\\.]+)TestRegex"}
	filterConfig.IncludeCommands.Config = filterset.Config{MatchType: filterset.Regexp}
	filter, err = createFilter(filterConfig)
	assert.Nil(t, err)

	match = filter.includeCommand("astring"+"TestRegex", "--command args")
	assert.True(t, match)

	match = filter.includeCommand("astring"+"FailTest", "--command args")
	assert.False(t, match)

	// Test Regular Expression with command line
	filterConfig.IncludeCommands.Commands = []string{"^([a-zA-Z0-9_\\-\\.]+)TestRegex"}
	filterConfig.IncludeCommands.Config = filterset.Config{MatchType: filterset.Regexp}
	filterConfig.IncludeCommandLines.CommandLines = []string{"pas*word"}
	filterConfig.IncludeCommandLines.Config = filterset.Config{MatchType: filterset.Regexp}
	filter, err = createFilter(filterConfig)
	assert.Nil(t, err)

	match = filter.includeCommand("astring"+"TestRegex", "passsssssword")
	assert.True(t, match)

	match = filter.includeCommand("NoMatchCommand", "passsssssword")
	assert.False(t, match)

	match = filter.includeCommand("astring"+"TestRegex", "NoMatchCommandLine")
	assert.False(t, match)

}

func TestExcludeCommandNameFilter(t *testing.T) {
	var filterConfig FilterConfig

	filterConfig.ExcludeCommands.Commands = []string{"stringMatch"}
	filterConfig.ExcludeCommands.Config = filterset.Config{MatchType: filterset.Strict}
	filter, err := createFilter(filterConfig)
	assert.Nil(t, err)

	match := filter.includeCommand("stringMatch", "")
	assert.False(t, match)

	match = filter.includeCommand("noMatch", "args")
	assert.True(t, match)

	// Test Regular Expression
	filterConfig.ExcludeCommands.Commands = []string{"^([a-zA-Z0-9_\\-\\.]+)TestRegex"}
	filterConfig.ExcludeCommands.Config = filterset.Config{MatchType: filterset.Regexp}
	filter, err = createFilter(filterConfig)
	assert.Nil(t, err)

	match = filter.includeCommand("astring"+"TestRegex", "--command args")
	assert.False(t, match)

	match = filter.includeCommand("astring"+"FailTest", "--command args")
	assert.True(t, match)

	// Test Regular Expression with quotes
	filterConfig.ExcludeCommands.Commands = []string{"^([a-zA-Z0-9_\\-\\.]+)TestRegex"}
	filterConfig.ExcludeCommands.Config = filterset.Config{MatchType: filterset.Regexp}
	filter, err = createFilter(filterConfig)
	assert.Nil(t, err)

	match = filter.includeCommand("astring"+"TestRegex", "--command args")
	assert.False(t, match)

	match = filter.includeCommand("astring"+"FailTest", "--command args")
	assert.True(t, match)

	// Test Regular Expression with command line
	filterConfig.ExcludeCommands.Commands = []string{"^([a-zA-Z0-9_\\-\\.]+)TestRegex"}
	filterConfig.ExcludeCommands.Config = filterset.Config{MatchType: filterset.Regexp}
	filterConfig.IncludeCommandLines.CommandLines = []string{"pas*word"}
	filterConfig.IncludeCommandLines.Config = filterset.Config{MatchType: filterset.Regexp}
	filter, err = createFilter(filterConfig)
	assert.Nil(t, err)

	match = filter.includeCommand("astring"+"TestRegex", "passsssssword")
	assert.False(t, match)

	match = filter.includeCommand("RandomCommand", "passsssssword")
	assert.True(t, match)

}

func TestPid(t *testing.T) {
	var filterConfig FilterConfig

	// test include
	filterConfig.IncludePids = []int32{123454}
	filter, err := createFilter(filterConfig)
	assert.Nil(t, err)

	match := filter.includePid(123454)
	assert.True(t, match)

	match = filter.includePid(11111)
	assert.False(t, match)

	// test exclude
	filterConfig.IncludePids = []int32{}
	filterConfig.ExcludePids = []int32{123454}
	filter, err = createFilter(filterConfig)
	assert.Nil(t, err)

	match = filter.includePid(123454)
	assert.False(t, match)

	match = filter.includePid(11111)
	assert.True(t, match)
}

func TestOwner(t *testing.T) {
	var filterConfig FilterConfig

	filterConfig.IncludeOwners.Owners = []string{"owner"}
	filterConfig.IncludeOwners.Config = filterset.Config{MatchType: filterset.Strict}
	filter, err := createFilter(filterConfig)
	assert.Nil(t, err)

	match := filter.includeOwner("owner")
	assert.True(t, match)

	match = filter.includeOwner("wrongowner")
	assert.False(t, match)

	filterConfig.IncludeOwners.Owners = []string{"^owner"}
	filterConfig.IncludeOwners.Config = filterset.Config{MatchType: filterset.Regexp}
	filter, err = createFilter(filterConfig)
	assert.Nil(t, err)

	match = filter.includeOwner("ownerOwner")
	assert.True(t, match)

	match = filter.includeOwner("notowner")
	assert.False(t, match)
}

func TestExecutable(t *testing.T) {
	var filterConfig FilterConfig

	filterConfig.IncludeExecutableNames.ExecutableNames = []string{"executableName"}
	filterConfig.IncludeExecutableNames.Config = filterset.Config{MatchType: filterset.Strict}
	filter, err := createFilter(filterConfig)
	assert.Nil(t, err)

	match := filter.includeExecutable("executableName", "//executable//path")
	assert.True(t, match)

	match = filter.includeExecutable("noMatch", "//executable//path")
	assert.False(t, match)

	filterConfig.IncludeExecutableNames = ExecutableNameMatchConfig{}
	filterConfig.IncludeExecutablePaths.ExecutablePaths = []string{"//executable//path"}
	filterConfig.IncludeExecutablePaths.Config = filterset.Config{MatchType: filterset.Strict}
	filter, err = createFilter(filterConfig)
	assert.Nil(t, err)

	match = filter.includeExecutable("executableName", "//executable//path")
	assert.True(t, match)

	match = filter.includeExecutable("executableName", "//nomatch//path")
	assert.False(t, match)

	filterConfig.IncludeExecutableNames.ExecutableNames = []string{"executableName"}
	filterConfig.IncludeExecutableNames.Config = filterset.Config{MatchType: filterset.Strict}
	filterConfig.IncludeExecutablePaths.ExecutablePaths = []string{"//executable//path"}
	filterConfig.IncludeExecutablePaths.Config = filterset.Config{MatchType: filterset.Strict}
	filter, err = createFilter(filterConfig)
	assert.Nil(t, err)

	match = filter.includeExecutable("executableName", "//executable//path")
	assert.True(t, match)

	match = filter.includeExecutable("executableName", "//nomatch//path")
	assert.False(t, match)

	match = filter.includeExecutable("noMatch", "//executable//path")
	assert.False(t, match)
}
