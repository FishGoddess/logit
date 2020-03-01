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
    "log"
    "sync"
)

// Logger is a struct based on "log.Logger".
type Logger struct {

    // logger is an inside logger for really logging.
    // Any operations do on the Logger will finally do on this logger.
    logger *log.Logger

    // level is the level representation of the Logger.
    // In this version of logit, there are four levels:
    //
    //  DebugLevel, InfoLevel, WarningLevel, ErrorLevel.
    //
    // The righter level has higher visibility which means
    // one debug message will not be logged in one Logger in InfoLevel.
    // That's we called level-based logging.
    level LogLevel

    // running is the status of the Logger.
    // true means the Logger is running, false means the Logger is shutdown.
    running bool

    // mu is for safe concurrency.
    mu sync.RWMutex
}

// NewLogger create one Logger with given out and level.
// The first parameter out is a writer for logging.
// The second parameter level is the level of this Logger.
// It returns a new running Logger holder.
func NewLogger(out io.Writer, level LogLevel) *Logger {
    return &Logger{
        logger:  log.New(out, "", log.LstdFlags|log.Lshortfile),
        level:   level,
        running: true,
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
func (l *Logger) ChangeLevelTo(newLevel LogLevel) {
    l.mu.Lock()
    defer l.mu.Unlock()
    l.level = newLevel
}

// log can output msg to l.writer, notices that level will affect the visibility of this msg.
func (l *Logger) log(level LogLevel, msg string) {

    // 加上读锁
    l.mu.RLock()

    // 以下两种条件直接返回，不记录日志：
    // 1. 日志处于禁用状态，也就是 l.running = false
    // 2. 日志记录器的日志级别高于这条记录的日志级别
    if !l.running || l.level > level {
        // 释放读锁
        l.mu.RUnlock()
        return
    }

    // 提前释放读锁，后续操作不需要加锁
    l.mu.RUnlock()

    // 记录日志
    // 这个 3 是 runtime.Caller 方法的参数，表示上面三层调用者信息
    // 第 0 层是 l.logger.Output 里面调用的 runtime.Caller 的那一行代码
    // 第 1 层是 l.logger.Output 这一行代码
    // 第 2 层是调用这个 log 方法那一层
    // 第 3 层是调用这个 log 方法的那个方法的外部调用上一层，比如调用 Debug 方法的那一层
    l.logger.Output(3, prefixOf(level)+msg)
}

// Debug will output msg as a debug message.
func (l *Logger) Debug(msg string) {
    l.log(DebugLevel, msg)
}

// Info will output msg as an info message.
func (l *Logger) Info(msg string) {
    l.log(InfoLevel, msg)
}

// Warning will output msg as a warning message.
func (l *Logger) Warning(msg string) {
    l.log(WarningLevel, msg)
}

// Error will output msg as an error message.
func (l *Logger) Error(msg string) {
    l.log(ErrorLevel, msg)
}
