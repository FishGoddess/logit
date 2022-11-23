// Copyright 2022 FishGoddess. All Rights Reserved.
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

package logit

import (
	"os"

	"github.com/FishGoddess/logit/support/global"
)

var (
	// globalLogger is a logger for global usage.
	globalLogger = NewLogger(
		Options().WithInfoLevel(),
		Options().WithWarnWriter(os.Stderr),
		Options().WithErrorWriter(os.Stderr),
		Options().WithCallerDepth(global.CallerDepth+1),
	)
)

// SetGlobal sets global logger to value returned from newLogger.
// We don't recommend you to call this function unless you really need to call.
// Instead, we recommend you to call logger.SetToGlobal to set one logger to global if you need.
func SetGlobal(logger *Logger) {
	globalLogger = logger
}

// Debug returns a Log with debug level if debug level is enabled.
func Debug(msg string, params ...interface{}) *Log {
	return globalLogger.Debug(msg, params...)
}

// Info returns a Log with info level if info level is enabled.
func Info(msg string, params ...interface{}) *Log {
	return globalLogger.Info(msg, params...)
}

// Warn returns a Log with warn level if warn level is enabled.
func Warn(msg string, params ...interface{}) *Log {
	return globalLogger.Warn(msg, params...)
}

// Error returns a Log with error level if error level is enabled.
func Error(err error, msg string, params ...interface{}) *Log {
	return globalLogger.Error(err, msg, params...)
}

// Printf prints a log if print level is enabled.
func Printf(format string, params ...interface{}) {
	globalLogger.Printf(format, params...)
}

// Print prints a log if print level is enabled.
func Print(params ...interface{}) {
	globalLogger.Print(params...)
}

// Println prints a log if print level is enabled.
func Println(params ...interface{}) {
	globalLogger.Println(params...)
}

// Sync syncs data storing in global logger's writer.
func Sync() error {
	return globalLogger.Sync()
}

// Close closes global logger and releases resources.
func Close() error {
	return globalLogger.Close()
}
