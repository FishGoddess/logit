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

import (
	"os"
)

var (
	// globalLogger is a logger for global usage.
	globalLogger *Logger
)

func init() {
	// 这边留一个注意点，由于这个 newGlobalLogger 需要用到注册的 handlers
	// 而这些 handlers 是使用 init 函数进行注册的，所以 newGlobalLogger 必须
	// 在注册 handlers 的 init 函数执行完之后才能执行，也就是说存在隐形的依赖关系
	// 这是由 Go 语言 init 函数的执行顺序决定的，至少在 Go v1.14 版本中这个执行顺序
	// 是我们想要的执行顺序，但是之后的版本就不好说了，所以这边记录着这个点，万一以后真的
	// 出现注册过的 handler 却提示找不到，就很可能是这个点引起的
	globalLogger = newGlobalLogger()
}

// newGlobalLogger returns a logger for global usage.
// Notice that it will try to load "./logit.conf" first, but a logger to console
// will be created if failed.
func newGlobalLogger() *Logger {
	file, err := os.Open("./logit.conf")
	if err != nil {
		return NewLogger(DebugLevel, NewConsoleHandler(TextEncoder(), DefaultTimeFormat))
	}
	defer file.Close()
	return NewLoggerFrom(file)
}

// Me returns globalLogger for more usages.
func Me() *Logger {
	return globalLogger
}

const (
	// callDepth is the depth of the method calling stack, which is about file name and line.
	callDepthOfGlobalLogger = 3
)

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

// Debugf will output msg as a debug message.
// The msg is the return value of generateMessage.
// This is a way to output a long log made from many variables.
// The msgFormat is the same as format in fmt.Printf, so you can use
// all format it supports, such as '%d'.
// msgParams is the params msgFormat needs, and it is variable-length, so
// you can add all your params here.
// You should know that this way to output msg is the most expensive way in time,
// but it's still faster than other logging libs. If you care about performance,
// than you should think about it, and if you don't, just use it without thinking.
func Debugf(msgFormat string, msgParams ...interface{}) {
	globalLogger.log(callDepth, DebugLevel, generateMessage(msgFormat, msgParams...))
}

// Infof will output msg as an info message.
// The msg is the return value of generateMessage.
// This is a way to output a long log made from many variables.
// The msgFormat is the same as format in fmt.Printf, so you can use
// all format it supports, such as '%d'.
// msgParams is the params msgFormat needs, and it is variable-length, so
// you can add all your params here.
// You should know that this way to output msg is the most expensive way in time,
// but it's still faster than other logging libs. If you care about performance,
// than you should think about it, and if you don't, just use it without thinking.
func Infof(msgFormat string, msgParams ...interface{}) {
	globalLogger.log(callDepth, InfoLevel, generateMessage(msgFormat, msgParams...))
}

// Warnf will output msg as a warn message.
// The msg is the return value of generateMessage.
// This is a way to output a long log made from many variables.
// The msgFormat is the same as format in fmt.Printf, so you can use
// all format it supports, such as '%d'.
// msgParams is the params msgFormat needs, and it is variable-length, so
// you can add all your params here.
// You should know that this way to output msg is the most expensive way in time,
// but it's still faster than other logging libs. If you care about performance,
// than you should think about it, and if you don't, just use it without thinking.
func Warnf(msgFormat string, msgParams ...interface{}) {
	globalLogger.log(callDepth, WarnLevel, generateMessage(msgFormat, msgParams...))
}

// Errorf will output msg as an error message.
// The msg is the return value of generateMessage.
// This is the better way to output a long log made from many variables.
// The msgFormat is the same as format in fmt.Printf, so you can use
// all format it supports, such as '%d'.
// msgParams is the params msgFormat needs, and it is variable-length, so
// you can add all your params here.
// You should know that this way to output msg is the most expensive way in time,
// but it's still faster than other logging libs. If you care about performance,
// than you should think about it, and if you don't, just use it without thinking.
func Errorf(msgFormat string, msgParams ...interface{}) {
	globalLogger.log(callDepth, ErrorLevel, generateMessage(msgFormat, msgParams...))
}
