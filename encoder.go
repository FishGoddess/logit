// Copyright 2020 Ye Zi Jie. All Rights Reserved.
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
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2020/04/14 21:06:56

package logit

import (
	"bytes"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

// Encoder encodes a log to bytes.
// No matter what you do, remember, return this log as bytes.
type Encoder interface {
	Encode(log *Log) []byte
}

type basedEncoder struct {
	timeFormat *atomic.Value
	buffers    *sync.Pool
}

func newBasedEncoder(timeFormat string) *basedEncoder {
	tf := &atomic.Value{}
	tf.Store(timeFormat)
	return &basedEncoder{
		timeFormat: tf,
		buffers: &sync.Pool{
			New: func() interface{} {
				return bytes.NewBuffer(make([]byte, 0, 64))
			},
		},
	}
}

func (be *basedEncoder) GetTimeFormat() string {
	return be.timeFormat.Load().(string)
}

func (be *basedEncoder) SetTimeFormat(timeFormat string) {
	be.timeFormat.Store(timeFormat)
}

func (be *basedEncoder) formatTime(t time.Time, quote bool) string {

	timeFormat := be.GetTimeFormat()
	if timeFormat == "" {
		return strconv.FormatInt(t.Unix(), 10)
	}

	result := t.Format(timeFormat)
	if quote {
		result = strconv.Quote(result)
	}
	return result
}

func (be *basedEncoder) newBuffer() *bytes.Buffer {
	result := be.buffers.Get().(*bytes.Buffer)
	result.Reset()
	return result
}

func (be *basedEncoder) releaseBuffer(buffer *bytes.Buffer) {
	be.buffers.Put(buffer)
}

func (be *basedEncoder) Encode(log *Log) []byte { return nil }

// =================================== text encoder ===================================

type textEncoder struct {
	*basedEncoder
}

func NewTextEncoder(timeFormat string) Encoder {
	return &textEncoder{
		basedEncoder: newBasedEncoder(timeFormat),
	}
}

func (te *textEncoder) Encode(log *Log) []byte {

	buffer := te.newBuffer()
	defer te.releaseBuffer(buffer)

	buffer.WriteString(te.formatTime(log.Time(), false))
	buffer.WriteByte('\t')
	buffer.WriteString(log.Level().String())
	buffer.WriteByte('\t')

	// Check caller
	if caller, ok := log.Caller(); ok {
		buffer.WriteString(caller.File + ":" + strconv.Itoa(caller.Line))
		buffer.WriteByte('\t')
	}

	buffer.WriteString(log.Msg())
	buffer.WriteString("\n")
	return buffer.Bytes()
}

// =================================== json encoder ===================================

type jsonEncoder struct {
	*basedEncoder
}

func NewJsonEncoder(timeFormat string) Encoder {
	return &jsonEncoder{
		basedEncoder: newBasedEncoder(timeFormat),
	}
}

// escapeString escapes string from special characters, such as double quotes.
// See issue: https://github.com/FishGoddess/logit/issues/1
func (je *jsonEncoder) escapeString(s string) string {

	buffer := bytes.NewBuffer(make([]byte, 0, 64))

	runes := []rune(s)
	for _, r := range runes {

		// The main character should be escaped is \ and " and ascii less than \u0020
		switch r {
		case '"', '\\':
			buffer.WriteRune('\\')
			buffer.WriteRune(r)
		default:
			// Notice: ascii < 16 needs to add \u000 to behind, ascii in [16, 32) needs to add \u00 to behind
			if r < 16 {
				buffer.WriteString("\\u000" + strconv.FormatInt(int64(r), 16))
			} else if r < 32 {
				buffer.WriteString("\\u00" + strconv.FormatInt(int64(r), 16))
			} else {
				buffer.WriteRune(r)
			}
		}
	}
	return buffer.String()
}

func (je *jsonEncoder) Encode(log *Log) []byte {

	buffer := je.newBuffer()
	defer je.releaseBuffer(buffer)

	buffer.WriteString(`{"level":"`)
	buffer.WriteString(log.Level().String())
	buffer.WriteString(`","time":`)
	buffer.WriteString(je.formatTime(log.Time(), true))

	// Check caller
	if caller, ok := log.Caller(); ok {
		buffer.WriteString(`,"file":"` + caller.File)
		buffer.WriteString(`","line":` + strconv.Itoa(caller.Line))
	}

	buffer.WriteString(`,"msg":"`)
	buffer.WriteString(je.escapeString(log.Msg()))
	buffer.WriteString("\"}\n")
	return buffer.Bytes()
}
