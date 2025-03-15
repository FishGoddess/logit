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
	"io"
	"log/slog"
	"os"
	"time"

	"github.com/FishGoddess/logit/handler"
)

type nilSyncer struct{}

func (nilSyncer) Sync() error {
	return nil
}

type nilCloser struct{}

func (nilCloser) Close() error {
	return nil
}

type config struct {
	level   slog.Level
	handler string

	newWriter  func() (io.Writer, error)
	wrapWriter func(io.Writer) io.Writer

	replaceAttr func(groups []string, attr slog.Attr) slog.Attr

	withSource bool
	withPID    bool

	syncTimer time.Duration
}

func newDefaultConfig() *config {
	newWriter := func() (io.Writer, error) {
		return os.Stdout, nil
	}

	conf := &config{
		level:       slog.LevelDebug,
		handler:     handler.Tape,
		newWriter:   newWriter,
		wrapWriter:  nil,
		replaceAttr: nil,
		withSource:  false,
		withPID:     false,
		syncTimer:   0,
	}

	return conf
}

func (c *config) newSyncer(handler slog.Handler, writer io.Writer) Syncer {
	if syncer, ok := handler.(Syncer); ok {
		return syncer
	}

	if syncer, ok := writer.(Syncer); ok {
		return syncer
	}

	return nilSyncer{}
}

func (c *config) newCloser(handler slog.Handler, writer io.Writer) io.Closer {
	if closer, ok := handler.(io.Closer); ok {
		return closer
	}

	if closer, ok := writer.(io.Closer); ok {
		return closer
	}

	return nilCloser{}
}

func (c *config) newHandlerOptions() *slog.HandlerOptions {
	opts := &slog.HandlerOptions{
		Level:       c.level,
		AddSource:   c.withSource,
		ReplaceAttr: c.replaceAttr,
	}

	return opts
}

func (c *config) newHandler() (slog.Handler, Syncer, io.Closer, error) {
	newHandler, err := handler.Get(c.handler)
	if err != nil {
		return nil, nil, nil, err
	}

	writer, err := c.newWriter()
	if err != nil {
		return nil, nil, nil, err
	}

	if c.wrapWriter != nil {
		writer = c.wrapWriter(writer)
	}

	opts := c.newHandlerOptions()
	handler := newHandler(writer, opts)
	syncer := c.newSyncer(handler, writer)
	closer := c.newCloser(handler, writer)

	return handler, syncer, closer, nil
}
