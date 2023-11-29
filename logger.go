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
	"io"
	"log/slog"
	"os"
	"runtime"
	"time"

	"github.com/FishGoddess/logit/defaults"
	"github.com/FishGoddess/logit/writer"
)

const (
	keyBad = "!BADKEY"
	keyPID = "pid"
)

var (
	pid = os.Getpid()
)

type Logger struct {
	handler slog.Handler

	withSource bool
	withPID    bool

	syncDuration time.Duration
}

// NewLogger creates a logger with given options or panics if failed.
// If you don't want to panic on failing, use NewLoggerGracefully instead.
func NewLogger(opts ...Option) *Logger {
	logger, err := NewLoggerGracefully(opts...)
	if err != nil {
		panic(err)
	}

	return logger
}

// NewLoggerGracefully creates a logger with given options or returns an error if failed.
// It's a more graceful way to create a logger than NewLogger function.
func NewLoggerGracefully(opts ...Option) (*Logger, error) {
	conf := newDefaultConfig()

	for _, opt := range opts {
		opt.applyTo(conf)
	}

	handler, err := conf.handler()
	if err != nil {
		return nil, err
	}

	logger := &Logger{
		handler:      handler,
		withSource:   conf.withSource,
		withPID:      conf.withPID,
		syncDuration: conf.syncDuration,
	}

	logger.runSyncTask()
	return logger, nil
}

func (l *Logger) runSyncTask() {
	if l.syncDuration <= 0 {
		return
	}

	go func() {
		for {
			time.Sleep(l.syncDuration)

			if err := l.Sync(); err != nil {
				defaults.HandleError("Logger.Sync", err)
			}
		}
	}()
}

func (l *Logger) clone() *Logger {
	newLogger := *l
	return &newLogger
}

func (l *Logger) squeezeAttr(args []any) (slog.Attr, []any) {
	// len of args must be > 0
	switch arg := args[0].(type) {
	case slog.Attr:
		return arg, args[1:]
	case string:
		if len(args) <= 1 {
			return slog.String(keyBad, arg), nil
		}

		return slog.Any(arg, args[1]), args[2:]
	default:
		return slog.Any(keyBad, arg), args[1:]
	}
}

func (l *Logger) newAttrs(args []any) (attrs []slog.Attr) {
	var attr slog.Attr
	for len(args) > 0 {
		attr, args = l.squeezeAttr(args)
		attrs = append(attrs, attr)
	}

	return attrs
}

// With returns a new logger with args.
// All logs from the new logger will carry the given args.
// See slog.Handler.WithAttrs.
func (l *Logger) With(args ...any) *Logger {
	if len(args) <= 0 {
		return l
	}

	attrs := l.newAttrs(args)
	if len(attrs) <= 0 {
		return l
	}

	newLogger := l.clone()
	newLogger.handler = l.handler.WithAttrs(attrs)

	return newLogger
}

// WithGroup returns a new logger with group name.
// All logs from the new logger will be grouped by the name.
// See slog.Handler.WithGroup.
func (l *Logger) WithGroup(name string) *Logger {
	if name == "" {
		return l
	}

	newLogger := l.clone()
	newLogger.handler = l.handler.WithGroup(name)

	return newLogger

}

// Enabled reports whether the logger should ignore logs whose level is lower.
func (l *Logger) Enabled(ctx context.Context, level Level) bool {
	if ctx == nil {
		ctx = context.Background()
	}

	return l.handler.Enabled(ctx, level.Peel())
}

func (l *Logger) newRecord(level Level, msg string, args []any) slog.Record {
	now := defaults.CurrentTime()

	var pc uintptr
	if l.withSource {
		var pcs [1]uintptr
		runtime.Callers(defaults.CallerDepth, pcs[:])
		pc = pcs[0]
	}

	record := slog.NewRecord(now, level.Peel(), msg, pc)
	if l.withPID {
		record.AddAttrs(slog.Int(keyPID, pid))
	}

	attrs := l.newAttrs(args)
	record.AddAttrs(attrs...)

	return record
}

func (l *Logger) log(ctx context.Context, level Level, msg string, args ...any) {
	if !l.Enabled(ctx, level) {
		return
	}

	if ctx == nil {
		ctx = context.Background()
	}

	// TODO 尝试用对象池优化
	record := l.newRecord(level, msg, args)

	if err := l.handler.Handle(ctx, record); err != nil {
		defaults.HandleError("Logger.handler.Handle", err)
	}
}

// Debug logs a log with msg and args in debug level.
func (l *Logger) Debug(msg string, args ...any) {
	l.log(context.Background(), LevelDebug, msg, args...)
}

// Info logs a log with msg and args in info level.
func (l *Logger) Info(msg string, args ...any) {
	l.log(context.Background(), LevelInfo, msg, args...)
}

// Warn logs a log with msg and args in warn level.
func (l *Logger) Warn(msg string, args ...any) {
	l.log(context.Background(), LevelWarn, msg, args...)
}

// Error logs a log with msg and args in error level.
func (l *Logger) Error(msg string, args ...any) {
	l.log(context.Background(), LevelError, msg, args...)
}

// DebugContext logs a log with ctx, msg and args in debug level.
func (l *Logger) DebugContext(ctx context.Context, msg string, args ...any) {
	l.log(ctx, LevelDebug, msg, args...)
}

// InfoContext logs a log with ctx, msg and args in info level.
func (l *Logger) InfoContext(ctx context.Context, msg string, args ...any) {
	l.log(ctx, LevelInfo, msg, args...)
}

// WarnContext logs a log with ctx, msg and args in warn level.
func (l *Logger) WarnContext(ctx context.Context, msg string, args ...any) {
	l.log(ctx, LevelWarn, msg, args...)
}

// ErrorContext logs a log with ctx, msg and args in error level.
func (l *Logger) ErrorContext(ctx context.Context, msg string, args ...any) {
	l.log(ctx, LevelError, msg, args...)
}

// Printf logs a log with format and args in print level.
// It a old-school way to log.
func (l *Logger) Printf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.log(context.Background(), LevelPrint, msg)
}

// Print logs a log with args in print level.
// It a old-school way to log.
func (l *Logger) Print(args ...interface{}) {
	msg := fmt.Sprint(args...)
	l.log(context.Background(), LevelPrint, msg)
}

// Println logs a log with args in print level.
// It a old-school way to log.
func (l *Logger) Println(args ...interface{}) {
	msg := fmt.Sprintln(args...)
	l.log(context.Background(), LevelPrint, msg)
}

// Sync syncs the logger and returns an error if failed.
func (l *Logger) Sync() error {
	if syncer, ok := l.handler.(writer.Syncer); ok {
		return syncer.Sync()
	}

	return nil
}

// Close closes the logger and returns an error if failed.
func (l *Logger) Close() error {
	if err := l.Sync(); err != nil {
		return err
	}

	if closer, ok := l.handler.(io.Closer); ok {
		return closer.Close()
	}

	return nil
}
