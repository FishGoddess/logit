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
// Created at 2020/04/24 23:37:56

package logit

import (
	"github.com/FishGoddess/logit/writer"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func init() {
	registerConsoleHandler()
	registerFileHandler()
	registerDurationRollingHandler()
	registerSizeRollingHandler()
}

// encoderAndTimeFormatOf returns an encoder and time format of this params.
func encoderAndTimeFormatOf(params map[string]interface{}) (Encoder, string) {

	// 如果编码器参数没有传递，就使用 text 编码器
	encoder := EncoderOf("text")
	if encoderName, ok := params["encoder"]; ok && strings.TrimSpace(encoderName.(string)) != "" {
		encoder = EncoderOf(encoderName.(string))
	}

	// 如果时间格式参数没有传递，就使用默认的时间格式
	timeFormat := DefaultTimeFormat
	if format, ok := params["timeFormat"]; ok && strings.TrimSpace(format.(string)) != "" {
		timeFormat = format.(string)
		// 如果参数是 unix，则直接使用空字符串
		if timeFormat == "unix" {
			timeFormat = ""
		}
	}

	return encoder, timeFormat
}

func limitAndDirectoryOf(params map[string]interface{}, defaultDuration int) io.Writer {

	// 滚动的时间间隔，单位是秒
	duration := 24 * 60 * 60 // 默认是一天
	if param, ok := params["limit"]; ok {
		duration = int(param.(float64))
	}

	// 保存日志的目标文件夹
	directory := "./" // 默认是当前目录
	if param, ok := params["directory"]; ok {
		directory = param.(string)
	}

	return writer.NewDurationRollingFile(time.Duration(duration)*time.Second, writer.NextFilename(directory))
}

// =============================== console handler ===============================

// registerConsoleHandler registers console handler.
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
func registerConsoleHandler() {
	RegisterHandler("console", func(params map[string]interface{}) Handler {
		encoder, timeFormat := encoderAndTimeFormatOf(params)
		return NewConsoleHandler(encoder, timeFormat)
	})
}

// =============================== file handler ===============================

func pathOf(params map[string]interface{}) string {

	// 写出的目标文件
	path := "./logit-" + strconv.FormatInt(time.Now().Unix(), 10) + writer.SuffixOfLogFile
	if param, ok := params["path"]; ok && strings.TrimSpace(param.(string)) != "" {
		path = param.(string)
	}

	return path
}

// registerFileHandler registers file handler.
// Generally speaking, encoding a log to bytes then written to file is a common thing.
// So we provide a file handler, which is only for config file.
// If you want to use it in your code, try logit.HandlerOf("file", map[string]interface{...})
//
// For config:
//
//         "handlers":{
//             "file":{
//                 "path":"D:/logit.log"
//             }
//         }
//
//     It will use logit.DefaultTimeFormat to format time in default, so if you want to
//     use your layout to format time, try this:
//
//         "handlers":{
//             "file":{
//                 "path":"D:/logit.log",
//                 "timeFormat": "2006-01-02"
//             }
//         }
//
//     Want a json string? Try this:
//
//         "handlers":{
//             "file":{
//                 "encoder": "json",
//                 "path":"D:/logit.log",
//                 "timeFormat": "2006-01-02"
//             }
//         }
//
func registerFileHandler() {
	RegisterHandler("file", func(params map[string]interface{}) Handler {
		encoder, timeFormat := encoderAndTimeFormatOf(params)
		return NewFileHandler(pathOf(params), encoder, timeFormat)
	})
}

// =============================== duration rolling handler ===============================

func durationRollingWriterOf(params map[string]interface{}) io.Writer {

	// 滚动的时间间隔，单位是秒
	duration := 24 * 60 * 60 // 默认是一天
	if param, ok := params["limit"]; ok {
		duration = int(param.(float64))
	}

	// 保存日志的目标文件夹
	directory := "./" // 默认是当前目录
	if param, ok := params["directory"]; ok {
		directory = param.(string)
	}

	return writer.NewDurationRollingFile(time.Duration(duration)*time.Second, writer.NextFilename(directory))
}

func registerDurationRollingHandler() {
	RegisterHandler("duration", func(params map[string]interface{}) Handler {
		encoder, timeFormat := encoderAndTimeFormatOf(params)
		return NewStandardHandler(durationRollingWriterOf(params), encoder, timeFormat)
	})
}

// =============================== size rolling handler ===============================

func sizeRollingWriterOf(params map[string]interface{}) io.Writer {

	// 滚动的文件大小，单位是 MB
	size := 64 // 默认是 64MB
	if param, ok := params["limit"]; ok {
		size = int(param.(float64))
	}

	// 写出的目标文件夹
	directory := "./" // 默认是当前目录
	if param, ok := params["directory"]; ok {
		directory = param.(string)
	}

	return writer.NewSizeRollingFile(int64(size)*writer.MB, writer.NextFilename(directory))
}

func registerSizeRollingHandler() {
	RegisterHandler("size", func(params map[string]interface{}) Handler {
		encoder, timeFormat := encoderAndTimeFormatOf(params)
		return NewStandardHandler(sizeRollingWriterOf(params), encoder, timeFormat)
	})
}

// =============================== for convenience ===============================

func NewConsoleHandler(encoder Encoder, timeFormat string) Handler {
	return NewStandardHandler(os.Stdout, encoder, timeFormat)
}

func NewFileHandler(path string, encoder Encoder, timeFormat string) Handler {
	file, err := writer.NewFile(path)
	if err != nil {
		panic(err)
	}
	return NewStandardHandler(file, encoder, timeFormat)
}

func NewDurationRollingHandler(limit int, directory string, encoder Encoder, timeFormat string) Handler {
	file := writer.NewDurationRollingFile(time.Duration(limit)*time.Second, writer.NextFilename(directory))
	return NewStandardHandler(file, encoder, timeFormat)
}

func NewSizeRollingHandler(limit int, directory string, encoder Encoder, timeFormat string) Handler {
	file := writer.NewSizeRollingFile(int64(limit)*writer.MB, writer.NextFilename(directory))
	return NewStandardHandler(file, encoder, timeFormat)
}
