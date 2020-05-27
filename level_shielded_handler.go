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
// Created at 2020/05/27 20:50:41

package logit

func init() {
	registerNonDebugLevelHandler()
	registerNonInfoLevelHandler()
	registerNonWarnLevelHandler()
	registerNonErrorLevelHandler()
}

// levelShieldedHandler is a level shielded handler.
// It will shield the specific level and handles logs in other levels, for example, you want debug/info/warn
// level logs are written to log file "xxx.log" and error level logs are written to log file "xxx.error.log",
// then you can use this handler to do it.
type levelShieldedHandler struct {

	// level is the level of log that shielded by this handler.
	// See logit.Level.
	level Level

	// handlers is all handlers used to handle logs in level.
	// See logit.Handler.
	handlers []Handler
}

// NewLevelShieldedHandler returns a handler handled logs in all levels except level.
// You can add more than one handler to this handler. This handler is just like a
// wrapper wrapping some handlers.
func NewLevelShieldedHandler(level Level, handlers ...Handler) Handler {
	return &levelShieldedHandler{
		level:    level,
		handlers: handlers,
	}
}

// Handle handles a log with handlers in lsh.
// Notice that the handling process will be interrupted if one of them
// returned false. However, this method will always return true, so the handlers
// after it will always be used.
func (lsh *levelShieldedHandler) Handle(log *Log) bool {
	if log.Level() != lsh.level {
		for _, handler := range lsh.handlers {
			if !handler.Handle(log) {
				break
			}
		}
	}
	return true
}

// ================================ non-debug level handler ================================

// registerNonDebugLevelHandler registers non-debug level handler which
// handles logs in all levels except debug level.
//
// For config:
//     If you want to use this handler in your logger by config file, try this:
//
//         "handlers": {
//             "!debug": {
//                 "console": {}
//             }
//         }
//
// You should always know that this handler is a wrapper, so you must add some other
// handlers to it otherwise it will do nothing. All registered handlers can be added
// to it, even one more non-debug level handler. Now, you know it is only a filter that filters
// debug level logs. As for what params should be written in handlers inside are dependent
// to different handlers. Check other handlers' documents to know more about information.
// See logit.Handler.
func registerNonDebugLevelHandler() {
	RegisterHandler("!debug", func(params map[string]interface{}) Handler {
		return NewLevelShieldedHandler(DebugLevel, handlersOf(params)...)
	})
}

// ================================ non-info level handler ================================

// registerNonInfoLevelHandler registers non-info level handler which
// handles logs in all levels except info level.
//
// For config:
//     If you want to use this handler in your logger by config file, try this:
//
//         "handlers": {
//             "!info": {
//                 "console": {}
//             }
//         }
//
// You should always know that this handler is a wrapper, so you must add some other
// handlers to it otherwise it will do nothing. All registered handlers can be added
// to it, even one more non-info level handler. Now, you know it is only a filter that filters
// info level logs. As for what params should be written in handlers inside are dependent
// to different handlers. Check other handlers' documents to know more about information.
// See logit.Handler.
func registerNonInfoLevelHandler() {
	RegisterHandler("!info", func(params map[string]interface{}) Handler {
		return NewLevelShieldedHandler(InfoLevel, handlersOf(params)...)
	})
}

// ================================ non-warn level handler ================================

// registerNonWarnLevelHandler registers non-warn level handler which
// handles logs in all levels except warn level.
//
// For config:
//     If you want to use this handler in your logger by config file, try this:
//
//         "handlers": {
//             "!warn": {
//                 "console": {}
//             }
//         }
//
// You should always know that this handler is a wrapper, so you must add some other
// handlers to it otherwise it will do nothing. All registered handlers can be added
// to it, even one more non-warn level handler. Now, you know it is only a filter that filters
// warn level logs. As for what params should be written in handlers inside are dependent
// to different handlers. Check other handlers' documents to know more about information.
// See logit.Handler.
func registerNonWarnLevelHandler() {
	RegisterHandler("!warn", func(params map[string]interface{}) Handler {
		return NewLevelShieldedHandler(WarnLevel, handlersOf(params)...)
	})
}

// ================================ non-error level handler ================================

// registerNonErrorLevelHandler registers non-error level handler which
// handles logs in all levels except error level.
//
// For config:
//     If you want to use this handler in your logger by config file, try this:
//
//         "handlers": {
//             "!error": {
//                 "console": {}
//             }
//         }
//
// You should always know that this handler is a wrapper, so you must add some other
// handlers to it otherwise it will do nothing. All registered handlers can be added
// to it, even one more non-error level handler. Now, you know it is only a filter that filters
// error level logs. As for what params should be written in handlers inside are dependent
// to different handlers. Check other handlers' documents to know more about information.
// See logit.Handler.
func registerNonErrorLevelHandler() {
	RegisterHandler("!error", func(params map[string]interface{}) Handler {
		return NewLevelShieldedHandler(ErrorLevel, handlersOf(params)...)
	})
}
