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
	"io"
	"os"
	"runtime"
	"sync"
	"time"
)

// Logger is the core type of logit. All functions is provided by it.
type Logger struct {

	// level is the level representation of the Logger.
	// In this version of logit, there are five levels:
	//
	//  DebugLevel, InfoLevel, WarnLevel, ErrorLevel, OffLevel.
	//
	// The righter level has higher visibility which means
	// one debug message will not be logged in one Logger in InfoLevel.
	// That's we called level-based logging.
	//
	// In particular, OffLevel is the highest level, so if you set one
	// logger to OffLevel, it will never log anything.
	level Level

	// handlers is the slice of log handlers.
	// You can add your handler for some situations.
	// See logit.Handler.
	handlers []Handler

	// needCaller is a flag to check if msg should contain caller info.
	// This step is useful but too expensive, so default is false.
	needCaller bool

	// logs is an object pool cache some Log holders.
	// Use a pool is for reducing memory allocation.
	logs *sync.Pool

	// mu is for safe concurrency.
	mu *sync.RWMutex
}

// NewLogger creates a logger with given level and handlers.
// The level is the level of this Logger.
// The handlers is all logger handlers for handling each log.
// Notice that at least one handler should be added or a panic will happen.
func NewLogger(level Level, handlers ...Handler) *Logger {

	// 至少添加一个日志处理器，否则直接报错
	if len(handlers) < 1 {
		panic("You must add at least one handler!")
	}

	// 创建 logger 对象
	logger := &Logger{
		level:      level,
		handlers:   handlers,
		needCaller: false,
		mu:         &sync.RWMutex{},
	}

	// 初始化 logs 对象池
	logger.logs = &sync.Pool{
		New: func() interface{} {
			return &Log{
				logger: logger,
			}
		},
	}

	return logger
}

// ChangeLevelTo will change the logger level of current logger to newLevel.
// It returns old level of current logger.
func (l *Logger) ChangeLevelTo(newLevel Level) Level {
	l.mu.Lock()
	defer l.mu.Unlock()
	oldLevel := l.level
	l.level = newLevel
	return oldLevel
}

// Level returns the logger level of current Logger.
func (l *Logger) Level() Level {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.level
}

// AddHandlers adds more handlers to current logger, and all handlers added before
// will be retained. If you want to remove all handlers, try l.SetHandlers().
// See logit.Handler.
func (l *Logger) AddHandlers(handlers ...Handler) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.handlers = append(l.handlers, handlers...)
}

// SetHandlers replaces l.handlers with newHandlers, all handlers added before will be removed.
// If you want to add more handlers rather than replace them, try l.AddHandlers.
// Notice that at least one handler should be added, so if len(newHandlers) < 1, it returns false
// which means setting failed. Return true if setting is successful.
// See logit.Handler.
func (l *Logger) SetHandlers(newHandlers ...Handler) bool {

	// 必须添加至少一个处理器
	if len(newHandlers) < 1 {
		return false
	}

	// 先清空原本的日志处理器，再添加新的日志处理器
	l.mu.Lock()
	defer l.mu.Unlock()
	l.handlers = nil
	l.handlers = append(l.handlers, newHandlers...)
	return true
}

// Handlers returns all handlers of current logger in a copy slice.
func (l *Logger) Handlers() []Handler {
	l.mu.RLock()
	defer l.mu.RUnlock()

	// 返回的是日志处理器的副本，防止被非法篡改
	handlers := make([]Handler, 0, len(l.handlers)+2)
	return append(handlers, l.handlers...)
}

// EnableFileInfo means every log will contain file info like line number.
// However, you should know that this is expensive in time.
// So be sure you really need it or keep it disabled.
func (l *Logger) EnableFileInfo() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.needCaller = true
}

// DisableFileInfo means every log will not contain file info like line number.
// If you want file info again, try l.EnableFileInfo().
func (l *Logger) DisableFileInfo() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.needCaller = false
}

