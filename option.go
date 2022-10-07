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
	"io"
	"time"

	"github.com/FishGoddess/logit/core/appender"
	"github.com/FishGoddess/logit/core/writer"
)

// Option is a function that applies to logger.
type Option func(logger *Logger)

// Apply applies option to logger.
func (o Option) Apply(logger *Logger) {
	o(logger)
}

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

// WithPrintLevel returns an option which sets logger to print level.
func (o *options) WithPrintLevel() Option {
	return func(logger *Logger) {
		logger.level = printLevel
	}
}

// WithOffLevel returns an option which sets logger to off level.
func (o *options) WithOffLevel() Option {
	return func(logger *Logger) {
		logger.level = offLevel
	}
}

// WithAppender returns an option which sets logger's appender to a new one.
func (o *options) WithAppender(appender appender.Appender) Option {
	return func(logger *Logger) {
		logger.debugAppender = appender
		logger.infoAppender = appender
		logger.warnAppender = appender
		logger.errorAppender = appender
		logger.printAppender = appender
	}
}

// WithDebugAppender returns an option which sets logger's debug appender to a new one.
func (o *options) WithDebugAppender(appender appender.Appender) Option {
	return func(logger *Logger) {
		logger.debugAppender = appender
	}
}

// WithInfoAppender returns an option which sets logger's info appender to a new one.
func (o *options) WithInfoAppender(appender appender.Appender) Option {
	return func(logger *Logger) {
		logger.infoAppender = appender
	}
}

// WithWarnAppender returns an option which sets logger's warn appender to a new one.
func (o *options) WithWarnAppender(appender appender.Appender) Option {
	return func(logger *Logger) {
		logger.warnAppender = appender
	}
}

// WithErrorAppender returns an option which sets logger's error appender to a new one.
func (o *options) WithErrorAppender(appender appender.Appender) Option {
	return func(logger *Logger) {
		logger.errorAppender = appender
	}
}

// WithPrintAppender returns an option which sets logger's print appender to a new one.
func (o *options) WithPrintAppender(appender appender.Appender) Option {
	return func(logger *Logger) {
		logger.printAppender = appender
	}
}

// WithWriter returns an option which sets logger's writer to a new one.
func (o *options) WithWriter(w io.Writer) Option {
	return func(logger *Logger) {
		wr := writer.Wrap(w)
		logger.debugWriter = wr
		logger.infoWriter = wr
		logger.warnWriter = wr
		logger.errorWriter = wr
		logger.printWriter = wr
	}
}

// WithBufferWriter returns an option which sets logger's writer to a buffer one.
func (o *options) WithBufferWriter(w io.Writer) Option {
	return func(logger *Logger) {
		bw := writer.Buffer(w)
		logger.debugWriter = bw
		logger.infoWriter = bw
		logger.warnWriter = bw
		logger.errorWriter = bw
		logger.printWriter = bw
	}
}

// WithBatchWriter returns an option which sets logger's writer to a batch one.
func (o *options) WithBatchWriter(w io.Writer) Option {
	return func(logger *Logger) {
		bw := writer.Batch(w)
		logger.debugWriter = bw
		logger.infoWriter = bw
		logger.warnWriter = bw
		logger.errorWriter = bw
		logger.printWriter = bw
	}
}

// WithDebugWriter returns an option which sets logger's debug writer to a new one.
func (o *options) WithDebugWriter(w io.Writer) Option {
	return func(logger *Logger) {
		logger.debugWriter = writer.Wrap(w)
	}
}

// WithInfoWriter returns an option which sets logger's info writer to a new one.
func (o *options) WithInfoWriter(w io.Writer) Option {
	return func(logger *Logger) {
		logger.infoWriter = writer.Wrap(w)
	}
}

// WithWarnWriter returns an option which sets logger's warn writer to a new one.
func (o *options) WithWarnWriter(w io.Writer) Option {
	return func(logger *Logger) {
		logger.warnWriter = writer.Wrap(w)
	}
}

// WithErrorWriter returns an option which sets logger's error writer to a new one.
func (o *options) WithErrorWriter(w io.Writer) Option {
	return func(logger *Logger) {
		logger.errorWriter = writer.Wrap(w)
	}
}

// WithPrintWriter returns an option which sets logger's print writer to a new one.
func (o *options) WithPrintWriter(w io.Writer) Option {
	return func(logger *Logger) {
		logger.printWriter = writer.Wrap(w)
	}
}

// WithPID returns an option which lets logs carry pid information.
func (o *options) WithPID() Option {
	return func(logger *Logger) {
		logger.withPID = true
	}
}

// WithCaller returns an option which lets logs carry caller information.
func (o *options) WithCaller() Option {
	return func(logger *Logger) {
		logger.withCaller = true
	}
}

// WithCallerDepth returns an option which sets caller depth in logs.
func (o *options) WithCallerDepth(depth int) Option {
	return func(logger *Logger) {
		logger.callerDepth = depth
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

// WithPIDKey returns an option which sets logger's pidKey to a new one.
func (o *options) WithPIDKey(key string) Option {
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

// WithFuncKey returns an option which sets logger's funcKey to a new one.
func (o *options) WithFuncKey(key string) Option {
	return func(logger *Logger) {
		logger.funcKey = key
	}
}

// WithTimeFormat returns an option which sets format of time in logs.
func (o *options) WithTimeFormat(format string) Option {
	return func(logger *Logger) {
		logger.timeFormat = format
	}
}

// WithInterceptors returns an option which sets interceptors to logger.
func (o *options) WithInterceptors(interceptors ...Interceptor) Option {
	return func(logger *Logger) {
		logger.interceptors = append(logger.interceptors, interceptors...)
	}
}

// WithAutoSync returns an option which do sync automatically at fixed frequency.
func (o *options) WithAutoSync(frequency time.Duration) Option {
	return func(logger *Logger) {
		go func() {
			ticker := time.NewTicker(frequency)
			defer ticker.Stop()

			select {
			case <-ticker.C:
				logger.Sync()
			}
		}()
	}
}
