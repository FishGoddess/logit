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
	"fmt"
	"log/slog"
	"strings"
)

const (
	levelDebug = slog.LevelDebug
	levelInfo  = slog.LevelInfo
	levelWarn  = slog.LevelWarn
	levelError = slog.LevelError

	// levelPrint is for some logging methods which are compatible with log package in Go.
	levelPrint = levelError + 1

	// levelOff is for disabling logging.
	levelOff = levelError + 2
)

var (
	levels = map[slog.Level]string{
		levelDebug: "debug",
		levelInfo:  "info",
		levelWarn:  "warn",
		levelError: "error",
		levelPrint: "print",
		levelOff:   "off",
	}
)

// stringifyLevel returns the string form of level.
func stringifyLevel(level slog.Level) string {
	if name, ok := levels[level]; ok {
		return name
	}

	return "unknown"
}

// ParseLevel parses str to level and returns an error if failed.
func ParseLevel(str string) (slog.Level, error) {
	str = strings.ToLower(str)

	for level, name := range levels {
		if str == name {
			return level, nil
		}
	}

	return 0, fmt.Errorf("logit: unknown level %s", str)
}
