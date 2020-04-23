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
	"errors"
	"io"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/FishGoddess/logit/writer"
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

func init() {

	// 注册默认日志处理器
	RegisterHandler("default", func(params map[string]interface{}) Handler {
		timeFormat := DefaultTimeFormat
		if format, ok := params["timeFormat"]; ok && format != "" {
			timeFormat = format.(string)
		}
		return NewDefaultHandler(WriterOf(params), timeFormat)
	})

	// 注册 Json 格式日志处理器
	RegisterHandler("json", func(params map[string]interface{}) Handler {
		timeFormat := ""
		if format, ok := params["timeFormat"]; ok {
			timeFormat = format.(string)
		}
		return NewJsonHandler(WriterOf(params), timeFormat)
	})
}

const (
	// DefaultTimeFormat is the default format for formatting time.
	DefaultTimeFormat = "2006-01-02 15:04:05"
)

var (
	// handlers stores all registered handlers.
	// mutexOfHandlers is for concurrency.
	handlers        = map[string]func(params map[string]interface{}) Handler{}
	mutexOfHandlers = &sync.RWMutex{}

	// HandlerIsExistedError is an error happens on repeating handler name.
	HandlerIsExistedError = errors.New("the name of handler you want to register already exists! May be you should give it an another name")
)

// RegisterHandler registers your handler to logit so that you can use them easily.
// Return an error if the name is existed, and you should change another name for your handler.
// Notice that newHandler has a parameter called params, which will be injected into newHandler
// by logit automatically. Different handler may have different params, so what params should
// be injected into newHandler is dependent to specific handler.
func RegisterHandler(name string, newHandler func(params map[string]interface{}) Handler) error {
	mutexOfHandlers.Lock()
	defer mutexOfHandlers.Unlock()
	if _, ok := handlers[name]; ok {
		return HandlerIsExistedError
	}
	handlers[name] = newHandler
	return nil
}

// HandlerOf returns handler whose name is given name and params.
// Different handler may have different params, so what params should
// be injected into newHandler is dependent to specific handler.
// Notice that we don't use an error mechanism or ok mechanism to check the name but
// a default handler returning mechanism. This is a more convenient way to use handlers (we think).
func HandlerOf(name string, params map[string]interface{}) Handler {
	mutexOfHandlers.RLock()
	defer mutexOfHandlers.RUnlock()
	newHandler, ok := handlers[name]
	if !ok {
		return NewDefaultHandler(os.Stdout, DefaultTimeFormat)
	}
	return newHandler(params)
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
//
// For config:
//     If you want to use this handler in your logger by config file, try this:
//
//         "handlers":{
//             "json":{
//
//             }
//         }
//
//     This config will not format time, and keep time in unix form. If you want to
//     use your layout to format time, try this:
//
//         "handlers":{
//             "json":{
//                 "timeFormat": "2006-01-02"
//             }
//         }
//
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
