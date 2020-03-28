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
// Created at 2020/03/25 21:59:32

package logit

import (
    "bytes"
    "strconv"
    "sync"
    "time"
)

// Encoder is how to encode a log to be writable.
// There are two encoders for you:
//     1. DefaultEncoder
//     2. JsonEncoder
//
// DefaultEncoder encodes a log to a plain string like "[Info] [2020-03-06 16:10:44] msg" in bytes.
// JsonEncoder encodes a log to a Json string like `{"level":"debug", "time":"2020-03-22 22:35:00", "msg":"log content..."}` in bytes.
// Of cause, you can implement Encoder interface to do you encoding job in you own way.
type Encoder interface {

    // Encode encode a log to bytes.
    Encode(log *Log) []byte
}

// DefaultEncoder is an encoder which encodes log to a plain string.
// The log encoded by this encoder will be like the string "[Info] [2020-03-06 16:10:44] msg" in bytes.
type DefaultEncoder struct {
    timeFormatter *timeCachedFormatter
}

// NewDefaultEncoder returns a DefaultEncoder with given timeFormat.
func NewDefaultEncoder(timeFormat string) Encoder {
    return &DefaultEncoder{
        timeFormatter: NewTimeCachedFormatter(timeFormat),
    }
}

// Encode encodes a log to a plain string like "[Info] [2020-03-06 16:10:44] msg" in bytes.
func (de *DefaultEncoder) Encode(log *Log) []byte {

    // 组装 log
    buffer := bytes.NewBuffer(make([]byte, 0, 64))
    buffer.WriteString("[")
    buffer.WriteString(log.Level().String())
    buffer.WriteString("] [")
    buffer.WriteString(de.timeFormatter.format(log.Now()))
    buffer.WriteString("] ")

    // 如果有文件信息，就把文件信息也加进去
    file, hasFile := log.extra["file"]
    line, hasLine := log.extra["line"]
    if hasFile && hasLine {
        buffer.WriteString("[")
        buffer.WriteString(file + ":" + line)
        buffer.WriteString("] ")
    }

    buffer.WriteString(log.Msg())
    buffer.WriteString("\n")
    return buffer.Bytes()
}

// JsonEncoder is an encoder which encodes log to a Json string.
// The log encoded by this encoder will be like the string `{"level":"debug", "time":"2020-03-22 22:35:00", "msg":"log content..."}` in bytes.
type JsonEncoder struct {
    needFormattingTime bool
    timeFormatter      *timeCachedFormatter
}

// NewJsonEncoder returns a JsonEncoder with given timeFormat.
// needFormattingTime is to check you want a formatted time or time in unix.
// Only needFormattingTime is true then timeFormat is valid.
func NewJsonEncoder(timeFormat string, needFormattingTime bool) Encoder {
    return &JsonEncoder{
        needFormattingTime: needFormattingTime,
        timeFormatter:      NewTimeCachedFormatter(timeFormat),
    }
}

// Encode encodes a log to a Json string like `{"level":"debug", "time":"2020-03-22 22:35:00", "msg":"log content..."}` in bytes.
func (je *JsonEncoder) Encode(log *Log) []byte {

    // 组装 log
    buffer := bytes.NewBuffer(make([]byte, 0, 64))
    buffer.WriteString(`{"level":"`)
    buffer.WriteString(log.Level().String())
    buffer.WriteString(`", "time":`)

    // 判断是否需要格式化时间
    if je.needFormattingTime {
        buffer.WriteString(strconv.Quote(je.timeFormatter.format(log.Now())))
    } else {
        buffer.WriteString(strconv.FormatInt(log.Now().Unix(), 10))
    }

    // 如果有文件信息，就把文件信息也加进去
    file, hasFile := log.extra["file"]
    line, hasLine := log.extra["line"]
    if hasFile && hasLine {
        buffer.WriteString(`, "file":"` + file)
        buffer.WriteString(`", "line":` + line)
    }

    buffer.WriteString(`, "msg":"`)
    buffer.WriteString(log.Msg())
    buffer.WriteString("\"}\n")
    return buffer.Bytes()
}

// **********************************************************
// For experiment.
// This is a time cache mechanism and you know that this is an experiment.
// We don't know if this mechanism is worth yet. You should know that time format operation
// takes lots of time, but concurrent competition does it, too. So is it worth to replace time
// format operation with concurrent competition? Only time will tell us!
// Notice that this struct might be removed someday, so be careful if you use it.
type timeCachedFormatter struct {
    timeInUnix    int64
    timeFormat    string
    timeFormatted string
    mu            *sync.Mutex
}

// NewTimeCachedFormatter returns a *timeCachedFormatter holder with given timeFormat.
// Notice that this method is for experiment, so you better not use it directly.
func NewTimeCachedFormatter(timeFormat string) *timeCachedFormatter {
    return &timeCachedFormatter{
        timeFormat: timeFormat,
        mu:         &sync.Mutex{},
    }
}

// format is for formatting time.
// Use a global formatted time to cache the current time
// for avoiding formatting too much times in the same second.
func (tcf *timeCachedFormatter) format(now time.Time) string {

    // 并发激烈的时候或许会成为瓶颈，反而得不偿失，这个需要经过大量试验才知道是否值得
    tcf.mu.Lock()
    defer tcf.mu.Unlock()

    // 使用 != 而不是 > 是为了防止时钟回拨时日志记录时间不正确的情况
    if tcf.timeInUnix != now.Unix() {
        tcf.timeInUnix = now.Unix()
        tcf.timeFormatted = now.Format(tcf.timeFormat)
    }
    return tcf.timeFormatted
}

// **********************************************************
