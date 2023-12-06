// Copyright 2023 FishGoddess. All Rights Reserved.
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

// go test -v -cover -run=^TestWithTimeFormat$
func TestWithTimeFormat(t *testing.T) {
	c := newDefaultConfig()
	c.timeFormat = ""

	WithTimeFormat("20060102").apply(&c)

	want := newDefaultConfig()
	want.timeFormat = "20060102"

	if c != want {
		t.Fatalf("c %+v != want %+v", c, want)
	}
}

// go test -v -cover -run=^TestWithMaxSize$
func TestWithMaxSize(t *testing.T) {
	c := newDefaultConfig()
	c.maxSize = 0

	WithMaxSize(4 * 1024).apply(&c)

	want := newDefaultConfig()
	want.maxSize = 4 * 1024

	if c != want {
		t.Fatalf("c %+v != want %+v", c, want)
	}
}

// go test -v -cover -run=^TestWithMaxAge$
func TestWithMaxAge(t *testing.T) {
	c := newDefaultConfig()
	c.maxAge = 0

	WithMaxAge(24 * time.Hour).apply(&c)

	want := newDefaultConfig()
	want.maxAge = 24 * time.Hour

	if c != want {
		t.Fatalf("c %+v != want %+v", c, want)
	}
}

// go test -v -cover -run=^TestWithMaxBackups$
func TestWithMaxBackups(t *testing.T) {
	c := newDefaultConfig()
	c.maxBackups = 0

	WithMaxBackups(30).apply(&c)

	want := newDefaultConfig()
	want.maxBackups = 30

	if c != want {
		t.Fatalf("c %+v != want %+v", c, want)
	}
}
