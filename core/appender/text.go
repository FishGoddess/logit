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
	"reflect"
	"strconv"
	"time"
)

const (
	textArrayBegin         = '['
	textArrayEnd           = ']'
	textArrayItemSeparator = '|'
	textItemSeparator      = '&'
	textKeyValueSeparator  = '='
	textNil                = "nil"
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
	dst = ta.appendKey(dst, key)
	return strconv.AppendBool(dst, value)
}

func (ta *textAppender) AppendByte(dst []byte, key string, value byte) []byte {
	dst = ta.appendKey(dst, key)
	return appendEscapedByte(dst, value)
}

func (ta *textAppender) AppendRune(dst []byte, key string, value rune) []byte {
	dst = ta.appendKey(dst, key)
	return appendEscapedRune(dst, value)
}

func (ta *textAppender) AppendInt(dst []byte, key string, value int) []byte {
	dst = ta.appendKey(dst, key)
	return strconv.AppendInt(dst, int64(value), 10)
}

func (ta *textAppender) AppendInt8(dst []byte, key string, value int8) []byte {
	dst = ta.appendKey(dst, key)
	return strconv.AppendInt(dst, int64(value), 10)
}

func (ta *textAppender) AppendInt16(dst []byte, key string, value int16) []byte {
	dst = ta.appendKey(dst, key)
	return strconv.AppendInt(dst, int64(value), 10)
}

func (ta *textAppender) AppendInt32(dst []byte, key string, value int32) []byte {
	dst = ta.appendKey(dst, key)
	return strconv.AppendInt(dst, int64(value), 10)
}

func (ta *textAppender) AppendInt64(dst []byte, key string, value int64) []byte {
	dst = ta.appendKey(dst, key)
	return strconv.AppendInt(dst, value, 10)
}

func (ta *textAppender) AppendUint(dst []byte, key string, value uint) []byte {
	dst = ta.appendKey(dst, key)
	return strconv.AppendUint(dst, uint64(value), 10)
}

func (ta *textAppender) AppendUint8(dst []byte, key string, value uint8) []byte {
	dst = ta.appendKey(dst, key)
	return strconv.AppendUint(dst, uint64(value), 10)
}

func (ta *textAppender) AppendUint16(dst []byte, key string, value uint16) []byte {
	dst = ta.appendKey(dst, key)
	return strconv.AppendUint(dst, uint64(value), 10)
}

func (ta *textAppender) AppendUint32(dst []byte, key string, value uint32) []byte {
	dst = ta.appendKey(dst, key)
	return strconv.AppendUint(dst, uint64(value), 10)
}

func (ta *textAppender) AppendUint64(dst []byte, key string, value uint64) []byte {
	dst = ta.appendKey(dst, key)
	return strconv.AppendUint(dst, value, 10)
}

func (ta *textAppender) AppendFloat32(dst []byte, key string, value float32) []byte {
	dst = ta.appendKey(dst, key)
	return strconv.AppendFloat(dst, float64(value), 'f', -1, 64)
}

func (ta *textAppender) AppendFloat64(dst []byte, key string, value float64) []byte {
	dst = ta.appendKey(dst, key)
	return strconv.AppendFloat(dst, value, 'f', -1, 64)
}

func (ta *textAppender) AppendString(dst []byte, key string, value string) []byte {
	dst = ta.appendKey(dst, key)
	return appendEscapedString(dst, value)
}

func (ta *textAppender) AppendTime(dst []byte, key string, value time.Time, format string) []byte {
	dst = ta.appendKey(dst, key)
	if format == UnixTime {
		return strconv.AppendInt(dst, value.Unix(), 10)
	}
	return value.AppendFormat(dst, format)
}

func (ta *textAppender) AppendError(dst []byte, key string, value error) []byte {
	if value == nil {
		return append(ta.appendKey(dst, key), textNil...)
	}
	return ta.AppendString(dst, key, value.Error())
}

func (ta *textAppender) AppendStringer(dst []byte, key string, value fmt.Stringer) []byte {

	val := reflect.ValueOf(value)
	if val.Kind() == reflect.Ptr && val.IsNil() {
		return append(dst, textNil...)
	}
	return ta.AppendString(dst, key, value.String())
}

