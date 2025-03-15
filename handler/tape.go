// Copyright 2025 FishGoddess. All Rights Reserved.
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

package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/FishGoddess/logit/defaults"
)

const (
	zero      = '0'
	lineBreak = '\n'

	dateConnector       = '-'
	clockConnector      = ':'
	timeConnector       = ' '
	timeMillisConnector = '.'

	keyValueConnector = '='
	sourceConnector   = ':'
	groupConnector    = "."
)

var (
	attrConnector = []byte(" ¦ ")
)

var (
	emptyAttr = slog.Attr{}
)

type tapeHandler struct {
	w    io.Writer
	opts slog.HandlerOptions

	group  string
	groups []string
	attrs  []slog.Attr

	lock *sync.Mutex
}

// NewTapeHandler creates a tape handler with w and opts.
// This handler is more readable and faster than slog's handlers.
func NewTapeHandler(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
	if opts == nil {
		opts = new(slog.HandlerOptions)
	}

	if opts.Level == nil {
		opts.Level = slog.LevelInfo
	}

	handler := &tapeHandler{
		w:    w,
		opts: *opts,
		lock: &sync.Mutex{},
	}

	return handler
}

// WithAttrs returns a new handler with attrs.
func (th *tapeHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	if len(attrs) <= 0 {
		return th
	}

	handler := *th
	handler.attrs = append(handler.attrs, attrs...)

	return &handler
}

// WithGroup returns a new handler with group.
func (th *tapeHandler) WithGroup(name string) slog.Handler {
	if name == "" {
		return th
	}

	handler := *th
	if handler.group != "" {
		handler.group = handler.group + groupConnector
	}

	handler.group = handler.group + name
	handler.groups = append(handler.groups, name)

	return &handler
}

func (th *tapeHandler) copyGroups(group string) []string {
	cap := cap(th.groups)
	if group != "" {
		cap += 1
	}

	groups := make([]string, 0, cap)
	groups = append(groups, th.groups...)

	if group != "" {
		groups = append(groups, group)
	}

	return groups
}

// Enabled reports whether the logger should ignore logs whose level is lower than passed level.
func (th *tapeHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= th.opts.Level.Level()
}

func (th *tapeHandler) appendKey(bs []byte, group string, key string) []byte {
	if key == "" {
		return bs
	}

	if th.group != "" {
		bs = appendEscapedString(bs, th.group)
		bs = append(bs, groupConnector...)
	}

	if group != "" {
		bs = appendEscapedString(bs, group)
		bs = append(bs, groupConnector...)
	}

	bs = appendEscapedString(bs, key)
	bs = append(bs, keyValueConnector)

	return bs
}

func (th *tapeHandler) appendBool(bs []byte, value bool) []byte {
	bs = strconv.AppendBool(bs, value)
	bs = append(bs, attrConnector...)

	return bs
}

func (th *tapeHandler) appendInt64(bs []byte, value int64) []byte {
	bs = strconv.AppendInt(bs, value, 10)
	bs = append(bs, attrConnector...)

	return bs
}

func (th *tapeHandler) appendUint64(bs []byte, value uint64) []byte {
	bs = strconv.AppendUint(bs, value, 10)
	bs = append(bs, attrConnector...)

	return bs
}

func (th *tapeHandler) appendFloat64(bs []byte, value float64) []byte {
	bs = strconv.AppendFloat(bs, value, 'f', -1, 64)
	bs = append(bs, attrConnector...)

	return bs
}

func (th *tapeHandler) appendString(bs []byte, value string) []byte {
	bs = appendEscapedString(bs, value)
	bs = append(bs, attrConnector...)

	return bs
}

func (th *tapeHandler) appendDuration(bs []byte, value time.Duration) []byte {
	bs = append(bs, value.String()...)
	bs = append(bs, attrConnector...)

	return bs
}

func (th *tapeHandler) appendTime(bs []byte, value time.Time) []byte {
	// Time format is an usual but expensive operation if using time.AppendFormat,
	// so we use a stupid but faster way to format time.
	// The result formatted is like "2006-01-02 15:04:05.000".
	year, month, day := value.Date()
	hour, minute, second := value.Clock()
	mircosecond := time.Duration(value.Nanosecond()) / time.Microsecond

	if year < 10 {
		bs = append(bs, zero, zero, zero)
	} else if year < 100 {
		bs = append(bs, zero, zero)
	} else if year < 1000 {
		bs = append(bs, zero)
	}

	bs = strconv.AppendInt(bs, int64(year), 10)
	bs = append(bs, dateConnector)

	if month < 10 {
		bs = append(bs, zero)
	}

	bs = strconv.AppendInt(bs, int64(month), 10)
	bs = append(bs, dateConnector)

	if day < 10 {
		bs = append(bs, zero)
	}

	bs = strconv.AppendInt(bs, int64(day), 10)
	bs = append(bs, timeConnector)

	if hour < 10 {
		bs = append(bs, zero)
	}

	bs = strconv.AppendInt(bs, int64(hour), 10)
	bs = append(bs, clockConnector)

	if minute < 10 {
		bs = append(bs, zero)
	}

	bs = strconv.AppendInt(bs, int64(minute), 10)
	bs = append(bs, clockConnector)

	if second < 10 {
		bs = append(bs, zero)
	}

	bs = strconv.AppendInt(bs, int64(second), 10)
	bs = append(bs, timeMillisConnector)

	if mircosecond < 10 {
		bs = append(bs, zero, zero, zero, zero, zero)
	} else if mircosecond < 100 {
		bs = append(bs, zero, zero, zero, zero)
	} else if mircosecond < 1000 {
		bs = append(bs, zero, zero, zero)
	} else if mircosecond < 10000 {
		bs = append(bs, zero, zero)
	} else if mircosecond < 100000 {
		bs = append(bs, zero)
	}

	bs = strconv.AppendInt(bs, int64(mircosecond), 10)
	bs = append(bs, attrConnector...)

	return bs
}

