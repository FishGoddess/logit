// Copyright 2022 FishGoddess. All Rights Reserved.
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

package file

import (
	"testing"

	"github.com/FishGoddess/logit/support/size"
)

// go test -v -cover -run=^TestNewDefaultConfig$
func TestNewDefaultConfig(t *testing.T) {
	c := newDefaultConfig()

	want := config{
		mode:       0644,
		dirMode:    0755,
		timeFormat: "20060102150405",
		maxSize:    256 * size.MB,
		maxAge:     0,
		maxBackups: 0,
	}

	if c != want {
		t.Errorf("c %+v != want %+v", c, want)
	}
}
