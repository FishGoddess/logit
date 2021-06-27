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
	"time"

	"github.com/FishGoddess/logit/appender"
)

type Logger struct {
	level    Level
	appender appender.Appender
	writer   io.Writer
	logPool  *sync.Pool
}

func NewLogger(level Level, appender appender.Appender, writer io.Writer) *Logger {

	logger := &Logger{
		level:    level,
		appender: appender,
		writer:   writer,
	}

	logger.logPool = &sync.Pool{
		New: func() interface{} {
			return newLog(logger)
		},
	}
	return logger
}

func (l *Logger) newLog() *Log {
	log := l.logPool.Get().(*Log)
	log.reset()
	return log
}

func (l *Logger) releaseLog(log *Log) {
	l.logPool.Put(log)
}

func (l *Logger) log(level Level) *Log {

	if level < l.level {
		return nil
	}
	return l.newLog().Str("level", level.String()).Time("time", time.Now(), "2006-01-02 15:04:05")
}

func (l *Logger) Debug() *Log {
	return l.log(DebugLevel)
}

func (l *Logger) Info() *Log {
	return l.log(InfoLevel)
}

func (l *Logger) Warn() *Log {
	return l.log(WarnLevel)
}

func (l *Logger) Error() *Log {
	return l.log(ErrorLevel)
}
