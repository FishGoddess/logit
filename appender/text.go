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
// Created at 2021/06/30 00:50:12

package appender

import (
	"fmt"
	"time"
)

const (
	textArrayBegin        = '['
	textArrayEnd          = ']'
	textItemSeparator     = '\t'
	textKeyValueSeparator = '='
)

type textAppender struct {
}

func (ta *textAppender) Begin(dst []byte) []byte {
	return dst
}

func (ta *textAppender) End(dst []byte) []byte {
	return append(dst, lineBreak)
}

func (ta *textAppender) appendKey(dst []byte, key string) []byte {

	if len(dst) > 0 {
		dst = append(dst, textItemSeparator)
	}

	dst = appendEscapedString(dst, key)
	return append(dst, textKeyValueSeparator)
}

func (ta *textAppender) AppendAny(dst []byte, key string, value interface{}) []byte {
	return append(ta.appendKey(dst, key), fmt.Sprintf(`%+v`, value)...)
}

func (ta *textAppender) AppendBool(dst []byte, key string, value bool) []byte {
	return dst
}

func (ta *textAppender) AppendByte(dst []byte, key string, value byte) []byte {
	return dst
}

func (ta *textAppender) AppendRune(dst []byte, key string, value rune) []byte {
	return dst
}

func (ta *textAppender) AppendInt(dst []byte, key string, value int) []byte {
	return dst
}

func (ta *textAppender) AppendInt8(dst []byte, key string, value int8) []byte {
	return dst
}

func (ta *textAppender) AppendInt16(dst []byte, key string, value int16) []byte {
	return dst
}

func (ta *textAppender) AppendInt32(dst []byte, key string, value int32) []byte {
	return dst
}

func (ta *textAppender) AppendInt64(dst []byte, key string, value int64) []byte {
	return dst
}

func (ta *textAppender) AppendUint(dst []byte, key string, value uint) []byte {
	return dst
}

func (ta *textAppender) AppendUint8(dst []byte, key string, value uint8) []byte {
	return dst
}

func (ta *textAppender) AppendUint16(dst []byte, key string, value uint16) []byte {
	return dst
}

func (ta *textAppender) AppendUint32(dst []byte, key string, value uint32) []byte {
	return dst
}

func (ta *textAppender) AppendUint64(dst []byte, key string, value uint64) []byte {
	return dst
}

func (ta *textAppender) AppendFloat32(dst []byte, key string, value float32) []byte {
	return dst
}

func (ta *textAppender) AppendFloat64(dst []byte, key string, value float64) []byte {
	return dst
}

func (ta *textAppender) AppendString(dst []byte, key string, value string) []byte {
	return dst
}

func (ta *textAppender) AppendTime(dst []byte, key string, value time.Time, format string) []byte {
	return dst
}

func (ta *textAppender) AppendError(dst []byte, key string, value error) []byte {
	return dst
}

func (ta *textAppender) AppendStringer(dst []byte, key string, value fmt.Stringer) []byte {
	return dst
}

func (ta *textAppender) AppendBools(dst []byte, key string, values []bool) []byte {
	return dst
}

func (ta *textAppender) AppendBytes(dst []byte, key string, values []byte) []byte {
	return dst
}

func (ta *textAppender) AppendRunes(dst []byte, key string, values []rune) []byte {
	return dst
}

func (ta *textAppender) AppendInts(dst []byte, key string, values []int) []byte {
	return dst
}

func (ta *textAppender) AppendInt8s(dst []byte, key string, values []int8) []byte {
	return dst
}

func (ta *textAppender) AppendInt16s(dst []byte, key string, values []int16) []byte {
	return dst
}

func (ta *textAppender) AppendInt32s(dst []byte, key string, values []int32) []byte {
	return dst
}

func (ta *textAppender) AppendInt64s(dst []byte, key string, values []int64) []byte {
	return dst
}

func (ta *textAppender) AppendUints(dst []byte, key string, values []uint) []byte {
	return dst
}

func (ta *textAppender) AppendUint8s(dst []byte, key string, values []uint8) []byte {
	return dst
}

func (ta *textAppender) AppendUint16s(dst []byte, key string, values []uint16) []byte {
	return dst
}

func (ta *textAppender) AppendUint32s(dst []byte, key string, values []uint32) []byte {
	return dst
}

func (ta *textAppender) AppendUint64s(dst []byte, key string, values []uint64) []byte {
	return dst
}

func (ta *textAppender) AppendFloat32s(dst []byte, key string, values []float32) []byte {
	return dst
}

func (ta *textAppender) AppendFloat64s(dst []byte, key string, values []float64) []byte {
	return dst
}

func (ta *textAppender) AppendStrings(dst []byte, key string, values []string) []byte {
	return dst
}

func (ta *textAppender) AppendTimes(dst []byte, key string, values []time.Time, format string) []byte {
	return dst
}

func (ta *textAppender) AppendErrors(dst []byte, key string, values []error) []byte {
	return dst
}

func (ta *textAppender) AppendStringers(dst []byte, key string, values []fmt.Stringer) []byte {
	return dst
}
