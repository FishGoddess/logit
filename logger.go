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

	"github.com/go-logit/logit/core/appender"
	"github.com/go-logit/logit/core/writer"
	"github.com/go-logit/logit/support/global"
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
	logPool sync.Pool
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
		debugWriter:   writer.Wrap(os.Stdout),
		infoWriter:    writer.Wrap(os.Stdout),
		warnWriter:    writer.Wrap(os.Stderr),
		errorWriter:   writer.Wrap(os.Stderr),
		printWriter:   writer.Wrap(os.Stdout),
		logPool: sync.Pool{
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

// SetToGlobal clones a new logger of current logger and sets it to global.
// Depth of caller will increase 1 due to wrapping functions.
func (l *Logger) SetToGlobal() {
	newLogger := &Logger{
		config:        l.config,
		debugAppender: l.debugAppender,
		infoAppender:  l.infoAppender,
		warnAppender:  l.warnAppender,
		errorAppender: l.errorAppender,
		printAppender: l.printAppender,
		debugWriter:   l.debugWriter,
		infoWriter:    l.infoWriter,
		warnWriter:    l.warnWriter,
		errorWriter:   l.errorWriter,
		printWriter:   l.printWriter,
		interceptors:  l.interceptors,
		logPool: sync.Pool{
			New: func() interface{} {
				return newLog()
			},
		},
	}

	// Increase depth so the caller is correct.
	newLogger.callerDepth++
	SetGlobal(newLogger)
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
	log.data = log.data[:0]
	log.ctx = context.Background()
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
		log = log.WithTime(l.timeKey, global.CurrentTime(), l.timeFormat)
	}

	if l.levelKey != "" {
		log = log.String(l.levelKey, level.String())
	}

	if l.withPID {
		log = log.WithPID()
	}

	if l.withCaller {
		log = log.withCaller(l.callerDepth)
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
	l.log(printLevel, format, params...).Log()
}

// Print prints a log if print level is enabled.
func (l *Logger) Print(params ...interface{}) {
	l.log(printLevel, fmt.Sprint(params...)).Log()
}

// Println prints a log if print level is enabled.
func (l *Logger) Println(params ...interface{}) {
	l.log(printLevel, fmt.Sprintln(params...)).Log()
}

// Sync syncs data storing in logger's writer.
// You can use an option to sync automatically, see options.
// Close a logger will also invoke Sync(), so you can use an option or Close() to sync instead.
// However, you still need to sync manually if you want your logs to be stored immediately.
func (l *Logger) Sync() error {
	var err error

	if e := l.printWriter.Sync(); e != nil {
		err = e
	}

	if e := l.errorWriter.Sync(); e != nil {
		err = e
	}

	if e := l.warnWriter.Sync(); e != nil {
		err = e
	}

	if e := l.infoWriter.Sync(); e != nil {
		err = e
	}

	if e := l.debugWriter.Sync(); e != nil {
		err = e
	}

	return err
}

// Close closes logger and releases resources.
// It will sync data and set level to offLevel.
// It will invoke close() if writer is io.Closer.
// So, it is recommended for you to invoke it habitually.
func (l *Logger) Close() error {
	l.level = offLevel // uint8 is safe-concurrent in assignment, but may cause dirty read?

	if err := l.Sync(); err != nil {
		return err
	}

	if err := l.printWriter.Close(); err != nil {
		return err
	}

	if err := l.errorWriter.Close(); err != nil {
		return err
	}

	if err := l.warnWriter.Close(); err != nil {
		return err
	}

	if err := l.infoWriter.Close(); err != nil {
		return err
	}

	return l.debugWriter.Close()
}