// newLog returns a Log holder from object pool.
// Notice that not every holder returned is new, as you know, that is why we use a pool.
func (l *Logger) newLog(level Level, msg string) *Log {
	log := l.logs.Get().(*Log)
	log.level = level
	log.now = time.Now()
	log.msg = msg
	return log
}

// releaseLog releases log to object pool so that this log can be reused next time.
func (l *Logger) releaseLog(log *Log) {
	log.file = ""
	log.line = 0
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
	l.mu.RLock()

	// 日志记录器的级别高于日志的级别，不进行记录
	if l.level > level {
		l.mu.RUnlock()
		return
	}

	// 提前释放读锁，后续操作非常消耗时间，可以不用加锁了，彻底释放并发的天性
	// 但是 needCaller 的获取需要保证并发安全，就在释放锁之前拷贝一份副本
	// 即使释放锁之后有人修改了这个属性，也和这里无关了，因为在执行这个 log 方法的时间点上，
	// 这个属性的值就已经确定了，并且不允许被修改了，这类似于 copy on write 的解决思路
	// 这个解决并发竞争的方案是否没有问题，需要时间的验证才知道
	needCaller := l.needCaller
	l.mu.RUnlock()

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
	for _, handler := range l.handlers {
		if !handler.Handle(log) {
			return
		}
	}
}

// wrapLogWithCaller wraps log with caller info.
// This function is too expensive because of runtime.Caller.
// Notice that callDepth is the depth of calling stack. See callDepth.
func wrapLogWithCaller(callDepth int, log *Log) {

	// 这个 callDepth 是 runtime.Caller 方法的参数，表示要获取第几层调用者的信息
	_, file, line, ok := runtime.Caller(callDepth)
	if !ok {
		log.file = "unknown file"
		log.line = -1
	}

	log.file = file
	log.line = line
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

// NewLoggerFrom returns a logger parsed from reader which returns a config.
// The reader is the reader of your config. A config looks like:
//
//     "level": "debug",
//     "caller": false,
//     "handlers": {
//         "console": {
//             # I am a comment...
//         }
//     }
//
// Check config file templates to know about more information.
func NewLoggerFrom(reader io.Reader) *Logger {

	// 解析配置
	conf, err := parseConfigFrom(reader)
	if err != nil {
		panic(err)
	}

	// 根据配置创建并初始化 logger
	logger := NewLogger(parseLevel(conf.Level), parseHandlersFrom(conf)...)
	if conf.Caller {
		logger.EnableFileInfo()
	}
	return logger
}

// NewLoggerFromPath returns a logger parsed from config file.
// pathOfConfigFile is the path of your config file. A config file
// is a file like "xxx.conf", and its content looks like:
//
//     "level": "debug",
//     "caller": false,
//     "handlers": {
//         "console": {
//             # I am a comment...
//         }
//     }
//
// Check config file templates to know about more information.
func NewLoggerFromPath(pathOfConfigFile string) *Logger {

	// 打开配置文件
	file, err := os.Open(pathOfConfigFile)
	if err != nil {
		panic(err)
	}

	return NewLoggerFrom(file)
}

// DebugFunc will output msg as a debug message.
// The msg is the return value of msgGenerator.
// This is the better way to output a long log made from many variables.
func (l *Logger) DebugFunc(msgGenerator func() string) {
	l.log(callDepth, DebugLevel, msgGenerator())
}

// InfoFunc will output msg as an info message.
// The msg is the return value of msgGenerator.
// This is the better way to output a long log made from many variables.
func (l *Logger) InfoFunc(msgGenerator func() string) {
	l.log(callDepth, InfoLevel, msgGenerator())
}

// WarnFunc will output msg as a warn message.
// The msg is the return value of msgGenerator.
// This is the better way to output a long log made from many variables.
func (l *Logger) WarnFunc(msgGenerator func() string) {
	l.log(callDepth, WarnLevel, msgGenerator())
}

// ErrorFunc will output msg as an error message.
// The msg is the return value of msgGenerator.
// This is the better way to output a long log made from many variables.
func (l *Logger) ErrorFunc(msgGenerator func() string) {
	l.log(callDepth, ErrorLevel, msgGenerator())
}
