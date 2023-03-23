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
	"fmt"
)

var (
	// globalLogger is a logger for global usages.
	globalLogger = NewLogger()
)

// Global returns the global logger.
func Global() *Logger {
	return globalLogger
}

// SetGlobal sets logger to global usages.
func SetGlobal(logger *Logger) *Logger {
	if logger != nil {
		globalLogger = logger
	}

	return globalLogger
}

// Debug returns a Log with debug level if debug level is enabled.
func Debug(msg string, params ...interface{}) *Log {
	return globalLogger.log(debugLevel, msg, params...)
}

// Info returns a Log with info level if info level is enabled.
func Info(msg string, params ...interface{}) *Log {
	return globalLogger.log(infoLevel, msg, params...)
}

// Warn returns a Log with warn level if warn level is enabled.
func Warn(msg string, params ...interface{}) *Log {
	return globalLogger.log(warnLevel, msg, params...)
}

// Error returns a Log with error level if error level is enabled.
func Error(err error, msg string, params ...interface{}) *Log {
	return globalLogger.log(errorLevel, msg, params...).WithError(err)
}

// Printf prints a log if print level is enabled.
func Printf(format string, params ...interface{}) {
	globalLogger.log(printLevel, format, params...).Log()
}

// Print prints a log if print level is enabled.
func Print(params ...interface{}) {
	globalLogger.log(printLevel, fmt.Sprint(params...)).Log()
}

// Println prints a log if print level is enabled.
func Println(params ...interface{}) {
	globalLogger.log(printLevel, fmt.Sprintln(params...)).Log()
}

// Sync syncs data storing in global logger's writer.
func Sync() error {
	return globalLogger.Sync()
}

// Close closes global logger and releases resources.
func Close() error {
	return globalLogger.Close()
}
