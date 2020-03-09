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
// Author: fish
// Email: fishinlove@163.com
// Created at 2020/02/29 15:39:02

package logit

import (
    "io"
    "runtime"
    "strconv"
    "sync"
    "time"
)

// Logger is a struct to log.
type Logger struct {

    // writer is the output of this Logger.
    writer io.Writer

    // handlers is the slice of log handlers.
    // You can add your handler for some situations.
    // See LoggerHandler.
    handlers []LoggerHandler

    // formatOfTime is the format for formatting time.
    // Default is "2006-01-02 15:04:05", see DefaultFormatOfTime.
    formatOfTime string

    // level is the level representation of the Logger.
    // In this version of logit, there are four levels:
    //
    //  DebugLevel, InfoLevel, WarnLevel, ErrorLevel.
    //
    // The righter level has higher visibility which means
    // one debug message will not be logged in one Logger in InfoLevel.
    // That's we called level-based logging.
    level LoggerLevel

    // running is the status of the Logger.
    // true means the Logger is running, false means the Logger is shutdown.
    running bool

    // needFileInfo is a flag to check if msg should contain file info.
    // This step is useful but too expensive, so default is false.
    needFileInfo bool

    // mu is for safe concurrency.
    mu sync.RWMutex
}

// DefaultFormatOfTime is the default format for formatting time.
const DefaultFormatOfTime = "2006-01-02 15:04:05"

// NewLogger creates one Logger with given out and level.
// The first parameter writer is the writer for logging.
// The second parameter level is the level of this Logger.
// It returns a new running Logger holder.
func NewLogger(writer io.Writer, level LoggerLevel) *Logger {
    return NewLoggerWithHandlers(writer, level, DefaultLoggerHandler)
}

// NewLoggerWithHandlers creates one Logger with given out and level and handlers.
// The first parameter writer is the writer for logging.
// The second parameter level is the level of this Logger.
// The third parameter handlers is all logger handlers for handling each log.
// It returns a new running Logger holder.
func NewLoggerWithHandlers(writer io.Writer, level LoggerLevel, handlers ...LoggerHandler) *Logger {

    // 至少添加一个日志处理器，否则直接报错
    if len(handlers) < 1 {
        panic("You must add at least one handler!")
    }

    return &Logger{
        writer:       writer,
        formatOfTime: DefaultFormatOfTime,
        handlers:     handlers,
        level:        level,
        running:      true,
        needFileInfo: false,
    }
}

// Enable sets l on running status.
func (l *Logger) Enable() {
    l.mu.Lock()
    defer l.mu.Unlock()
    l.running = true
}

// Disable sets l on shutdown status.
func (l *Logger) Disable() {
    l.mu.Lock()
    defer l.mu.Unlock()
    l.running = false
}

// ChangeLevelTo will change the level of current Logger to newLevel.
func (l *Logger) ChangeLevelTo(newLevel LoggerLevel) {
    l.mu.Lock()
    defer l.mu.Unlock()
    l.level = newLevel
}

// Level returns the logger level of l.
func (l *Logger) Level() LoggerLevel {
    l.mu.RLock()
    defer l.mu.RUnlock()
    return l.level
}

// EnableFileInfo means every log will contain file info like line number.
// However, you should know that this is expensive in time.
// So be sure you really need it or keep it disabled.
func (l *Logger) EnableFileInfo() {
    l.mu.Lock()
    defer l.mu.Unlock()
    l.needFileInfo = true
}

// DisableFileInfo means every log will not contain file info like line number.
// If you want file info again, try l.EnableFileInfo().
func (l *Logger) DisableFileInfo() {
    l.mu.Lock()
    defer l.mu.Unlock()
    l.needFileInfo = false
}

// AddHandlers adds more handlers to l, and all handlers added before will be retained.
// If you want to remove all handlers, try l.SetHandlers().
// See logit.DefaultLoggerHandler.
func (l *Logger) AddHandlers(handlers ...LoggerHandler) {
    l.mu.Lock()
    defer l.mu.Unlock()
    l.handlers = append(l.handlers, handlers...)
}

