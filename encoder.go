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
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

type buffer struct {
	buff []byte
}

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
				return &buffer{buff: make([]byte, 0, 256)}
			},
		},
	}
}

func (be *basedEncoder) TimeFormat() string {
	return be.timeFormat.Load().(string)
}

func (be *basedEncoder) SetTimeFormat(timeFormat string) {
	be.timeFormat.Store(timeFormat)
}

func (be *basedEncoder) formatTime(t time.Time, quote bool) string {

	timeFormat := be.TimeFormat()
	if timeFormat == "" {
		return strconv.FormatInt(t.Unix(), 10)
	}

	// TODO this should be checked for usage
	result := string(t.AppendFormat(make([]byte, 0, 24), timeFormat))
	if quote {
		result = strconv.Quote(result)
	}
	return result
}

func (be *basedEncoder) newBuffer() *buffer {
	result := be.buffers.Get().(*buffer)
	result.buff = result.buff[:0]
	return result
}

func (be *basedEncoder) releaseBuffer(buffer *buffer) {
	be.buffers.Put(buffer)
}

func (be *basedEncoder) Encode(log *Log) []byte { return nil }

// =================================== text encoder ===================================

type TextEncoder struct {
	*basedEncoder
}

func NewTextEncoder(timeFormat string) *TextEncoder {
	return &TextEncoder{
		basedEncoder: newBasedEncoder(timeFormat),
	}
}

func (te *TextEncoder) Encode(log *Log) []byte {

	buffer := te.newBuffer()
	defer te.releaseBuffer(buffer)

	if te.TimeFormat() == "" {
		buffer.buff = strconv.AppendInt(buffer.buff, log.Time().Unix(), 10)
	} else {
		buffer.buff = log.Time().AppendFormat(buffer.buff, te.TimeFormat())
	}
	buffer.buff = append(buffer.buff, '\t')
	buffer.buff = append(buffer.buff, log.Level().String()...)
	buffer.buff = append(buffer.buff, '\t')

	// Check caller
	if caller, ok := log.Caller(); ok {
		buffer.buff = append(buffer.buff, caller.File...)
		buffer.buff = append(buffer.buff, ':')
		buffer.buff = strconv.AppendInt(buffer.buff, int64(caller.Line), 10)
		buffer.buff = append(buffer.buff, '\t')
	}

	// TODO need optimization
	for k, v := range log.KVs() {
		buffer.buff = append(buffer.buff, fmt.Sprintf("%s=%+v", k, v)...)
		buffer.buff = append(buffer.buff, '\t')
	}

	buffer.buff = append(buffer.buff, log.Msg()...)
	buffer.buff = append(buffer.buff, '\n')
	return buffer.buff
}

// =================================== json encoder ===================================

type JsonEncoder struct {
	*basedEncoder
}

func NewJsonEncoder(timeFormat string) *JsonEncoder {
	return &JsonEncoder{
		basedEncoder: newBasedEncoder(timeFormat),
	}
}

// escapeString escapes string from special characters, such as double quotes.
// See issue: https://github.com/FishGoddess/logit/issues/1
func (je *JsonEncoder) escapeString(s string) string {

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

func (je *JsonEncoder) Encode(log *Log) []byte {

	buffer := je.newBuffer()
	defer je.releaseBuffer(buffer)

	buffer.buff = append(buffer.buff, `{"level":"`...)
	buffer.buff = append(buffer.buff, log.Level().String()...)
	buffer.buff = append(buffer.buff, `","time":`...)
	buffer.buff = append(buffer.buff, je.formatTime(log.Time(), true)...)

	// Check caller
	if caller, ok := log.Caller(); ok {
		buffer.buff = append(buffer.buff, `,"file":"`...)
		buffer.buff = append(buffer.buff, caller.File...)
		buffer.buff = append(buffer.buff, `","line":`...)
		buffer.buff = strconv.AppendInt(buffer.buff, int64(caller.Line), 10)
	}

	// TODO encode kvs

	buffer.buff = append(buffer.buff, `,"msg":"`...)
	buffer.buff = append(buffer.buff, je.escapeString(log.Msg())...)
	buffer.buff = append(buffer.buff, "\"}\n"...)
	return buffer.buff
}
