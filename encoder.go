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
    "strconv"
    "sync"
    "time"
)

type Encoder interface {
    Encode(log *Log) []byte
}

// DefaultEncoder is an encoder which encodes log to a plain string.
// The log encoded by this encoder will be like the string "[Info] [2020-03-06 16:10:44] msg" in bytes.
type DefaultEncoder struct {
    timeFormatter *timeCachedFormatter
}

func NewDefaultEncoder(timeFormat string) Encoder {
    return &DefaultEncoder{
        timeFormatter: NewTimeCachedFormatter(timeFormat),
    }
}

// Encode encodes log to a plain string like "[Info] [2020-03-06 16:10:44] msg" in bytes.
func (de *DefaultEncoder) Encode(log *Log) []byte {
    return []byte("[" + log.Level().String() + "] [" + de.timeFormatter.format(log.Now()) + "] " + log.Msg() + "\n")
}

// JsonEncoder is an encoder which encodes log to a Json string.
// The log encoded by this encoder will be like the string `{"level":"debug", "time":"2020-03-22 22:35:00", "msg":"log content..."}` in bytes.
type JsonEncoder struct {
    needFormattingTime bool
    timeFormatter      *timeCachedFormatter
}

func NewJsonEncoder(timeFormat string, needFormattingTime bool) Encoder {
    return &JsonEncoder{
        needFormattingTime: needFormattingTime,
        timeFormatter:      NewTimeCachedFormatter(timeFormat),
    }
}

// Encode encodes log to a Json string like `{"level":"debug", "time":"2020-03-22 22:35:00", "msg":"log content..."}` in bytes.
func (je *JsonEncoder) Encode(log *Log) []byte {

    if je.needFormattingTime {
        return []byte(`{"level":"` + log.Level().String() + `", "time":"` + je.timeFormatter.format(log.Now()) + `", "msg":"` + log.Msg() + `"}` + "\n")
    }
    return []byte(`{"level":"` + log.Level().String() + `", "time":` + strconv.FormatInt(log.Now().Unix(), 10) + `, "msg":"` + log.Msg() + `"}` + "\n")
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
