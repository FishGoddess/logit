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
	"time"
)

const (
	nan  = `"NaN"`
	pInf = `"+Inf"`
	nInf = `"-Inf"`

	lineBreak = '\n'
)

type appender interface {
	Begin(dst []byte) []byte
	End(dst []byte) []byte
	AppendAny(dst []byte, key string, value interface{}) []byte

	AppendByte(dst, key string, value byte) []byte
	AppendBool(dst []byte, key string, value bool) []byte
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
	AppendDuration(dst []byte, key string, value time.Duration) []byte
	AppendTime(dst []byte, key string, value time.Time, format string) []byte

	AppendBools(dst []byte, key string, values []bool) []byte
	AppendBytes(dst []byte, key string, values []byte) []byte
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
	AppendDurations(dst []byte, key string, values []time.Duration) []byte
	AppendTimes(dst []byte, key string, values []time.Time, format string) []byte
}
