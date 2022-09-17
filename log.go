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

package logit

import (
	"context"
	"fmt"
	"time"

	"github.com/go-logit/logit/support/global"
	"github.com/go-logit/logit/support/runtime"

	"github.com/go-logit/logit/core/appender"
	"github.com/go-logit/logit/core/writer"
)

// Log stores data of a whole logging message.
// Notice: All functions in Log is unsafe-concurrent.
type Log struct {
	// logger is the owner of the log.
	logger *Logger

	// appender is an appender appending entries to the log.
	appender appender.Appender

	// writer writes the log to somewhere.
	writer writer.Writer

	// data stores all entries in log.
	data []byte

	// ctx is the context of this log.
	ctx context.Context
}

// newLog returns a new Log with pre-malloc memory.
// The default pre-malloc size is better to not-long logs.
// So if your logs are extremely long, you can set LogMallocSize to larger to avoid re-malloc.
func newLog() *Log {
	return &Log{
		data: make([]byte, 0, global.LogMallocSize),
		ctx:  context.Background(),
	}
}

// begin do some initializations of l.
func (l *Log) begin() *Log {
	if l == nil {
		return l
	}

	l.data = l.appender.Begin(l.data)
	return l
}

// end ends a log with writing and releasing.
func (l *Log) end() {
	if l == nil {
		return
	}

	defer l.logger.releaseLog(l)
	_, err := l.writer.Write(l.appender.End(l.data))

	global.HandleError("Log.writer.Write", err)
}

// Any adds an entry which key is string and value is interface{} type to l.
func (l *Log) Any(key string, value interface{}) *Log {
	if l == nil {
		return nil
	}

	l.data = l.appender.AppendAny(l.data, key, value)
	return l
}

// Json adds an entry which key is string and value is marshaled to a json string to l.
func (l *Log) Json(key string, value interface{}) *Log {
	if l == nil {
		return nil
	}

	l.data = l.appender.AppendJson(l.data, key, value)
	return l
}

// Bool adds an entry which key is string and value is bool type to l.
func (l *Log) Bool(key string, value bool) *Log {
	if l == nil {
		return nil
	}

	l.data = l.appender.AppendBool(l.data, key, value)
	return l
}

// Byte adds an entry which key is string and value is byte type to l.
func (l *Log) Byte(key string, value byte) *Log {
	if l == nil {
		return nil
	}

	l.data = l.appender.AppendByte(l.data, key, value)
	return l
}

// Rune adds an entry which key is string and value is rune type to l.
func (l *Log) Rune(key string, value rune) *Log {
	if l == nil {
		return nil
	}

	l.data = l.appender.AppendRune(l.data, key, value)
	return l
}

// Int adds an entry which key is string and value is int type to l.
func (l *Log) Int(key string, value int) *Log {
	if l == nil {
		return nil
	}

	l.data = l.appender.AppendInt(l.data, key, value)
	return l
}

// Int8 adds an entry which key is string and value is int8 type to l.
func (l *Log) Int8(key string, value int8) *Log {
	if l == nil {
		return nil
	}

	l.data = l.appender.AppendInt8(l.data, key, value)
	return l
}

// Int16 adds an entry which key is string and value is int16 type to l.
func (l *Log) Int16(key string, value int16) *Log {
	if l == nil {
		return nil
	}

	l.data = l.appender.AppendInt16(l.data, key, value)
	return l
}

// Int32 adds an entry which key is string and value is int32 type to l.
func (l *Log) Int32(key string, value int32) *Log {
	if l == nil {
		return nil
	}

	l.data = l.appender.AppendInt32(l.data, key, value)
	return l
}

// Int64 adds an entry which key is string and value is int64 type to l.
func (l *Log) Int64(key string, value int64) *Log {
	if l == nil {
		return nil
	}

	l.data = l.appender.AppendInt64(l.data, key, value)
	return l
}

// Uint adds an entry which key is string and value is uint type to l.
func (l *Log) Uint(key string, value uint) *Log {
	if l == nil {
		return nil
	}

	l.data = l.appender.AppendUint(l.data, key, value)
	return l
}

// Uint8 adds an entry which key is string and value is uint8 type to l.
func (l *Log) Uint8(key string, value uint8) *Log {
	if l == nil {
		return nil
	}

	l.data = l.appender.AppendUint8(l.data, key, value)
	return l
}

// Uint16 adds an entry which key is string and value is uint16 type to l.
func (l *Log) Uint16(key string, value uint16) *Log {
	if l == nil {
		return nil
	}

	l.data = l.appender.AppendUint16(l.data, key, value)
	return l
}

