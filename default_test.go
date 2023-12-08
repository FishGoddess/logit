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
	"testing"
)

// go test -v -cover -count=1 -test.cpu=1 -run=^TestSetDefault$
func TestSetDefault(t *testing.T) {
	defaultLogger.Store(NewLogger())

	logger := NewLogger()
	SetDefault(logger)

	gotLogger, ok := defaultLogger.Load().(*Logger)
	if !ok {
		t.Fatalf("logger type %T is wrong", defaultLogger.Load())
	}

	if gotLogger != logger {
		t.Fatalf("gotLogger %+v != logger %+v", gotLogger, logger)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestDefault$
func TestDefault(t *testing.T) {
	logger := NewLogger()
	defaultLogger.Store(logger)

	gotLogger := Default()
	if gotLogger != logger {
		t.Fatalf("gotLogger %+v != logger %+v", gotLogger, logger)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestDefaultLogger$
func TestDefaultLogger(t *testing.T) {
	newHandler := func(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
		return slog.NewTextHandler(w, opts)
	}

	ctx := context.Background()
	buffer := bytes.NewBuffer(make([]byte, 0, 1024))
	logger := NewLogger(
		WithDebugLevel(), WithHandler(newHandler), WithWriter(buffer), WithSource(), WithPID(),
	)

	SetDefault(logger)

	Default().Debug("debug msg", "key1", 1)
	Default().Info("info msg", "key2", 2)
	Default().Warn("warn msg", "key3", 3)
	Default().Error("error msg", "key4", 4)

	Default().DebugContext(ctx, "debug msg with context", "key5", 5)
	Default().InfoContext(ctx, "info msg with context", "key6", 6)
	Default().WarnContext(ctx, "warn msg with context", "key7", 7)
	Default().ErrorContext(ctx, "error msg with context", "key8", 8)

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
