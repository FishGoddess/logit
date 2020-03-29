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
    "bytes"
    "io"
    "strconv"
)

// Handler is an interface representation of log handler.
// Every log will be handled by handler, and you can customize your own handler
// to handle logs in your way. The return value is meaningful, false means
// next handler will not be used, only true will go on handling process.
// Notice that if one handler returns false, then all handlers after it
// will not be used anymore.
type Handler interface {

    // Handle should handle this log in someway.
    // If you don't want next handler to be used, just return false.
    // Then all handlers after current handler will not be used.
    Handle(log *Log) bool
}

// EncodeToText encodes a log to a plain string like "[Info] [2020-03-06 16:10:44] msg" in bytes.
func EncodeToText(log *Log, timeFormat string) []byte {

    // 组装 log
    buffer := bytes.NewBuffer(make([]byte, 0, 64))
    buffer.WriteString("[")
    buffer.WriteString(log.Level().String())
    buffer.WriteString("] [")
    buffer.WriteString(log.Now().Format(timeFormat))
    buffer.WriteString("] ")

    // 如果有文件信息，就把文件信息也加进去
    if log.file != "" && log.Line() != 0 {
        buffer.WriteString("[")
        buffer.WriteString(log.File() + ":" + strconv.Itoa(log.Line()))
        buffer.WriteString("] ")
    }

    buffer.WriteString(log.Msg())
    buffer.WriteString("\n")
    return buffer.Bytes()
}

// EncodeToJson encodes a log to a Json string like `{"level":"debug", "time":"2020-03-22 22:35:00", "msg":"log content..."}` in bytes.
// If timeFormat == "", then it will not format time and keep time in unix form.
func EncodeToJson(log *Log, timeFormat string) []byte {

    // 组装 log
    buffer := bytes.NewBuffer(make([]byte, 0, 64))
    buffer.WriteString(`{"level":"`)
    buffer.WriteString(log.Level().String())
    buffer.WriteString(`", "time":`)

    // 判断是否需要格式化时间
    if timeFormat != "" {
        buffer.WriteString(strconv.Quote(log.Now().Format(timeFormat)))
    } else {
        buffer.WriteString(strconv.FormatInt(log.Now().Unix(), 10))
    }

    // 如果有文件信息，就把文件信息也加进去
    if log.file != "" && log.Line() != 0 {
        buffer.WriteString(`, "file":"` + log.File())
        buffer.WriteString(`", "line":` + strconv.Itoa(log.Line()))
    }

    buffer.WriteString(`, "msg":"`)
    buffer.WriteString(log.Msg())
    buffer.WriteString("\"}\n")
    return buffer.Bytes()
}

// DefaultHandler is a default handler for use.
// Generally speaking, encoding a log to bytes then written by writer is the most of
// handlers do. So we provide a default handler, which only need a writer and an encoder.
type DefaultHandler struct {
    writer     io.Writer
    timeFormat string
}

// NewDefaultHandler returns a DefaultHandler holder with given writer.
func NewDefaultHandler(writer io.Writer, timeFormat string) Handler {
    return &DefaultHandler{
        writer:     writer,
        timeFormat: timeFormat,
    }
}

// Handle will encode log and write log by internal writer.
// Return true so that handlers after it will be used.
func (dh *DefaultHandler) Handle(log *Log) bool {
    dh.writer.Write(EncodeToText(log, dh.timeFormat))
    return true
}

// JsonHandler is a json handler for use.
type JsonHandler struct {
    writer     io.Writer
    timeFormat string
}

// NewJsonHandler returns a JsonHandler holder with given writer.
// If timeFormat == "", then it will not format time and keep time in unix form.
func NewJsonHandler(writer io.Writer, timeFormat string) Handler {
    return &JsonHandler{
        writer:     writer,
        timeFormat: timeFormat,
    }
}

// Handle will encode log and write log by internal writer.
// Return true so that handlers after it will be used.
func (jh *JsonHandler) Handle(log *Log) bool {
    jh.writer.Write(EncodeToJson(log, jh.timeFormat))
    return true
}
