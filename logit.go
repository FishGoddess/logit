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
// Created at 2020/02/29 22:55:31

package logit

// Logger holder for global usage.
var defaultLogger = NewStdoutLogger(InfoLevel)

// ChangeLevelTo will change the level of logit to newLevel.
func ChangeLevelTo(level LogLevel) {
    defaultLogger.ChangeLevelTo(level)
}

// Enable sets logit on running status.
func Enable() {
    defaultLogger.Enable()
}

// Disable sets logit on shutdown status.
func Disable() {
    defaultLogger.Disable()
}

// callDepth is the depth of the method calling stack, which is about file name and line.
const callDepthOfDefaultLogger = 3

// Debug will output msg as a debug message.
func Debug(msg string, args ...interface{}) {
    defaultLogger.log(callDepthOfDefaultLogger, DebugLevel, formatMessage(msg, args...))
}

// Info will output msg as an info message.
func Info(msg string, args ...interface{}) {
    defaultLogger.log(callDepthOfDefaultLogger, InfoLevel, formatMessage(msg, args...))
}

// Warning will output msg as a warning message.
func Warning(msg string, args ...interface{}) {
    defaultLogger.log(callDepthOfDefaultLogger, WarningLevel, formatMessage(msg, args...))
}

// Error will output msg as an error message.
func Error(msg string, args ...interface{}) {
    defaultLogger.log(callDepthOfDefaultLogger, ErrorLevel, formatMessage(msg, args...))
}
