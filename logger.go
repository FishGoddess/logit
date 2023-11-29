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

// New creates a logger with given options or panics if failed.
// If you don't want to panic on failing, use NewGracefully instead.
func New(opts ...Option) *Logger {
	logger, err := NewGracefully(opts...)
	if err != nil {
		panic(err)
	}

	return logger
}

// NewProduction creates a logger with production options we think and given options you like.
// We recommend you to use some options in production, so we provide this way to create a logger.
func NewProduction(opts ...Option) *Logger {
	usingOpts := []Option{
		WithInfoLevel(), WithSource(), WithPID(),
		WithBatch(16), WithRotateFile("./logit.log"),
	}

	usingOpts = append(usingOpts, opts...)
	return New(usingOpts...)
}

// NewGracefully creates a logger with given options or returns an error if failed.
// It's a more graceful way to create a logger than New function.
func NewGracefully(opts ...Option) (*Logger, error) {
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

func (l *Logger) WithGroup(name string) *Logger {
	if name == "" {
		return l
	}

	newLogger := l.clone()
	newLogger.handler = l.handler.WithGroup(name)
	return newLogger

}

func (l *Logger) Enabled(ctx context.Context, level slog.Level) bool {
	if ctx == nil {
		ctx = context.Background()
	}

	return l.handler.Enabled(ctx, level)
}

func (l *Logger) newRecord(level slog.Level, msg string, args []any) slog.Record {
	now := defaults.CurrentTime()

	var pc uintptr
	if l.withSource {
		var pcs [1]uintptr
		runtime.Callers(defaults.CallerDepth, pcs[:])
		pc = pcs[0]
	}

	record := slog.NewRecord(now, level, msg, pc)
	if l.withPID {
		record.AddAttrs(slog.Int(keyPID, pid))
	}

	attrs := l.newAttrs(args)
	record.AddAttrs(attrs...)

	return record
}

func (l *Logger) log(ctx context.Context, level Level, msg string, args ...any) {
	slogLevel := level.Peel()

	if !l.Enabled(ctx, slogLevel) {
		return
	}

	if ctx == nil {
		ctx = context.Background()
	}

	// TODO 尝试用对象池优化
	record := l.newRecord(slogLevel, msg, args)

	if err := l.handler.Handle(ctx, record); err != nil {
		defaults.HandleError("Logger.handler.Handle", err)
	}
}

func (l *Logger) Debug(msg string, args ...any) {
	l.log(context.Background(), levelDebug, msg, args...)
}

func (l *Logger) Info(msg string, args ...any) {
	l.log(context.Background(), levelInfo, msg, args...)
}

func (l *Logger) Warn(msg string, args ...any) {
	l.log(context.Background(), levelWarn, msg, args...)
}

func (l *Logger) Error(msg string, args ...any) {
	l.log(context.Background(), levelError, msg, args...)
}

func (l *Logger) DebugContext(ctx context.Context, msg string, args ...any) {
	l.log(ctx, levelDebug, msg, args...)
}

func (l *Logger) InfoContext(ctx context.Context, msg string, args ...any) {
	l.log(ctx, levelInfo, msg, args...)
}

func (l *Logger) WarnContext(ctx context.Context, msg string, args ...any) {
	l.log(ctx, levelWarn, msg, args...)
}

func (l *Logger) ErrorContext(ctx context.Context, msg string, args ...any) {
	l.log(ctx, levelError, msg, args...)
}

// Printf prints a log if print level is enabled.
func (l *Logger) Printf(format string, params ...interface{}) {
	msg := fmt.Sprintf(format, params...)
	l.log(context.Background(), levelPrint, msg)
}

// Print prints a log if print level is enabled.
func (l *Logger) Print(params ...interface{}) {
	msg := fmt.Sprint(params...)
	l.log(context.Background(), levelPrint, msg)
}

// Println prints a log if print level is enabled.
func (l *Logger) Println(params ...interface{}) {
	msg := fmt.Sprintln(params...)
	l.log(context.Background(), levelPrint, msg)
}

func (l *Logger) Sync() error {
	if syncer, ok := l.handler.(writer.Syncer); ok {
		return syncer.Sync()
	}

	return nil
}

func (l *Logger) Close() error {
	if err := l.Sync(); err != nil {
		return err
	}

	if closer, ok := l.handler.(io.Closer); ok {
		return closer.Close()
	}

	return nil
}
