// Copyright 2020 Ye Zi Jie. All Rights Reserved.
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
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2020/03/01 14:18:33

package logit

import (
	"fmt"
	"math"
	"os"
)

// Level is the type representation of the logger level.
type Level uint8

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel

	// OffLevel is for disabling a logger
	OffLevel = math.MaxUint8
)

var (
	// levels store the names of all level provided.
	levels = map[Level]string{
		DebugLevel: "debug",
		InfoLevel:  "info",
		WarnLevel:  "warn",
		ErrorLevel: "error",
		OffLevel:   "off",
	}
)

// parseLevel parses level and returns the Level of it.
// If the level doesn't exist, a tip will be printed and
// the program will exit with status code 3.
func parseLevel(level string) Level {
	for k, v := range levels {
		if v == level {
			return k
		}
	}
	fmt.Fprintf(os.Stderr, "Error: Level \"%s\" doesn't exist! Be sure your level is one of them: debug, info, warn, error, off\n", level)
	os.Exit(3)
	return OffLevel
}

// String returns the name of Level ll.
// This method will be called when using printing operations like fmt.Println.
func (ll Level) String() string {
	return levels[ll]
}
