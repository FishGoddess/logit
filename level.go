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
	LevelDebug = Level(slog.LevelDebug)
	LevelInfo  = Level(slog.LevelInfo)
	LevelWarn  = Level(slog.LevelWarn)
	LevelError = Level(slog.LevelError)

	// LevelPrint is for some logging methods which are compatible with log package in Go.
	LevelPrint Level = LevelError + 1

	// LevelOff is for disabling logging.
	LevelOff Level = LevelError + 2
)

var (
	levels = map[Level]string{
		LevelDebug: "debug",
		LevelInfo:  "info",
		LevelWarn:  "warn",
		LevelError: "error",
		LevelPrint: "print",
		LevelOff:   "off",
	}
)

// Level is an alias to slog.Level.
// We extends level in order to fit our logging methods.
type Level slog.Level

// Peel returns the level in slog.Level form.
func (l Level) Peel() slog.Level {
	return slog.Level(l)
}

// String returns the string form of level.
func (l Level) String() string {
	if name, ok := levels[l]; ok {
		return name
	}

	return "unknown"
}

// ParseLevel parses str to level and returns an error if failed.
func ParseLevel(str string) (Level, error) {
	str = strings.ToLower(str)

	for level, name := range levels {
		if str == name {
			return level, nil
		}
	}

	return 0, fmt.Errorf("logit: unknown level %s", str)
}
