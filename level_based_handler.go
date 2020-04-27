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
// Email: fishinlove@163.com
// Created at 2020/04/27 22:44:04

package logit

func init() {
	registerDebugLevelHandler()
	registerInfoLevelHandler()
	registerWarnLevelHandler()
	registerErrorLevelHandler()
}

type levelBasedHandler struct {
	level    Level
	handlers []Handler
}

func NewLevelBasedHandler(level Level, handlers ...Handler) Handler {
	return &levelBasedHandler{
		level:    level,
		handlers: handlers,
	}
}

func (lbh *levelBasedHandler) Handle(log *Log) bool {
	if log.Level() == lbh.level {
		for _, handler := range lbh.handlers {
			if !handler.Handle(log) {
				break
			}
		}
	}
	return true
}

// ================================ debug level handler ================================

func registerDebugLevelHandler() {
	RegisterHandler("debug", func(params map[string]interface{}) Handler {
		handlers := make([]Handler, 0, len(params)+2)
		for name, paramsOfHandler := range params {
			handlers = append(handlers, handlerOf(name, paramsOfHandler.(map[string]interface{})))
		}
		return NewLevelBasedHandler(DebugLevel, handlers...)
	})
}

// ================================  info level handler ================================

func registerInfoLevelHandler() {
	RegisterHandler("info", func(params map[string]interface{}) Handler {
		handlers := make([]Handler, 0, len(params)+2)
		for name, paramsOfHandler := range params {
			handlers = append(handlers, handlerOf(name, paramsOfHandler.(map[string]interface{})))
		}
		return NewLevelBasedHandler(InfoLevel, handlers...)
	})
}

// ================================  warn level handler ================================

func registerWarnLevelHandler() {
	RegisterHandler("warn", func(params map[string]interface{}) Handler {
		handlers := make([]Handler, 0, len(params)+2)
		for name, paramsOfHandler := range params {
			handlers = append(handlers, handlerOf(name, paramsOfHandler.(map[string]interface{})))
		}
		return NewLevelBasedHandler(WarnLevel, handlers...)
	})
}

// ================================ error level handler ================================

func registerErrorLevelHandler() {
	RegisterHandler("error", func(params map[string]interface{}) Handler {
		handlers := make([]Handler, 0, len(params)+2)
		for name, paramsOfHandler := range params {
			handlers = append(handlers, handlerOf(name, paramsOfHandler.(map[string]interface{})))
		}
		return NewLevelBasedHandler(ErrorLevel, handlers...)
	})
}
