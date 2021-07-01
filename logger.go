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
	"os"
	"sync"
	"time"

	"github.com/FishGoddess/logit/core/appender"
	"github.com/FishGoddess/logit/core/writer"
)

type Logger struct {
	level    level
	appender appender.Appender
	writer   writer.Writer
	logPool  *sync.Pool
}

func NewLogger(options ...Option) *Logger {

	logger := new(Logger)
	logger.level = debugLevel
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

func (l *Logger) newLog() *Log {
	log := l.logPool.Get().(*Log)
	log.initialize()
	return log
}

func (l *Logger) releaseLog(log *Log) {
	l.logPool.Put(log)
}

func (l *Logger) log(level level) *Log {

	if level < l.level {
		return nil
	}
	return l.newLog().Time("log.time", time.Now(), "2006-01-02 15:04:05").String("log.level", level.String())
}

func (l *Logger) Debug() *Log {
	return l.log(debugLevel)
}

func (l *Logger) Info() *Log {
	return l.log(infoLevel)
}

func (l *Logger) Warn() *Log {
	return l.log(warnLevel)
}

func (l *Logger) Error() *Log {
	return l.log(errorLevel)
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
