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
// Created at 2020/04/24 11:20:24

package logit

import (
	"io"
	"os"
	"strconv"
	"time"

	"github.com/FishGoddess/logit/writer"
)

// WriterOf returns a writer implement with given params.
// Different writer implement may have different params, so what params should
// be used is dependent to specific writer implement.
// An example of this method in config:
//
//     "writer": {
//         "rolling": "duration",
//         "duration": 1,
//         "directory": "D:/"
//     }
//
func WriterOf(params map[string]interface{}) io.Writer {

	// 默认使用 os.Stdout
	w := os.Stdout
	param, ok := params["writer"]
	if !ok {
		return w
	}

	// 下面这段代码有些 “肮脏”，但是大张旗鼓地重构这个又有点主次不分的感觉，所以先保持这样，后续再考虑这个点
	writerConfig := param.(map[string]interface{})
	rolling, ok := writerConfig["rolling"]
	if !ok {
		return w
	}

	switch rolling {

	// 以时间间隔进行滚动的日志写出器
	case "duration":

		// 滚动的时间间隔，单位是秒
		duration := 24 * 60 * 60 // 一天
		if param, ok := writerConfig["duration"]; ok {
			duration = int(param.(float64))
		}

		// 写出的目标文件夹
		directory := "./"
		if param, ok := writerConfig["directory"]; ok {
			directory = param.(string)
		}

		return writer.NewDurationRollingFile(time.Duration(duration)*time.Second, writer.NextFilename(directory))

	// 以文件大小进行滚动的日志写出器
	case "size":

		// 滚动的文件大小，单位是 MB
		size := 64 // 64MB
		if param, ok := writerConfig["size"]; ok {
			size = int(param.(float64))
		}

		// 写出的目标文件夹
		directory := "./"
		if param, ok := writerConfig["directory"]; ok {
			directory = param.(string)
		}

		return writer.NewSizeRollingFile(int64(size)*writer.MB, writer.NextFilename(directory))

	// 不滚动的日志写出器
	case "off":

		// 写出的目标文件
		if param, ok := writerConfig["file"]; ok {
			file, err := writer.NewFile(param.(string))
			if err != nil {
				panic(err)
			}
			return file
		}

		file, err := writer.NewFile("./logit-" + strconv.FormatInt(time.Now().Unix(), 10) + writer.SuffixOfLogFile)
		if err != nil {
			panic(err)
		}
		return file
	}

	return w
}

// Register file handler.
// Generally speaking, encoding a log to bytes then written to file is a common thing.
// So we provide a file handler, which is only for config file.
// If you want to use it in your code, try logit.HandlerOf("file", map[string]interface{...})
//
// For config:
//
//         "handlers":{
//             "file":{
//                 "writer":{
//                     "rolling":"off",
//                     "file":"D:/logit.log"
//                 }
//             }
//         }
//
//     It will use logit.DefaultTimeFormat to format time in default, so if you want to
//     use your layout to format time, try this:
//
//         "handlers":{
//             "file":{
//                 "timeFormat": "2006-01-02"
//                 "writer":{
//                     "rolling":"off",
//                     "file":"D:/logit.log"
//                 }
//             }
//         }
//
//     Want a json string? Try this:
//
//         "handlers":{
//             "file":{
//                 "timeFormat": "2006-01-02",
//                 "encoder": "json"
//                 "writer":{
//                     "rolling":"off",
//                     "file":"D:/logit.log"
//                 }
//             }
//         }
//
func init() {
	// 注册日志处理器
	RegisterHandler("file", func(params map[string]interface{}) Handler {

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

		return NewStandardHandler(WriterOf(params), encoder, timeFormat)
	})
}
