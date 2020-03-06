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
// Created at 2020/03/06 13:36:28

package logit

import (
    "fmt"
    "time"
)

// TODO 注释
type LoggerHandler func(logger *Logger, level LoggerLevel, now time.Time, msg string) bool

// TODO 注释
func (lh LoggerHandler) handle(logger *Logger, level LoggerLevel, now time.Time, msg string) bool {
    return lh(logger, level, now, msg)
}

// TODO 注释
func DefaultLoggerHandler(logger *Logger, level LoggerLevel, now time.Time, msg string) bool {
    fmt.Fprintf(logger.writer, "[%s] [%s] %s\n", prefixOf(level), now.Format(logger.formatOfTime), msg)
    return true
}
