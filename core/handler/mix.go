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

	keyValueSeparator = '='
	sourceSeparator   = ':'
	groupSeparator    = "."
)

var (
	attrSeparator = []byte(" ¦ ")
)

var (
	emptyAttr = slog.Attr{}
)

type mixHandler struct {
	w    io.Writer
	opts slog.HandlerOptions

	group      string
	attrsBytes []byte

	lock *sync.Mutex
}

// NewMixHandler creates a mix handler with w and opts.
// This handler is more readable and faster than slog's handlers.
func NewMixHandler(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
	if opts == nil {
		opts = &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}
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
	handler.attrsBytes = mh.appendGroupAttrs(handler.attrsBytes, mh.group, attrs)

	return &handler
}

// WithGroup returns a new handler with group.
func (mh *mixHandler) WithGroup(name string) slog.Handler {
	if name == "" {
		return mh
	}

	handler := *mh
	handler.group = name

	return &handler
}

// Enabled reports whether the logger should ignore logs whose level is lower than passed level.
func (mh *mixHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= mh.opts.Level.Level()
}

func (mh *mixHandler) appendKey(bs []byte, key string) []byte {
	if key != "" {
		bs = appendEscapedString(bs, key)
		bs = append(bs, keyValueSeparator)
	}

	return bs
}

func (mh *mixHandler) appendString(bs []byte, value string) []byte {
	bs = appendEscapedString(bs, value)
	bs = append(bs, attrSeparator...)

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
	bs = append(bs, attrSeparator...)

	return bs
}

func (mh *mixHandler) appendAny(bs []byte, value any) []byte {
	if err, ok := value.(error); ok {
		bs = append(bs, err.Error()...)
		bs = append(bs, attrSeparator...)

		return bs
	}

	if stringer, ok := value.(fmt.Stringer); ok {
		bs = append(bs, stringer.String()...)
		bs = append(bs, attrSeparator...)

		return bs
	}

	marshaled, err := json.Marshal(value)
	if err == nil {
		bs = append(bs, marshaled...)
		bs = append(bs, attrSeparator...)

		return bs
	}

	defaults.HandleError("json.Marshal", err)

	bs = fmt.Appendf(bs, "%+v", value)
	bs = append(bs, attrSeparator...)

	return bs
}

func (mh *mixHandler) appendAttr(bs []byte, attr slog.Attr) []byte {
	// Resolve the Attr's value before doing anything else.
	attr.Value = attr.Value.Resolve()

	if attr.Equal(emptyAttr) {
		return bs
	}

	bs = mh.appendKey(bs, attr.Key)

	switch attr.Value.Kind() {
	case slog.KindGroup:
		bs = mh.appendGroupAttrs(bs, attr.Key, attr.Value.Group())
	case slog.KindTime:
		bs = mh.appendTime(bs, attr.Value.Time())
	case slog.KindAny:
		bs = mh.appendAny(bs, attr.Value.Any())
	default:
		bs = mh.appendString(bs, attr.Value.String())
	}

	return bs
}

func (mh *mixHandler) appendGroupAttrs(bs []byte, group string, attrs []slog.Attr) []byte {
	for _, groupAttr := range attrs {
		if group != "" {
			groupAttr.Key = group + groupSeparator + groupAttr.Key
		}

		bs = mh.appendAttr(bs, groupAttr)
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
	bs = append(bs, keyValueSeparator)
	bs = appendEscapedString(bs, frame.File)
	bs = append(bs, sourceSeparator)
	bs = strconv.AppendInt(bs, int64(frame.Line), 10)
	bs = append(bs, attrSeparator...)

	return bs
}

// Handle handles one record and returns an error if failed.
func (mh *mixHandler) Handle(ctx context.Context, record slog.Record) error {
	// Setup a buffer for handling record.
	bufferPtr := newBuffer()
	buffer := *bufferPtr

	defer func() {
		bufferPtr = &buffer
		freeBuffer(bufferPtr)
	}()

	// Handling record.
	buffer = mh.appendTime(buffer, record.Time)
	buffer = mh.appendString(buffer, record.Level.String())
	buffer = mh.appendString(buffer, record.Message)
	buffer = mh.appendSource(buffer, record.PC)

	buffer = append(buffer, mh.attrsBytes...)
	if record.NumAttrs() > 0 {
		record.Attrs(func(attr slog.Attr) bool {
			buffer = mh.appendAttr(buffer, attr)
			return true
		})
	}

	buffer = bytes.TrimSuffix(buffer, attrSeparator)
	buffer = append(buffer, lineBreak)

	// Write handled record.
	mh.lock.Lock()
	defer mh.lock.Unlock()

	_, err := mh.w.Write(buffer)
	return err
}
