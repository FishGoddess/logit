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

// levelBasedHandler is a level sensitive handler.
// It only handles the specific level of this handler, so you can use different handlers
// to handle logs in different levels. For example, you want debug/info/warn level logs are
// written to log file "xxx.log" and error level logs are written to log file "xxx.error.log",
// maybe written to log server, too, then you can use this handler to do it.
type levelBasedHandler struct {

	// level is the level of log that can be handled by this handler.
	// See logit.Level.
	level Level

	// handlers is all handlers used to handle logs in level.
	// See logit.Handler.
	handlers []Handler
}

// NewLevelBasedHandler returns a handler handled logs in level.
// You can add more than one handler to this handler. This handler is just like a
// wrapper wrapping some handlers.
func NewLevelBasedHandler(level Level, handlers ...Handler) Handler {
	return &levelBasedHandler{
		level:    level,
		handlers: handlers,
	}
}

// Handle handles a log with handlers in lbh.
// Notice that the handle process will be interrupted if one of them
// returned false. However, this method will always return true, so the handlers
// after it will always be used.
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

// handlersOf returns handlers parsed from params.
func handlersOf(params map[string]interface{}) []Handler {
	handlers := make([]Handler, 0, len(params)+2)
	for name, paramsOfHandler := range params {
		handlers = append(handlers, handlerOf(name, paramsOfHandler.(map[string]interface{})))
	}
	return handlers
}

// ================================ debug level handler ================================

// registerDebugLevelHandler registers debug level handler which
// only handles logs in debug level.
//
// For config:
//     If you want to use this handler in your logger by config file, try this:
//
//         "handlers": {
//             "debug": {
//                 "console": {}
//             }
//         }
//
// You should always know that this handler is a wrapper, so you must add some other
// handlers to it otherwise it will do nothing. All registered handlers can be added
// to it, even one more debug level handler. Now, you know it is only a filter that filters
// non-debug level logs. As for what params should be written in handlers inside are dependent
// to different handlers. Check other handlers' documents to know more about information.
// See logit.Handler.
func registerDebugLevelHandler() {
	RegisterHandler("debug", func(params map[string]interface{}) Handler {
		return NewLevelBasedHandler(DebugLevel, handlersOf(params)...)
	})
}

// ================================  info level handler ================================

// registerInfoLevelHandler registers info level handler which
// only handles logs in info level.
//
// For config:
//     If you want to use this handler in your logger by config file, try this:
//
//         "handlers": {
//             "info": {
//                 "console": {}
//             }
//         }
//
// You should always know that this handler is a wrapper, so you must add some other
// handlers to it otherwise it will do nothing. All registered handlers can be added
// to it, even one more info level handler. Now, you know it is only a filter that filters
// non-info level logs. As for what params should be written in handlers inside are dependent
// to different handlers. Check other handlers' documents to know more about information.
// See logit.Handler.
func registerInfoLevelHandler() {
	RegisterHandler("info", func(params map[string]interface{}) Handler {
		return NewLevelBasedHandler(InfoLevel, handlersOf(params)...)
	})
}

// ================================  warn level handler ================================

// registerWarnLevelHandler registers warn level handler which
// only handles logs in warn level.
//
// For config:
//     If you want to use this handler in your logger by config file, try this:
//
//         "handlers": {
//             "warn": {
//                 "console": {}
//             }
//         }
//
// You should always know that this handler is a wrapper, so you must add some other
// handlers to it otherwise it will do nothing. All registered handlers can be added
// to it, even one more warn level handler. Now, you know it is only a filter that filters
// non-warn level logs. As for what params should be written in handlers inside are dependent
// to different handlers. Check other handlers' documents to know more about information.
// See logit.Handler.
func registerWarnLevelHandler() {
	RegisterHandler("warn", func(params map[string]interface{}) Handler {
		return NewLevelBasedHandler(WarnLevel, handlersOf(params)...)
	})
}

// ================================ error level handler ================================

// registerErrorLevelHandler registers error level handler which
// only handles logs in error level.
//
// For config:
//     If you want to use this handler in your logger by config file, try this:
//
//         "handlers": {
//             "error": {
//                 "console": {}
//             }
//         }
//
// You should always know that this handler is a wrapper, so you must add some other
// handlers to it otherwise it will do nothing. All registered handlers can be added
// to it, even one more error level handler. Now, you know it is only a filter that filters
// non-error level logs. As for what params should be written in handlers inside are dependent
// to different handlers. Check other handlers' documents to know more about information.
// See logit.Handler.
func registerErrorLevelHandler() {
	RegisterHandler("error", func(params map[string]interface{}) Handler {
		return NewLevelBasedHandler(ErrorLevel, handlersOf(params)...)
	})
}