func (ta *textAppender) appendArray(dst []byte, key string, length int, fn func(source []byte, index int) []byte) []byte {

	dst = ta.appendKey(dst, key)
	dst = append(dst, textArrayBegin)
	for i := 0; i < length; i++ {

		if dst[len(dst)-1] != textArrayBegin {
			dst = append(dst, textArrayItemSeparator)
		}
		dst = fn(dst, i)
	}
	dst = append(dst, textArrayEnd)
	return dst
}

func (ta *textAppender) AppendBytes(dst []byte, key string, values []byte) []byte {
	return ta.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		return appendEscapedByte(source, values[index])
	})
}

func (ta *textAppender) AppendRunes(dst []byte, key string, values []rune) []byte {
	return ta.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		return appendEscapedRune(source, values[index])
	})
}

func (ta *textAppender) AppendBools(dst []byte, key string, values []bool) []byte {
	return ta.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		return strconv.AppendBool(source, values[index])
	})
}

func (ta *textAppender) AppendInts(dst []byte, key string, values []int) []byte {
	return ta.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		return strconv.AppendInt(source, int64(values[index]), 10)
	})
}

func (ta *textAppender) AppendInt8s(dst []byte, key string, values []int8) []byte {
	return ta.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		return strconv.AppendInt(source, int64(values[index]), 10)
	})
}

func (ta *textAppender) AppendInt16s(dst []byte, key string, values []int16) []byte {
	return ta.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		return strconv.AppendInt(source, int64(values[index]), 10)
	})
}

func (ta *textAppender) AppendInt32s(dst []byte, key string, values []int32) []byte {
	return ta.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		return strconv.AppendInt(source, int64(values[index]), 10)
	})
}

func (ta *textAppender) AppendInt64s(dst []byte, key string, values []int64) []byte {
	return ta.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		return strconv.AppendInt(source, values[index], 10)
	})
}

func (ta *textAppender) AppendUints(dst []byte, key string, values []uint) []byte {
	return ta.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		return strconv.AppendUint(source, uint64(values[index]), 10)
	})
}

func (ta *textAppender) AppendUint8s(dst []byte, key string, values []uint8) []byte {
	return ta.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		return strconv.AppendUint(source, uint64(values[index]), 10)
	})
}

func (ta *textAppender) AppendUint16s(dst []byte, key string, values []uint16) []byte {
	return ta.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		return strconv.AppendUint(source, uint64(values[index]), 10)
	})
}

func (ta *textAppender) AppendUint32s(dst []byte, key string, values []uint32) []byte {
	return ta.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		return strconv.AppendUint(source, uint64(values[index]), 10)
	})
}

func (ta *textAppender) AppendUint64s(dst []byte, key string, values []uint64) []byte {
	return ta.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		return strconv.AppendUint(source, values[index], 10)
	})
}

func (ta *textAppender) AppendFloat32s(dst []byte, key string, values []float32) []byte {
	return ta.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		return strconv.AppendFloat(source, float64(values[index]), 'f', -1, 64)
	})
}

func (ta *textAppender) AppendFloat64s(dst []byte, key string, values []float64) []byte {
	return ta.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		return strconv.AppendFloat(source, values[index], 'f', -1, 64)
	})
}

func (ta *textAppender) AppendStrings(dst []byte, key string, values []string) []byte {
	return ta.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		return appendEscapedString(source, values[index])
	})
}

func (ta *textAppender) AppendTimes(dst []byte, key string, values []time.Time, format string) []byte {
	return ta.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		if format == UnixTime {
			return strconv.AppendInt(source, values[index].Unix(), 10)
		}
		return values[index].AppendFormat(source, format)
	})
}

func (ta *textAppender) AppendErrors(dst []byte, key string, values []error) []byte {
	return ta.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		if values[index] == nil {
			return append(source, textNil...)
		}
		return ta.AppendString(source, key, values[index].Error())
	})
}

func (ta *textAppender) AppendStringers(dst []byte, key string, values []fmt.Stringer) []byte {
	return ta.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		val := reflect.ValueOf(values[index])
		if val.Kind() == reflect.Ptr && val.IsNil() {
			return append(source, textNil...)
		}
		return ta.AppendString(source, key, values[index].String())
	})
}
