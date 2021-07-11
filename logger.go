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
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/FishGoddess/logit/core/appender"
	"github.com/FishGoddess/logit/core/writer"
	"github.com/FishGoddess/logit/lib"
)

type Logger struct {
	*config
	appender appender.Appender
	writer   writer.Writer
	logPool  *sync.Pool
}

func NewLogger(options ...Option) *Logger {

	logger := new(Logger)
	logger.config = newDefaultConfig()
	logger.appender = appender.Text()
	logger.writer = writer.Wrapped(os.Stdout)
	logger.logPool = &sync.Pool{
		New: func() interface{} {
			return newLog(logger)
		},
	}

	for _, applyOption := range options {
		applyOption(logger)
	}
	return logger
}

func (l *Logger) getLog() *Log {
	return l.logPool.Get().(*Log)
}

func (l *Logger) releaseLog(log *Log) {
	l.logPool.Put(log)
}

func (l *Logger) log(level level, msg string, params ...interface{}) *Log {

	if level < l.level {
		return nil
	}

	log := l.getLog().begin()
	if l.timeKey != "" {
		log.Time(l.timeKey, time.Now(), l.timeFormat)
	}

	if l.levelKey != "" {
		log.String(l.levelKey, level.String())
	}

	if l.needPid && l.pidKey != "" {
		log.Int(l.pidKey, lib.Pid())
	}

	if l.needCaller && l.fileKey != "" && l.lineKey != "" {
		file, line := lib.Caller(3)
		log.String(l.fileKey, file).Int(l.lineKey, line)
	}

	if len(params) > 0 {
		msg = fmt.Sprintf(msg, params...)
	}
	log.String(l.msgKey, msg)
	return log
}

func (l *Logger) Debug(msg string, params ...interface{}) *Log {
	return l.log(debugLevel, msg, params...)
}

func (l *Logger) Info(msg string, params ...interface{}) *Log {
	return l.log(infoLevel, msg, params...)
}

func (l *Logger) Warn(msg string, params ...interface{}) *Log {
	return l.log(warnLevel, msg, params...)
}

func (l *Logger) Error(msg string, params ...interface{}) *Log {
	return l.log(errorLevel, msg, params...)
}

func (l *Logger) Flush() (n int, err error) {
	if flusher, ok := l.writer.(writer.Flusher); ok {
		return flusher.Flush()
	}
	return 0, nil
}

func (l *Logger) Close() error {
	if closer, ok := l.writer.(io.Closer); ok {
		return closer.Close()
	}
	l.level = offLevel
	return nil
}
