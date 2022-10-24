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

package logit

import (
	"testing"

	"github.com/FishGoddess/logit/support/global"
)

// go test -v -cover -run=^TestNewDefaultConfig$
func TestNewDefaultConfig(t *testing.T) {
	defaultConfig := newDefaultConfig()

	cfg := config{
		level:       debugLevel,
		withPID:     false,
		withCaller:  false,
		msgKey:      "log.msg",
		timeKey:     "log.time",
		levelKey:    "log.level",
		pidKey:      "log.pid",
		fileKey:     "log.file",
		lineKey:     "log.line",
		funcKey:     "log.func",
		errorKey:    "log.err",
		timeFormat:  "2006-01-02 15:04:05",
		callerDepth: global.CallerDepth,
	}

	if defaultConfig != cfg {
		t.Errorf("default config %+v is wrong", defaultConfig)
	}
}
