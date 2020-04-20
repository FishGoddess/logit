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
// Author: fish
// Email: fishinlove@163.com
// Created at 2020/03/01 14:18:33

package logit

import (
	"errors"
	"math"
)

// Level is the type representation of the level.
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
	// levels provides the name of all supported level.
	levels = map[Level]string{
		DebugLevel: "debug",
		InfoLevel:  "info",
		WarnLevel:  "warn",
		ErrorLevel: "error",
		OffLevel:   "off",
	}

	// LevelIsNotExisted is an error representation of a level is not existed.
	LevelIsNotExisted = errors.New("level is not existed, be sure your level is one of them: debug, info, warn, error, off")
)

// ParseLevel parses level and returns the Level of it.
// Return LevelIsNotExisted if level is not existed.
func ParseLevel(level string) (Level, error) {
	for k, v := range levels {
		if v == level {
			return k, nil
		}
	}
	return 0, LevelIsNotExisted
}

// The String method is used to print values passed as an operand
// to any format that accepts a string or to an printer without format such as Print.
func (ll Level) String() string {
	return levels[ll]
}
