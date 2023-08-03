// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	_ "embed"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"go.opentelemetry.io/collector/confmap/provider/fileprovider"
)

const codeownersHeader = `# Code generated by githubgen. DO NOT EDIT.
#####################################################
#
# List of approvers for OpenTelemetry Collector Contrib
#
#####################################################
#
# Learn about membership in OpenTelemetry community:
# https://github.com/open-telemetry/community/blob/main/community-membership.md
#
#
# Learn about CODEOWNERS file format:
# https://help.github.com/en/articles/about-code-owners
#
# NOTE: Lines should be entered in the following format:
# <component_path_relative_from_project_root>/<min_1_space><owner_1><space><owner_2><space>..<owner_n>
# extension/oauth2clientauthextension/                 @open-telemetry/collector-contrib-approvers @pavankrish123 @jpkrohling
# Path separator and minimum of 1 space between component path and owners is
# important for validation steps
#

* @open-telemetry/collector-contrib-approvers
`

const unmaintainedHeader = `

## UNMAINTAINED components
## The Github issue template generation code needs this to generate the corresponding labels.

`

const allowlistHeader = `# Code generated by githubgen. DO NOT EDIT.
#####################################################
#
# List of components in OpenTelemetry Collector Contrib
# waiting on owners to be assigned
#
#####################################################
#
# Learn about membership in OpenTelemetry community:
#  https://github.com/open-telemetry/community/blob/main/community-membership.md
#
#
# Learn about CODEOWNERS file format:
#  https://help.github.com/en/articles/about-code-owners
#

## 
# NOTE: New components MUST have a codeowner. Add new components to the CODEOWNERS file instead of here.
##

## COMMON & SHARED components
internal/common

`

const unmaintainedStatus = "unmaintained"

//go:embed members.txt
var membersData string

//go:embed allowlist.txt
var allowlistData string
var members []string
var allowlist []string

func init() {
	members = strings.Split(membersData, "\n")

	allowlist = strings.Split(allowlistData, "\n")
}

// Generates files specific to Github according to status metadata:
// .github/CODEOWNERS
// .github/ALLOWLIST
// Usage: go run cmd/githubgen/main.go <repository root> [--check]
func main() {
	flag.Parse()
	folder := flag.Arg(0)
	if err := run(folder, flag.Arg(1) == "--check"); err != nil {
		log.Fatal(err)
	}
}

type Codeowners struct {
	// Active codeowners
	Active []string `mapstructure:"active"`
	// Emeritus codeowners
	Emeritus []string `mapstructure:"emeritus"`
}
type Status struct {
	Stability     map[string][]string `mapstructure:"stability"`
	Distributions []string            `mapstructure:"distributions"`
	Class         string              `mapstructure:"class"`
	Warnings      []string            `mapstructure:"warnings"`
	Codeowners    *Codeowners         `mapstructure:"codeowners"`
}
type metadata struct {
	// Type of the component.
	Type string `mapstructure:"type"`
	// Type of the parent component (applicable to subcomponents).
	Parent string `mapstructure:"parent"`
	// Status information for the component.
	Status *Status `mapstructure:"status"`
}

func loadMetadata(filePath string) (metadata, error) {
	cp, err := fileprovider.New().Retrieve(context.Background(), "file:"+filePath, nil)
	if err != nil {
		return metadata{}, err
	}

	conf, err := cp.AsConf()
	if err != nil {
		return metadata{}, err
	}

	md := metadata{}
	if err := conf.Unmarshal(&md); err != nil {
		return md, err
	}

	return md, nil
}

