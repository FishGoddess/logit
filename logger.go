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
	"os"
	"sync"
	"time"

	"github.com/go-logit/logit/core/appender"
	"github.com/go-logit/logit/core/writer"
)

// Interceptor intercepts log with context.
type Interceptor = func(ctx context.Context, log *Log)

// Logger is the core of logging operations.
type Logger struct {
	// config stores all configurations of logger.
	config

	// debugAppender, infoAppender, warnAppender, errorAppender, printAppender is an appender appending entries to debug, info, warn, error, print logs.
	debugAppender appender.Appender
	infoAppender  appender.Appender
	warnAppender  appender.Appender
	errorAppender appender.Appender
	printAppender appender.Appender

	// debugWriter, infoWriter, warnWriter, errorWriter, printWriter writes debug, info, warn, error, print logs to somewhere.
	debugWriter writer.Writer
	infoWriter  writer.Writer
	warnWriter  writer.Writer
	errorWriter writer.Writer
	printWriter writer.Writer

	// interceptors stores all interceptors.
	interceptors []Interceptor

	// logPool is for reusing logs.
	logPool *sync.Pool
}

// NewLogger returns a new Logger created with options.
func NewLogger(options ...Option) *Logger {
	logger := &Logger{
		config:        newDefaultConfig(),
		debugAppender: appender.Text(),
		infoAppender:  appender.Text(),
		warnAppender:  appender.Text(),
		errorAppender: appender.Text(),
		printAppender: appender.Text(),
		debugWriter:   writer.Wrapped(os.Stdout),
		infoWriter:    writer.Wrapped(os.Stdout),
		warnWriter:    writer.Wrapped(os.Stderr),
		errorWriter:   writer.Wrapped(os.Stderr),
		printWriter:   writer.Wrapped(os.Stdout),
		logPool: &sync.Pool{
			New: func() interface{} {
				return newLog()
			},
		},
	}

	for _, applyOption := range options {
		applyOption(logger)
	}

	return logger
}

// appenderOf returns the appender of level.
func (l *Logger) appenderOf(level level) appender.Appender {
	switch level {
	case printLevel:
		return l.printAppender
	case errorLevel:
		return l.errorAppender
	case warnLevel:
		return l.warnAppender
	case infoLevel:
		return l.infoAppender
	default:
		return l.debugAppender
	}
}

// writerOf returns the writer of level.
func (l *Logger) writerOf(level level) writer.Writer {
	switch level {
	case printLevel:
		return l.printWriter
	case errorLevel:
		return l.errorWriter
	case warnLevel:
		return l.warnWriter
	case infoLevel:
		return l.infoWriter
	default:
		return l.debugWriter
	}
}

// getLog returns a Log instance from pool.
// This is a better way to memory.
func (l *Logger) getLog(level level) *Log {
	log := l.logPool.Get().(*Log)
	log.logger = l
	log.appender = l.appenderOf(level)
	log.writer = l.writerOf(level)
	return log
}

// releaseLog releases a Log instance to pool.
func (l *Logger) releaseLog(log *Log) {
	l.logPool.Put(log)
}

// log returns a Log instance with level and msg.
// Check Log for more information.
func (l *Logger) log(level level, msg string, params ...interface{}) *Log {
	if level < l.level {
		return nil
	}

	log := l.getLog(level).begin()
	if l.timeKey != "" {
		log.Time(l.timeKey, time.Now(), l.timeFormat)
	}

	if l.levelKey != "" {
		log.String(l.levelKey, level.String())
	}

	if l.needPid {
		log.WithPid()
	}

	if l.needCaller {
		log.withCaller(l.callerDepth)
	}

	if len(params) > 0 {
		msg = fmt.Sprintf(msg, params...)
	}

	return log.String(l.msgKey, msg)
}

// Debug returns a Log with debug level if debug level is enabled.
func (l *Logger) Debug(msg string, params ...interface{}) *Log {
	return l.log(debugLevel, msg, params...)
}

// Info returns a Log with info level if info level is enabled.
func (l *Logger) Info(msg string, params ...interface{}) *Log {
	return l.log(infoLevel, msg, params...)
}

// Warn returns a Log with warn level if warn level is enabled.
func (l *Logger) Warn(msg string, params ...interface{}) *Log {
	return l.log(warnLevel, msg, params...)
}

// Error returns a Log with error level if error level is enabled.
func (l *Logger) Error(msg string, params ...interface{}) *Log {
	return l.log(errorLevel, msg, params...)
}

// Printf prints a log if print level is enabled.
func (l *Logger) Printf(format string, params ...interface{}) {
	l.log(printLevel, format, params...).End()
}

// Print prints a log if print level is enabled.
func (l *Logger) Print(params ...interface{}) {
	l.log(printLevel, fmt.Sprint(params...)).End()
}

// Println prints a log if print level is enabled.
func (l *Logger) Println(params ...interface{}) {
	l.log(printLevel, fmt.Sprintln(params...)).End()
}

// Flush flushes data storing in logger's writer.
// This isn't necessary for all writers, but buffered writer needs.
// Actually, you can use an option to flush automatically, see options.
// Close a logger will also invoke Flush(), so you can use an option or Close() to flush instead.
// However, you still need to flush manually if you want your logs store immediately.
func (l *Logger) Flush() (n int, err error) {
	i, e := l.printWriter.Flush()
	if e != nil {
		err = e
	}

	n += i

	i, e = l.errorWriter.Flush()
	if e != nil {
		err = e
	}

	n += i

	i, e = l.warnWriter.Flush()
	if e != nil {
		err = e
	}

	n += i

	i, e = l.infoWriter.Flush()
	if e != nil {
		err = e
	}

	n += i

	i, e = l.debugWriter.Flush()
	if e != nil {
		err = e
	}

	n += i
	return n, err
}

// Close closes logger and releases resources.
// It will flush data and set level to offLevel.
// It will invoke close() if writer is io.Closer.
// So, it is recommended for you to invoke it habitually.
func (l *Logger) Close() error {
	l.level = offLevel

	_, err := l.Flush()
	if err != nil {
		return err
	}

	err = l.printWriter.Close()
	if err != nil {
		return err
	}

	err = l.errorWriter.Close()
	if err != nil {
		return err
	}

	err = l.warnWriter.Close()
	if err != nil {
		return err
	}

	err = l.infoWriter.Close()
	if err != nil {
		return err
	}

	return l.debugWriter.Close()
}
