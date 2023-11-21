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
	newTextHandler := func(writer io.Writer, opts *slog.HandlerOptions) slog.Handler {
		return slog.NewTextHandler(writer, opts)
	}

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

type Option func(conf *config)

func WithDebugLevel() Option {
	return func(conf *config) {
		conf.level = slog.LevelDebug
	}
}

func WithInfoLevel() Option {
	return func(conf *config) {
		conf.level = slog.LevelInfo
	}
}

func WithWarnLevel() Option {
	return func(conf *config) {
		conf.level = slog.LevelWarn
	}
}

func WithErrorLevel() Option {
	return func(conf *config) {
		conf.level = slog.LevelError
	}
}

func WithNewHandler(newHandler func(writer io.Writer, opts *slog.HandlerOptions) slog.Handler) Option {
	return func(conf *config) {
		conf.newHandler = newHandler
	}
}

func WithJsonHandler() Option {
	newJsonHandler := func(writer io.Writer, opts *slog.HandlerOptions) slog.Handler {
		return slog.NewJSONHandler(writer, opts)
	}

	return WithNewHandler(newJsonHandler)
}

func WithReplaceAttr(replaceAttr func(groups []string, attr slog.Attr) slog.Attr) Option {
	return func(conf *config) {
		conf.replaceAttr = replaceAttr
	}
}

func WithSource() Option {
	return func(conf *config) {
		conf.withSource = true
	}
}

func WithPID() Option {
	return func(conf *config) {
		conf.withPID = true
	}
}
