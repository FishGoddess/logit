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
	"math"
	"reflect"
	"strconv"
	"time"

	"github.com/go-logit/logit/support/global"
)

const (
	jsonBegin             = '{'    // jsonBegin is the begin character of Json.
	jsonEnd               = '}'    // jsonEnd is the end character of Json.
	jsonArrayBegin        = '['    // jsonArrayBegin is the begin character of Json array.
	jsonArrayEnd          = ']'    // jsonArrayEnd is the end character of Json array.
	jsonItemSeparator     = ','    // jsonItemSeparator is the character between items in Json.
	jsonKeyValueSeparator = ':'    // jsonKeyValueSeparator is the character between key and value of Json.
	jsonStringQuotation   = '"'    // jsonStringQuotation is the quotation character of Json string.
	jsonNull              = "null" // jsonNull is the null characters of Json.
)

// jsonAppender is a Json appender.
type jsonAppender struct{}

// Begin appends begin character to dst.
func (ja *jsonAppender) Begin(dst []byte) []byte {
	return append(dst, jsonBegin)
}

// End appends end character to dst.
func (ja *jsonAppender) End(dst []byte) []byte {
	return append(dst, jsonEnd, lineBreak)
}

// appendKey appends key to dst.
func (ja *jsonAppender) appendKey(dst []byte, key string) []byte {
	if dst[len(dst)-1] != jsonBegin {
		dst = append(dst, jsonItemSeparator)
	}

	dst = append(dst, jsonStringQuotation)
	dst = appendEscapedString(dst, key)
	dst = append(dst, jsonStringQuotation)
	return append(dst, jsonKeyValueSeparator)
}

// AppendAny appends any entries to dst.
func (ja *jsonAppender) AppendAny(dst []byte, key string, value interface{}) []byte {
	dst = ja.appendKey(dst, key)

	valueBytes, err := global.MarshalToJson(value)
	if err != nil {
		dst = append(dst, jsonStringQuotation)
		dst = appendEscapedString(dst, err.Error())
		return append(dst, jsonStringQuotation)
	}

	return append(dst, valueBytes...)
}

// AppendJson appends any entries as Json to dst.
func (ja *jsonAppender) AppendJson(dst []byte, key string, value interface{}) []byte {
	valueBytes, err := global.MarshalToJson(value)
	if err != nil {
		return ja.AppendString(dst, key, err.Error())
	}
	return append(ja.appendKey(dst, key), valueBytes...)
}

// AppendBool appends a bool entry to dst.
func (ja *jsonAppender) AppendBool(dst []byte, key string, value bool) []byte {
	dst = ja.appendKey(dst, key)
	return strconv.AppendBool(dst, value)
}

// AppendByte appends a byte entry to dst.
func (ja *jsonAppender) AppendByte(dst []byte, key string, value byte) []byte {
	dst = ja.appendKey(dst, key)
	dst = append(dst, jsonStringQuotation)
	dst = appendEscapedByte(dst, value)
	dst = append(dst, jsonStringQuotation)
	return dst
}

// AppendRune appends a rune entry to dst.
func (ja *jsonAppender) AppendRune(dst []byte, key string, value rune) []byte {
	dst = ja.appendKey(dst, key)
	dst = append(dst, jsonStringQuotation)
	dst = appendEscapedRune(dst, value)
	dst = append(dst, jsonStringQuotation)
	return dst
}

// AppendInt appends an int entry to dst.
func (ja *jsonAppender) AppendInt(dst []byte, key string, value int) []byte {
	dst = ja.appendKey(dst, key)
	return strconv.AppendInt(dst, int64(value), 10)
}

// AppendInt8 appends an int8 entry to dst.
func (ja *jsonAppender) AppendInt8(dst []byte, key string, value int8) []byte {
	dst = ja.appendKey(dst, key)
	return strconv.AppendInt(dst, int64(value), 10)
}

// AppendInt16 appends an int16 entry to dst.
func (ja *jsonAppender) AppendInt16(dst []byte, key string, value int16) []byte {
	dst = ja.appendKey(dst, key)
	return strconv.AppendInt(dst, int64(value), 10)
}

// AppendInt32 appends an int32 entry to dst.
func (ja *jsonAppender) AppendInt32(dst []byte, key string, value int32) []byte {
	dst = ja.appendKey(dst, key)
	return strconv.AppendInt(dst, int64(value), 10)
}

// AppendInt64 appends an int64 entry to dst.
func (ja *jsonAppender) AppendInt64(dst []byte, key string, value int64) []byte {
	dst = ja.appendKey(dst, key)
	return strconv.AppendInt(dst, value, 10)
}

