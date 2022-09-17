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
	"time"

	"github.com/go-logit/logit/support/size"
)

// go test -v -cover -run=^TestWithMode$
func TestWithMode(t *testing.T) {
	c := newDefaultConfig()
	c.mode = 0

	WithMode(0600).Apply(&c)

	want := newDefaultConfig()
	want.mode = 0600

	if c != want {
		t.Errorf("c %+v != want %+v", c, want)
	}
}

// go test -v -cover -run=^TestWithDirMode$
func TestWithDirMode(t *testing.T) {
	c := newDefaultConfig()
	c.dirMode = 0

	WithDirMode(0700).Apply(&c)

	want := newDefaultConfig()
	want.dirMode = 0700

	if c != want {
		t.Errorf("c %+v != want %+v", c, want)
	}
}

// go test -v -cover -run=^TestWithTimeFormat$
func TestWithTimeFormat(t *testing.T) {
	c := newDefaultConfig()
	c.timeFormat = ""

	WithTimeFormat("20060102").Apply(&c)

	want := newDefaultConfig()
	want.timeFormat = "20060102"

	if c != want {
		t.Errorf("c %+v != want %+v", c, want)
	}
}

// go test -v -cover -run=^TestWithMaxSize$
func TestWithMaxSize(t *testing.T) {
	c := newDefaultConfig()
	c.maxSize = 0

	WithMaxSize(4 * size.KB).Apply(&c)

	want := newDefaultConfig()
	want.maxSize = 4 * size.KB

	if c != want {
		t.Errorf("c %+v != want %+v", c, want)
	}
}

// go test -v -cover -run=^TestWithMaxAge$
func TestWithMaxAge(t *testing.T) {
	c := newDefaultConfig()
	c.maxAge = 0

	WithMaxAge(24 * time.Hour).Apply(&c)

	want := newDefaultConfig()
	want.maxAge = 24 * time.Hour

	if c != want {
		t.Errorf("c %+v != want %+v", c, want)
	}
}

// go test -v -cover -run=^TestWithMaxBackups$
func TestWithMaxBackups(t *testing.T) {
	c := newDefaultConfig()
	c.maxBackups = 0

	WithMaxBackups(30).Apply(&c)

	want := newDefaultConfig()
	want.maxBackups = 30

	if c != want {
		t.Errorf("c %+v != want %+v", c, want)
	}
}
