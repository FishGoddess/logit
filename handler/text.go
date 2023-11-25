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

package handler

import (
	"context"
	"io"
	"log/slog"
)

type TextHandler struct {
	handler slog.Handler
}

func newTextHandler(w io.Writer, opts *slog.HandlerOptions) *TextHandler {
	handler := &TextHandler{
		handler: slog.NewTextHandler(w, opts),
	}

	return handler
}

func (th *TextHandler) WithGroup(name string) slog.Handler {
	return th.handler.WithGroup(name)
}

func (th *TextHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return th.handler.WithAttrs(attrs)
}

func (th *TextHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return th.handler.Enabled(ctx, level)
}

func (th *TextHandler) Handle(ctx context.Context, record slog.Record) error {
	return th.handler.Handle(ctx, record)
}