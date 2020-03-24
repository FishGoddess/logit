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
// Created at 2020/03/06 13:36:28

package logit

import (
    "os"
    "sync"
    "time"
)

type Handler struct {
    Encode func(logger *Logger, level Level, now time.Time, msg string) []byte

    Output func(logger *Logger, data []byte) bool
}

func (h *Handler) handle(logger *Logger, level Level, now time.Time, msg string) bool {
    return h.Output(logger, h.Encode(logger, level, now, msg))
}

func StandardEncode(logger *Logger, level Level, now time.Time, msg string) []byte {
    return []byte("[" + level.String() + "] [" + formatTime(now, logger.FormatOfTime()) + "] " + msg + "\n")
}

func JsonEncode(logger *Logger, level Level, now time.Time, msg string) []byte {
    return []byte(`{"level":"` + level.String() + `", "time":"` + formatTime(now, logger.FormatOfTime()) + `", "msg":"` + msg + `"}` + "\n")
}

// TODO 使用闭包来实现获取 logger.Writer
func StandardOutput(logger *Logger, data []byte) bool {
    logger.Writer().Write(data)
    return true
}

func ConsoleOutput(logger *Logger, data []byte) bool {
    os.Stdout.Write(data)
    return true
}

// **************************************************************************
// LoggerHandler is a struct representation of log handler.
// Every log will handle by this handler, and you can customize your own handler
// to handle logs in your way. The return value is meaningful, false means
// next handler will not be handled, only true will go on handling process.
// Notice that if one handler returns false, then all handlers after it
// will not use anymore.
type LoggerHandler func(logger *Logger, level Level, now time.Time, msg string) bool

// handle will call lh as a function. It is just a proxy method.
func (lh LoggerHandler) handle(logger *Logger, level Level, now time.Time, msg string) bool {
    return lh(logger, level, now, msg)
}

// DefaultLoggerHandler is the default handler in logit.
// The log handled by this handler will be like "[Info] [2020-03-06 16:10:44] msg".
// If you want to customize, just code your own handler, then replace it!
func DefaultLoggerHandler(logger *Logger, level Level, now time.Time, msg string) bool {
    //logger.Writer().Write([]byte("[" + PrefixOf(level) + "] [" + now.Format(logger.FormatOfTime()) + "] " + msg + "\n"))
    logger.Writer().Write([]byte("[" + level.String() + "] [" + formatTime(now, logger.FormatOfTime()) + "] " + msg + "\n"))
    return true
}

// JsonLoggerHandler is the handler which handles log as a Json string.
// The log handled by this handler will be like `{"level":"debug", "time":"2020-03-22 22:35:00", "msg":"log content..."}`.
func JsonLoggerHandler(logger *Logger, level Level, now time.Time, msg string) bool {
    //logger.Writer().Write([]byte(`{"level":"` + PrefixOf(level) + `", "time":"` + now.Format(logger.FormatOfTime()) + `", "msg":"` + msg + `"}` + "\n"))
    logger.Writer().Write([]byte(`{"level":"` + level.String() + `", "time":"` + formatTime(now, logger.FormatOfTime()) + `", "msg":"` + msg + `"}` + "\n"))
    return true
}

// **********************************************************
// For experiment.
// This is a time cache mechanism and you know that this is an experiment.
// We don't know if this mechanism is worth yet. You should know that time format operation
// takes lots of time, but concurrent competition does it, too. So is it worth to replace time
// format operation with concurrent competition? Only time will tell us!
var nowInUnix = time.Unix(0, 0).Unix()
var nowFormatOfTime = "0"
var nowFormatted = "0"
var mutexForNow = &sync.Mutex{}

// formatTime is for formatting time.
// Use a global formatted time to cache the current time
// for avoiding formatting too much times in the same second.
func formatTime(now time.Time, formatOfTime string) string {
    mutexForNow.Lock()
    defer mutexForNow.Unlock()

    // 使用 != 而不是 > 是为了防止时钟回拨时日志记录时间不正确的情况
    if nowInUnix != now.Unix() || nowFormatOfTime != formatOfTime {
        nowInUnix = now.Unix()
        nowFormatOfTime = formatOfTime
        nowFormatted = now.Format(formatOfTime)
    }
    return nowFormatted
}

// **********************************************************
