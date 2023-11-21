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
	"log/slog"
	"testing"
)

// go test -v -cover -run=^TestWithDebugLevel$
func TestWithDebugLevel(t *testing.T) {
	conf := newDefaultConfig()
	conf.level = slog.LevelError
	conf.Accept(WithDebugLevel())

	if conf.level != slog.LevelDebug {
		t.Errorf("conf's level %+v is wrong", conf.level)
	}
}

// go test -v -cover -run=^TestWithInfoLevel$
func TestWithInfoLevel(t *testing.T) {
	conf := newDefaultConfig()
	conf.level = slog.LevelDebug
	conf.Accept(WithInfoLevel())

	if conf.level != slog.LevelInfo {
		t.Errorf("conf's level %+v is wrong", conf.level)
	}
}

// go test -v -cover -run=^TestWithWarnLevel$
func TestWithWarnLevel(t *testing.T) {
	conf := newDefaultConfig()
	conf.level = slog.LevelDebug
	conf.Accept(WithWarnLevel())

	if conf.level != slog.LevelWarn {
		t.Errorf("conf's level %+v is wrong", conf.level)
	}
}

// go test -v -cover -run=^TestWithErrorLevel$
func TestWithErrorLevel(t *testing.T) {
	conf := newDefaultConfig()
	conf.level = slog.LevelDebug
	conf.Accept(WithErrorLevel())

	if conf.level != slog.LevelError {
		t.Errorf("conf's level %+v is wrong", conf.level)
	}
}

// go test -v -cover -run=^TestWithSource$
func TestWithSource(t *testing.T) {
	conf := newDefaultConfig()
	conf.withSource = false
	conf.Accept(WithSource())

	if conf.withSource != true {
		t.Errorf("conf's withSource %+v is wrong", conf.withSource)
	}
}

// go test -v -cover -run=^TestWithPID$
func TestWithPID(t *testing.T) {
	conf := newDefaultConfig()
	conf.withPID = false
	conf.Accept(WithPID())

	if conf.withPID != true {
		t.Errorf("conf's withPID %+v is wrong", conf.withPID)
	}
}
