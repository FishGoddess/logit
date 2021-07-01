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

	"github.com/FishGoddess/logit/core/appender"
	"github.com/FishGoddess/logit/core/writer"
)

var (
	globalOptions = (*options)(nil)
)

type Option func(logger *Logger)

type options struct {
}

func Options() *options {
	return globalOptions
}

func (o *options) WithDebug() Option {
	return func(logger *Logger) {
		logger.level = debugLevel
	}
}

func (o *options) WithInfo() Option {
	return func(logger *Logger) {
		logger.level = infoLevel
	}
}

func (o *options) WithWarn() Option {
	return func(logger *Logger) {
		logger.level = warnLevel
	}
}

func (o *options) WithError() Option {
	return func(logger *Logger) {
		logger.level = errorLevel
	}
}

func (o *options) WithAppender(appender appender.Appender) Option {
	return func(logger *Logger) {
		logger.appender = appender
	}
}

func (o *options) WithWriter(w io.Writer) Option {
	return func(logger *Logger) {
		logger.writer = writer.Wrapped(w)
	}
}

func (o *options) WithBuffered(w io.Writer) Option {
	return func(logger *Logger) {
		logger.writer = writer.Buffered(w)
	}
}

func (o *options) WithCaller() Option {
	return func(logger *Logger) {
		logger.needCaller = true
	}
}

func (o *options) WithMsgKey(key string) Option {
	return func(logger *Logger) {
		logger.msgKey = key
	}
}

func (o *options) WithTimeKey(key string) Option {
	return func(logger *Logger) {
		logger.timeKey = key
	}
}

func (o *options) WithLevelKey(key string) Option {
	return func(logger *Logger) {
		logger.levelKey = key
	}
}

func (o *options) WithFileKey(key string) Option {
	return func(logger *Logger) {
		logger.fileKey = key
	}
}

func (o *options) WithLineKey(key string) Option {
	return func(logger *Logger) {
		logger.lineKey = key
	}
}

func (o *options) WithTimeFormat(format string) Option {
	return func(logger *Logger) {
		logger.timeFormat = format
	}
}
