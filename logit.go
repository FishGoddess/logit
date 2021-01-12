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
// Created at 2020/02/29 22:55:31

package logit

const (
	// callDepth is the depth of calling stack, which is about file name and line number.
	callDepthOfGlobalLogger = 3
)

var (
	// globalLogger is a logger for global usage.
	globalLogger = NewLogger()
)

// Me returns globalLogger for more usages.
func Me() *Logger {
	return globalLogger
}

// Debug will output msg as a debug message.
func Debug(msg string, params ...interface{}) {
	globalLogger.log(callDepthOfGlobalLogger, DebugLevel, msg, params...)
}

// Info will output msg as an info message.
func Info(msg string, params ...interface{}) {
	globalLogger.log(callDepthOfGlobalLogger, InfoLevel, msg, params...)
}

// Warn will output msg as a warn message.
func Warn(msg string, params ...interface{}) {
	globalLogger.log(callDepthOfGlobalLogger, WarnLevel, msg, params...)
}

// Error will output msg as an error message.
func Error(msg string, params ...interface{}) {
	globalLogger.log(callDepthOfGlobalLogger, ErrorLevel, msg, params...)
}
