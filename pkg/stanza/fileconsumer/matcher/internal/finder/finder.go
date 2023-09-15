// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package finder // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/matcher/internal/finder"

import (
	"fmt"

	"github.com/bmatcuk/doublestar/v4"
)

func Validate(globs []string) error {
	for _, glob := range globs {
		_, err := doublestar.PathMatch(glob, "matchstring")
		if err != nil {
			return fmt.Errorf("parse glob: %w", err)
		}
	}
	return nil
}

// FindFiles gets a list of paths given an array of glob patterns to include and exclude
func FindFiles(includes []string, excludes []string) ([]string, error) {
	all := make([]string, 0, len(includes))
	for _, include := range includes {
		matches, err := doublestar.FilepathGlob(include, doublestar.WithFilesOnly(), doublestar.WithFailOnIOErrors())
		if err != nil {
			return all, err
		}
	INCLUDE:
		for _, match := range matches {
			for _, exclude := range excludes {
				if itMatches, _ := doublestar.PathMatch(exclude, match); itMatches {
					continue INCLUDE
				}
			}

			for _, existing := range all {
				if existing == match {
					continue INCLUDE
				}
			}

			all = append(all, match)
		}
	}

	return all, nil
}
