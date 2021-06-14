// Copyright 2020 Ye Zi Jie. All Rights Reserved.
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
// Created at 2020/02/29 15:39:02

package logit

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

const (
	// callerDepth is the depth of calling stack, which is about file name and line number.
	callerDepth = 4

	// TimeFormat is the format of time.
	TimeFormat = "2006-01-02 15:04:05"
)

// Logger is the type of logging output.
type Logger struct {

	// core is the core of this logger.
	// Some settings are set in this core, such as level of logger and need caller or not.
	*core

	// kvs stores all extra key-values of logger.
	// Every logs logged by this logger will carry this key-values.
	kvs M

	mPool *sync.Pool

	// logPool is a log pool caching some log instances.
	// It is for reducing memory allocation.
	logPool *sync.Pool

	// lock is for safe concurrency.
	lock *sync.RWMutex
}

// NewLogger returns a new logger instance with default settings.
func NewLogger(options ...Options) *Logger {

	c := newCore(NewTextEncoder(TimeFormat), os.Stdout).SetLevel(InfoLevel).SetNeedCaller(false)
	c.Writers().SetErrorWriter(os.Stderr)
	result := &Logger{
		core: c,
		mPool: &sync.Pool{
			New: func() interface{} {
				return M{}
			},
		},
		logPool: &sync.Pool{
			New: func() interface{} {
				return newLog()
			},
		},
		lock: &sync.RWMutex{},
	}

	for _, opt := range options {
		opt.Apply(result)
	}
	return result
}

func (l *Logger) prepareM(ms ...M) M {

	if len(l.kvs) <= 0 && len(ms) <= 0 {
		return nil
	}

	result := l.mPool.Get().(M)
	for k, v := range l.kvs {
		result[k] = v
	}

	for _, kvs := range ms {
		for k, v := range kvs {
			result[k] = v
		}
	}
	return result
}

// releaseM releases m to object pool so that this m can be reused next time.
func (l *Logger) releaseM(m M) {
	if m != nil {
		m.reset()
		l.mPool.Put(m)
	}
}

// prepareLog prepares a Log holder for use.
func (l *Logger) prepareLog(level Level, msg string, ms ...M) *Log {
	result := l.logPool.Get().(*Log)
	result.msg = msg
	result.level = level
	result.time = time.Now()
	result.kvs = l.prepareM(ms...)
	return result
}

// releaseLog releases log to object pool so that this log can be reused next time.
func (l *Logger) releaseLog(log *Log) {
	l.releaseM(log.kvs)
	log.reset()
	l.logPool.Put(log)
}

// filledWithCallerIfNeed fills log with caller if needCaller in logger is true.
// This function is too expensive because of runtime.Caller.
// Notice that callerDepth is the depth of calling stack. See callerDepth.
func (l *Logger) filledWithCallerIfNeed(log *Log) {

	if !l.NeedCaller() {
		return
	}

	if _, file, line, ok := runtime.Caller(callerDepth); ok {
		log.setCaller(file, line)
	}
}

// handleLog handles log with encoders and writers.
func (l *Logger) handleLog(log *Log) {
	l.filledWithCallerIfNeed(log)
	encoder := l.Encoders().of(log.level)
	writer := l.Writers().of(log.level)
	writer.Write(encoder.Encode(log))
}

// ======================================== struct log ========================================

// log handles msg by l.handlers, and level will affect the visibility of this msg.
func (l *Logger) log(level Level, msg string, values ...M) {

	if l.Level() > level {
		return
	}

	log := l.prepareLog(level, msg, values...)
	defer l.releaseLog(log)
	l.handleLog(log)
}

// Debug will output msg as a debug message.
func (l *Logger) Debug(msg string, values ...M) {
	l.log(DebugLevel, msg, values...)
}

// Info will output msg as an info message.
func (l *Logger) Info(msg string, values ...M) {
	l.log(InfoLevel, msg, values...)
}

// Warn will output msg as a warn message.
func (l *Logger) Warn(msg string, values ...M) {
	l.log(WarnLevel, msg, values...)
}

// Error will output msg as an error message.
func (l *Logger) Error(msg string, values ...M) {
	l.log(ErrorLevel, msg, values...)
}

// ======================================== format log ========================================

func formatMsg(msg string, params ...interface{}) string {

	if len(params) > 0 {
		return fmt.Sprintf(msg, params...)
	}
	return msg
}

func (l *Logger) logF(level Level, msg string, params ...interface{}) {

	if l.Level() > level {
		return
	}

	log := l.prepareLog(level, formatMsg(msg, params...))
	defer l.releaseLog(log)
	l.handleLog(log)
}

// DebugF will output msg as a debug message.
func (l *Logger) DebugF(msg string, params ...interface{}) {
	l.logF(DebugLevel, msg, params...)
}

// InfoF will output msg as an info message.
func (l *Logger) InfoF(msg string, params ...interface{}) {
	l.logF(InfoLevel, msg, params...)
}

// WarnF will output msg as a warn message.
func (l *Logger) WarnF(msg string, params ...interface{}) {
	l.logF(WarnLevel, msg, params...)
}

// ErrorF will output msg as an error message.
func (l *Logger) ErrorF(msg string, params ...interface{}) {
	l.logF(ErrorLevel, msg, params...)
}