// AppendUint appends an uint entry to dst.
func (ja *jsonAppender) AppendUint(dst []byte, key string, value uint) []byte {
	dst = ja.appendKey(dst, key)
	return strconv.AppendUint(dst, uint64(value), 10)
}

// AppendUint8 appends an uin8 entry to dst.
func (ja *jsonAppender) AppendUint8(dst []byte, key string, value uint8) []byte {
	dst = ja.appendKey(dst, key)
	return strconv.AppendUint(dst, uint64(value), 10)
}

// AppendUint16 appends an uint16 entry to dst.
func (ja *jsonAppender) AppendUint16(dst []byte, key string, value uint16) []byte {
	dst = ja.appendKey(dst, key)
	return strconv.AppendUint(dst, uint64(value), 10)
}

// AppendUint32 appends an uint32 entry to dst.
func (ja *jsonAppender) AppendUint32(dst []byte, key string, value uint32) []byte {
	dst = ja.appendKey(dst, key)
	return strconv.AppendUint(dst, uint64(value), 10)
}

// AppendUint64 appends an uint64 entry to dst.
func (ja *jsonAppender) AppendUint64(dst []byte, key string, value uint64) []byte {
	dst = ja.appendKey(dst, key)
	return strconv.AppendUint(dst, value, 10)
}

// AppendFloat32 appends a float32 entry to dst.
func (ja *jsonAppender) AppendFloat32(dst []byte, key string, value float32) []byte {
	// Standard json doesn't support NaN and Inf, so append a null.
	dst = ja.appendKey(dst, key)

	value64 := float64(value)
	if math.IsNaN(value64) || math.IsInf(value64, 0) {
		return append(dst, jsonNull...)
	}

	return strconv.AppendFloat(dst, value64, 'f', -1, 64)
}

// AppendFloat64 appends a float64 entry to dst.
func (ja *jsonAppender) AppendFloat64(dst []byte, key string, value float64) []byte {
	// Standard json doesn't support NaN and Inf, so append a null.
	dst = ja.appendKey(dst, key)

	if math.IsNaN(value) || math.IsInf(value, 0) {
		return append(dst, jsonNull...)
	}

	return strconv.AppendFloat(dst, value, 'f', -1, 64)
}

// AppendString appends a string entry to dst.
func (ja *jsonAppender) AppendString(dst []byte, key string, value string) []byte {
	dst = ja.appendKey(dst, key)
	dst = append(dst, jsonStringQuotation)
	dst = appendEscapedString(dst, value)
	dst = append(dst, jsonStringQuotation)
	return dst
}

// AppendTime appends a time.Time entry formatted with format to dst.
func (ja *jsonAppender) AppendTime(dst []byte, key string, value time.Time, format string) []byte {
	dst = ja.appendKey(dst, key)
	if format == UnixTimeFormat {
		return strconv.AppendInt(dst, value.Unix(), 10)
	}

	dst = append(dst, jsonStringQuotation)
	dst = value.AppendFormat(dst, format)
	dst = append(dst, jsonStringQuotation)
	return dst
}

// AppendError appends an error entry to dst.
func (ja *jsonAppender) AppendError(dst []byte, key string, value error) []byte {
	if value == nil {
		return append(ja.appendKey(dst, key), jsonNull...)
	}
	return ja.AppendString(dst, key, value.Error())
}

// AppendStringer appends an fmt.Stringer entry to dst.
func (ja *jsonAppender) AppendStringer(dst []byte, key string, value fmt.Stringer) []byte {
	val := reflect.ValueOf(value)
	if val.Kind() == reflect.Ptr && val.IsNil() {
		return append(dst, jsonNull...)
	}
	return ja.AppendString(dst, key, value.String())
}

// appendArray appends array to dst.
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

// AppendBools appends a []bool entry to dst.
func (ja *jsonAppender) AppendBools(dst []byte, key string, values []bool) []byte {
	return ja.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		return strconv.AppendBool(source, values[index])
	})
}

// AppendBytes appends a []byte entry to dst.
func (ja *jsonAppender) AppendBytes(dst []byte, key string, values []byte) []byte {
	return ja.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		source = append(source, jsonStringQuotation)
		source = appendEscapedByte(source, values[index])
		source = append(source, jsonStringQuotation)
		return source
	})
}

// AppendRunes appends a []rune entry to dst.
func (ja *jsonAppender) AppendRunes(dst []byte, key string, values []rune) []byte {
	return ja.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		source = append(source, jsonStringQuotation)
		source = appendEscapedRune(source, values[index])
		source = append(source, jsonStringQuotation)
		return source
	})
}

// AppendInts appends an []int entry to dst.
func (ja *jsonAppender) AppendInts(dst []byte, key string, values []int) []byte {
	return ja.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		return strconv.AppendInt(source, int64(values[index]), 10)
	})
}

