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
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/FishGoddess/logit/core/handler"
	"github.com/FishGoddess/logit/core/rotate"
	"github.com/FishGoddess/logit/core/writer"
	"github.com/FishGoddess/logit/defaults"
)

var (
	_ NewHandlerFunc = handler.NewTextHandler
	_ NewHandlerFunc = handler.NewJsonHandler
)

type NewHandlerFunc func(w io.Writer, opts *slog.HandlerOptions) slog.Handler

type ReplaceAttrFunc func(groups []string, attr slog.Attr) slog.Attr

type LoggerBuilder struct {
	level Level

	newWriter  func() (io.Writer, error)
	wrapWriter func(io.Writer) writer.Writer

	newHandler  NewHandlerFunc
	replaceAttr ReplaceAttrFunc

	withSource bool
	withPID    bool
}

func Builder() *LoggerBuilder {
	newWriter := func() (io.Writer, error) {
		return os.Stdout, nil
	}

	builder := &LoggerBuilder{
		level:       levelDebug,
		newWriter:   newWriter,
		wrapWriter:  writer.Wrap,
		newHandler:  handler.NewTextHandler,
		replaceAttr: nil,
		withSource:  false,
		withPID:     false,
	}

	return builder
}

func (lb *LoggerBuilder) SetLevel(level Level) *LoggerBuilder {
	lb.level = level
	return lb
}

func (lb *LoggerBuilder) WriteTo(w io.Writer) *LoggerBuilder {
	newWriter := func() (io.Writer, error) {
		return w, nil
	}

	lb.newWriter = newWriter
	return lb
}

func (lb *LoggerBuilder) WriteToStdout() *LoggerBuilder {
	newWriter := func() (io.Writer, error) {
		return os.Stdout, nil
	}

	lb.newWriter = newWriter
	return lb
}

func (lb *LoggerBuilder) WriteToStderr() *LoggerBuilder {
	newWriter := func() (io.Writer, error) {
		return os.Stderr, nil
	}

	lb.newWriter = newWriter
	return lb
}

func (lb *LoggerBuilder) WriteToFile(path string) *LoggerBuilder {
	newWriter := func() (io.Writer, error) {
		dir := filepath.Dir(path)
		if err := os.MkdirAll(dir, defaults.FileDirMode); err != nil {
			return nil, err
		}

		return defaults.OpenFile(path, defaults.FileMode)
	}

	lb.newWriter = newWriter
	return lb
}

func (lb *LoggerBuilder) WriteToRotateFile(path string, opts ...rotate.Option) *LoggerBuilder {
	newWriter := func() (io.Writer, error) {
		return rotate.New(path, opts...)
	}

	lb.newWriter = newWriter
	return lb
}

func (lb *LoggerBuilder) runWriterSyncTask(w writer.Writer, frequency time.Duration) {
	if frequency <= 0 {
		return
	}

	go func() {
		for {
			time.Sleep(frequency)

			if err := w.Sync(); err != nil {
				defaults.HandleError("writer.sync", err)
			}
		}
	}()
}

func (lb *LoggerBuilder) UseBuffer(bufferSize uint64, syncFrequency time.Duration) *LoggerBuilder {
	wrapWriter := func(w io.Writer) writer.Writer {
		ww := writer.Buffer(w, bufferSize)

		lb.runWriterSyncTask(ww, syncFrequency)
		return ww
	}

	lb.wrapWriter = wrapWriter
	return lb
}

func (lb *LoggerBuilder) UseBatch(batchSize uint64, syncFrequency time.Duration) *LoggerBuilder {
	wrapWriter := func(w io.Writer) writer.Writer {
		ww := writer.Batch(w, batchSize)

		lb.runWriterSyncTask(ww, syncFrequency)
		return ww
	}

	lb.wrapWriter = wrapWriter
	return lb
}

func (lb *LoggerBuilder) UseHandler(newHandler NewHandlerFunc) *LoggerBuilder {
	lb.newHandler = newHandler
	return lb
}

func (lb *LoggerBuilder) UseTextHandler() *LoggerBuilder {
	lb.newHandler = handler.NewTextHandler
	return lb
}

func (lb *LoggerBuilder) UseJsonHandler() *LoggerBuilder {
	lb.newHandler = handler.NewJsonHandler
	return lb
}

func (lb *LoggerBuilder) WithSource() *LoggerBuilder {
	lb.withSource = true
	return lb
}

func (lb *LoggerBuilder) WithPID() *LoggerBuilder {
	lb.withPID = true
	return lb
}

func (lb *LoggerBuilder) newHandlerOptions() *slog.HandlerOptions {
	opts := &slog.HandlerOptions{
		Level:       lb.level,
		AddSource:   lb.withSource,
		ReplaceAttr: lb.replaceAttr,
	}

	return opts
}

func (lb *LoggerBuilder) Build() (*Logger, error) {
	w, err := lb.newWriter()
	if err != nil {
		return nil, err
	}

	ww := lb.wrapWriter(w)
	opts := lb.newHandlerOptions()

	logger := &Logger{
		handler:    lb.newHandler(ww, opts),
		withSource: lb.withSource,
		withPID:    lb.withPID,
	}

	return logger, nil
}

func (lb *LoggerBuilder) MustBuild() *Logger {
	logger, err := lb.Build()
	if err != nil {
		panic(err)
	}

	return logger
}
