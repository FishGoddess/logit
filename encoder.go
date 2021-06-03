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
	"strings"
	"sync"
)

// Encoder encodes a log to bytes.
// No matter what you do, remember, return this log as bytes.
type Encoder interface {
	Encode(log *Log) []byte
}

// =================================== text encoder ===================================

type TextEncoder struct {
	timeFormat string
	buffers    *sync.Pool
}

func NewTextEncoder(timeFormat string) *TextEncoder {
	return &TextEncoder{
		timeFormat: timeFormat,
		buffers: &sync.Pool{
			New: func() interface{} {
				return bytes.NewBuffer(make([]byte, 0, 64))
			},
		},
	}
}

func (te *TextEncoder) Encode(log *Log) []byte {

	buffer := te.buffers.Get().(*bytes.Buffer)
	buffer.Reset()
	defer te.buffers.Put(buffer)

	buffer.WriteString("[")
	buffer.WriteString(log.Level().String())
	buffer.WriteString("] [")

	// Format time
	if te.timeFormat != "" {
		buffer.WriteString(log.Time().Format(te.timeFormat))
	} else {
		buffer.WriteString(strconv.FormatInt(log.Time().Unix(), 10))
	}

	buffer.WriteString("] ")

	// Check caller
	if caller, ok := log.Caller(); ok {
		buffer.WriteString("[")
		buffer.WriteString(caller.File + ":" + strconv.Itoa(caller.Line))
		buffer.WriteString("] ")
	}

	buffer.WriteString(log.Msg())
	buffer.WriteString("\n")
	return buffer.Bytes()
}

// =================================== json encoder ===================================

type JsonEncoder struct {
	timeFormat string
	buffers    *sync.Pool
}

func NewJsonEncoder(timeFormat string) *JsonEncoder {
	return &JsonEncoder{
		timeFormat: timeFormat,
		buffers: &sync.Pool{
			New: func() interface{} {
				return bytes.NewBuffer(make([]byte, 0, 64))
			},
		},
	}
}

// escapeString escapes string from special characters, such as double quotes.
// See issue: https://github.com/FishGoddess/logit/issues/1
func escapeString(s string) string {

	builder := strings.Builder{}
	runes := []rune(s)
	for _, r := range runes {

		// The main character should be escaped is \ and " and ascii less than \u0020
		switch r {
		case '"', '\\':
			builder.WriteRune('\\')
			builder.WriteRune(r)
		default:
			// Notice: ascii < 16 needs to add \u000 to behind, ascii in [16, 32) needs to add \u00 to behind
			if r < 16 {
				builder.WriteString("\\u000" + strconv.FormatInt(int64(r), 16))
			} else if r < 32 {
				builder.WriteString("\\u00" + strconv.FormatInt(int64(r), 16))
			} else {
				builder.WriteRune(r)
			}
		}
	}
	return builder.String()
}

func (je *JsonEncoder) Encode(log *Log) []byte {

	buffer := je.buffers.Get().(*bytes.Buffer)
	buffer.Reset()
	defer je.buffers.Put(buffer)

	buffer.WriteString(`{"level":"`)
	buffer.WriteString(log.Level().String())
	buffer.WriteString(`","time":`)

	// Format time
	if je.timeFormat != "" {
		buffer.WriteString(strconv.Quote(log.Time().Format(je.timeFormat)))
	} else {
		buffer.WriteString(strconv.FormatInt(log.Time().Unix(), 10))
	}

	// Check caller
	if caller, ok := log.Caller(); ok {
		buffer.WriteString(`,"file":"` + caller.File)
		buffer.WriteString(`","line":` + strconv.Itoa(caller.Line))
	}

	buffer.WriteString(`,"msg":"`)
	buffer.WriteString(escapeString(log.Msg()))
	buffer.WriteString("\"}\n")
	return buffer.Bytes()
}