// AppendInt8s appends an []int8 entry to dst.
func (ja *jsonAppender) AppendInt8s(dst []byte, key string, values []int8) []byte {
	return ja.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		return strconv.AppendInt(source, int64(values[index]), 10)
	})
}

// AppendInt16s appends an []int16 entry to dst.
func (ja *jsonAppender) AppendInt16s(dst []byte, key string, values []int16) []byte {
	return ja.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		return strconv.AppendInt(source, int64(values[index]), 10)
	})
}

// AppendInt32s appends an []int32 entry to dst.
func (ja *jsonAppender) AppendInt32s(dst []byte, key string, values []int32) []byte {
	return ja.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		return strconv.AppendInt(source, int64(values[index]), 10)
	})
}

// AppendInt64s appends an []int64 entry to dst.
func (ja *jsonAppender) AppendInt64s(dst []byte, key string, values []int64) []byte {
	return ja.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		return strconv.AppendInt(source, values[index], 10)
	})
}

// AppendUints appends an []uint entry to dst.
func (ja *jsonAppender) AppendUints(dst []byte, key string, values []uint) []byte {
	return ja.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		return strconv.AppendUint(source, uint64(values[index]), 10)
	})
}

// AppendUint8s appends an []uint8 entry to dst.
func (ja *jsonAppender) AppendUint8s(dst []byte, key string, values []uint8) []byte {
	return ja.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		return strconv.AppendUint(source, uint64(values[index]), 10)
	})
}

// AppendUint16s appends an []uint16 entry to dst.
func (ja *jsonAppender) AppendUint16s(dst []byte, key string, values []uint16) []byte {
	return ja.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		return strconv.AppendUint(source, uint64(values[index]), 10)
	})
}

// AppendUint32s appends an []uint32 entry to dst.
func (ja *jsonAppender) AppendUint32s(dst []byte, key string, values []uint32) []byte {
	return ja.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		return strconv.AppendUint(source, uint64(values[index]), 10)
	})
}

// AppendUint64s appends an []uint64 entry to dst.
func (ja *jsonAppender) AppendUint64s(dst []byte, key string, values []uint64) []byte {
	return ja.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		return strconv.AppendUint(source, values[index], 10)
	})
}

// AppendFloat32s appends a []float32 entry to dst.
func (ja *jsonAppender) AppendFloat32s(dst []byte, key string, values []float32) []byte {
	return ja.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		value64 := float64(values[index])
		if math.IsNaN(value64) || math.IsInf(value64, 0) {
			return append(dst, jsonNull...)
		}
		return strconv.AppendFloat(source, value64, 'f', -1, 64)
	})
}

// AppendFloat64s appends a []float64 entry to dst.
func (ja *jsonAppender) AppendFloat64s(dst []byte, key string, values []float64) []byte {
	return ja.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		if math.IsNaN(values[index]) || math.IsInf(values[index], 0) {
			return append(dst, jsonNull...)
		}
		return strconv.AppendFloat(source, values[index], 'f', -1, 64)
	})
}

// AppendStrings appends a []string entry to dst.
func (ja *jsonAppender) AppendStrings(dst []byte, key string, values []string) []byte {
	return ja.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		source = append(source, jsonStringQuotation)
		source = appendEscapedString(source, values[index])
		source = append(source, jsonStringQuotation)
		return source
	})
}

// AppendTimes appends a []time.Time entry formatted with format to dst.
func (ja *jsonAppender) AppendTimes(dst []byte, key string, values []time.Time, format string) []byte {
	return ja.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		if format == UnixTimeFormat {
			return strconv.AppendInt(source, values[index].Unix(), 10)
		}

		source = append(source, jsonStringQuotation)
		source = values[index].AppendFormat(source, format)
		source = append(source, jsonStringQuotation)
		return source
	})
}

// AppendErrors appends an []error entry to dst.
func (ja *jsonAppender) AppendErrors(dst []byte, key string, values []error) []byte {
	return ja.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		if values[index] == nil {
			return append(source, jsonNull...)
		}

		source = append(source, jsonStringQuotation)
		source = appendEscapedString(source, values[index].Error())
		source = append(source, jsonStringQuotation)
		return source
	})
}

// AppendStringers appends a []fmt.Stringer entry to dst.
func (ja *jsonAppender) AppendStringers(dst []byte, key string, values []fmt.Stringer) []byte {
	return ja.appendArray(dst, key, len(values), func(source []byte, index int) []byte {
		val := reflect.ValueOf(values[index])
		if val.Kind() == reflect.Ptr && val.IsNil() {
			return append(source, jsonNull...)
		}

		source = append(source, jsonStringQuotation)
		source = appendEscapedString(source, values[index].String())
		source = append(source, jsonStringQuotation)
		return source
	})
}
