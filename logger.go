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
// Created at 2021/06/27 16:40:31

package logit

import (
	"io"
	"sync"
)

// []byte => interface{} will cause extra memory, so we wraps []byte into a struct which can be passed by pointer.
type buffer struct {
	data []byte
}

type Logger struct {
	level      Level
	encoder    Encoder
	writer     io.Writer
	bufferPool *sync.Pool
}

func NewLogger(level Level, encoder Encoder, writer io.Writer) *Logger {
	return &Logger{
		level:   level,
		encoder: encoder,
		writer:  writer,
		bufferPool: &sync.Pool{
			New: func() interface{} {
				return &buffer{data: make([]byte, 0, 512)}
			},
		},
	}
}

func (l *Logger) newBuffer() *buffer {
	buff := l.bufferPool.Get().(*buffer)
	buff.data = buff.data[:0]
	return buff
}

func (l *Logger) freeBuffer(buff *buffer) {
	l.bufferPool.Put(buff)
}

func (l *Logger) log(level Level, msg string, kvs []interface{}) {

	if level < l.level {
		return
	}

	kvsLen := len(kvs)
	if kvsLen > 0 && kvsLen&1 != 0 {
		kvs = append(kvs, nil)
	}

	buff := l.newBuffer()
	defer l.freeBuffer(buff)

	buff.data = l.encoder.Begin(buff.data)
	buff.data = l.encoder.AppendString(buff.data, "level", level.String())
	buff.data = l.encoder.AppendString(buff.data, "msg", msg)
	for i := 0; i < kvsLen; i += 2 {
		buff.data = l.encoder.Append(buff.data, kvs[i].(string), kvs[i+1])
	}
	buff.data = l.encoder.End(buff.data)

	l.writer.Write(buff.data)
}

func (l *Logger) Debug(msg string, kvs ...interface{}) {
	l.log(DebugLevel, msg, kvs)
}

func (l *Logger) Info(msg string, kvs ...interface{}) {
	l.log(InfoLevel, msg, kvs)
}

func (l *Logger) Warn(msg string, kvs ...interface{}) {
	l.log(WarnLevel, msg, kvs)
}

func (l *Logger) Error(msg string, kvs ...interface{}) {
	l.log(ErrorLevel, msg, kvs)
}
