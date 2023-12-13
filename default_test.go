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

// go test -v -cover -count=1 -test.cpu=1 -run=^TestSetDefault$
func TestSetDefault(t *testing.T) {
	defaultLogger.Store(NewLogger())

	logger := NewLogger()
	SetDefault(logger)

	gotLogger, ok := defaultLogger.Load().(*Logger)
	if !ok {
		t.Fatalf("logger type %T is wrong", defaultLogger.Load())
	}

	if gotLogger != logger {
		t.Fatalf("gotLogger %+v != logger %+v", gotLogger, logger)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestDefault$
func TestDefault(t *testing.T) {
	logger := NewLogger()
	defaultLogger.Store(logger)

	gotLogger := Default()
	if gotLogger != logger {
		t.Fatalf("gotLogger %+v != logger %+v", gotLogger, logger)
	}
}
