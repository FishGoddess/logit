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
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"testing"

	"github.com/FishGoddess/logit/rotate"
	"github.com/FishGoddess/logit/writer"
)

// go test -v -cover -run=^TestWithDebugLevel$
func TestWithDebugLevel(t *testing.T) {
	conf := &config{level: slog.LevelError}
	WithDebugLevel().applyTo(conf)

	if conf.level != slog.LevelDebug {
		t.Fatalf("conf.level %+v != slog.LevelDebug", conf.level)
	}
}

// go test -v -cover -run=^TestWithInfoLevel$
func TestWithInfoLevel(t *testing.T) {
	conf := &config{level: slog.LevelError}
	WithInfoLevel().applyTo(conf)

	if conf.level != slog.LevelInfo {
		t.Fatalf("conf.level %+v != slog.LevelInfo", conf.level)
	}
}

// go test -v -cover -run=^TestWithWarnLevel$
func TestWithWarnLevel(t *testing.T) {
	conf := &config{level: slog.LevelError}
	WithWarnLevel().applyTo(conf)

	if conf.level != slog.LevelWarn {
		t.Fatalf("conf.level %+v != slog.LevelWarn", conf.level)
	}
}

// go test -v -cover -run=^TestWithErrorLevel$
func TestWithErrorLevel(t *testing.T) {
	conf := &config{level: slog.LevelDebug}
	WithErrorLevel().applyTo(conf)

	if conf.level != slog.LevelError {
		t.Fatalf("conf.level %+v != slog.LevelError", conf.level)
	}
}

// go test -v -cover -run=^TestWithWriter$
func TestWithWriter(t *testing.T) {
	conf := &config{newWriter: nil}
	WithWriter(os.Stdout).applyTo(conf)

	w, err := conf.newWriter()
	if err != nil {
		t.Fatal(err)
	}

	if w != os.Stdout {
		t.Fatalf("w %+v != os.Stdout", w)
	}
}

// go test -v -cover -run=^TestWithStdout$
func TestWithStdout(t *testing.T) {
	conf := &config{newWriter: nil}
	WithStdout().applyTo(conf)

	w, err := conf.newWriter()
	if err != nil {
		t.Fatal(err)
	}

	if w != os.Stdout {
		t.Fatalf("w %+v != os.Stdout", w)
	}
}

// go test -v -cover -run=^TestWithStderr$
func TestWithStderr(t *testing.T) {
	conf := &config{newWriter: nil}
	WithStderr().applyTo(conf)

	w, err := conf.newWriter()
	if err != nil {
		t.Fatal(err)
	}

	if w != os.Stderr {
		t.Fatalf("w %+v != os.Stderr", w)
	}
}

