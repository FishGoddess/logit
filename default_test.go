// Copyright 2024 FishGoddess. All Rights Reserved.
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
	"io"
	"log/slog"
	"testing"

	"github.com/FishGoddess/logit/handler"
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
	handlerName := t.Name()

	newHandler := func(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
		return slog.NewTextHandler(w, opts)
	}

	handler.Register(handlerName, newHandler)

	buffer := bytes.NewBuffer(make([]byte, 0, 1024))
	logger := NewLogger(
		WithDebugLevel(), WithHandler(handlerName), WithWriter(buffer), WithSource(), WithPID(),
	)

	SetDefault(logger)
	Debug("debug msg", "key1", 1)
	Info("info msg", "key2", 2)
	Warn("warn msg", "key3", 3)
	Error("error msg", "key4", 4)

	opts := &slog.HandlerOptions{AddSource: true, Level: slog.LevelDebug}
	wantBuffer := bytes.NewBuffer(make([]byte, 0, 1024))
	slogLogger := slog.New(newHandler(wantBuffer, opts)).With(keyPID, pid)

	slogLogger.Debug("debug msg", "key1", 1)
	slogLogger.Info("info msg", "key2", 2)
	slogLogger.Warn("warn msg", "key3", 3)
	slogLogger.Error("error msg", "key4", 4)

	got := removeTimeAndSource(buffer.String())
	want := removeTimeAndSource(wantBuffer.String())

	if got != want {
		t.Fatalf("got %s != want %s", got, want)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestDefaultLoggerSync$
func TestDefaultLoggerSync(t *testing.T) {
	syncer := &testSyncer{
		synced: false,
	}

	logger := &Logger{
		syncer: syncer,
	}

	SetDefault(logger)
	Sync()

	if !syncer.synced {
		t.Fatal("syncer.synced is wrong")
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestDefaultLoggerClose$
func TestDefaultLoggerClose(t *testing.T) {
	syncer := &testSyncer{
		synced: false,
	}

	closer := &testCloser{
		closed: false,
	}

	logger := &Logger{
		syncer: syncer,
		closer: closer,
	}

	SetDefault(logger)
	Close()

	if !syncer.synced {
		t.Fatal("syncer.synced is wrong")
	}

	if !closer.closed {
		t.Fatal("closer.closed is wrong")
	}
}
