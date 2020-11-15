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

	// encoder is used to encode a log to bytes.
	encoder Encoder

	// writer is used to output an encoded log.
	writer io.Writer

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
		level:      InfoLevel,
		encoder:    TextEncoder(),
		writer:     os.Stdout,
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

// ChangeLevelTo will change the logger level of current logger to newLevel.
// It returns old level of current logger.
func (l *Logger) ChangeLevelTo(newLevel Level) Level {
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

// EnableFileInfo means every log will contain file info like line number.
// However, you should know that this is expensive in time.
// So be sure you really need it or keep it disabled.
func (l *Logger) EnableFileInfo() {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.needCaller = true
}

// DisableFileInfo means every log will not contain file info like line number.
// If you want file info again, try l.EnableFileInfo().
func (l *Logger) DisableFileInfo() {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.needCaller = false
}

// newLog returns a Log holder from object pool.
// Notice that not every holder returned is new, as you know, that is why we use a pool.
func (l *Logger) newLog(level Level, msg string) *Log {
	log := l.logs.Get().(*Log)
	log.level = level
	log.time = time.Now()
	log.msg = msg
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

const (
	// callDepth is the depth of the method calling stack, which is about file name and line number.
	callDepth = 3
)

// log handles msg by l.handlers, and level will affect the visibility of this msg.
// Notice that callDepth is caller sensitive.
func (l *Logger) log(callDepth int, level Level, msg string) {

	// 加上读锁
	l.lock.RLock()

	// 日志记录器的级别高于日志的级别，不进行记录
	if l.level > level {
		l.lock.RUnlock()
		return
	}

	// 提前释放读锁，后续操作非常消耗时间，可以不用加锁了，彻底释放并发的天性
	// 但是 needCaller 的获取需要保证并发安全，就在释放锁之前拷贝一份副本
	// 即使释放锁之后有人修改了这个属性，也和这里无关了，因为在执行这个 log 方法的时间点上，
	// 这个属性的值就已经确定了，并且不允许被修改了，这类似于 copy on write 的解决思路
	// 这个解决并发竞争的方案是否没有问题，需要时间的验证才知道
	needCaller := l.needCaller
	l.lock.RUnlock()

	// 处理日志
	log := l.newLog(level, msg)
	defer l.releaseLog(log)

	// 如果需要调用者的信息，对当前的 msg 进行包装
	if needCaller {
		wrapLogWithCaller(callDepth, log)
	}
	l.handleLog(log)
}

// handleLog handles log with l.handlers.
// Notice that if one handler returns false, then all handlers after it
// will not be used anymore.
func (l *Logger) handleLog(log *Log) {
	// TODO handle...
}

// wrapLogWithCaller wraps log with caller info.
// This function is too expensive because of runtime.Caller.
// Notice that callDepth is the depth of calling stack. See callDepth.
func wrapLogWithCaller(callDepth int, log *Log) {

	log.caller = &caller{
		File: "unknown file",
		Line: -1,
	}

	// 这个 callDepth 是 runtime.Caller 方法的参数，表示要获取第几层调用者的信息
	if _, file, line, ok := runtime.Caller(callDepth); ok {
		log.caller.File = file
		log.caller.Line = line
	}
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

// ================================== extension ==================================

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
