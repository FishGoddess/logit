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

// DebugF will output msg as a debug message.
// The msg is the return value of generateMessage.
// This is a way to output a long log made from many variables.
// The msgFormat is the same as format in fmt.Printf, so you can use
// all format it supports, such as '%d'.
// msgParams is the params msgFormat needs, and it is variable-length, so
// you can add all your params here.
// You should know that this way to output msg is the most expensive way in time,
// but it's still faster than other logging libs. If you care about performance,
// than you should think about it, and if you don't, just use it without thinking.
func DebugF(msgFormat string, msgParams ...interface{}) {
	globalLogger.log(callDepth, DebugLevel, generateMessage(msgFormat, msgParams...))
}

// InfoF will output msg as an info message.
// The msg is the return value of generateMessage.
// This is a way to output a long log made from many variables.
// The msgFormat is the same as format in fmt.Printf, so you can use
// all format it supports, such as '%d'.
// msgParams is the params msgFormat needs, and it is variable-length, so
// you can add all your params here.
// You should know that this way to output msg is the most expensive way in time,
// but it's still faster than other logging libs. If you care about performance,
// than you should think about it, and if you don't, just use it without thinking.
func InfoF(msgFormat string, msgParams ...interface{}) {
	globalLogger.log(callDepth, InfoLevel, generateMessage(msgFormat, msgParams...))
}

// WarnF will output msg as a warn message.
// The msg is the return value of generateMessage.
// This is a way to output a long log made from many variables.
// The msgFormat is the same as format in fmt.Printf, so you can use
// all format it supports, such as '%d'.
// msgParams is the params msgFormat needs, and it is variable-length, so
// you can add all your params here.
// You should know that this way to output msg is the most expensive way in time,
// but it's still faster than other logging libs. If you care about performance,
// than you should think about it, and if you don't, just use it without thinking.
func WarnF(msgFormat string, msgParams ...interface{}) {
	globalLogger.log(callDepth, WarnLevel, generateMessage(msgFormat, msgParams...))
}

// ErrorF will output msg as an error message.
// The msg is the return value of generateMessage.
// This is the better way to output a long log made from many variables.
// The msgFormat is the same as format in fmt.Printf, so you can use
// all format it supports, such as '%d'.
// msgParams is the params msgFormat needs, and it is variable-length, so
// you can add all your params here.
// You should know that this way to output msg is the most expensive way in time,
// but it's still faster than other logging libs. If you care about performance,
// than you should think about it, and if you don't, just use it without thinking.
func ErrorF(msgFormat string, msgParams ...interface{}) {
	globalLogger.log(callDepth, ErrorLevel, generateMessage(msgFormat, msgParams...))
}
