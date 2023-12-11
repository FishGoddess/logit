// Copyright 2023 FishGoddess. All Rights Reserved.
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
	"context"
	"fmt"
	"log/slog"
	"sync/atomic"

	"github.com/FishGoddess/logit/defaults"
)

var defaultLogger atomic.Value

func init() {
	SetDefault(NewLogger())
}

// SetDefault sets logger as the default logger.
func SetDefault(logger *Logger) {
	defaultLogger.Store(logger)
}

// Default returns the default logger.
func Default() *Logger {
	return defaultLogger.Load().(*Logger)
}

// Debug logs a log with msg and args in debug level.
func Debug(msg string, args ...any) {
	Default().log(context.Background(), slog.LevelDebug, msg, args...)
}

// Info logs a log with msg and args in info level.
func Info(msg string, args ...any) {
	Default().log(context.Background(), slog.LevelInfo, msg, args...)
}

// Warn logs a log with msg and args in warn level.
func Warn(msg string, args ...any) {
	Default().log(context.Background(), slog.LevelWarn, msg, args...)
}

// Error logs a log with msg and args in error level.
func Error(msg string, args ...any) {
	Default().log(context.Background(), slog.LevelError, msg, args...)
}

// DebugContext logs a log with ctx, msg and args in debug level.
func DebugContext(ctx context.Context, msg string, args ...any) {
	Default().log(ctx, slog.LevelDebug, msg, args...)
}

// InfoContext logs a log with ctx, msg and args in info level.
func InfoContext(ctx context.Context, msg string, args ...any) {
	Default().log(ctx, slog.LevelInfo, msg, args...)
}

// WarnContext logs a log with ctx, msg and args in warn level.
func WarnContext(ctx context.Context, msg string, args ...any) {
	Default().log(ctx, slog.LevelWarn, msg, args...)
}

// ErrorContext logs a log with ctx, msg and args in error level.
func ErrorContext(ctx context.Context, msg string, args ...any) {
	Default().log(ctx, slog.LevelError, msg, args...)
}

// Printf logs a log with format and args in print level.
// It a old-school way to log.
func Printf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	Default().log(context.Background(), defaults.LevelPrint, msg)
}

// Print logs a log with args in print level.
// It a old-school way to log.
func Print(args ...interface{}) {
	msg := fmt.Sprint(args...)
	Default().log(context.Background(), defaults.LevelPrint, msg)
}

// Println logs a log with args in print level.
// It a old-school way to log.
func Println(args ...interface{}) {
	msg := fmt.Sprintln(args...)
	Default().log(context.Background(), defaults.LevelPrint, msg)
}