func (th *tapeHandler) appendAny(bs []byte, value any) []byte {
	if err, ok := value.(error); ok {
		bs = append(bs, err.Error()...)
		bs = append(bs, attrConnector...)

		return bs
	}

	if stringer, ok := value.(fmt.Stringer); ok {
		bs = append(bs, stringer.String()...)
		bs = append(bs, attrConnector...)

		return bs
	}

	marshaled, err := json.Marshal(value)
	if err == nil {
		bs = append(bs, marshaled...)
		bs = append(bs, attrConnector...)

		return bs
	}

	defaults.HandleError("json.Marshal", err)

	bs = fmt.Appendf(bs, "%+v", value)
	bs = append(bs, attrConnector...)

	return bs
}

func (th *tapeHandler) appendAttr(bs []byte, group string, attr slog.Attr) []byte {
	kind := attr.Value.Kind()
	replaceAttr := th.opts.ReplaceAttr

	if replaceAttr != nil && kind != slog.KindGroup {
		groups := th.copyGroups(group)
		attr.Value = attr.Value.Resolve()
		attr = replaceAttr(groups, attr)
	}

	// Resolve the Attr's value before doing anything else.
	attr.Value = attr.Value.Resolve()

	if attr.Equal(emptyAttr) {
		return bs
	}

	if kind == slog.KindGroup {
		bs = th.appendAttrs(bs, attr.Key, attr.Value.Group())

		return bs
	}

	bs = th.appendKey(bs, group, attr.Key)

	switch kind {
	case slog.KindBool:
		bs = th.appendBool(bs, attr.Value.Bool())
	case slog.KindInt64:
		bs = th.appendInt64(bs, attr.Value.Int64())
	case slog.KindUint64:
		bs = th.appendUint64(bs, attr.Value.Uint64())
	case slog.KindFloat64:
		bs = th.appendFloat64(bs, attr.Value.Float64())
	case slog.KindDuration:
		bs = th.appendDuration(bs, attr.Value.Duration())
	case slog.KindTime:
		bs = th.appendTime(bs, attr.Value.Time())
	case slog.KindAny:
		bs = th.appendAny(bs, attr.Value.Any())
	default:
		bs = th.appendString(bs, attr.Value.String())
	}

	return bs
}

func (th *tapeHandler) appendAttrs(bs []byte, group string, attrs []slog.Attr) []byte {
	for _, attr := range attrs {
		bs = th.appendAttr(bs, group, attr)
	}

	return bs
}

func (th *tapeHandler) appendSource(bs []byte, pc uintptr) []byte {
	if !th.opts.AddSource || pc == 0 {
		return bs
	}

	frames := runtime.CallersFrames([]uintptr{pc})
	frame, _ := frames.Next()

	bs = append(bs, slog.SourceKey...)
	bs = append(bs, keyValueConnector)
	bs = appendEscapedString(bs, frame.File)
	bs = append(bs, sourceConnector)
	bs = strconv.AppendInt(bs, int64(frame.Line), 10)
	bs = append(bs, attrConnector...)

	return bs
}

// Handle handles one record and returns an error if failed.
func (th *tapeHandler) Handle(ctx context.Context, record slog.Record) error {
	// Setup a buffer for handling record.
	buffer := newBuffer()
	bs := buffer.bs

	defer func() {
		buffer.bs = bs
		freeBuffer(buffer)
	}()

	// Handling record.
	bs = th.appendTime(bs, record.Time)
	bs = th.appendString(bs, record.Level.String())
	bs = th.appendString(bs, record.Message)
	bs = th.appendSource(bs, record.PC)
	bs = th.appendAttrs(bs, "", th.attrs)

	if record.NumAttrs() > 0 {
		record.Attrs(func(attr slog.Attr) bool {
			bs = th.appendAttr(bs, "", attr)
			return true
		})
	}

	bs = bytes.TrimSuffix(bs, attrConnector)
	bs = append(bs, lineBreak)

	// Write handled record.
	th.lock.Lock()
	defer th.lock.Unlock()

	_, err := th.w.Write(bs)
	return err
}
