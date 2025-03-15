// Copyright 2025 FishGoddess. All Rights Reserved.
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
)

const (
	keyBad = "!BADKEY"
	keyPID = "pid"
)

var (
	pid = os.Getpid()
)

// Syncer is an interface that syncs data to somewhere.
type Syncer interface {
	Sync() error
}

// Logger is the entry of logging in logit.
// It has several levels including debug, info, warn and error.
// It's also a syncer or closer if handler is a syncer or closer.
type Logger struct {
	handler slog.Handler

	syncer Syncer
	closer io.Closer

	withSource bool
	withPID    bool
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

	handler, syncer, closer, err := conf.newHandler()
	if err != nil {
		return nil, err
	}

	logger := &Logger{
		handler:    handler,
		syncer:     syncer,
		closer:     closer,
		withSource: conf.withSource,
		withPID:    conf.withPID,
	}

	if conf.syncTimer > 0 {
		go logger.runSyncTimer(conf.syncTimer)
	}

	return logger, nil
}

func (l *Logger) runSyncTimer(d time.Duration) {
	timer := time.NewTimer(d)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			if err := l.Sync(); err != nil {
				defaults.HandleError("Logger.Sync", err)
			}
		}
	}
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

// enabled reports whether the logger should ignore logs whose level is lower.
func (l *Logger) enabled(level slog.Level) bool {
	return l.handler.Enabled(context.Background(), level)
}

// DebugEnabled reports whether the logger should ignore logs whose level is lower than debug.
func (l *Logger) DebugEnabled() bool {
	return l.enabled(slog.LevelDebug)
}

// InfoEnabled reports whether the logger should ignore logs whose level is lower than info.
func (l *Logger) InfoEnabled() bool {
	return l.enabled(slog.LevelInfo)
}

// WarnEnabled reports whether the logger should ignore logs whose level is lower than warn.
func (l *Logger) WarnEnabled() bool {
	return l.enabled(slog.LevelWarn)
}

// ErrorEnabled reports whether the logger should ignore logs whose level is lower than error.
func (l *Logger) ErrorEnabled() bool {
	return l.enabled(slog.LevelError)
}

// PrintEnabled reports whether the logger should ignore logs whose level is lower than print.
func (l *Logger) PrintEnabled() bool {
	return l.enabled(defaults.LevelPrint)
}

func (l *Logger) newRecord(level slog.Level, msg string, args []any) slog.Record {
	var pc uintptr
	if l.withSource {
		var pcs [1]uintptr
		runtime.Callers(defaults.CallerDepth, pcs[:])
		pc = pcs[0]
	}

	now := defaults.CurrentTime()
	record := slog.NewRecord(now, level, msg, pc)

	if l.withPID {
		record.AddAttrs(slog.Int(keyPID, pid))
	}

	var attr slog.Attr
	for len(args) > 0 {
		attr, args = l.squeezeAttr(args)
		record.AddAttrs(attr)
	}

	return record
}

func (l *Logger) log(level slog.Level, msg string, args ...any) {
	if !l.enabled(level) {
		return
	}

	record := l.newRecord(level, msg, args)

	if err := l.handler.Handle(context.Background(), record); err != nil {
		defaults.HandleError("Logger.handler.Handle", err)
	}
}

// Debug logs a log with msg and args in debug level.
func (l *Logger) Debug(msg string, args ...any) {
	l.log(slog.LevelDebug, msg, args...)
}

// Info logs a log with msg and args in info level.
func (l *Logger) Info(msg string, args ...any) {
	l.log(slog.LevelInfo, msg, args...)
}

// Warn logs a log with msg and args in warn level.
func (l *Logger) Warn(msg string, args ...any) {
	l.log(slog.LevelWarn, msg, args...)
}

// Error logs a log with msg and args in error level.
func (l *Logger) Error(msg string, args ...any) {
	l.log(slog.LevelError, msg, args...)
}

// Printf logs a log with format and args in print level.
// It a old-school way to log.
func (l *Logger) Printf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.log(defaults.LevelPrint, msg)
}

// Print logs a log with args in print level.
// It a old-school way to log.
func (l *Logger) Print(args ...interface{}) {
	msg := fmt.Sprint(args...)
	l.log(defaults.LevelPrint, msg)
}

// Println logs a log with args in print level.
// It a old-school way to log.
func (l *Logger) Println(args ...interface{}) {
	msg := fmt.Sprintln(args...)
	l.log(defaults.LevelPrint, msg)
}

// Sync syncs the logger and returns an error if failed.
func (l *Logger) Sync() error {
	return l.syncer.Sync()
}

// Close closes the logger and returns an error if failed.
func (l *Logger) Close() error {
	if err := l.Sync(); err != nil {
		return err
	}

	return l.closer.Close()
}
