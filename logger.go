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
	"io"
	"os"
	"runtime"
	"sync"
	"time"
)

const (
	// callDepth is the depth of calling stack, which is about file name and line number.
	callDepth = 3
)

// Logger is the type of logging output.
type Logger struct {

	// level is the position of log.
	// In this version of logit, there are five levels:
	//
	//  DebugLevel, InfoLevel, WarnLevel, ErrorLevel, OffLevel.
	//
	// Higher level has higher visibility which means
	// one debug log will not be logged in one Logger set to InfoLevel.
	// That's we called level-based logger.
	//
	// In particular, OffLevel is the highest level, so if you set one
	// logger to OffLevel, it will shut up and log nothing.
	level Level

	// encoders are used to encode a log to bytes.
	// Every level has own encoder.
	encoders map[Level]Encoder

	// writers are used to output an encoded log.
	// Every level has own writer.
	writers map[Level]io.Writer

	// needCaller is a flag to check if logs should contain caller's info or not.
	// This feature is useful but expensive in performance, so set to false if you don't need it.
	needCaller bool

	// timeFormat is used to format time.
	timeFormat string

	// logs is a log pool caching some log instances.
	// It is for reducing memory allocation.
	logs *sync.Pool

	// lock is for safe concurrency.
	lock *sync.RWMutex
}

// NewLogger returns a new logger instance with default settings.
func NewLogger() *Logger {
	return &Logger{
		level: InfoLevel,
		encoders: map[Level]Encoder{
			DebugLevel: TextEncoder(),
			InfoLevel:  TextEncoder(),
			WarnLevel:  TextEncoder(),
			ErrorLevel: TextEncoder(),
		},
		writers: map[Level]io.Writer{
			DebugLevel: os.Stdout,
			InfoLevel:  os.Stdout,
			WarnLevel:  os.Stdout,
			ErrorLevel: os.Stderr,
		},
		needCaller: false,
		timeFormat: "2006-01-02 15:04:05",
		logs: &sync.Pool{
			New: func() interface{} {
				return &Log{}
			},
		},
		lock: &sync.RWMutex{},
	}
}

// SetLevel will change the logging level to newLevel.
// It returns the old level.
func (l *Logger) SetLevel(newLevel Level) Level {
	l.lock.Lock()
	defer l.lock.Unlock()
	oldLevel := l.level
	l.level = newLevel
	return oldLevel
}

// Level returns the logger level of current Logger.
func (l *Logger) Level() Level {
	l.lock.RLock()
	defer l.lock.RUnlock()
	return l.level
}

// SetEncoder sets encoder to new one.
// This encoder will apply to all levels.
func (l *Logger) SetEncoder(encoder Encoder) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.encoders[DebugLevel] = encoder
	l.encoders[InfoLevel] = encoder
	l.encoders[WarnLevel] = encoder
	l.encoders[ErrorLevel] = encoder
}

// SetDebugEncoder sets encoder of debug to new one.
func (l *Logger) SetDebugEncoder(encoder Encoder) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.encoders[DebugLevel] = encoder
}

// SetInfoEncoder sets encoder of info to new one.
func (l *Logger) SetInfoEncoder(encoder Encoder) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.encoders[InfoLevel] = encoder
}

// SetWarnEncoder sets encoder of warn to new one.
func (l *Logger) SetWarnEncoder(encoder Encoder) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.encoders[WarnLevel] = encoder
}

// SetErrorEncoder sets encoder of error to new one.
func (l *Logger) SetErrorEncoder(encoder Encoder) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.encoders[ErrorLevel] = encoder
}

// SetWriter sets writer to new one.
// This writer will apply to all levels.
func (l *Logger) SetWriter(writer io.Writer) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.writers[DebugLevel] = writer
	l.writers[InfoLevel] = writer
	l.writers[WarnLevel] = writer
	l.writers[ErrorLevel] = writer
}

// SetDebugWriter sets writer of debug to new one.
func (l *Logger) SetDebugWriter(writer io.Writer) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.writers[DebugLevel] = writer
}

// SetInfoWriter sets writer of info to new one.
func (l *Logger) SetInfoWriter(writer io.Writer) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.writers[InfoLevel] = writer
}

// SetWarnWriter sets writer of warn to new one.
func (l *Logger) SetWarnWriter(writer io.Writer) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.writers[WarnLevel] = writer
}

// SetErrorWriter sets writer of error to new one.
func (l *Logger) SetErrorWriter(writer io.Writer) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.writers[ErrorLevel] = writer
}

// NeedCaller sets needCaller to new one.
// If true, then every log will contain file name and line number.
// However, you should know that this is expensive in time.
// So be sure you really need it or keep it false.
func (l *Logger) NeedCaller(needCaller bool) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.needCaller = needCaller
}

// TimeFormat sets timeFormat to new one.
// This format follows the format in time package of Go.
// Yep, it is "2006-01-02 15:04:05".
func (l *Logger) TimeFormat(timeFormat string) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.timeFormat = timeFormat
}

