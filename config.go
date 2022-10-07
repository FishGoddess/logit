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
	"github.com/FishGoddess/logit/support/global"
)

// config stores all configurations used in Logger.
type config struct {
	level       level  // The level of a logger.
	withPID     bool   // Logs will carry pid if withPID is true.
	withCaller  bool   // Logs will carry caller information if withCaller is true.
	msgKey      string // The key of message in a log.
	timeKey     string // The key of time in a log.
	levelKey    string // The key of level in a log.
	pidKey      string // The key of pid in a log.
	fileKey     string // The key of caller's file in a log.
	lineKey     string // The key of caller's line in a log.
	funcKey     string // The key of caller's function in a log.
	timeFormat  string // The format of time in a log.
	callerDepth int    // The depth of caller.
}

// newDefaultConfig returns a default config.
func newDefaultConfig() config {
	return config{
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
		timeFormat:  "2006-01-02 15:04:05",
		callerDepth: global.CallerDepth,
	}
}
