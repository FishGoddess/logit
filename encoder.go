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
)

// Encoder is for encoding a log to bytes with timeFormat.
// No matter what you do, remember, return this log as bytes.
type Encoder func(log *Log, timeFormat string) []byte

// Encode encodes a log to bytes with timeFormat.
// This is Encoder's substitute, and only for more code-readable.
func (e Encoder) Encode(log *Log, timeFormat string) []byte {
	return e(log, timeFormat)
}

// =================================== text encoder ===================================

// TextEncoder encodes a log to a plain string like "[Info] [2020-03-06 16:10:44] msg" in bytes.
// If timeFormat == "", then it will not format time and keep time in unix form.
func TextEncoder() Encoder {
	return func(log *Log, timeFormat string) []byte {

		buffer := bytes.NewBuffer(make([]byte, 0, 64))
		buffer.WriteString("[")
		buffer.WriteString(log.Level().String())
		buffer.WriteString("] [")

		// format time
		if timeFormat != "" {
			buffer.WriteString(log.Time().Format(timeFormat))
		} else {
			buffer.WriteString(strconv.FormatInt(log.Time().Unix(), 10))
		}

		buffer.WriteString("] ")

		// Add caller information if need
		if caller, ok := log.Caller(); ok {
			buffer.WriteString("[")
			buffer.WriteString(caller.File + ":" + strconv.Itoa(caller.Line))
			buffer.WriteString("] ")
		}

		buffer.WriteString(log.Msg())
		buffer.WriteString("\n")
		return buffer.Bytes()
	}
}

// =================================== json encoder ===================================

// JsonEncoder encodes a log to a Json string in bytes.
// If timeFormat == "", then it will not format time and keep time in unix form.
// The result looks like `{"level":"debug", "time":"2020-03-22 22:35:00", "msg":"log content..."}`.
func JsonEncoder() Encoder {
	return func(log *Log, timeFormat string) []byte {

		buffer := bytes.NewBuffer(make([]byte, 0, 64))
		buffer.WriteString(`{"level":"`)
		buffer.WriteString(log.Level().String())
		buffer.WriteString(`","time":`)

		// format time
		if timeFormat != "" {
			buffer.WriteString(strconv.Quote(log.Time().Format(timeFormat)))
		} else {
			buffer.WriteString(strconv.FormatInt(log.Time().Unix(), 10))
		}

		// Add caller information if need
		if caller, ok := log.Caller(); ok {
			buffer.WriteString(`,"file":"` + caller.File)
			buffer.WriteString(`","line":` + strconv.Itoa(caller.Line))
		}

		buffer.WriteString(`,"msg":"`)
		buffer.WriteString(escapeString(log.Msg()))
		buffer.WriteString("\"}\n")
		return buffer.Bytes()
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
