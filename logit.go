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
var defaultLogger = NewStdoutLogger(InfoLevel)

// ChangeLevelTo will change the level of logit to newLevel.
func ChangeLevelTo(level LoggerLevel) {
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
func AddHandlers(handlers ...LoggerHandler) {
    defaultLogger.AddHandlers(handlers...)
}

// SetHandlers replaces logit's handlers with handlers, all handlers added before will be removed.
// If you want to add more handlers rather than replace them, try logit.AddHandlers().
// Notice that at least one handler should be added, so if len(handlers) < 1, it returns false
// which means setting failed. Return true if setting is successful.
// See logit.DefaultLoggerHandler.
func SetHandlers(handlers ...LoggerHandler) bool {
    return defaultLogger.SetHandlers(handlers...)
}

// SetFormatOfTime sets format of time as you want.
// Default is "2006-01-02 15:04:05", see logit.DefaultFormatOfTime.
func SetFormatOfTime(formatOfTime string) {
    defaultLogger.SetFormatOfTime(formatOfTime)
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
