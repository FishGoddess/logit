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
// Created at 2021/06/27 16:37:21

package appender

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"time"
)

const (
	jsonBegin             = '{'
	jsonEnd               = '}'
	jsonArrayBegin        = '['
	jsonArrayEnd          = ']'
	jsonItemSeparator     = ','
	jsonKeyValueSeparator = ':'
	jsonStringQuotation   = '"'
)

type jsonAppender struct {
}

func (ja *jsonAppender) Begin(dst []byte) []byte {
	return append(dst, jsonBegin)
}

func (ja *jsonAppender) End(dst []byte) []byte {
	return append(dst, jsonEnd, lineBreak)
}

func (ja *jsonAppender) appendKey(dst []byte, key string) []byte {

	if dst[len(dst)-1] != jsonBegin {
		dst = append(dst, jsonItemSeparator)
	}

	dst = append(dst, jsonStringQuotation)
	dst = appendEscapedString(dst, key)
	dst = append(dst, jsonStringQuotation)
	return append(dst, jsonKeyValueSeparator)
}

func (ja *jsonAppender) AppendAny(dst []byte, key string, value interface{}) []byte {

	dst = ja.appendKey(dst, key)

	valueBytes, err := json.Marshal(value)
	if err != nil {
		return append(dst, fmt.Sprintf(`"%+v"`, value)...)
	}
	return append(dst, valueBytes...)
}

func (ja *jsonAppender) AppendBool(dst []byte, key string, value bool) []byte {
	dst = ja.appendKey(dst, key)
	return strconv.AppendBool(dst, value)
}

func (ja *jsonAppender) AppendByte(dst []byte, key string, value byte) []byte {
	dst = ja.appendKey(dst, key)
	dst = append(dst, jsonStringQuotation)
	dst = appendEscapedByte(dst, value)
	dst = append(dst, jsonStringQuotation)
	return dst
}

func (ja *jsonAppender) AppendRune(dst []byte, key string, value rune) []byte {
	dst = ja.appendKey(dst, key)
	dst = append(dst, jsonStringQuotation)
	dst = appendEscapedRune(dst, value)
	dst = append(dst, jsonStringQuotation)
	return dst
}

func (ja *jsonAppender) AppendInt(dst []byte, key string, value int) []byte {
	dst = ja.appendKey(dst, key)
	return strconv.AppendInt(dst, int64(value), 10)
}

func (ja *jsonAppender) AppendInt8(dst []byte, key string, value int8) []byte {
	dst = ja.appendKey(dst, key)
	return strconv.AppendInt(dst, int64(value), 10)
}

func (ja *jsonAppender) AppendInt16(dst []byte, key string, value int16) []byte {
	dst = ja.appendKey(dst, key)
	return strconv.AppendInt(dst, int64(value), 10)
}

func (ja *jsonAppender) AppendInt32(dst []byte, key string, value int32) []byte {
	dst = ja.appendKey(dst, key)
	return strconv.AppendInt(dst, int64(value), 10)
}

func (ja *jsonAppender) AppendInt64(dst []byte, key string, value int64) []byte {
	dst = ja.appendKey(dst, key)
	return strconv.AppendInt(dst, value, 10)
}

func (ja *jsonAppender) AppendUint(dst []byte, key string, value uint) []byte {
	dst = ja.appendKey(dst, key)
	return strconv.AppendUint(dst, uint64(value), 10)
}

func (ja *jsonAppender) AppendUint8(dst []byte, key string, value uint8) []byte {
	dst = ja.appendKey(dst, key)
	return strconv.AppendUint(dst, uint64(value), 10)
}

func (ja *jsonAppender) AppendUint16(dst []byte, key string, value uint16) []byte {
	dst = ja.appendKey(dst, key)
	return strconv.AppendUint(dst, uint64(value), 10)
}

func (ja *jsonAppender) AppendUint32(dst []byte, key string, value uint32) []byte {
	dst = ja.appendKey(dst, key)
	return strconv.AppendUint(dst, uint64(value), 10)
}

func (ja *jsonAppender) AppendUint64(dst []byte, key string, value uint64) []byte {
	dst = ja.appendKey(dst, key)
	return strconv.AppendUint(dst, value, 10)
}

func (ja *jsonAppender) AppendFloat32(dst []byte, key string, value float32) []byte {

	// Standard json doesn't support NaN and Inf, so coverts them to string
	// NaN => "NaN"
	// +inf => "+inf"
	// -inf => "-inf"
	dst = ja.appendKey(dst, key)

	value64 := float64(value)
	if math.IsNaN(value64) {
		return append(dst, nan...)
	}

	if math.IsInf(value64, 1) {
		return append(dst, pInf...)
	}

	if math.IsInf(value64, -1) {
		return append(dst, nInf...)
	}
	return strconv.AppendFloat(dst, value64, 'f', -1, 64)
}

func (ja *jsonAppender) AppendFloat64(dst []byte, key string, value float64) []byte {

	// Standard json doesn't support NaN and Inf, so coverts them to string
	// NaN => "NaN"
	// +inf => "+inf"
	// -inf => "-inf"
	dst = ja.appendKey(dst, key)

	if math.IsNaN(value) {
		return append(dst, nan...)
	}

	if math.IsInf(value, 1) {
		return append(dst, pInf...)
	}

	if math.IsInf(value, -1) {
		return append(dst, nInf...)
	}
	return strconv.AppendFloat(dst, value, 'f', -1, 64)
}