func run(folder string, checkMembers bool) error {
	components := map[string]metadata{}
	foldersList := []string{}
	maxLength := 0
	allCodeowners := make(map[string]struct{}, 24)
	err := filepath.Walk(folder, func(path string, info fs.FileInfo, err error) error {
		if info.Name() == "metadata.yaml" {
			m, err := loadMetadata(path)
			if err != nil {
				return err
			}
			if m.Status == nil {
				return nil
			}
			key := filepath.Dir(path) + "/"
			components[key] = m
			foldersList = append(foldersList, key)
			for stability := range m.Status.Stability {
				if stability == unmaintainedStatus {
					// do not account for unmaintained status to change the max length of the component line.
					return nil
				}
			}
			for _, id := range m.Status.Codeowners.Active {
				allCodeowners[id] = struct{}{}
			}
			if len(key) > maxLength {
				maxLength = len(key)
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	sort.Strings(foldersList)
	var missingCodeowners []string
	for codeowner := range allCodeowners {
		present, err2 := hasMember(codeowner, checkMembers)
		if err2 != nil {
			return err2
		}

		if !present {
			allowed := inAllowlist(codeowner) || strings.HasPrefix(codeowner, "open-telemetry/")
			if !allowed {
				missingCodeowners = append(missingCodeowners, codeowner)
			}
		}
	}
	if len(missingCodeowners) > 0 {
		sort.Strings(missingCodeowners)
		return fmt.Errorf("codeowners are not members: %s", strings.Join(missingCodeowners, ", "))
	}
	if checkMembers {
		var list []string
		for codeowner := range allCodeowners {
			if strings.HasPrefix(codeowner, "open-telemetry/") {
				continue
			}
			list = append(list, codeowner)
		}
		sort.Strings(list)
		err = os.WriteFile(filepath.Join("cmd", "githubgen", "members.txt"), []byte(strings.Join(list, "\n")), 0600)
		if err != nil {
			return err
		}
	}

	codeowners := codeownersHeader
	deprecatedList := "## DEPRECATED components\n"
	unmaintainedList := "\n## UNMAINTAINED components\n"

	unmaintainedCodeowners := unmaintainedHeader
	currentFirstSegment := ""
LOOP:
	for _, key := range foldersList {
		m := components[key]
		for stability := range m.Status.Stability {
			if stability == unmaintainedStatus {
				unmaintainedList += key + "\n"
				unmaintainedCodeowners += fmt.Sprintf("%s%s @open-telemetry/collector-contrib-approvers \n", key, strings.Repeat(" ", maxLength-len(key)))
				continue LOOP
			}
			if stability == "deprecated" && (m.Status.Codeowners == nil || len(m.Status.Codeowners.Active) == 0) {
				deprecatedList += key + "\n"
			}
		}

		if m.Status.Codeowners != nil {
			parts := strings.Split(key, string(os.PathSeparator))
			firstSegment := parts[0]
			if firstSegment != currentFirstSegment {
				currentFirstSegment = firstSegment
				codeowners += "\n"
			}
			owners := ""
			for _, owner := range m.Status.Codeowners.Active {
				owners += " "
				owners += "@" + owner
			}
			codeowners += fmt.Sprintf("%s%s @open-telemetry/collector-contrib-approvers%s\n", key, strings.Repeat(" ", maxLength-len(key)), owners)
		}
	}

	err = os.WriteFile(filepath.Join(".github", "CODEOWNERS"), []byte(codeowners+unmaintainedCodeowners), 0600)
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath.Join(".github", "ALLOWLIST"), []byte(allowlistHeader+deprecatedList+unmaintainedList), 0600)
	if err != nil {
		return err
	}

	return nil
}

func hasMember(id string, checkGithub bool) (bool, error) {
	if checkGithub {
		present, err := checkGithubMembership(id)
		return present, err
	}
	for _, m := range members {
		if id == m {
			return true, nil
		}
	}
	return false, nil
}

func inAllowlist(id string) bool {
	for _, m := range allowlist {
		if id == m {
			return true
		}
	}
	return false
}

func checkGithubMembership(id string) (bool, error) {
	req, err := http.NewRequest("GET", "https://api.github.com/orgs/open-telemetry/members/"+id, nil)
	if err != nil {
		return false, err
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	if token := os.Getenv("GITHUB_TOKEN"); token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	}
	httpClient := http.DefaultClient

	res, err := httpClient.Do(req)
	if err == nil {
		_, _ = io.Copy(io.Discard, res.Body)
		_ = res.Body.Close()
	}
	return res.StatusCode == 204, err
}
