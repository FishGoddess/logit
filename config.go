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

	"github.com/FishGoddess/logit/handler"
)

type config struct {
	level  slog.Level
	writer io.Writer

	newHandler  func(writer io.Writer, opts *slog.HandlerOptions) slog.Handler
	replaceAttr func(groups []string, attr slog.Attr) slog.Attr

	withSource bool
	withPID    bool
}

func newDefaultConfig() config {
	return config{
		level:       slog.LevelDebug,
		writer:      os.Stdout,
		newHandler:  newTextHandler,
		replaceAttr: nil,
		withSource:  false,
		withPID:     false,
	}
}

func (c *config) Accept(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
}

func (c *config) NewHandler() slog.Handler {
	opts := &slog.HandlerOptions{
		AddSource:   c.withSource,
		Level:       c.level,
		ReplaceAttr: c.replaceAttr,
	}

	return c.newHandler(c.writer, opts)
}

func newTextHandler(writer io.Writer, opts *slog.HandlerOptions) slog.Handler {
	return handler.NewTextHandler(writer, opts)
}

func newJsonHandler(writer io.Writer, opts *slog.HandlerOptions) slog.Handler {
	return handler.NewJsonHandler(writer, opts)
}
