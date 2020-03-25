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

// defaultLogger is a Logger holder for global usage.
// Default level is info level.
var defaultLogger = NewDevelopLogger()

// ChangeLevelTo will change the level of logit to newLevel.
func ChangeLevelTo(level Level) {
    defaultLogger.ChangeLevelTo(level)
}

// EnableFileInfo means every log will contain file info like line number.
// However, you should know that this is expensive in time.
// So be sure you really need it or keep it disabled.
func EnableFileInfo() {
    defaultLogger.EnableFileInfo()
}

// DisableFileInfo means every log will not contain file info like line number.
// If you want file info again, try logit.EnableFileInfo().
func DisableFileInfo() {
    defaultLogger.DisableFileInfo()
}

// AddHandlers adds more handlers to logit, and all handlers added before will be retained.
// If you want to remove all handlers, try logit.SetHandlers().
// See logit.DefaultLoggerHandler.
func AddHandlers(handlers ...Handler) {
    defaultLogger.AddHandlers(handlers...)
}

// SetHandlers replaces logit's handlers with handlers, all handlers added before will be removed.
// If you want to add more handlers rather than replace them, try logit.AddHandlers().
// Notice that at least one handler should be added, so if len(handlers) < 1, it returns false
// which means setting failed. Return true if setting is successful.
// See logit.DefaultLoggerHandler.
func SetHandlers(handlers ...Handler) bool {
    return defaultLogger.SetHandlers(handlers...)
}

// callDepth is the depth of the method calling stack, which is about file name and line.
const callDepthOfDefaultLogger = 3

// Debug will output msg as a debug message.
func Debug(msg string) {
    defaultLogger.log(callDepthOfDefaultLogger, DebugLevel, msg)
}

// Info will output msg as an info message.
func Info(msg string) {
    defaultLogger.log(callDepthOfDefaultLogger, InfoLevel, msg)
}

// Warn will output msg as a warn message.
func Warn(msg string) {
    defaultLogger.log(callDepthOfDefaultLogger, WarnLevel, msg)
}

// Error will output msg as an error message.
func Error(msg string) {
    defaultLogger.log(callDepthOfDefaultLogger, ErrorLevel, msg)
}

// DebugFunc will output msg as a debug message.
// The msg is the return value of msgGenerator.
// This is a better way to output a long log made of many variables.
func DebugFunc(msgGenerator func() string) {
    defaultLogger.DebugFunc(msgGenerator)
}

// InfoFunc will output msg as an info message.
// The msg is the return value of msgGenerator.
// This is a better way to output a long log made of many variables.
func InfoFunc(msgGenerator func() string) {
    defaultLogger.InfoFunc(msgGenerator)
}

// WarnFunc will output msg as a warn message.
// The msg is the return value of msgGenerator.
// This is a better way to output a long log made of many variables.
func WarnFunc(msgGenerator func() string) {
    defaultLogger.WarnFunc(msgGenerator)
}

// ErrorFunc will output msg as an error message.
// The msg is the return value of msgGenerator.
// This is a better way to output a long log made of many variables.
func ErrorFunc(msgGenerator func() string) {
    defaultLogger.ErrorFunc(msgGenerator)
}
