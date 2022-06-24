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

package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	changelogMD = "CHANGELOG.md"
	chlogDir    = "changelog"
	exampleYAML = "EXAMPLE.yaml"
	tmpMD       = "changelog.tmp.MD"

	breaking     = "breaking"
	deprecation  = "deprecation"
	newComponent = "new_component"
	enhancement  = "enhancement"
	bugFix       = "bug_fix"

	breakingLabel     = "## 🛑 Breaking changes 🛑"
	deprecationLabel  = "### 🚩 Deprecations 🚩"
	newComponentLabel = "### 🚀 New components 🚀"
	enhancementLabel  = "### 💡 Enhancements 💡"
	bugFixLabel       = "### 🧰 Bug fixes 🧰"

	insertPoint = "<!-- next version -->"
)

func main() {
	if len(os.Args) != 2 {
		usage()
		os.Exit(1)
	}

	switch command := os.Args[1]; command {
	case "validate":
		if err := validate(); err != nil {
			fmt.Printf("FAIL: validate: %v\n", err)
			os.Exit(1)
		}
	case "preview":
		if err := preview(); err != nil {
			fmt.Printf("FAIL: preview: %v\n", err)
			os.Exit(1)
		}
	case "update":
		if err := update(); err != nil {
			fmt.Printf("FAIL: update: %v\n", err)
			os.Exit(1)
		}
	default:
		usage()
		os.Exit(1)
	}
}

func usage() {
	fmt.Println("usage: chloggen validate")
	fmt.Println("       chloggen preview")
	fmt.Println("       chloggen update")
}

func validate() error {
	entries, err := readEntries(false)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if err = entry.Validate(); err != nil {
			return err
		}
	}
	fmt.Printf("PASS: all files in ./%s/ are valid\n", chlogDir)
	return nil
}

func preview() error {
	entries, err := readEntries(true)
	if err != nil {
		return err
	}

	if len(entries) == 0 {
		fmt.Println("No changelog updates")
		return nil
	}

	chlogUpdate := newEntriesByTypeMap(entries).toChlogString()
	fmt.Printf("Generated changelog updates:")
	fmt.Println(chlogUpdate)
	return nil
}

func update() error {
	entries, err := readEntries(true)
	if err != nil {
		return err
	}

	if len(entries) == 0 {
		return fmt.Errorf("no entries to add to the changelog")
	}

	chlogUpdate := newEntriesByTypeMap(entries).toChlogString()

	oldChlogBytes, err := os.ReadFile(changelogMD)
	if err != nil {
		return err
	}
	chlogParts := bytes.Split(oldChlogBytes, []byte(insertPoint))
	if len(chlogParts) != 2 {
		return fmt.Errorf("expected one instance of %s", insertPoint)
	}

	chlogHeader, chlogHistory := string(chlogParts[0]), string(chlogParts[1])

	var chlogBuilder strings.Builder
	chlogBuilder.WriteString(chlogHeader)
	chlogBuilder.WriteString(insertPoint)
	chlogBuilder.WriteString(chlogUpdate)
	chlogBuilder.WriteString(chlogHistory)

	tmpChlogFile, err := os.Create(tmpMD)
	if err != nil {
		return err
	}
	defer tmpChlogFile.Close()

	tmpChlogFile.WriteString(chlogBuilder.String())

	if err := os.Rename(tmpMD, changelogMD); err != nil {
		return err
	}

	fmt.Printf("Finished updating %s\n", changelogMD)

	return deleteEntries()
}

func readEntries(excludeExample bool) ([]*Entry, error) {
	entryFiles, err := filepath.Glob(filepath.Join(chlogDir, "*.yaml"))
	if err != nil {
		return nil, err
	}

	// TODO aggregate errors and return valid entries
	entries := make([]*Entry, 0, len(entryFiles))
	for _, entryFile := range entryFiles {
		if excludeExample && filepath.Base(entryFile) == exampleYAML {
			continue
		}

		fileBytes, err := os.ReadFile(entryFile)
		if err != nil {
			return nil, err
		}

		entry := &Entry{}
		if err = yaml.Unmarshal(fileBytes, entry); err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}
	return entries, nil
}

func deleteEntries() error {
	entryFiles, err := filepath.Glob(filepath.Join(chlogDir, "*.yaml"))
	if err != nil {
		return err
	}

	for _, entryFile := range entryFiles {
		if filepath.Base(entryFile) == exampleYAML {
			continue
		}

		if err := os.Remove(entryFile); err != nil {
			fmt.Printf("Failed to delete: %s\n", entryFile)
		}
	}
	return nil
}

type entriesByTypeMap map[string][]*Entry

func newEntriesByTypeMap(entries []*Entry) entriesByTypeMap {
	entriesByType := make(map[string][]*Entry)
	for _, entry := range entries {
		switch entry.ChangeType {
		case breaking:
			entriesByType[breaking] = append(entriesByType[breaking], entry)
		case deprecation:
			entriesByType[deprecation] = append(entriesByType[deprecation], entry)
		case newComponent:
			entriesByType[newComponent] = append(entriesByType[newComponent], entry)
		case enhancement:
			entriesByType[enhancement] = append(entriesByType[enhancement], entry)
		case bugFix:
			entriesByType[bugFix] = append(entriesByType[bugFix], entry)
		}
	}
	return entriesByType
}

func (m entriesByTypeMap) toChlogString() string {
	var sb strings.Builder

	sb.WriteString("\n\n## vTODO")

	if len(m[breaking]) > 0 {
		sb.WriteString("\n\n")
		sb.WriteString(breakingLabel)
		sb.WriteString("\n")
		for _, entry := range m[breaking] {
			sb.WriteString("\n")
			sb.WriteString(entry.String())
		}
	}

	if len(m[deprecation]) > 0 {
		sb.WriteString("\n\n")
		sb.WriteString(deprecationLabel)
		sb.WriteString("\n")
		for _, entry := range m[deprecation] {
			sb.WriteString("\n")
			sb.WriteString(entry.String())
		}
	}

	if len(m[newComponent]) > 0 {
		sb.WriteString("\n\n")
		sb.WriteString(newComponentLabel)
		sb.WriteString("\n")
		for _, entry := range m[newComponent] {
			sb.WriteString("\n")
			sb.WriteString(entry.String())
		}
	}

	if len(m[enhancement]) > 0 {
		sb.WriteString("\n\n")
		sb.WriteString(enhancementLabel)
		sb.WriteString("\n")
		for _, entry := range m[enhancement] {
			sb.WriteString("\n")
			sb.WriteString(entry.String())
		}
	}

	if len(m[bugFix]) > 0 {
		sb.WriteString("\n\n")
		sb.WriteString(bugFixLabel)
		sb.WriteString("\n")
		for _, entry := range m[bugFix] {
			sb.WriteString("\n")
			sb.WriteString(entry.String())
		}
	}

	return sb.String()
}
