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
// Created at 2020/04/23 23:03:56

package logit

import (
	"io"
)

const (
	// DefaultTimeFormat is the default format for formatting time.
	DefaultTimeFormat = "2006-01-02 15:04:05"
)

func init() {
	// 注册日志处理器
	RegisterHandler("default", func(params map[string]interface{}) Handler {

		// 如果时间格式参数没有传递，就使用默认的时间格式
		timeFormat := DefaultTimeFormat
		if format, ok := params["timeFormat"]; ok && format.(string) != "" {
			timeFormat = format.(string)
		}

		// 如果编码器参数没有传递，就使用 text 编码器
		encoder := EncoderOf("text")
		if encoderName, ok := params["encoder"]; ok && encoderName.(string) != "" {
			encoder = EncoderOf(encoderName.(string))
		}

		return NewDefaultHandler(WriterOf(params), encoder, timeFormat)
	})
}

// DefaultHandler is a default handler for use.
// Generally speaking, encoding a log to bytes then written by writer is the most of
// handlers do. So we provide a default handler, which only need a writer and an encoder.
//
// For config:
//     If you want to use this handler in your logger by config file, try this:
//
//         "handlers":{
//             "default":{
//
//             }
//         }
//
//     This will use logit.DefaultTimeFormat to format time, and if you want to
//     use your layout to format time, try this:
//
//         "handlers":{
//             "default":{
//                 "timeFormat": "2006-01-02"
//             }
//         }
//
type DefaultHandler struct {
	writer     io.Writer
	encoder    Encoder
	timeFormat string
}

// NewDefaultHandlerWithoutEncoder returns a DefaultHandler holder with given writer.
func NewDefaultHandlerWithoutEncoder(writer io.Writer, timeFormat string) Handler {
	return &DefaultHandler{
		writer:     writer,
		encoder:    EncoderOf("text"),
		timeFormat: timeFormat,
	}
}

// NewDefaultHandler returns a DefaultHandler holder with given writer and encoder.
func NewDefaultHandler(writer io.Writer, encoder Encoder, timeFormat string) Handler {
	return &DefaultHandler{
		writer:     writer,
		encoder:    encoder,
		timeFormat: timeFormat,
	}
}

// Handle will encode log and write log by internal writer.
// Return true so that handlers after it will be used.
func (dh *DefaultHandler) Handle(log *Log) bool {
	dh.writer.Write(dh.encoder.Encode(log, dh.timeFormat))
	return true
}
