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

package filestorage // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage/filestorage"

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"time"

	"go.opentelemetry.io/collector/config"
)

// Config defines configuration for file storage extension.
type Config struct {
	config.ExtensionSettings `mapstructure:",squash"`

	Directory string        `mapstructure:"directory,omitempty"`
	Timeout   time.Duration `mapstructure:"timeout,omitempty"`

	Compaction *CompactionConfig `mapstructure:"compaction,omitempty"`
}

// CompactionConfig defines configuration for optional file storage compaction.
type CompactionConfig struct {
	// OnStart specifies that compaction is attempted each time on start
	OnStart bool `mapstructure:"on_start,omitempty"`
	// OnRebound specifies that compaction is attempted online, when rebound conditions are met.
	// This typically happens when storage usage has increased, which caused increase in space allocation
	// and afterwards it had most items removed. We want to run the compaction online only when there are
	// not too many elements still being stored (which is an indication that "heavy usage" period is over)
	// so compaction should be relatively fast and at the same time there is relatively large volume of space
	// that might be reclaimed.
	OnRebound bool `mapstructure:"on_rebound,omitempty"`
	// Directory specifies where the temporary files for compaction will be stored
	Directory string `mapstructure:"directory,omitempty"`
	// ReboundSizeBelowMiB specifies the maximum actually used size for online compaction to happen
	ReboundSizeBelowMiB int64 `mapstructure:"rebound_size_below_mib"`
	// ReboundTotalSizeAboveMiB specifies the minimum total allocated size (both used and empty) for online compaction
	ReboundTotalSizeAboveMiB int64 `mapstructure:"rebound_total_size_above_mib"`
	// MaxTransactionSize specifies the maximum number of items that might be present in single compaction iteration
	MaxTransactionSize int64 `mapstructure:"max_transaction_size,omitempty"`
}

func (cfg *Config) Validate() error {
	var dirs []string
	if cfg.Compaction.OnStart {
		dirs = []string{cfg.Directory, cfg.Compaction.Directory}
	} else {
		dirs = []string{cfg.Directory}
	}
	for _, dir := range dirs {
		info, err := os.Stat(dir)
		if err != nil {
			if os.IsNotExist(err) {
				return fmt.Errorf("directory must exist: %v", err)
			}
			if fsErr, ok := err.(*fs.PathError); ok {
				return fmt.Errorf(
					"problem accessing configured directory: %s, err: %v",
					dir, fsErr,
				)
			}

		}
		if !info.IsDir() {
			return fmt.Errorf("%s is not a directory", dir)
		}
	}

	if cfg.Compaction.MaxTransactionSize < 0 {
		return errors.New("max transaction size for compaction cannot be less than 0")
	}

	return nil
}