// SetHandlers replaces l.handlers with handlers, all handlers added before will be removed.
// If you want to add more handlers rather than replace them, try l.AddHandlers.
// Notice that at least one handler should be added, so if len(handlers) < 1, it returns false
// which means setting failed. Return true if setting is successful.
// See logit.DefaultLoggerHandler.
func (l *Logger) SetHandlers(handlers ...LoggerHandler) bool {

    // 必须添加至少一个处理器
    if len(handlers) < 1 {
        return false
    }

    // 先清空原本的日志处理器，再添加新的日志处理器
    l.mu.Lock()
    defer l.mu.Unlock()
    l.handlers = nil
    l.handlers = append(l.handlers, handlers...)
    return true
}

// Handlers returns all handlers of l in a copy slice.
func (l *Logger) Handlers() []LoggerHandler {
    l.mu.RLock()
    defer l.mu.RUnlock()

    // 返回的是日志处理器的副本，防止被非法篡改
    handlers := make([]LoggerHandler, 0, len(l.handlers))
    return append(handlers, l.handlers...)
}

// SetFormatOfTime sets format of time as you want.
// Default is "2006-01-02 15:04:05", see DefaultFormatOfTime.
func (l *Logger) SetFormatOfTime(formatOfTime string) {
    l.mu.Lock()
    defer l.mu.Unlock()
    l.formatOfTime = formatOfTime
}

// FormatOfTime returns the format of time in l.
func (l *Logger) FormatOfTime() string {
    l.mu.RLock()
    defer l.mu.RUnlock()
    return l.formatOfTime
}

// ChangeWriterTo changes current writer to newWriter.
func (l *Logger) ChangeWriterTo(newWriter io.Writer) {
    l.mu.Lock()
    defer l.mu.Unlock()
    l.writer = newWriter
}

// Writer returns the writer of l.
func (l *Logger) Writer() io.Writer {
    l.mu.RLock()
    defer l.mu.RUnlock()
    return l.writer
}

// callDepth is the depth of the method calling stack, which is about file name and line.
const callDepth = 3

// log can output msg to l.writer, notices that level will affect the visibility of this msg.
// Notice that callDepth is caller sensitive, and the value is about file name and line.
func (l *Logger) log(callDepth int, level LoggerLevel, msg string) {

    // 加上读锁
    l.mu.RLock()

    // 以下两种条件直接返回，不记录日志：
    // 1. 日志处于禁用状态，也就是 l.running = false
    // 2. 日志记录器的日志级别高于这条记录的日志级别
    if !l.running || l.level > level {
        l.mu.RUnlock()
        return
    }

    // 提前释放读锁，后续操作非常消耗时间，可以不用加锁了，彻底释放并发的天性
    // 但是 needFileInfo 的获取需要保证并发安全，就在释放锁之前拷贝一份副本，
    // 即使释放锁之后有人修改了这个属性，也和这里无关了，因为在执行这个 log 方法的时间点上
    // 这个属性的值就已经确定了，并且不允许被修改了，这类似于 copy on write 的解决思路。
    // 这个解决并发竞争的方案是否没有问题，需要时间的验证才知道
    needFileInfo := l.needFileInfo
    l.mu.RUnlock()

    // 如果需要文件信息，对当前的 msg 进行包装
    if needFileInfo {
        msg = wrapMessageWithFileInfo(callDepth, msg)
    }

    // 处理日志
    l.handleLog(level, time.Now(), msg)
}

// handleLog handles log with l.handlers.
// Notice that if one handler returns false, then all handlers after it
// will not use anymore.
func (l *Logger) handleLog(level LoggerLevel, now time.Time, msg string) {
    for _, handler := range l.handlers {
        if !handler.handle(l, level, now, msg) {
            break
        }
    }
}

// wrapMessageWithFileInfo wraps msg with file info.
// This function is too expensive because of runtime.Caller.
// Notice that callDepth is the depth of calling stack. See callDepth.
func wrapMessageWithFileInfo(callDepth int, msg string) string {

    // 这个 callDepth 是 runtime.Caller 方法的参数，表示上面第几层调用者信息
    _, file, line, ok := runtime.Caller(callDepth)
    if !ok {
        return "[unknown file:unknown line] " + msg
    }

    return "[" + file + ":" + strconv.Itoa(line) + "] " + msg
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