// Uint32 adds an entry which key is string and value is uint32 type to l.
func (l *Log) Uint32(key string, value uint32) *Log {
	if l == nil {
		return nil
	}

	l.data = l.appender.AppendUint32(l.data, key, value)
	return l
}

// Uint64 adds an entry which key is string and value is uint64 type to l.
func (l *Log) Uint64(key string, value uint64) *Log {
	if l == nil {
		return nil
	}

	l.data = l.appender.AppendUint64(l.data, key, value)
	return l
}

// Float32 adds an entry which key is string and value is float32 type to l.
func (l *Log) Float32(key string, value float32) *Log {
	if l == nil {
		return nil
	}

	l.data = l.appender.AppendFloat32(l.data, key, value)
	return l
}

// Float64 adds an entry which key is string and value is float64 type to l.
func (l *Log) Float64(key string, value float64) *Log {
	if l == nil {
		return nil
	}

	l.data = l.appender.AppendFloat64(l.data, key, value)
	return l
}

// String adds an entry which key is string and value is string type to l.
func (l *Log) String(key string, value string) *Log {
	if l == nil {
		return nil
	}

	l.data = l.appender.AppendString(l.data, key, value)
	return l
}

// Time adds an entry which key is string and value is time.Time type to l.
func (l *Log) Time(key string, value time.Time) *Log {
	if l == nil {
		return nil
	}

	if global.TimeLocation != nil {
		value = value.In(global.TimeLocation)
	}

	if l.logger != nil {
		l.data = l.appender.AppendTime(l.data, key, value, l.logger.timeFormat)
	} else {
		l.data = l.appender.AppendTime(l.data, key, value, global.UnixTimeFormat)
	}

	return l
}

// Error adds an entry which key is string and value is error type to l.
func (l *Log) Error(key string, value error) *Log {
	if l == nil {
		return nil
	}

	l.data = l.appender.AppendError(l.data, key, value)
	return l
}

// Stringer adds an entry which key is string and value is fmt.Stringer type to l.
func (l *Log) Stringer(key string, value fmt.Stringer) *Log {
	if l == nil {
		return nil
	}

	l.data = l.appender.AppendStringer(l.data, key, value)
	return l
}

// Bools adds an entry which key is string and value is []bool type to l.
func (l *Log) Bools(key string, value []bool) *Log {
	if l == nil {
		return nil
	}

	l.data = l.appender.AppendBools(l.data, key, value)
	return l
}

// Bytes adds an entry which key is string and value is []byte type to l.
func (l *Log) Bytes(key string, value []byte) *Log {
	if l == nil {
		return nil
	}

	l.data = l.appender.AppendBytes(l.data, key, value)
	return l
}

// Runes adds an entry which key is string and value is []rune type to l.
func (l *Log) Runes(key string, value []rune) *Log {
	if l == nil {
		return nil
	}

	l.data = l.appender.AppendRunes(l.data, key, value)
	return l
}

// Ints adds an entry which key is string and value is []int type to l.
func (l *Log) Ints(key string, value []int) *Log {
	if l == nil {
		return nil
	}

	l.data = l.appender.AppendInts(l.data, key, value)
	return l
}

// Int8s adds an entry which key is string and value is []int8 type to l.
func (l *Log) Int8s(key string, value []int8) *Log {
	if l == nil {
		return nil
	}

	l.data = l.appender.AppendInt8s(l.data, key, value)
	return l
}

// Int16s adds an entry which key is string and value is []int16 type to l.
func (l *Log) Int16s(key string, value []int16) *Log {
	if l == nil {
		return nil
	}

	l.data = l.appender.AppendInt16s(l.data, key, value)
	return l
}

// Int32s adds an entry which key is string and value is []int32 type to l.
func (l *Log) Int32s(key string, value []int32) *Log {
	if l == nil {
		return nil
	}

	l.data = l.appender.AppendInt32s(l.data, key, value)
	return l
}

// Int64s adds an entry which key is string and value is []int64 type to l.
func (l *Log) Int64s(key string, value []int64) *Log {
	if l == nil {
		return nil
	}

	l.data = l.appender.AppendInt64s(l.data, key, value)
	return l
}

// Uints adds an entry which key is string and value is []uint type to l.
func (l *Log) Uints(key string, value []uint) *Log {
	if l == nil {
		return nil
	}

	l.data = l.appender.AppendUints(l.data, key, value)
	return l
}

// Uint8s adds an entry which key is string and value is []uint8 type to l.
func (l *Log) Uint8s(key string, value []uint8) *Log {
	if l == nil {
		return nil
	}

	l.data = l.appender.AppendUint8s(l.data, key, value)
	return l
}

