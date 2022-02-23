// Copyright 2022 FishGoddess. All Rights Reserved.
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

package appender

import (
	"fmt"
	"strconv"
	"time"
	"unicode/utf8"
)

const (
	// lineBreak is the break between lines.
	lineBreak = '\n'

	// UnixTime is a flag that keeps time in unix format.
	UnixTime = ""
)

var (
	globalTextAppender Appender = (*textAppender)(nil) // Check and use.
	globalJsonAppender Appender = (*jsonAppender)(nil) // Check and use.
)

// Appender is an interface of appending entries to a byte slice.
// Although we keep it public, we don't recommend you to implement it.
type Appender interface {
	Begin(dst []byte) []byte                                    // Begin appends begin character to dst.
	End(dst []byte) []byte                                      // End appends end character to dst.
	AppendAny(dst []byte, key string, value interface{}) []byte // AppendAny appends any entries to dst.

	AppendBool(dst []byte, key string, value bool) []byte                     // AppendBool appends a bool entry to dst.
	AppendByte(dst []byte, key string, value byte) []byte                     // AppendByte appends a byte entry to dst.
	AppendRune(dst []byte, key string, value rune) []byte                     // AppendRune appends a rune entry to dst.
	AppendInt(dst []byte, key string, value int) []byte                       // AppendInt appends an int entry to dst.
	AppendInt8(dst []byte, key string, value int8) []byte                     // AppendInt8 appends an int8 entry to dst.
	AppendInt16(dst []byte, key string, value int16) []byte                   // AppendInt16 appends an int16 entry to dst.
	AppendInt32(dst []byte, key string, value int32) []byte                   // AppendInt32 appends an int32 entry to dst.
	AppendInt64(dst []byte, key string, value int64) []byte                   // AppendInt64 appends an int64 entry to dst.
	AppendUint(dst []byte, key string, value uint) []byte                     // AppendUint appends an uint entry to dst.
	AppendUint8(dst []byte, key string, value uint8) []byte                   // AppendUint8 appends an uin8 entry to dst.
	AppendUint16(dst []byte, key string, value uint16) []byte                 // AppendUint16 appends an uint16 entry to dst.
	AppendUint32(dst []byte, key string, value uint32) []byte                 // AppendUint32 appends an uint32 entry to dst.
	AppendUint64(dst []byte, key string, value uint64) []byte                 // AppendUint64 appends an uint64 entry to dst.
	AppendFloat32(dst []byte, key string, value float32) []byte               // AppendFloat32 appends a float32 entry to dst.
	AppendFloat64(dst []byte, key string, value float64) []byte               // AppendFloat64 appends a float64 entry to dst.
	AppendString(dst []byte, key string, value string) []byte                 // AppendString appends a string entry to dst.
	AppendTime(dst []byte, key string, value time.Time, format string) []byte // AppendTime appends a time.Time entry formatted with format to dst.
	AppendError(dst []byte, key string, value error) []byte                   // AppendError appends an error entry to dst.
	AppendStringer(dst []byte, key string, value fmt.Stringer) []byte         // AppendStringer appends an fmt.Stringer entry to dst.

	AppendBools(dst []byte, key string, values []bool) []byte                     // AppendBools appends a []bool entry to dst.
	AppendBytes(dst []byte, key string, values []byte) []byte                     // AppendBytes appends a []byte entry to dst.
	AppendRunes(dst []byte, key string, values []rune) []byte                     // AppendRunes appends a []rune entry to dst.
	AppendInts(dst []byte, key string, values []int) []byte                       // AppendInts appends an []int entry to dst.
	AppendInt8s(dst []byte, key string, values []int8) []byte                     // AppendInt8s appends an []int8 entry to dst.
	AppendInt16s(dst []byte, key string, values []int16) []byte                   // AppendInt16s appends an []int16 entry to dst.
	AppendInt32s(dst []byte, key string, values []int32) []byte                   // AppendInt32s appends an []int32 entry to dst.
	AppendInt64s(dst []byte, key string, values []int64) []byte                   // AppendInt64s appends an []int64 entry to dst.
	AppendUints(dst []byte, key string, values []uint) []byte                     // AppendUints appends an []uint entry to dst.
	AppendUint8s(dst []byte, key string, values []uint8) []byte                   // AppendUint8s appends an []uint8 entry to dst.
	AppendUint16s(dst []byte, key string, values []uint16) []byte                 // AppendUint16s appends an []uint16 entry to dst.
	AppendUint32s(dst []byte, key string, values []uint32) []byte                 // AppendUint32s appends an []uint32 entry to dst.
	AppendUint64s(dst []byte, key string, values []uint64) []byte                 // AppendUint64s appends an []uint64 entry to dst.
	AppendFloat32s(dst []byte, key string, values []float32) []byte               // AppendFloat32s appends a []float32 entry to dst.
	AppendFloat64s(dst []byte, key string, values []float64) []byte               // AppendFloat64s appends a []float64 entry to dst.
	AppendStrings(dst []byte, key string, values []string) []byte                 // AppendStrings appends a []string entry to dst.
	AppendTimes(dst []byte, key string, values []time.Time, format string) []byte // AppendTimes appends a []time.Time entry formatted with format to dst.
	AppendErrors(dst []byte, key string, values []error) []byte                   // AppendErrors appends an []error entry to dst.
	AppendStringers(dst []byte, key string, values []fmt.Stringer) []byte         // AppendStringers appends a []fmt.Stringer entry to dst.
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
	switch value {
	case '"', '\\':
		return append(dst, '\\', value)
	case '\b':
		return append(dst, '\\', 'b')
	case '\f':
		return append(dst, '\\', 'f')
	case '\n':
		return append(dst, '\\', 'n')
	case '\r':
		return append(dst, '\\', 'r')
	case '\t':
		return append(dst, '\\', 't')
	default:
		// ASCii < 16 needs to add \u000 to behind.
		if value < 16 {
			return strconv.AppendInt(append(dst, '\\', 'u', '0', '0', '0'), int64(value), 16)
		}

		// ASCii in [16, 32) needs to add \u00 to behind.
		if value < 32 {
			return strconv.AppendInt(append(dst, '\\', 'u', '0', '0'), int64(value), 16)
		}
		return append(dst, value)
	}
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
		// Encountered a byte that need escaping, so we appended bytes behinds it and appended it escaped.
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

	// There is no need for escaping, just appending like bytes.
	return append(dst, value...)
}

// Text returns an Text appender.
func Text() Appender {
	return globalTextAppender
}

// Json returns a Json appender.
func Json() Appender {
	return globalJsonAppender
}