func (ja *jsonAppender) AppendString(dst []byte, key string, value string) []byte {
	dst = ja.appendKey(dst, key)
	dst = append(dst, jsonStringQuotation)
	dst = appendEscapedString(dst, value)
	dst = append(dst, jsonStringQuotation)
	return dst
}

func (ja *jsonAppender) AppendTime(dst []byte, key string, value time.Time, format string) []byte {

	dst = ja.appendKey(dst, key)
	if format == UnixTime {
		return strconv.AppendInt(dst, value.Unix(), 10)
	}
	dst = append(dst, jsonStringQuotation)
	dst = value.AppendFormat(dst, format)
	dst = append(dst, jsonStringQuotation)
	return dst
}

func (ja *jsonAppender) appendArray(dst []byte, key string, length int, fn func(source []byte, index int) []byte) []byte {

	dst = ja.appendKey(dst, key)
	dst = append(dst, jsonArrayBegin)
	for i := 0; i < length; i++ {

		if dst[len(dst)-1] != jsonArrayBegin {
			dst = append(dst, jsonItemSeparator)
		}
		dst = fn(dst, i)
	}
	dst = append(dst, jsonArrayEnd)
	return dst
}

func (ja *jsonAppender) AppendBytes(dst []byte, key string, values []byte) []byte {

	return ja.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		source = append(source, jsonStringQuotation)
		source = appendEscapedByte(source, values[index])
		source = append(source, jsonStringQuotation)
		return source
	})
}

func (ja *jsonAppender) AppendRunes(dst []byte, key string, values []rune) []byte {

	return ja.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		source = append(source, jsonStringQuotation)
		source = appendEscapedRune(source, values[index])
		source = append(source, jsonStringQuotation)
		return source
	})
}

func (ja *jsonAppender) AppendBools(dst []byte, key string, values []bool) []byte {
	return ja.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		return strconv.AppendBool(source, values[index])
	})
}

func (ja *jsonAppender) AppendInts(dst []byte, key string, values []int) []byte {
	return ja.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		return strconv.AppendInt(source, int64(values[index]), 10)
	})
}

func (ja *jsonAppender) AppendInt8s(dst []byte, key string, values []int8) []byte {
	return ja.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		return strconv.AppendInt(source, int64(values[index]), 10)
	})
}

func (ja *jsonAppender) AppendInt16s(dst []byte, key string, values []int16) []byte {
	return ja.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		return strconv.AppendInt(source, int64(values[index]), 10)
	})
}

func (ja *jsonAppender) AppendInt32s(dst []byte, key string, values []int32) []byte {
	return ja.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		return strconv.AppendInt(source, int64(values[index]), 10)
	})
}

func (ja *jsonAppender) AppendInt64s(dst []byte, key string, values []int64) []byte {
	return ja.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		return strconv.AppendInt(source, values[index], 10)
	})
}

func (ja *jsonAppender) AppendUints(dst []byte, key string, values []uint) []byte {
	return ja.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		return strconv.AppendUint(source, uint64(values[index]), 10)
	})
}

func (ja *jsonAppender) AppendUint8s(dst []byte, key string, values []uint8) []byte {
	return ja.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		return strconv.AppendUint(source, uint64(values[index]), 10)
	})
}

func (ja *jsonAppender) AppendUint16s(dst []byte, key string, values []uint16) []byte {
	return ja.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		return strconv.AppendUint(source, uint64(values[index]), 10)
	})
}

func (ja *jsonAppender) AppendUint32s(dst []byte, key string, values []uint32) []byte {
	return ja.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		return strconv.AppendUint(source, uint64(values[index]), 10)
	})
}

func (ja *jsonAppender) AppendUint64s(dst []byte, key string, values []uint64) []byte {
	return ja.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		return strconv.AppendUint(source, values[index], 10)
	})
}

func (ja *jsonAppender) AppendFloat32s(dst []byte, key string, values []float32) []byte {
	return ja.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		return strconv.AppendFloat(source, float64(values[index]), 'f', -1, 64)
	})
}

func (ja *jsonAppender) AppendFloat64s(dst []byte, key string, values []float64) []byte {
	return ja.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		return strconv.AppendFloat(source, values[index], 'f', -1, 64)
	})
}

func (ja *jsonAppender) AppendStrings(dst []byte, key string, values []string) []byte {
	return ja.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		source = append(source, jsonStringQuotation)
		source = appendEscapedString(source, values[index])
		source = append(source, jsonStringQuotation)
		return source
	})
}

func (ja *jsonAppender) AppendTimes(dst []byte, key string, values []time.Time, format string) []byte {
	return ja.appendArray(dst, key, len(values), func(source []byte, index int) []byte {

		if format == UnixTime {
			return strconv.AppendInt(source, values[index].Unix(), 10)
		}
		source = append(source, jsonStringQuotation)
		source = values[index].AppendFormat(source, format)
		source = append(source, jsonStringQuotation)
		return source
	})
}
