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
// Created at 2021/07/01 23:01:38

package logit

import (
	"io"
	"time"

	"github.com/FishGoddess/logit/core/appender"
	"github.com/FishGoddess/logit/core/writer"
)

// Option is a function that applies to logger.
type Option func(logger *Logger)

// options stores all provided options.
type options struct{}

// Options returns all options provided.
func Options() *options {
	return (*options)(nil)
}

// WithDebugLevel returns an option which sets logger to debug level.
func (o *options) WithDebugLevel() Option {
	return func(logger *Logger) {
		logger.level = debugLevel
	}
}

// WithInfoLevel returns an option which sets logger to info level.
func (o *options) WithInfoLevel() Option {
	return func(logger *Logger) {
		logger.level = infoLevel
	}
}

// WithWarnLevel returns an option which sets logger to warn level.
func (o *options) WithWarnLevel() Option {
	return func(logger *Logger) {
		logger.level = warnLevel
	}
}

// WithErrorLevel returns an option which sets logger to error level.
func (o *options) WithErrorLevel() Option {
	return func(logger *Logger) {
		logger.level = errorLevel
	}
}

// WithAppender returns an option which sets logger's appender to a new one.
func (o *options) WithAppender(appender appender.Appender) Option {
	return func(logger *Logger) {
		logger.appender = appender
	}
}

// WithWriter returns an option which sets logger's writer to a new one.
func (o *options) WithWriter(w io.Writer) Option {
	return func(logger *Logger) {
		logger.writer = writer.Wrapped(w)
	}
}

// WithBuffered returns an option which sets logger's writer to a new one with buffer.
func (o *options) WithBuffered(w io.Writer) Option {
	return func(logger *Logger) {
		logger.writer = writer.Buffered(w)
	}
}

// WithPid returns an option which lets logs carry pid information.
func (o *options) WithPid() Option {
	return func(logger *Logger) {
		logger.needPid = true
	}
}

// WithCaller returns an option which lets logs carry caller information.
func (o *options) WithCaller() Option {
	return func(logger *Logger) {
		logger.needCaller = true
	}
}

// WithMsgKey returns an option which sets logger's msgKey to a new one.
func (o *options) WithMsgKey(key string) Option {
	return func(logger *Logger) {
		logger.msgKey = key
	}
}

// WithTimeKey returns an option which sets logger's timeKey to a new one.
func (o *options) WithTimeKey(key string) Option {
	return func(logger *Logger) {
		logger.timeKey = key
	}
}

// WithLevelKey returns an option which sets logger's levelKey to a new one.
func (o *options) WithLevelKey(key string) Option {
	return func(logger *Logger) {
		logger.levelKey = key
	}
}

// WithPidKey returns an option which sets logger's pidKey to a new one.
func (o *options) WithPidKey(key string) Option {
	return func(logger *Logger) {
		logger.pidKey = key
	}
}

// WithFileKey returns an option which sets logger's fileKey to a new one.
func (o *options) WithFileKey(key string) Option {
	return func(logger *Logger) {
		logger.fileKey = key
	}
}

// WithLineKey returns an option which sets logger's lineKey to a new one.
func (o *options) WithLineKey(key string) Option {
	return func(logger *Logger) {
		logger.lineKey = key
	}
}

// WithTimeFormat returns an option which sets format of time in logs.
func (o *options) WithTimeFormat(format string) Option {
	return func(logger *Logger) {
		logger.timeFormat = format
	}
}

// WithAutoFlush returns an option which do flush automatically at fixed frequency.
func (o *options) WithAutoFlush(frequency time.Duration) Option {
	return func(logger *Logger) {
		ticker := time.NewTicker(frequency)
		go func() {
			select {
			case <-ticker.C:
				logger.Flush()
			}
		}()
	}
}

// WithAutoClose returns an option which do close automatically at some signal happened.
func (o *options) WithAutoClose() Option {
	return func(logger *Logger) {
		// TODO 监听特定的信号，当信号发生时执行 logger.Close()
	}
}
