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
// Email: fishinlove@163.com
// Created at 2020/02/29 22:55:31

package logit

// globalLogger is a logger for global usage.
var globalLogger = newGlobalLogger()

// newGlobalLogger returns a logger for global usage.
// Default level is debug level.
func newGlobalLogger() *Logger {
	return NewLogger(DebugLevel, NewConsoleHandler(TextEncoder(), DefaultTimeFormat))
}

// Me returns globalLogger for more usages.
func Me() *Logger {
	return globalLogger
}

// callDepth is the depth of the method calling stack, which is about file name and line.
const callDepthOfGlobalLogger = 3

// Debug will output msg as a debug message.
func Debug(msg string) {
	globalLogger.log(callDepthOfGlobalLogger, DebugLevel, msg)
}

// Info will output msg as an info message.
func Info(msg string) {
	globalLogger.log(callDepthOfGlobalLogger, InfoLevel, msg)
}

// Warn will output msg as a warn message.
func Warn(msg string) {
	globalLogger.log(callDepthOfGlobalLogger, WarnLevel, msg)
}

// Error will output msg as an error message.
func Error(msg string) {
	globalLogger.log(callDepthOfGlobalLogger, ErrorLevel, msg)
}

// DebugFunc will output msg as a debug message.
// The msg is the return value of msgGenerator.
// This is the better way to output a long log made from many variables.
func DebugFunc(msgGenerator func() string) {
	globalLogger.log(callDepthOfGlobalLogger, DebugLevel, msgGenerator())
}

// InfoFunc will output msg as an info message.
// The msg is the return value of msgGenerator.
// This is the better way to output a long log made from many variables.
func InfoFunc(msgGenerator func() string) {
	globalLogger.log(callDepthOfGlobalLogger, InfoLevel, msgGenerator())
}

// WarnFunc will output msg as a warn message.
// The msg is the return value of msgGenerator.
// This is the better way to output a long log made from many variables.
func WarnFunc(msgGenerator func() string) {
	globalLogger.log(callDepthOfGlobalLogger, WarnLevel, msgGenerator())
}

// ErrorFunc will output msg as an error message.
// The msg is the return value of msgGenerator.
// This is the better way to output a long log made from many variables.
func ErrorFunc(msgGenerator func() string) {
	globalLogger.log(callDepthOfGlobalLogger, ErrorLevel, msgGenerator())
}
