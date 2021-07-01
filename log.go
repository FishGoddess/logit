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
// Created at 2021/06/27 23:54:11

package logit

import (
	"fmt"
	"time"
)

type Log struct {
	logger *Logger
	data   []byte
}

func newLog(logger *Logger) *Log {
	return &Log{
		logger: logger,
		data:   make([]byte, 0, 512),
	}
}

func (l *Log) initialize() {
	l.data = l.data[:0]
	l.data = l.logger.appender.Begin(l.data)
}

func (l *Log) Record() {

	if l == nil {
		return
	}

	defer l.logger.releaseLog(l)
	l.logger.writer.Write(l.logger.appender.End(l.data))
}

func (l *Log) Msg(msg string, params ...interface{}) {

	if l == nil {
		return
	}

	if len(params) > 0 {
		msg = fmt.Sprintf(msg, params...)
	}
	l.data = l.logger.appender.AppendString(l.data, "log.msg", msg)
	l.Record()
}

func (l *Log) Any(key string, value interface{}) *Log {

	if l == nil {
		return nil
	}
	l.data = l.logger.appender.AppendAny(l.data, key, value)
	return l
}

func (l *Log) Bool(key string, value bool) *Log {

	if l == nil {
		return nil
	}
	l.data = l.logger.appender.AppendBool(l.data, key, value)
	return l
}

func (l *Log) Byte(key string, value byte) *Log {

	if l == nil {
		return nil
	}
	l.data = l.logger.appender.AppendByte(l.data, key, value)
	return l
}

func (l *Log) Rune(key string, value rune) *Log {

	if l == nil {
		return nil
	}
	l.data = l.logger.appender.AppendRune(l.data, key, value)
	return l
}

func (l *Log) Int(key string, value int) *Log {

	if l == nil {
		return nil
	}
	l.data = l.logger.appender.AppendInt(l.data, key, value)
	return l
}

func (l *Log) Int8(key string, value int8) *Log {

	if l == nil {
		return nil
	}
	l.data = l.logger.appender.AppendInt8(l.data, key, value)
	return l
}

func (l *Log) Int16(key string, value int16) *Log {

	if l == nil {
		return nil
	}
	l.data = l.logger.appender.AppendInt16(l.data, key, value)
	return l
}

func (l *Log) Int32(key string, value int32) *Log {

	if l == nil {
		return nil
	}
	l.data = l.logger.appender.AppendInt32(l.data, key, value)
	return l
}

func (l *Log) Int64(key string, value int64) *Log {

	if l == nil {
		return nil
	}
	l.data = l.logger.appender.AppendInt64(l.data, key, value)
	return l
}

func (l *Log) Uint(key string, value uint) *Log {

	if l == nil {
		return nil
	}
	l.data = l.logger.appender.AppendUint(l.data, key, value)
	return l
}

func (l *Log) Uint8(key string, value uint8) *Log {

	if l == nil {
		return nil
	}
	l.data = l.logger.appender.AppendUint8(l.data, key, value)
	return l
}

func (l *Log) Uint16(key string, value uint16) *Log {

	if l == nil {
		return nil
	}
	l.data = l.logger.appender.AppendUint16(l.data, key, value)
	return l
}

func (l *Log) Uint32(key string, value uint32) *Log {

	if l == nil {
		return nil
	}
	l.data = l.logger.appender.AppendUint32(l.data, key, value)
	return l
}

func (l *Log) Uint64(key string, value uint64) *Log {

	if l == nil {
		return nil
	}
	l.data = l.logger.appender.AppendUint64(l.data, key, value)
	return l
}

func (l *Log) Float32(key string, value float32) *Log {

	if l == nil {
		return nil
	}
	l.data = l.logger.appender.AppendFloat32(l.data, key, value)
	return l
}

func (l *Log) Float64(key string, value float64) *Log {

	if l == nil {
		return nil
	}
	l.data = l.logger.appender.AppendFloat64(l.data, key, value)
	return l
}

func (l *Log) String(key string, value string) *Log {

	if l == nil {
		return nil
	}
	l.data = l.logger.appender.AppendString(l.data, key, value)
	return l
}

