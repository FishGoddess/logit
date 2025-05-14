// Copyright 2025 FishGoddess. All Rights Reserved.
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
	"time"
)

// go test -v -cover -count=1 -test.cpu=1 -run=^TestWithMaxSize$
func TestWithMaxSize(t *testing.T) {
	conf := newDefaultConfig(t.Name())
	conf.maxSize = 0

	WithMaxSize(4 * 1024).applyTo(conf)

	want := newDefaultConfig(t.Name())
	want.maxSize = 4 * 1024

	if *conf != *want {
		t.Fatalf("conf %+v != want %+v", conf, want)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestWithMaxAge$
func TestWithMaxAge(t *testing.T) {
	conf := newDefaultConfig(t.Name())
	conf.maxAge = 0

	WithMaxAge(24 * time.Hour).applyTo(conf)

	want := newDefaultConfig(t.Name())
	want.maxAge = 24 * time.Hour

	if *conf != *want {
		t.Fatalf("conf %+v != want %+v", conf, want)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestWithMaxBackups$
func TestWithMaxBackups(t *testing.T) {
	conf := newDefaultConfig(t.Name())
	conf.maxBackups = 0

	WithMaxBackups(30).applyTo(conf)

	want := newDefaultConfig(t.Name())
	want.maxBackups = 30

	if *conf != *want {
		t.Fatalf("conf %+v != want %+v", conf, want)
	}
}
