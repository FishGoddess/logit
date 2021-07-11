// Copyright 2021 Ye Zi Jie. All Rights Reserved.
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
// Created at 2021/06/27 16:36:33

package appender

import (
	"fmt"
	"strconv"
	"time"
	"unicode/utf8"
)

const (
	lineBreak = '\n'

	UnixTime = ""
)

type Appender interface {
	Begin(dst []byte) []byte
	End(dst []byte) []byte
	AppendAny(dst []byte, key string, value interface{}) []byte

	AppendBool(dst []byte, key string, value bool) []byte
	AppendByte(dst []byte, key string, value byte) []byte
	AppendRune(dst []byte, key string, value rune) []byte
	AppendInt(dst []byte, key string, value int) []byte
	AppendInt8(dst []byte, key string, value int8) []byte
	AppendInt16(dst []byte, key string, value int16) []byte
	AppendInt32(dst []byte, key string, value int32) []byte
	AppendInt64(dst []byte, key string, value int64) []byte
	AppendUint(dst []byte, key string, value uint) []byte
	AppendUint8(dst []byte, key string, value uint8) []byte
	AppendUint16(dst []byte, key string, value uint16) []byte
	AppendUint32(dst []byte, key string, value uint32) []byte
	AppendUint64(dst []byte, key string, value uint64) []byte
	AppendFloat32(dst []byte, key string, value float32) []byte
	AppendFloat64(dst []byte, key string, value float64) []byte
	AppendString(dst []byte, key string, value string) []byte
	AppendTime(dst []byte, key string, value time.Time, format string) []byte
	AppendError(dst []byte, key string, value error) []byte
	AppendStringer(dst []byte, key string, value fmt.Stringer) []byte

	AppendBools(dst []byte, key string, values []bool) []byte
	AppendBytes(dst []byte, key string, values []byte) []byte
	AppendRunes(dst []byte, key string, values []rune) []byte
	AppendInts(dst []byte, key string, values []int) []byte
	AppendInt8s(dst []byte, key string, values []int8) []byte
	AppendInt16s(dst []byte, key string, values []int16) []byte
	AppendInt32s(dst []byte, key string, values []int32) []byte
	AppendInt64s(dst []byte, key string, values []int64) []byte
	AppendUints(dst []byte, key string, values []uint) []byte
	AppendUint8s(dst []byte, key string, values []uint8) []byte
	AppendUint16s(dst []byte, key string, values []uint16) []byte
	AppendUint32s(dst []byte, key string, values []uint32) []byte
	AppendUint64s(dst []byte, key string, values []uint64) []byte
	AppendFloat32s(dst []byte, key string, values []float32) []byte
	AppendFloat64s(dst []byte, key string, values []float64) []byte
	AppendStrings(dst []byte, key string, values []string) []byte
	AppendTimes(dst []byte, key string, values []time.Time, format string) []byte
	AppendErrors(dst []byte, key string, values []error) []byte
	AppendStringers(dst []byte, key string, values []fmt.Stringer) []byte
}

// The main character should be escaped is ascii less than \u0020 and \ and ".
func needEscapedByte(value byte) bool {
	return value < 32 || value == '"' || value == '\\'
}

// The main character should be escaped is ascii less than \u0020 and \ and ".
func needEscapedRune(value rune) bool {
	return value < utf8.RuneSelf && needEscapedByte(byte(value))
}

// The main character should be escaped is ascii less than \u0020 and \ and ".
func appendEscapedByte(dst []byte, value byte) []byte {

	// ASCii < 16 needs to add \u000 to behind
	if value < 16 {
		return strconv.AppendInt(append(dst, '\\', 'u', '0', '0', '0'), int64(value), 16)
	}

	// ASCii in [16, 32) needs to add \u00 to behind
	if value < 32 {
		return strconv.AppendInt(append(dst, '\\', 'u', '0', '0'), int64(value), 16)
	}

	if value == '"' || value == '\\' {
		return append(dst, '\\', value)
	}
	return append(dst, value)
}

// The main character should be escaped is ascii less than \u0020 and \ and ".
func appendEscapedRune(dst []byte, value rune) []byte {

	if needEscapedRune(value) {
		return appendEscapedByte(dst, byte(value))
	}
	return append(dst, string(value)...)
}

// The main character should be escaped is ascii less than \u0020 and \ and ".
func appendEscapedString(dst []byte, value string) []byte {

	start := 0
	escaped := false
	for i := 0; i < len(value); i++ {
		// Encountered a byte that need escaping, so we appended bytes behinds it and appended it escaped
		if utf8.RuneStart(value[i]) && needEscapedByte(value[i]) {
			dst = append(dst, value[start:i]...)
			dst = appendEscapedByte(dst, value[i])
			start = i + 1
			escaped = true
		}
	}

	if escaped {
		return append(dst, value[start:]...)
	}

	// There is no need for escaping, just appending like bytes
	return append(dst, value...)
}

func Text() Appender {
	return (*textAppender)(nil)
}

func Json() Appender {
	return (*jsonAppender)(nil)
}