// go test -v -cover -run=^TestWithFile$
func TestWithFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), t.Name())

	conf := &config{newWriter: nil}
	WithFile(path).applyTo(conf)

	w, err := conf.newWriter()
	if err != nil {
		t.Fatal(err)
	}

	file, ok := w.(*os.File)
	if !ok {
		t.Fatalf("writer type %T is wrong", w)
	}

	text := t.Name()
	if _, err = w.Write([]byte(text)); err != nil {
		t.Fatal(err)
	}

	if err = file.Close(); err != nil {
		t.Fatal(err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	if string(data) != text {
		t.Fatalf("string(data) %s != text %s", string(data), text)
	}
}

// go test -v -cover -run=^TestWithRotateFile$
func TestWithRotateFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), t.Name())

	conf := &config{newWriter: nil}
	WithRotateFile(path).applyTo(conf)

	w, err := conf.newWriter()
	if err != nil {
		t.Fatal(err)
	}

	file, ok := w.(*rotate.File)
	if !ok {
		t.Fatalf("writer type %T is wrong", w)
	}

	text := t.Name()
	if _, err = w.Write([]byte(text)); err != nil {
		t.Fatal(err)
	}

	if err = file.Close(); err != nil {
		t.Fatal(err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	if string(data) != text {
		t.Fatalf("string(data) %s != text %s", string(data), text)
	}
}

// go test -v -cover -run=^TestWithBuffer$
func TestWithBuffer(t *testing.T) {
	conf := &config{wrapWriter: nil}
	WithBuffer(64).applyTo(conf)

	buffer := bytes.NewBuffer(make([]byte, 0, 128))
	w := conf.wrapWriter(buffer)

	ww, ok := w.(*writer.BufferWriter)
	if !ok {
		t.Fatalf("writer type %T is wrong", w)
	}

	text := string(make([]byte, 63))
	if _, err := ww.Write([]byte(text)); err != nil {
		t.Fatal(err)
	}

	data := buffer.Bytes()
	if buffer.Len() > 0 {
		t.Fatalf("buffer.Len() %d > 0", buffer.Len())
	}

	if _, err := ww.Write([]byte(text)); err != nil {
		t.Fatal(err)
	}

	data = buffer.Bytes()
	if string(data) != text {
		t.Fatalf("string(data) %s != text %s", string(data), text)
	}

	if err := ww.Sync(); err != nil {
		t.Fatal(err)
	}

	data = buffer.Bytes()
	text = text + text

	if string(data) != text {
		t.Fatalf("string(data) %s != text %s", string(data), text)
	}
}

// go test -v -cover -run=^TestWithBatch$
func TestWithBatch(t *testing.T) {
	conf := &config{wrapWriter: nil}
	WithBatch(16).applyTo(conf)

	buffer := bytes.NewBuffer(make([]byte, 0, 256))
	w := conf.wrapWriter(buffer)

	bw, ok := w.(*writer.BatchWriter)
	if !ok {
		t.Fatalf("writer type %T is wrong", w)
	}

	text := string(make([]byte, 4))
	for i := 0; i < 15; i++ {
		if _, err := bw.Write([]byte(text)); err != nil {
			t.Fatal(err)
		}
	}

	data := buffer.Bytes()
	if buffer.Len() > 0 {
		t.Fatalf("buffer.Len() %d > 0", buffer.Len())
	}

	for i := 0; i < 15; i++ {
		if _, err := bw.Write([]byte(text)); err != nil {
			t.Fatal(err)
		}
	}

	data = buffer.Bytes()
	want := ""
	for i := 0; i < 16; i++ {
		want = want + text
	}

	if string(data) != want {
		t.Fatalf("string(data) %s != want %s", string(data), want)
	}

	if err := bw.Sync(); err != nil {
		t.Fatal(err)
	}

	data = buffer.Bytes()
	for i := 0; i < 14; i++ {
		want = want + text
	}

	if string(data) != want {
		t.Fatalf("string(data) %s != want %s", string(data), want)
	}
}

// go test -v -cover -run=^TestWithHandler$
func TestWithHandler(t *testing.T) {
	newHandler := func(w io.Writer, opts *slog.HandlerOptions) slog.Handler { return nil }

	conf := &config{newHandler: nil}
	WithHandler(newHandler).applyTo(conf)

	if fmt.Sprintf("%p", conf.newHandler) != fmt.Sprintf("%p", newHandler) {
		t.Fatal("conf.newHandler is wrong")
	}
}

// go test -v -cover -run=^TestWithReplaceAttr$
func TestWithReplaceAttr(t *testing.T) {
	replaceAttr := func(groups []string, attr slog.Attr) slog.Attr { return slog.Attr{} }

	conf := &config{replaceAttr: nil}
	WithReplaceAttr(replaceAttr).applyTo(conf)

	if fmt.Sprintf("%p", conf.replaceAttr) != fmt.Sprintf("%p", replaceAttr) {
		t.Fatal("conf.replaceAttr is wrong")
	}
}

// go test -v -cover -run=^TestWithSource$
func TestWithSource(t *testing.T) {
	conf := &config{withSource: false}
	WithSource().applyTo(conf)

	if !conf.withSource {
		t.Fatal("conf.withSource is wrong")
	}
}

// go test -v -cover -run=^TestWithPID$
func TestWithPID(t *testing.T) {
	conf := &config{withPID: false}
	WithPID().applyTo(conf)

	if !conf.withPID {
		t.Fatal("conf.withPID is wrong")
	}
}
