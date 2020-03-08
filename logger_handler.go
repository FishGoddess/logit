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
    "time"
)

// LoggerHandler is a struct representation of log handler.
// Every log will handle by this handler, and you can customize your own handler
// to handle logs in your way. The return value is meaningful, false means
// next handler will not be handled, only true will go on handling process.
// Notice that if one handler returns false, then all handlers after it
// will not use anymore.
type LoggerHandler func(logger *Logger, level LoggerLevel, now time.Time, msg string) bool

// handle will call lh as a function. It is just a proxy method.
func (lh LoggerHandler) handle(logger *Logger, level LoggerLevel, now time.Time, msg string) bool {
    return lh(logger, level, now, msg)
}

// DefaultLoggerHandler is the default handler in logit.
// The log handled by this handler will be like "[Info] [2020-03-06 16:10:44] msg".
// If you want to customize, just code your own handler, then replace it!
func DefaultLoggerHandler(logger *Logger, level LoggerLevel, now time.Time, msg string) bool {
    logger.Writer().Write([]byte("[" + PrefixOf(level) + "] [" + now.Format(logger.formatOfTime) + "] " + msg + "\n"))
    return true
}
