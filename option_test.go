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

package logit

import (
	"testing"
)

// go test -v -cover -run=^TestWithDebugLevel$
func TestWithDebugLevel(t *testing.T) {
	conf := &config{level: LevelError}
	WithDebugLevel().applyTo(conf)

	if conf.level != LevelDebug {
		t.Errorf("conf.level %+v != LevelDebug", conf.level)
	}
}

// go test -v -cover -run=^TestWithInfoLevel$
func TestWithInfoLevel(t *testing.T) {
	conf := &config{level: LevelError}
	WithInfoLevel().applyTo(conf)

	if conf.level != LevelInfo {
		t.Errorf("conf.level %+v != LevelInfo", conf.level)
	}
}

// go test -v -cover -run=^TestWithWarnLevel$
func TestWithWarnLevel(t *testing.T) {
	conf := &config{level: LevelError}
	WithWarnLevel().applyTo(conf)

	if conf.level != LevelWarn {
		t.Errorf("conf.level %+v != LevelWarn", conf.level)
	}
}

// go test -v -cover -run=^TestWithErrorLevel$
func TestWithErrorLevel(t *testing.T) {
	conf := &config{level: LevelDebug}
	WithErrorLevel().applyTo(conf)

	if conf.level != LevelError {
		t.Errorf("conf.level %+v != LevelError", conf.level)
	}
}

// go test -v -cover -run=^TestWithPrintLevel$
func TestWithPrintLevel(t *testing.T) {
	conf := &config{level: LevelError}
	WithPrintLevel().applyTo(conf)

	if conf.level != LevelPrint {
		t.Errorf("conf.level %+v != LevelPrint", conf.level)
	}
}

// go test -v -cover -run=^TestWithOffLevel$
func TestWithOffLevel(t *testing.T) {
	conf := &config{level: LevelError}
	WithOffLevel().applyTo(conf)

	if conf.level != LevelOff {
		t.Errorf("conf.level %+v != LevelOff", conf.level)
	}
}