// newLog returns a Log holder from object pool.
// Notice that not every holder returned is new, as you know, that is why we use a pool.
func (l *Logger) newLog(level Level, msg string) *Log {
	log := l.logs.Get().(*Log)
	log.level = level
	log.msg = msg
	log.time = time.Now()
	return log
}

// releaseLog releases log to object pool so that this log can be reused next time.
func (l *Logger) releaseLog(log *Log) {
	if caller, ok := log.Caller(); ok {
		caller.File = ""
		caller.Line = 0
	}
	l.logs.Put(log)
}

// wrapLogWithCaller wraps log with caller.
// This function is too expensive because of runtime.Caller.
// Notice that callDepth is the depth of calling stack. See callDepth.
func wrapLogWithCaller(log *Log, callDepth int) {

	log.caller = &caller{
		File: "unknown file",
		Line: -1,
	}

	if _, file, line, ok := runtime.Caller(callDepth); ok {
		log.caller.File = file
		log.caller.Line = line
	}
}

// handleLog handles log with encoders and writers.
func (l *Logger) handleLog(log *Log) {
	encoder := l.encoders[log.level]
	writer := l.writers[log.level]
	writer.Write(encoder.Encode(log, l.timeFormat))
}

// log handles msg by l.handlers, and level will affect the visibility of this msg.
// Notice that callDepth is caller sensitive.
func (l *Logger) log(callDepth int, level Level, msg string) {

	l.lock.RLock()

	if l.level > level {
		l.lock.RUnlock()
		return
	}

	// Use copy-on-write way to keep high performance
	needCaller := l.needCaller
	l.lock.RUnlock()

	log := l.newLog(level, msg)
	defer l.releaseLog(log)

	if needCaller {
		wrapLogWithCaller(log, callDepth)
	}
	l.handleLog(log)
}

// Debug will output msg as a debug message.
func (l *Logger) Debug(msg string) {
	l.log(callDepth, DebugLevel, msg)
}

// Info will output msg as an info message.
func (l *Logger) Info(msg string) {
	l.log(callDepth, InfoLevel, msg)
}

// Warn will output msg as a warn message.
func (l *Logger) Warn(msg string) {
	l.log(callDepth, WarnLevel, msg)
}

// Error will output msg as an error message.
func (l *Logger) Error(msg string) {
	l.log(callDepth, ErrorLevel, msg)
}

// generateMessage generates a message from format and params.
func generateMessage(format string, params ...interface{}) string {
	return fmt.Sprintf(format, params...)
}

// DebugF will output msg as a debug message.
// The msg is the return value of generateMessage.
// This is a way to output a long log made from many variables.
// The msgFormat is the same as format in fmt.Printf, so you can use
// all format it supports, such as '%d'.
// msgParams is the params msgFormat needs, and it is variable-length, so
// you can add all your params here.
// You should know that this way to output msg is the most expensive way in time,
// but it's still faster than other logging libs. If you care about performance,
// than you should think about it, and if you don't, just use it without thinking.
func (l *Logger) DebugF(msgFormat string, msgParams ...interface{}) {
	l.log(callDepth, DebugLevel, generateMessage(msgFormat, msgParams...))
}

// InfoF will output msg as an info message.
// The msg is the return value of generateMessage.
// This is a way to output a long log made from many variables.
// The msgFormat is the same as format in fmt.Printf, so you can use
// all format it supports, such as '%d'.
// msgParams is the params msgFormat needs, and it is variable-length, so
// you can add all your params here.
// You should know that this way to output msg is the most expensive way in time,
// but it's still faster than other logging libs. If you care about performance,
// than you should think about it, and if you don't, just use it without thinking.
func (l *Logger) InfoF(msgFormat string, msgParams ...interface{}) {
	l.log(callDepth, InfoLevel, generateMessage(msgFormat, msgParams...))
}

// WarnF will output msg as a warn message.
// The msg is the return value of generateMessage.
// This is a way to output a long log made from many variables.
// The msgFormat is the same as format in fmt.Printf, so you can use
// all format it supports, such as '%d'.
// msgParams is the params msgFormat needs, and it is variable-length, so
// you can add all your params here.
// You should know that this way to output msg is the most expensive way in time,
// but it's still faster than other logging libs. If you care about performance,
// than you should think about it, and if you don't, just use it without thinking.
func (l *Logger) WarnF(msgFormat string, msgParams ...interface{}) {
	l.log(callDepth, WarnLevel, generateMessage(msgFormat, msgParams...))
}

// ErrorF will output msg as an error message.
// The msg is the return value of generateMessage.
// This is the better way to output a long log made from many variables.
// The msgFormat is the same as format in fmt.Printf, so you can use
// all format it supports, such as '%d'.
// msgParams is the params msgFormat needs, and it is variable-length, so
// you can add all your params here.
// You should know that this way to output msg is the most expensive way in time,
// but it's still faster than other logging libs. If you care about performance,
// than you should think about it, and if you don't, just use it without thinking.
func (l *Logger) ErrorF(msgFormat string, msgParams ...interface{}) {
	l.log(callDepth, ErrorLevel, generateMessage(msgFormat, msgParams...))
}
