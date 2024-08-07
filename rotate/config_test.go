// Copyright 2024 FishGoddess. All Rights Reserved.
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

package rotate

import (
	"testing"
)

// go test -v -cover -count=1 -test.cpu=1 -run=^TestNewDefaultConfig$
func TestNewDefaultConfig(t *testing.T) {
	c := newDefaultConfig()

	want := config{
		timeFormat: "20060102150405",
		maxSize:    128 * MB,
		maxAge:     60 * Day,
		maxBackups: 90,
	}

	if c != want {
		t.Fatalf("c %+v != want %+v", c, want)
	}
}
