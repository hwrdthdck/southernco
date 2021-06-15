// Copyright  OpenTelemetry Authors
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

package testutils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"

	cinfo "github.com/google/cadvisor/info/v1"
)

func LoadContainerInfo(file string) []*cinfo.ContainerInfo {
	info, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("Fail to read file content: %s", err)
	}

	var result []*cinfo.ContainerInfo
	containers := map[string]*cinfo.ContainerInfo{}
	err = json.Unmarshal(info, &containers)

	if err != nil {
		log.Printf("Fail to parse json string: %s", err)
	}

	for _, containerInfo := range containers {
		result = append(result, containerInfo)
	}

	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	enc.Encode(result)
	return result
}

type MockCPUMemInfo struct {
}

func (m MockCPUMemInfo) GetNumCores() int64 {
	return 2
}

func (m MockCPUMemInfo) GetMemoryCapacity() int64 {
	return 1073741824
}

type MockHostInfo struct {
	MockCPUMemInfo
	ClusterName string
}

func (m MockHostInfo) GetClusterName() string {
	return m.ClusterName
}
