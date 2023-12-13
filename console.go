// Copyright 2023 FichGoddess. All Rights Reserved.
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
	"runtime"
	"slices"
	"strconv"
	"sync"
	"time"
)

const (
	minBufferSize = 1 * 1024  // 1KB
	maxBufferSize = 16 * 1024 // 16KB
)

var (
	attrSeparator = []byte("¦")
)

var bufferPool = sync.Pool{
	New: func() any {
		bs := make([]byte, 0, minBufferSize)
		return &bs
	},
}

func newBuffer() *[]byte {
	return bufferPool.Get().(*[]byte)
}

func freeBuffer(b *[]byte) {
	// Return only smaller buffers for reducing peak allocation.
	if cap(*b) <= maxBufferSize {
		*b = (*b)[:0]
		bufferPool.Put(b)
	}
}

type consoleHandler struct {
	w    io.Writer
	opts slog.HandlerOptions

	group string
	attrs []slog.Attr

	mu *sync.Mutex
}

func newConsoleHandler(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
	h := &consoleHandler{
		w:     w,
		group: "-",
		mu:    &sync.Mutex{},
	}

	if opts != nil {
		h.opts = *opts
	}

	if h.opts.Level == nil {
		h.opts.Level = slog.LevelInfo
	}

	return h
}

func (ch *consoleHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= ch.opts.Level.Level()
}

func (ch *consoleHandler) WithGroup(name string) slog.Handler {
	if name == "" {
		return ch
	}

	h2 := *ch
	h2.group = name

	return &h2
}

func (ch *consoleHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	if len(attrs) <= 0 {
		return ch
	}

	h2 := *ch

	h2.attrs = slices.Clip(h2.attrs)
	h2.attrs = append(h2.attrs, attrs...)

	return &h2
}

func (ch *consoleHandler) Handle(ctx context.Context, record slog.Record) error {
	bufferPtr := newBuffer()
	buffer := *bufferPtr

	defer func() {
		*bufferPtr = buffer
		freeBuffer(bufferPtr)
	}()

	if !record.Time.IsZero() {
		buffer = record.Time.AppendFormat(buffer, time.RFC3339)
		buffer = append(buffer, attrSeparator...)
	}

	buffer = append(buffer, record.Level.String()...)
	buffer = append(buffer, attrSeparator...)
	buffer = append(buffer, ch.group...)
	buffer = append(buffer, attrSeparator...)
	buffer = append(buffer, record.Message...)
	buffer = append(buffer, attrSeparator...)

	if record.PC != 0 {
		fs := runtime.CallersFrames([]uintptr{record.PC})
		f, _ := fs.Next()

		// Optimize to minimize allocation.
		smallBuf := newBuffer()
		defer freeBuffer(smallBuf)

		*smallBuf = append(*smallBuf, f.File...)
		*smallBuf = append(*smallBuf, ':')
		*smallBuf = strconv.AppendInt(*smallBuf, int64(f.Line), 10)

		buffer = ch.appendAttr(buffer, slog.String(slog.SourceKey, string(*smallBuf)))
	}

	for _, attr := range ch.attrs {
		buffer = ch.appendAttr(buffer, attr)
	}

	if record.NumAttrs() > 0 {
		record.Attrs(func(attr slog.Attr) bool {
			buffer = ch.appendAttr(buffer, attr)
			return true
		})
	}

	buffer = bytes.TrimSuffix(buffer, attrSeparator)
	buffer = append(buffer, '\n')

	ch.mu.Lock()
	defer ch.mu.Unlock()

	_, err := ch.w.Write(buffer)
	return err
}

func (ch *consoleHandler) appendAttr(bs []byte, attr slog.Attr) []byte {
	// Resolve the Attr's value before doing anything else.
	attr.Value = attr.Value.Resolve()

	// Ignore empty Attrs.
	if attr.Equal(slog.Attr{}) {
		return bs
	}

	switch attr.Value.Kind() {
	case slog.KindTime:
		// Write times in a standard way, without the monotonic time.
		bs = append(bs, attr.Key...)
		bs = append(bs, '=')
		bs = attr.Value.Time().AppendFormat(bs, time.RFC3339)
		bs = append(bs, attrSeparator...)
	case slog.KindGroup:
		attrs := attr.Value.Group()

		// Ignore empty groups.
		if len(attrs) == 0 {
			return bs
		}

		for _, gAttr := range attrs {
			gAttr.Key = attr.Key + "." + gAttr.Key
			bs = ch.appendAttr(bs, gAttr)
		}
	default:
		bs = append(bs, attr.Key...)
		bs = append(bs, '=')
		bs = append(bs, attr.Value.String()...)
		bs = append(bs, attrSeparator...)
	}

	return bs
}
