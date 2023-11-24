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
	"strings"
)

const (
	debugLevel level = slog.LevelDebug
	infoLevel  level = slog.LevelInfo
	warnLevel  level = slog.LevelWarn
	errorLevel level = slog.LevelError
	printLevel level = errorLevel + 1
	offLevel   level = errorLevel + 2
)

var (
	levels = map[string]level{
		"debug": debugLevel,
		"info":  infoLevel,
		"warn":  warnLevel,
		"error": errorLevel,
		"print": printLevel,
		"off":   offLevel,
	}
)

type level = slog.Level

func parseLevel(str string) level {
	str = strings.TrimSpace(str)
	str = strings.ToLower(str)

	switch str {
	case "debug":
		return debugLevel
	case "info":
		return infoLevel
	case "warn":
		return warnLevel
	case "error":
		return errorLevel
	case "print":
		return printLevel
	case "off":
		return offLevel
	default:
		return debugLevel
	}
}
