/*
Copyright The ORAS Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package platform

import (
	"encoding/json"
	"testing"

	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
)

func TestMatchPlatform(t *testing.T) {
	tests := []struct {
		curr   ocispec.Platform
		target ocispec.Platform
		want   bool
	}{{
		ocispec.Platform{Architecture: "amd64", OS: "linux"},
		ocispec.Platform{Architecture: "amd64", OS: "linux"},
		true,
	}, {
		ocispec.Platform{Architecture: "amd64", OS: "linux"},
		ocispec.Platform{Architecture: "amd64", OS: "LINUX"},
		false,
	}, {
		ocispec.Platform{Architecture: "amd64", OS: "linux"},
		ocispec.Platform{Architecture: "arm64", OS: "linux"},
		false,
	}, {
		ocispec.Platform{Architecture: "arm", OS: "linux"},
		ocispec.Platform{Architecture: "arm", OS: "linux", Variant: "v7"},
		false,
	}, {
		ocispec.Platform{Architecture: "arm", OS: "linux", Variant: "v7"},
		ocispec.Platform{Architecture: "arm", OS: "linux"},
		true,
	}, {
		ocispec.Platform{Architecture: "arm", OS: "linux", Variant: "v7"},
		ocispec.Platform{Architecture: "arm", OS: "linux", Variant: "v7"},
		true,
	}, {
		ocispec.Platform{Architecture: "amd64", OS: "windows", OSVersion: "10.0.20348.768"},
		ocispec.Platform{Architecture: "amd64", OS: "windows", OSVersion: "10.0.20348.700"},
		false,
	}, {
		ocispec.Platform{Architecture: "amd64", OS: "windows"},
		ocispec.Platform{Architecture: "amd64", OS: "windows", OSVersion: "10.0.20348.768"},
		false,
	}, {
		ocispec.Platform{Architecture: "amd64", OS: "windows", OSVersion: "10.0.20348.768"},
		ocispec.Platform{Architecture: "amd64", OS: "windows"},
		true,
	}, {
		ocispec.Platform{Architecture: "amd64", OS: "windows", OSVersion: "10.0.20348.768"},
		ocispec.Platform{Architecture: "amd64", OS: "windows", OSVersion: "10.0.20348.768"},
		true,
	}, {
		ocispec.Platform{Architecture: "arm", OS: "linux", OSFeatures: []string{"a", "d"}},
		ocispec.Platform{Architecture: "arm", OS: "linux", OSFeatures: []string{"a", "c"}},
		false,
	}, {
		ocispec.Platform{Architecture: "arm", OS: "linux"},
		ocispec.Platform{Architecture: "arm", OS: "linux", OSFeatures: []string{"a"}},
		false,
	}, {
		ocispec.Platform{Architecture: "arm", OS: "linux", OSFeatures: []string{"a"}},
		ocispec.Platform{Architecture: "arm", OS: "linux"},
		true,
	}, {
		ocispec.Platform{Architecture: "arm", OS: "linux", OSFeatures: []string{"a", "b"}},
		ocispec.Platform{Architecture: "arm", OS: "linux", OSFeatures: []string{"a", "b"}},
		true,
	}, {
		ocispec.Platform{Architecture: "arm", OS: "linux", OSFeatures: []string{"a", "d", "c", "b"}},
		ocispec.Platform{Architecture: "arm", OS: "linux", OSFeatures: []string{"d", "c", "a", "b"}},
		true,
	}}

	for _, tt := range tests {
		currPlatforJSON, _ := json.Marshal(tt.curr)
		targetPlatforJSON, _ := json.Marshal(tt.target)
		name := string(currPlatforJSON) + string(targetPlatforJSON)
		t.Run(name, func(t *testing.T) {
			if got := MatchPlatform(&tt.curr, &tt.target); got != tt.want {
				t.Errorf("MatchPlatform() = %v, want %v", got, tt.want)
			}
		})
	}
}
