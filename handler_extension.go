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
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2020/04/24 23:37:56

package logit

import (
	"github.com/FishGoddess/logit/files"
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

// registerConsoleHandler registers console handler.
// Actually, output a log to console is the most of things loggers do when developing.
// So we provide a console handler, which only need an encoder.
// If you want to use it in code, see logit.NewConsoleHandler.
//
// For config:
//     If you want to use this handler in your logger by config file, try this:
//
//         "handlers": {
//             "console": {
//
//             }
//         }
//
// It will use logit.DefaultTimeFormat to format time in default, so if you want to
// use your layout to format time, try this:
//
//         "handlers": {
//             "console": {
//                 "timeFormat": "2006-01-02"
//             }
//         }
//
// Want a Json string? Try this:
//
//         "handlers": {
//             "console": {
//                 "encoder": "json",
//                 "timeFormat": "2006-01-02"
//             }
//         }
//
func registerConsoleHandler() {
	RegisterHandler("console", func(params map[string]interface{}) Handler {
		encoder, timeFormat := encoderAndTimeFormatOf(params, TextEncoder(), DefaultTimeFormat)
		return NewConsoleHandler(encoder, timeFormat)
	})
}

// registerFileHandler registers file handler.
// Generally speaking, encoding a log to bytes then written to file is a common thing.
// So we provide a file handler, which writes logs to file.
// If you want to use it in code, see logit.NewFileHandler.
//
// For config:
//
//         "handlers": {
//             "file": {
//                 "path": "D:/logit.log"
//             }
//         }
//
// The path is where the logs should be written. It is a file, and will be created
// by logit automatically. It will use logit.DefaultTimeFormat to format time in default,
// so if you want to use your layout to format time, try this:
//
//         "handlers": {
//             "file": {
//                 "path": "D:/logit.log",
//                 "timeFormat": "2006-01-02"
//             }
//         }
//
// Want a Json string? Try this:
//
//         "handlers": {
//             "file": {
//                 "encoder": "json",
//                 "path": "D:/logit.log",
//                 "timeFormat": "2006-01-02"
//             }
//         }
//
func registerFileHandler() {
	RegisterHandler("file", func(params map[string]interface{}) Handler {
		path := pathOf(params, "./logit-"+strconv.FormatInt(time.Now().Unix(), 10)+files.SuffixOfLogFile)
		encoder, timeFormat := encoderAndTimeFormatOf(params, TextEncoder(), DefaultTimeFormat)
		return NewFileHandler(path, encoder, timeFormat)
	})
}

// registerDurationRollingHandler registers duration rolling handler.
// Sometimes we want each day has its own log file, or each fixed duration
// has its own log file. That means the log file should switch to a new one
// after duration. That why we provide a duration rolling handler!
// If you want to use it in code, see logit.NewDurationRollingHandler.
//
// For config:
//
//         "handlers": {
//             "duration": {
//                 "limit": 60,
//                 "directory": "D:/logs"
//             }
//         }
//
// You can point limit and directory here. The limit is the duration, the unit is second.
// In demo, "limit": 60 means each 60 seconds has its log file, and one day will be set
// in default. The directory is where the logs store, and "./" will be set in default.
// It will use logit.DefaultTimeFormat to format time in default, so if you want to
// use your layout to format time, try this:
//
//         "handlers": {
//             "duration": {
//                 "limit": 60,
//                 "directory": "D:/logs",
//                 "timeFormat": "2006-01-02"
//             }
//         }
//
// Want a Json string? Try this:
//
//         "handlers":{
//             "duration":{
//                 "limit": 60,
//                 "encoder": "json",
//                 "directory": "D:/logs",
//                 "timeFormat": "2006-01-02"
//             }
//         }
//
func registerDurationRollingHandler() {
	RegisterHandler("duration", func(params map[string]interface{}) Handler {
		// 滚动的时间间隔，单位是秒，默认是 1 天
		limit, directory := limitAndDirectoryOf(params, 24*60*60, "./")
		encoder, timeFormat := encoderAndTimeFormatOf(params, TextEncoder(), DefaultTimeFormat)
		options := rollingHandlerOptionsOf(params, files.DefaultNameGenerator(), files.NewDefaultRollingHook())
		return NewDurationRollingHandlerWithOptions(directory, time.Duration(limit)*time.Second, encoder, timeFormat, options)
	})
}

// registerSizeRollingHandler registers size rolling handler.
// Sometimes we want each log file has a max size, which means the log file should
// switch to a new one after reaching to max size. That why we provide a size rolling handler!
// If you want to use it in code, see logit.NewSizeRollingHandler.
//
// For config:
//
//         "handlers": {
//             "size": {
//                 "limit": 16,
//                 "directory": "D:/logs"
//             }
//         }
//
// You can point limit and directory here. The limit is the max size, the unit is MB.
// In demo, "limit": 16 means the max size of each log file is 16 MB, and 64 MB will be set
// in default. The directory is where the logs store, and "./" will be set in default.
// It will use logit.DefaultTimeFormat to format time in default, so if you want to
// use your layout to format time, try this:
//
//         "handlers": {
//             "duration": {
//                 "limit": 16,
//                 "directory": "D:/logs",
//                 "timeFormat": "2006-01-02"
//             }
//         }
//
// Want a Json string? Try this:
//
//         "handlers":{
//             "duration":{
//                 "limit": 16,
//                 "encoder": "json",
//                 "directory": "D:/logs",
//                 "timeFormat": "2006-01-02"
//             }
//         }
//
func registerSizeRollingHandler() {
	RegisterHandler("size", func(params map[string]interface{}) Handler {
		// 滚动的文件大小，单位是 MB，默认是 64 MB
		limit, directory := limitAndDirectoryOf(params, 64, "./")
		encoder, timeFormat := encoderAndTimeFormatOf(params, TextEncoder(), DefaultTimeFormat)
		options := rollingHandlerOptionsOf(params, files.DefaultNameGenerator(), files.NewDefaultRollingHook())
		return NewSizeRollingHandlerWithOptions(directory, int64(limit)*files.MB, encoder, timeFormat, options)
	})
}

// =============================== for convenience ===============================

// encoderAndTimeFormatOf returns encoder and time format in this params.
// defaultEncoder and defaultTimeFormat will be used if you don't set to params.
func encoderAndTimeFormatOf(params map[string]interface{}, defaultEncoder Encoder, defaultTimeFormat string) (Encoder, string) {

	// 日志编码器参数
	encoder := defaultEncoder
	if encoderName, ok := params["encoder"]; ok && strings.TrimSpace(encoderName.(string)) != "" {
		encoder = encoderOf(encoderName.(string))
	}

	// 时间格式化参数
	timeFormat := defaultTimeFormat
	if format, ok := params["timeFormat"]; ok && strings.TrimSpace(format.(string)) != "" {
		timeFormat = format.(string)
		// 如果参数是 unix，则直接使用空字符串
		if timeFormat == "unix" {
			timeFormat = ""
		}
	}

	return encoder, timeFormat
}

// pathOf returns path in this params.
// defaultPath will be used if you don't set to params.
func pathOf(params map[string]interface{}, defaultPath string) string {

	// 日志输出的目标文件
	path := defaultPath
	if param, ok := params["path"]; ok && strings.TrimSpace(param.(string)) != "" {
		path = param.(string)
	}

	return path
}

// limitAndDirectoryOf returns limit and directory in this params.
// defaultLimit and defaultDirectory will be used if you don't set to params.
func limitAndDirectoryOf(params map[string]interface{}, defaultLimit int, defaultDirectory string) (int, string) {

	// 限制属性的参数
	limit := defaultLimit
	if param, ok := params["limit"]; ok {
		limit = int(param.(float64))
	}

	// 保存日志的目标文件夹
	directory := defaultDirectory // 默认是当前目录
	if param, ok := params["directory"]; ok {
		directory = param.(string)
	}

	return limit, directory
}

// rollingHandlerOptionsOf returns rolling handler options in this params.
// defaultNameGenerator and defaultRollingHook will be used if you don't set to params.
func rollingHandlerOptionsOf(params map[string]interface{}, defaultNameGenerator files.NameGenerator, defaultRollingHook files.RollingHook) RollingHandlerOptions {

	// 获取名字生成器配置
	nameGenerator := defaultNameGenerator
	if param, ok := params["nameGenerator"]; ok {
		nameGenerator = files.NameGeneratorOf(param.(string))
	}

	// 获取滚动钩子配置
	rollingHook := defaultRollingHook
	if _, ok := params["rollingHook"]; ok {
		param := params["rollingHook"].(map[string]map[string]interface{})
		for rollingHookName, rollingHookParams := range param {
			rollingHook = files.RollingHookOf(rollingHookName, rollingHookParams)
			break
		}
	}

	// 返回配置结果
	return RollingHandlerOptions{
		nameGenerator: nameGenerator,
		rollingHook:   rollingHook,
	}
}

// =============================== for public users ===============================

// NewConsoleHandler returns a handler for console.
// This handler will write logs to console by os.Stdout.
// See logit.Encoder, logit.TextEncoder, logit.JsonEncoder.
func NewConsoleHandler(encoder Encoder, timeFormat string) Handler {
	return NewStandardHandler(os.Stdout, encoder, timeFormat)
}

// NewFileHandler returns a handler which writes logs to a file.
// You can point a path (the path of log file) to be used to write logs.
// If the file of this path doesn't exist, a new file will be created.
// See logit.Encoder, logit.TextEncoder, logit.JsonEncoder.
func NewFileHandler(path string, encoder Encoder, timeFormat string) Handler {
	file, err := files.CreateFileOf(path)
	if err != nil {
		panic(err)
	}
	return NewStandardHandler(file, encoder, timeFormat)
}

// NewDurationRollingHandler returns a handler which uses
// a duration rolling file to write logs. The limit is duration, and
// each duration has its own log file. Also you can point a directory
// to be used to store all created log files.
// See logit.Encoder, logit.TextEncoder, logit.JsonEncoder.
// See files.NewDurationRollingFile.
func NewDurationRollingHandler(directory string, limit time.Duration, encoder Encoder, timeFormat string) Handler {
	return NewDurationRollingHandlerWithOptions(directory, limit, encoder, timeFormat, RollingHandlerOptions{})
}

// NewSizeRollingHandler returns a handler which uses
// a size rolling file to write logs. The limit is the max size of log file,
// and the log file will switch to a new one after reaching to max size.
// Also you can point a directory to be used to store all created log files.
// See logit.Encoder, logit.TextEncoder, logit.JsonEncoder.
// See files.NewSizeRollingFile.
func NewSizeRollingHandler(directory string, limit int64, encoder Encoder, timeFormat string) Handler {
	return NewSizeRollingHandlerWithOptions(directory, limit, encoder, timeFormat, RollingHandlerOptions{})
}

// RollingHandlerOptions includes two options for creating duration rolling handler
// and size rolling handler. NameGenerator is for generating the name of created files,
// and rollingHook is a hook which will be triggered on rolling to next file.
// See files.NameGenerator and files.RollingHook.
type RollingHandlerOptions struct {
	nameGenerator files.NameGenerator
	rollingHook   files.RollingHook
}

// filledDurationRollingFileWithOptions fills file with non-nil options.
func filledDurationRollingFileWithOptions(file *files.DurationRollingFile, options RollingHandlerOptions) {
	if options.nameGenerator != nil {
		file.SetNameGenerator(options.nameGenerator)
	}
	if options.rollingHook != nil {
		file.SetRollingHook(options.rollingHook)
	}
}

// filledSizeRollingFileWithOptions fills file with non-nil options.
func filledSizeRollingFileWithOptions(file *files.SizeRollingFile, options RollingHandlerOptions) {
	if options.nameGenerator != nil {
		file.SetNameGenerator(options.nameGenerator)
	}
	if options.rollingHook != nil {
		file.SetRollingHook(options.rollingHook)
	}
}

// NewDurationRollingHandlerWithOptions returns a handler which uses
// a duration rolling file to write logs. The limit is duration, and
// each duration has its own log file. Also you can point a directory
// to be used to store all created log files. Notice that you should
// point an options object which includes nameGenerator and rollingHook.
// NameGenerator is for generating the name of created files, and rollingHook
// is a hook which will be triggered on rolling to next file. However,
// both of them is not necessary, so if one of them is nil then this "one"
// option will be ignored. See logit.RollingHandlerOptions.
// See logit.Encoder, logit.TextEncoder, logit.JsonEncoder.
// See files.NewDurationRollingFile.
func NewDurationRollingHandlerWithOptions(directory string, limit time.Duration, encoder Encoder, timeFormat string, options RollingHandlerOptions) Handler {
	file := files.NewDurationRollingFile(directory, limit)
	filledDurationRollingFileWithOptions(file, options)
	return NewStandardHandler(file, encoder, timeFormat)
}

// NewSizeRollingHandlerWithOptions returns a handler which uses
// a size rolling file to write logs. The limit is the max size of log file,
// and the log file will switch to a new one after reaching to max size.
// Also you can point a directory to be used to store all created log files.
// Notice that you should point an options object which includes nameGenerator
// and rollingHook. NameGenerator is for generating the name of created files,
// and rollingHook is a hook which will be triggered on rolling to next file.
// However, both of them is not necessary, so if one of them is nil then this "one"
// option will be ignored. See logit.RollingHandlerOptions.
// See logit.Encoder, logit.TextEncoder, logit.JsonEncoder.
// See files.NewSizeRollingFile.
func NewSizeRollingHandlerWithOptions(directory string, limit int64, encoder Encoder, timeFormat string, options RollingHandlerOptions) Handler {
	file := files.NewSizeRollingFile(directory, limit)
	filledSizeRollingFileWithOptions(file, options)
	return NewStandardHandler(file, encoder, timeFormat)
}
