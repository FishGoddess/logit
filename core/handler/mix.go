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

type mixHandler struct {
	w    io.Writer
	opts slog.HandlerOptions

	group  string
	groups []string
	attrs  []slog.Attr

	lock *sync.Mutex
}

// NewMixHandler creates a mix handler with w and opts.
// This handler is more readable and faster than slog's handlers.
func NewMixHandler(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
	if opts == nil {
		opts = new(slog.HandlerOptions)
	}

	if opts.Level == nil {
		opts.Level = slog.LevelInfo
	}

	handler := &mixHandler{
		w:    w,
		opts: *opts,
		lock: &sync.Mutex{},
	}

	return handler
}

// WithAttrs returns a new handler with attrs.
func (mh *mixHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	if len(attrs) <= 0 {
		return mh
	}

	handler := *mh
	handler.attrs = append(handler.attrs, attrs...)

	return &handler
}

// WithGroup returns a new handler with group.
func (mh *mixHandler) WithGroup(name string) slog.Handler {
	if name == "" {
		return mh
	}

	handler := *mh
	if handler.group != "" {
		handler.group = handler.group + groupConnector
	}

	handler.group = handler.group + name
	handler.groups = append(handler.groups, name)

	return &handler
}

func (mh *mixHandler) copyGroups(group string) []string {
	cap := cap(mh.groups)
	if group != "" {
		cap += 1
	}

	groups := make([]string, 0, cap)
	groups = append(groups, mh.groups...)

	if group != "" {
		groups = append(groups, group)
	}

	return groups
}

// Enabled reports whether the logger should ignore logs whose level is lower than passed level.
func (mh *mixHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= mh.opts.Level.Level()
}

func (mh *mixHandler) appendKey(bs []byte, group string, key string) []byte {
	if key == "" {
		return bs
	}

	if mh.group != "" {
		bs = appendEscapedString(bs, mh.group)
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

func (mh *mixHandler) appendBool(bs []byte, value bool) []byte {
	bs = strconv.AppendBool(bs, value)
	bs = append(bs, attrConnector...)

	return bs
}

func (mh *mixHandler) appendInt64(bs []byte, value int64) []byte {
	bs = strconv.AppendInt(bs, value, 10)
	bs = append(bs, attrConnector...)

	return bs
}

func (mh *mixHandler) appendUint64(bs []byte, value uint64) []byte {
	bs = strconv.AppendUint(bs, value, 10)
	bs = append(bs, attrConnector...)

	return bs
}

func (mh *mixHandler) appendFloat64(bs []byte, value float64) []byte {
	bs = strconv.AppendFloat(bs, value, 'f', -1, 64)
	bs = append(bs, attrConnector...)

	return bs
}

func (mh *mixHandler) appendString(bs []byte, value string) []byte {
	bs = appendEscapedString(bs, value)
	bs = append(bs, attrConnector...)

	return bs
}

func (mh *mixHandler) appendDuration(bs []byte, value time.Duration) []byte {
	bs = append(bs, value.String()...)
	bs = append(bs, attrConnector...)

	return bs
}

func (mh *mixHandler) appendTime(bs []byte, value time.Time) []byte {
	// Time format is an usual but expensive operation if using time.AppendFormat,
	// so we use a stupid but faster way to format time.
	// The result formatted is like "2006-01-02 15:04:05.000".
	year, month, day := value.Date()
	hour, minute, second := value.Clock()
	millisecond := time.Duration(value.Nanosecond()) / time.Millisecond

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

	if millisecond < 10 {
		bs = append(bs, zero, zero)
	} else if millisecond < 100 {
		bs = append(bs, zero)
	}

	bs = strconv.AppendInt(bs, int64(millisecond), 10)
	bs = append(bs, attrConnector...)

	return bs
}

func (mh *mixHandler) appendAny(bs []byte, value any) []byte {
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

func (mh *mixHandler) appendAttr(bs []byte, group string, attr slog.Attr) []byte {
	kind := attr.Value.Kind()
	replaceAttr := mh.opts.ReplaceAttr

	if replaceAttr != nil && kind != slog.KindGroup {
		groups := mh.copyGroups(group)
		attr.Value = attr.Value.Resolve()
		attr = replaceAttr(groups, attr)
	}

	// Resolve the Attr's value before doing anything else.
	attr.Value = attr.Value.Resolve()

	if attr.Equal(emptyAttr) {
		return bs
	}

	if kind == slog.KindGroup {
		bs = mh.appendAttrs(bs, attr.Key, attr.Value.Group())

		return bs
	}

	bs = mh.appendKey(bs, group, attr.Key)

	switch kind {
	case slog.KindBool:
		bs = mh.appendBool(bs, attr.Value.Bool())
	case slog.KindInt64:
		bs = mh.appendInt64(bs, attr.Value.Int64())
	case slog.KindUint64:
		bs = mh.appendUint64(bs, attr.Value.Uint64())
	case slog.KindFloat64:
		bs = mh.appendFloat64(bs, attr.Value.Float64())
	case slog.KindDuration:
		bs = mh.appendDuration(bs, attr.Value.Duration())
	case slog.KindTime:
		bs = mh.appendTime(bs, attr.Value.Time())
	case slog.KindAny:
		bs = mh.appendAny(bs, attr.Value.Any())
	default:
		bs = mh.appendString(bs, attr.Value.String())
	}

	return bs
}

func (mh *mixHandler) appendAttrs(bs []byte, group string, attrs []slog.Attr) []byte {
	for _, attr := range attrs {
		bs = mh.appendAttr(bs, group, attr)
	}

	return bs
}

func (mh *mixHandler) appendSource(bs []byte, pc uintptr) []byte {
	if !mh.opts.AddSource || pc == 0 {
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
func (mh *mixHandler) Handle(ctx context.Context, record slog.Record) error {
	// Setup a buffer for handling record.
	buffer := newBuffer()
	bs := buffer.bs

	defer func() {
		buffer.bs = bs
		freeBuffer(buffer)
	}()

	// Handling record.
	bs = mh.appendTime(bs, record.Time)
	bs = mh.appendString(bs, record.Level.String())
	bs = mh.appendString(bs, record.Message)
	bs = mh.appendSource(bs, record.PC)
	bs = mh.appendAttrs(bs, "", mh.attrs)

	if record.NumAttrs() > 0 {
		record.Attrs(func(attr slog.Attr) bool {
			bs = mh.appendAttr(bs, "", attr)
			return true
		})
	}

	bs = bytes.TrimSuffix(bs, attrConnector)
	bs = append(bs, lineBreak)

	// Write handled record.
	mh.lock.Lock()
	defer mh.lock.Unlock()

	_, err := mh.w.Write(bs)
	return err
}
