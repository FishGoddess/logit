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
	"time"

	"github.com/FishGoddess/logit/core/handler"
	"github.com/FishGoddess/logit/core/writer"
)

var (
	_ NewHandlerFunc = handler.NewTextHandler
	_ NewHandlerFunc = handler.NewJsonHandler
)

// NewHandlerFunc is a function creating a slog.Handler instance with w and opts.
type NewHandlerFunc func(w io.Writer, opts *slog.HandlerOptions) slog.Handler

// ReplaceAttrFunc is a function replacing attr of groups.
// See slog.HandlerOptions.ReplaceAttr.
type ReplaceAttrFunc func(groups []string, attr slog.Attr) slog.Attr

type config struct {
	level Level

	newWriter  func() (io.Writer, error)
	wrapWriter func(io.Writer) writer.Writer

	newHandler  NewHandlerFunc
	replaceAttr ReplaceAttrFunc

	withSource bool
	withPID    bool

	syncDuration time.Duration
}

func newDefaultConfig() *config {
	newWriter := func() (io.Writer, error) {
		return os.Stdout, nil
	}

	conf := &config{
		level:        levelDebug,
		newWriter:    newWriter,
		wrapWriter:   writer.Wrap,
		newHandler:   handler.NewTextHandler,
		replaceAttr:  nil,
		withSource:   false,
		withPID:      false,
		syncDuration: 0,
	}

	return conf
}

func (c *config) handlerOptions() *slog.HandlerOptions {
	opts := &slog.HandlerOptions{
		Level:       c.level,
		AddSource:   c.withSource,
		ReplaceAttr: c.replaceAttr,
	}

	return opts
}

func (c *config) handler() (slog.Handler, error) {
	w, err := c.newWriter()
	if err != nil {
		return nil, err
	}

	ww := c.wrapWriter(w)
	opts := c.handlerOptions()
	handler := c.newHandler(ww, opts)

	return handler, nil
}
