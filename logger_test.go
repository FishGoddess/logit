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
	"bytes"
	"context"
	"io"
	"log/slog"
	"strings"
	"testing"
)

type testLoggerHandler struct {
	slog.TextHandler
	w    io.Writer
	opts slog.HandlerOptions
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestNewLogger$
func TestNewLogger(t *testing.T) {
	handler := &testLoggerHandler{}

	RegisterHandler(t.Name(), func(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
		return handler
	})

	logger := NewLogger(WithHandler(t.Name()))
	if logger.handler != handler {
		t.Fatalf("logger.handler %+v != handler %+v", logger.handler, handler)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestLoggerClone$
func TestLoggerClone(t *testing.T) {
	logger := NewLogger()
	newLogger := logger.clone()

	if logger == newLogger {
		t.Fatalf("logger %+v == newLogger %+v", logger, newLogger)
	}

	if logger.handler != newLogger.handler {
		t.Fatalf("logger.handler %+v != newLogger.handler %+v", logger.handler, newLogger.handler)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestLoggerNewAttrs$
func TestLoggerNewAttrs(t *testing.T) {
	logger := NewLogger()

	args := []any{
		"key1", 123, "key2", "456", slog.Bool("key3", true), 666, "key4",
	}

	attrs := logger.newAttrs(args)
	if len(attrs) != 5 {
		t.Fatalf("len(attrs) %d != 5", len(attrs))
	}

	if attrs[0].String() != "key1=123" {
		t.Fatalf("attrs[0] %s is wrong", attrs[0])
	}

	if attrs[1].String() != "key2=456" {
		t.Fatalf("attrs[1] %s is wrong", attrs[1])
	}

	if attrs[2].String() != "key3=true" {
		t.Fatalf("attrs[2] %s is wrong", attrs[2])
	}

	if attrs[3].String() != keyBad+"=666" {
		t.Fatalf("attrs[3] %s is wrong", attrs[3])
	}

	if attrs[4].String() != keyBad+"=key4" {
		t.Fatalf("attrs[4] %s is wrong", attrs[4])
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestLoggerWith$
func TestLoggerWith(t *testing.T) {
	logger := NewLogger()
	newLogger := logger.With()

	if logger != newLogger {
		t.Fatalf("logger %+v != newLogger %+v", logger, newLogger)
	}

	if logger.handler != newLogger.handler {
		t.Fatalf("logger.handler %+v != newLogger.handler %+v", logger.handler, newLogger.handler)
	}

	newLogger = logger.With("key", 123)

	if logger == newLogger {
		t.Fatalf("logger %+v == newLogger %+v", logger, newLogger)
	}

	if logger.handler == newLogger.handler {
		t.Fatalf("logger.handler %+v == newLogger.handler %+v", logger.handler, newLogger.handler)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestLoggerWithGroup$
func TestLoggerWithGroup(t *testing.T) {
	logger := NewLogger()
	newLogger := logger.WithGroup("")

	if logger != newLogger {
		t.Fatalf("logger %+v != newLogger %+v", logger, newLogger)
	}

	if logger.handler != newLogger.handler {
		t.Fatalf("logger.handler %+v != newLogger.handler %+v", logger.handler, newLogger.handler)
	}

	newLogger = logger.WithGroup("xxx")

	if logger == newLogger {
		t.Fatalf("logger %+v == newLogger %+v", logger, newLogger)
	}

	if logger.handler == newLogger.handler {
		t.Fatalf("logger.handler %+v == newLogger.handler %+v", logger.handler, newLogger.handler)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestLoggerEnabled$
func TestLoggerEnabled(t *testing.T) {
	ctx := context.Background()
	logger := NewLogger(WithErrorLevel())

	if logger.enabled(ctx, slog.LevelDebug) {
		t.Fatal("logger enabled debug")
	}

	if logger.enabled(ctx, slog.LevelInfo) {
		t.Fatal("logger enabled info")
	}

	if logger.enabled(ctx, slog.LevelWarn) {
		t.Fatal("logger enabled warn")
	}

	if !logger.enabled(ctx, slog.LevelError) {
		t.Fatal("logger enabled error")
	}
}

func removeTimeAndSource(str string) string {
	str = strings.ReplaceAll(str, "\n", " ")
	strs := strings.Split(str, " ")

	var removed strings.Builder
	for _, s := range strs {
		if strings.HasPrefix(s, slog.TimeKey) {
			continue
		}

		if strings.HasPrefix(s, slog.SourceKey) {
			continue
		}

		removed.WriteString(s)
		removed.WriteString(" ")
	}

	return removed.String()
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestLogger$
func TestLogger(t *testing.T) {
	newHandler := func(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
		return slog.NewTextHandler(w, opts)
	}

	RegisterHandler(t.Name(), newHandler)

	ctx := context.Background()
	buffer := bytes.NewBuffer(make([]byte, 0, 1024))
	logger := NewLogger(
		WithDebugLevel(), WithHandler(t.Name()), WithWriter(buffer), WithSource(), WithPID(),
	)

	logger.Debug("debug msg", "key1", 1)
	logger.Info("info msg", "key2", 2)
	logger.Warn("warn msg", "key3", 3)
	logger.Error("error msg", "key4", 4)

	logger.DebugContext(ctx, "debug msg with context", "key5", 5)
	logger.InfoContext(ctx, "info msg with context", "key6", 6)
	logger.WarnContext(ctx, "warn msg with context", "key7", 7)
	logger.ErrorContext(ctx, "error msg with context", "key8", 8)

	opts := &slog.HandlerOptions{AddSource: true, Level: slog.LevelDebug}
	wantBuffer := bytes.NewBuffer(make([]byte, 0, 1024))
	slogLogger := slog.New(newHandler(wantBuffer, opts)).With(keyPID, pid)

	slogLogger.Debug("debug msg", "key1", 1)
	slogLogger.Info("info msg", "key2", 2)
	slogLogger.Warn("warn msg", "key3", 3)
	slogLogger.Error("error msg", "key4", 4)

	slogLogger.DebugContext(ctx, "debug msg with context", "key5", 5)
	slogLogger.InfoContext(ctx, "info msg with context", "key6", 6)
	slogLogger.WarnContext(ctx, "warn msg with context", "key7", 7)
	slogLogger.ErrorContext(ctx, "error msg with context", "key8", 8)

	got := removeTimeAndSource(buffer.String())
	want := removeTimeAndSource(wantBuffer.String())

	if got != want {
		t.Fatalf("got %s != want %s", got, want)
	}
}
