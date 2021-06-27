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

func (l *Log) reset() {
	l.data = l.data[:0]
	l.data = l.logger.appender.Begin(l.data)
}

func (l *Log) Int(key string, value int) *Log {

	if l == nil {
		return nil
	}
	l.data = l.logger.appender.AppendInt(l.data, key, value)
	return l
}

func (l *Log) Float64(key string, value float64) *Log {

	if l == nil {
		return nil
	}
	l.data = l.logger.appender.AppendFloat64(l.data, key, value)
	return l
}

func (l *Log) Str(key string, value string) *Log {

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

func (l *Log) Record() {

	if l == nil {
		return
	}

	defer l.logger.releaseLog(l)
	l.logger.writer.Write(l.logger.appender.End(l.data))
}

func (l *Log) Msg(msg string) {

	if l == nil {
		return
	}
	l.data = l.logger.appender.AppendString(l.data, "msg", msg)
	l.Record()
}

func (l *Log) MsgF(msg string, params ...interface{}) {

	if len(params) > 0 {
		msg = fmt.Sprintf(msg, params...)
	}
	l.Msg(msg)
}
