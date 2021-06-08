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

	// values stores all extra values of logger.
	// Every logs logged by this logger will carry this values.
	values M

	// logs is a log pool caching some log instances.
	// It is for reducing memory allocation.
	logs *sync.Pool

	// lock is for safe concurrency.
	lock *sync.RWMutex
}

// NewLogger returns a new logger instance with default settings.
func NewLogger(values ...M) *Logger {

	c := newCore(NewTextEncoder(TimeFormat), os.Stdout)
	c.SetLevel(InfoLevel)
	c.SetNeedCaller(false)
	c.Writers().SetErrorWriter(os.Stderr)
	return &Logger{
		core:   c,
		values: mergeValues(values...),
		logs: &sync.Pool{
			New: func() interface{} {
				return newLog()
			},
		},
		lock: &sync.RWMutex{},
	}
}

// filledWithCaller fills log with caller.
// This function is too expensive because of runtime.Caller.
// Notice that callerDepth is the depth of calling stack. See callerDepth.
func (l *Logger) filledWithCaller(log *Log) {

	if !l.NeedCaller() {
		return
	}

	if _, file, line, ok := runtime.Caller(callerDepth); ok {
		log.setCaller(file, line)
	}
}

// prepareLog prepares a Log holder for use.
func (l *Logger) prepareLog(level Level, msg string, values M) *Log {
	result := l.logs.Get().(*Log)
	result.msg = msg
	result.level = level
	result.time = time.Now()
	result.values = values
	return result
}

// releaseLog releases log to object pool so that this log can be reused next time.
func (l *Logger) releaseLog(log *Log) {
	log.reset()
	l.logs.Put(log)
}

// handleLog handles log with encoders and writers.
func (l *Logger) handleLog(log *Log) {
	l.filledWithCaller(log)
	encoder := l.Encoders().of(log.level)
	writer := l.Writers().of(log.level)
	writer.Write(encoder.Encode(log))
}

// log handles msg by l.handlers, and level will affect the visibility of this msg.
func (l *Logger) log(level Level, msg string, values ...M) {

	if l.Level() > level {
		return
	}

	log := l.prepareLog(level, msg, l.values.clone(values...))
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

// logF handles msg by l.handlers, and level will affect the visibility of this msg.
func (l *Logger) logF(level Level, msg string, params ...interface{}) {

	if l.Level() > level {
		return
	}

	if len(params) > 0 {
		msg = fmt.Sprintf(msg, params...)
	}

	log := l.prepareLog(level, msg, l.values.clone())
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
