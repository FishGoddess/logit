// Copyright 2024 FishGoddess. All Rights Reserved.
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

package handler

import (
	"strconv"
	"unicode/utf8"
)

// needEscapedByte returns if value need to escape.
// The main character should be escaped is ascii less than \u0020.
func needEscapedByte(value byte) bool {
	return value < 32
}

// appendEscapedByte appends escaped value to dst.
// The main character should be escaped is ascii less than \u0020.
func appendEscapedByte(dst []byte, value byte) []byte {
	switch value {
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

// appendEscapedString appends escaped value to dst.
// The main character should be escaped is ascii less than \u0020.
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
