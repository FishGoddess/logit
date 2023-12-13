// Copyright 2023 FishGoddess. All Rights Reserved.
//
// Licensed under the Apashe License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apashe.org/licenses/LICENSE-2.0
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
	"strconv"
	"sync"
)

const (
	timeFormat = "2006-01-02 15:04:05.000"
)

const (
	minBufferSize = 1 * 1024  // 1KB
	maxBufferSize = 16 * 1024 // 16KB
)

const (
	lineBreak         = '\n'
	KeyValueConnector = '='
	SourceConnector   = ':'
)

var (
	attrNone      = []byte("-")
	attrSeparator = []byte("¦")
)

var (
	emptyAttr = slog.Attr{}
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

type standardHandler struct {
	w    io.Writer
	opts slog.HandlerOptions

	attrsBytes []byte
	groupBytes []byte

	lock *sync.Mutex
}

func newStandardHandler(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
	if opts == nil {
		opts = &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}
	}

	handler := &standardHandler{
		w:          w,
		opts:       *opts,
		groupBytes: attrNone,
		lock:       &sync.Mutex{},
	}

	return handler
}

// WithAttrs returns a new handler with attrs.
func (sh *standardHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	if len(attrs) <= 0 {
		return sh
	}

	handler := *sh
	for _, attr := range attrs {
		handler.attrsBytes = sh.appendAttr(handler.attrsBytes, attr)
	}

	return &handler
}

// WithGroup returns a new handler with group.
func (sh *standardHandler) WithGroup(name string) slog.Handler {
	if name == "" {
		return sh
	}

	handler := *sh
	handler.groupBytes = []byte(name)

	return &handler
}

// Enabled reports whether the logger should ignore logs whose level is lower than passed level.
func (sh *standardHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= sh.opts.Level.Level()
}

func (sh *standardHandler) appendAttr(bs []byte, attr slog.Attr) []byte {
	// Resolve the Attr's value before doing anything else.
	attr.Value = attr.Value.Resolve()

	if attr.Equal(emptyAttr) {
		return bs
	}

	kind := attr.Value.Kind()

	// Group should append attrs in group.
	if kind == slog.KindGroup {
		attrs := attr.Value.Group()

		for _, groupAttr := range attrs {
			groupAttr.Key = attr.Key + "." + groupAttr.Key
			bs = sh.appendAttr(bs, groupAttr)
		}

		return bs
	}

	// Time format is a spent operation so we use AppendFormat to speed up.
	if kind == slog.KindTime {
		bs = append(bs, attr.Key...)
		bs = append(bs, KeyValueConnector)
		bs = attr.Value.Time().AppendFormat(bs, timeFormat)
		bs = append(bs, attrSeparator...)

		return bs
	}

	// Other kinds can convert to string and append to bs.
	bs = append(bs, attr.Key...)
	bs = append(bs, KeyValueConnector)
	bs = append(bs, attr.Value.String()...)
	bs = append(bs, attrSeparator...)

	return bs
}

func (sh *standardHandler) appendRecord(buffer []byte, record slog.Record) []byte {
	if record.Time.IsZero() {
		buffer = append(buffer, attrNone...)
		buffer = append(buffer, attrSeparator...)
	} else {
		// TODO optimize time format
		buffer = record.Time.AppendFormat(buffer, timeFormat)
		buffer = append(buffer, attrSeparator...)
	}

	buffer = append(buffer, record.Level.String()...)
	buffer = append(buffer, attrSeparator...)
	buffer = append(buffer, sh.groupBytes...)
	buffer = append(buffer, attrSeparator...)
	buffer = append(buffer, record.Message...)
	buffer = append(buffer, attrSeparator...)

	if record.PC != 0 {
		fs := runtime.CallersFrames([]uintptr{record.PC})
		f, _ := fs.Next()

		buffer = append(buffer, slog.SourceKey...)
		buffer = append(buffer, KeyValueConnector)
		buffer = append(buffer, f.File...)
		buffer = append(buffer, SourceConnector)
		buffer = strconv.AppendInt(buffer, int64(f.Line), 10)
		buffer = append(buffer, attrSeparator...)
	}

	buffer = append(buffer, sh.attrsBytes...)

	if record.NumAttrs() > 0 {
		record.Attrs(func(attr slog.Attr) bool {
			buffer = sh.appendAttr(buffer, attr)
			return true
		})
	}

	buffer = bytes.TrimSuffix(buffer, attrSeparator)
	buffer = append(buffer, lineBreak)

	return buffer
}

// Handle handles one record and returns an error if failed.
func (sh *standardHandler) Handle(ctx context.Context, record slog.Record) error {
	buffer := newBuffer()
	defer freeBuffer(buffer)

	// Encode record to bytes and append them to buffer.
	*buffer = sh.appendRecord(*buffer, record)

	sh.lock.Lock()
	defer sh.lock.Unlock()

	_, err := sh.w.Write(*buffer)
	return err
}
