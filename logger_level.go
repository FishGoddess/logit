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

// LoggerLevel is the type representation of the level.
type LoggerLevel uint8

// Constants about logger level.
const (
    DebugLevel LoggerLevel = iota
    InfoLevel
    WarnLevel
    ErrorLevel
)

// prefixOfLevels provides a prefix of one level.
var prefixOfLevels = map[LoggerLevel]string{
    DebugLevel: "Debug",
    InfoLevel:  "Info",
    WarnLevel:  "Warn",
    ErrorLevel: "Error",
}

// The String method is used to print values passed as an operand
// to any format that accepts a string or to an printer without format such as Print.
func (ll LoggerLevel) String() string {
    return prefixOfLevels[ll]
}

// prefixOf gets the prefix of this level.
func PrefixOf(level LoggerLevel) string {
    return level.String()
}
