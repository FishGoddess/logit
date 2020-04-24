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
// Created at 2020/04/24 10:51:51

package logit

import (
	"os"
)

func init() {
	// 注册日志处理器
	RegisterHandler("console", func(params map[string]interface{}) Handler {

		// 如果时间格式参数没有传递，就使用默认的时间格式
		timeFormat := DefaultTimeFormat
		if format, ok := params["timeFormat"]; ok && format.(string) != "" {
			timeFormat = format.(string)
			// 如果参数是 unix，则直接使用空字符串
			if timeFormat == "unix" {
				timeFormat = ""
			}
		}

		// 如果编码器参数没有传递，就使用 text 编码器
		encoder := EncoderOf("text")
		if encoderName, ok := params["encoder"]; ok && encoderName.(string) != "" {
			encoder = EncoderOf(encoderName.(string))
		}

		return NewConsoleHandler(encoder, timeFormat)
	})
}

// ConsoleHandler is a console handler for use.
// Actually, output a log to console is the most of things loggers do when developing.
// So we provide a console handler, which only need an encoder.
// 
// For config:
//     If you want to use this handler in your logger by config file, try this:
//
//         "handlers":{
//             "console":{
//
//             }
//         }
//
//     It will use logit.DefaultTimeFormat to format time in default, so if you want to
//     use your layout to format time, try this:
//
//         "handlers":{
//             "console":{
//                 "timeFormat": "2006-01-02"
//             }
//         }
//
//     Want a json string? Try this:
//
//         "handlers":{
//             "console":{
//                 "timeFormat": "2006-01-02",
//                 "encoder": "json"
//             }
//         }
//
type ConsoleHandler struct {
	encoder    Encoder
	timeFormat string
}

// NewConsoleHandler returns a ConsoleHandler holder with given encoder.
func NewConsoleHandler(encoder Encoder, timeFormat string) Handler {
	return &ConsoleHandler{
		encoder:    encoder,
		timeFormat: timeFormat,
	}
}

// Handle will encode log and write log by stdout or stderr.
// Return true so that handlers after it will be used.
// Notice that an error log will be written to stderr.
func (ch *ConsoleHandler) Handle(log *Log) bool {
	// 错误日志通过 stderr 进行输出
	if log.Level() == ErrorLevel {
		os.Stderr.Write(ch.encoder.Encode(log, ch.timeFormat))
		return true
	}

	// 非错误日志通过 stdout 进行输出
	os.Stdout.Write(ch.encoder.Encode(log, ch.timeFormat))
	return true
}