func (l *Log) Time(key string, value time.Time, format string) *Log {

	if l == nil {
		return nil
	}
	l.data = l.logger.appender.AppendTime(l.data, key, value, format)
	return l
}

func (l *Log) Error(key string, value error) *Log {

	if l == nil {
		return nil
	}
	l.data = l.logger.appender.AppendError(l.data, key, value)
	return l
}

func (l *Log) Stringer(key string, value fmt.Stringer) *Log {

	if l == nil {
		return nil
	}
	l.data = l.logger.appender.AppendStringer(l.data, key, value)
	return l
}

func (l *Log) Bools(key string, value []bool) *Log {

	if l == nil {
		return nil
	}
	l.data = l.logger.appender.AppendBools(l.data, key, value)
	return l
}

func (l *Log) Bytes(key string, value []byte) *Log {

	if l == nil {
		return nil
	}
	l.data = l.logger.appender.AppendBytes(l.data, key, value)
	return l
}

func (l *Log) Runes(key string, value []rune) *Log {

	if l == nil {
		return nil
	}
	l.data = l.logger.appender.AppendRunes(l.data, key, value)
	return l
}

func (l *Log) Ints(key string, value []int) *Log {

	if l == nil {
		return nil
	}
	l.data = l.logger.appender.AppendInts(l.data, key, value)
	return l
}

func (l *Log) Int8s(key string, value []int8) *Log {

	if l == nil {
		return nil
	}
	l.data = l.logger.appender.AppendInt8s(l.data, key, value)
	return l
}

func (l *Log) Int16s(key string, value []int16) *Log {

	if l == nil {
		return nil
	}
	l.data = l.logger.appender.AppendInt16s(l.data, key, value)
	return l
}

func (l *Log) Int32s(key string, value []int32) *Log {

	if l == nil {
		return nil
	}
	l.data = l.logger.appender.AppendInt32s(l.data, key, value)
	return l
}

func (l *Log) Int64s(key string, value []int64) *Log {

	if l == nil {
		return nil
	}
	l.data = l.logger.appender.AppendInt64s(l.data, key, value)
	return l
}

func (l *Log) Uints(key string, value []uint) *Log {

	if l == nil {
		return nil
	}
	l.data = l.logger.appender.AppendUints(l.data, key, value)
	return l
}

func (l *Log) Uint8s(key string, value []uint8) *Log {

	if l == nil {
		return nil
	}
	l.data = l.logger.appender.AppendUint8s(l.data, key, value)
	return l
}

func (l *Log) Uint16s(key string, value []uint16) *Log {

	if l == nil {
		return nil
	}
	l.data = l.logger.appender.AppendUint16s(l.data, key, value)
	return l
}

func (l *Log) Uint32s(key string, value []uint32) *Log {

	if l == nil {
		return nil
	}
	l.data = l.logger.appender.AppendUint32s(l.data, key, value)
	return l
}

func (l *Log) Uint64s(key string, value []uint64) *Log {

	if l == nil {
		return nil
	}
	l.data = l.logger.appender.AppendUint64s(l.data, key, value)
	return l
}

func (l *Log) Float32s(key string, value []float32) *Log {

	if l == nil {
		return nil
	}
	l.data = l.logger.appender.AppendFloat32s(l.data, key, value)
	return l
}

func (l *Log) Float64s(key string, value []float64) *Log {

	if l == nil {
		return nil
	}
	l.data = l.logger.appender.AppendFloat64s(l.data, key, value)
	return l
}

func (l *Log) Strings(key string, value []string) *Log {

	if l == nil {
		return nil
	}
	l.data = l.logger.appender.AppendStrings(l.data, key, value)
	return l
}

func (l *Log) Times(key string, value []time.Time, format string) *Log {

	if l == nil {
		return nil
	}
	l.data = l.logger.appender.AppendTimes(l.data, key, value, format)
	return l
}

func (l *Log) Errors(key string, value []error) *Log {

	if l == nil {
		return nil
	}
	l.data = l.logger.appender.AppendErrors(l.data, key, value)
	return l
}

func (l *Log) Stringers(key string, value []fmt.Stringer) *Log {

	if l == nil {
		return nil
	}
	l.data = l.logger.appender.AppendStringers(l.data, key, value)
	return l
}