// Uint16s adds an entry which key is string and value is []uint16 type to l.
func (l *Log) Uint16s(key string, value []uint16) *Log {
	if l == nil {
		return nil
	}

	l.data = l.appender.AppendUint16s(l.data, key, value)
	return l
}

// Uint32s adds an entry which key is string and value is []uint32 type to l.
func (l *Log) Uint32s(key string, value []uint32) *Log {
	if l == nil {
		return nil
	}

	l.data = l.appender.AppendUint32s(l.data, key, value)
	return l
}

// Uint64s adds an entry which key is string and value is []uint64 type to l.
func (l *Log) Uint64s(key string, value []uint64) *Log {
	if l == nil {
		return nil
	}

	l.data = l.appender.AppendUint64s(l.data, key, value)
	return l
}

// Float32s adds an entry which key is string and value is []float32 type to l.
func (l *Log) Float32s(key string, value []float32) *Log {
	if l == nil {
		return nil
	}

	l.data = l.appender.AppendFloat32s(l.data, key, value)
	return l
}

// Float64s adds an entry which key is string and value is []float64 type to l.
func (l *Log) Float64s(key string, value []float64) *Log {
	if l == nil {
		return nil
	}

	l.data = l.appender.AppendFloat64s(l.data, key, value)
	return l
}

// Strings adds an entry which key is string and value is []string type to l.
func (l *Log) Strings(key string, value []string) *Log {
	if l == nil {
		return nil
	}

	l.data = l.appender.AppendStrings(l.data, key, value)
	return l
}

// Times adds an entry which key is string and value is []time.Time type to l.
func (l *Log) Times(key string, value []time.Time) *Log {
	if l == nil {
		return nil
	}

	if l.logger != nil {
		l.data = l.appender.AppendTimes(l.data, key, value, l.logger.timeFormat)
	} else {
		l.data = l.appender.AppendTimes(l.data, key, value, global.UnixTimeFormat)
	}

	return l
}

// Errors adds an entry which key is string and value is []error type to l.
func (l *Log) Errors(key string, value []error) *Log {
	if l == nil {
		return nil
	}

	l.data = l.appender.AppendErrors(l.data, key, value)
	return l
}

// Stringers adds an entry which key is string and value is []fmt.Stringer type to l.
func (l *Log) Stringers(key string, value []fmt.Stringer) *Log {
	if l == nil {
		return nil
	}

	l.data = l.appender.AppendStringers(l.data, key, value)
	return l
}

// WithTime adds an entry which key is string and value is time.Time formatted type to l.
func (l *Log) WithTime(key string, value time.Time, format string) *Log {
	if l == nil {
		return nil
	}

	l.data = l.appender.AppendTime(l.data, key, value, format)
	return l
}

// WithPID adds one entry about pid information to l.
func (l *Log) withPID() *Log {
	if l == nil {
		return l
	}

	if l.logger.pidKey != "" {
		l.Int(l.logger.pidKey, runtime.PID())
	}

	return l
}

// WithPID adds one entry about pid information to l.
func (l *Log) WithPID() *Log {
	if l == nil || l.logger.withPID {
		return l
	}

	return l.withPID()
}

// WithCaller adds some entries about caller information to l.
func (l *Log) withCaller(depth int) *Log {
	if l == nil {
		return nil
	}

	file, line, function := runtime.Caller(depth)
	if l.logger.fileKey != "" {
		l.String(l.logger.fileKey, file)
	}

	if l.logger.lineKey != "" {
		l.Int(l.logger.lineKey, line)
	}

	if l.logger.funcKey != "" {
		l.String(l.logger.funcKey, function)
	}

	return l
}

// WithCallerOf adds some entries about caller information to l.
func (l *Log) WithCallerOf(depth int) *Log {
	if l == nil || l.logger.withCaller {
		return l
	}

	return l.withCaller(depth)
}

// WithCaller adds some entries about caller information to l.
func (l *Log) WithCaller() *Log {
	return l.WithCallerOf(global.CallerDepth)
}

// WithContext sets ctx to l.
func (l *Log) WithContext(ctx context.Context) *Log {
	if l == nil {
		return nil
	}

	l.ctx = ctx
	return l
}

// Intercept intercepts l with interceptors.
func (l *Log) Intercept(interceptors ...Interceptor) *Log {
	if l == nil {
		return nil
	}

	for _, interceptor := range interceptors {
		interceptor(l.ctx, l)
	}

	return l
}

// Log logs l.
func (l *Log) Log() {
	if l == nil {
		return
	}

	for _, interceptor := range l.logger.interceptors {
		interceptor(l.ctx, l)
	}

	l.end()
}
