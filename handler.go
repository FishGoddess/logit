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
	"sync"
)

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

	// HandlerIsNotExistedError is an error happens on failing to find the handler.
	HandlerIsNotExistedError = errors.New("the handler you pointed is not existed! Please check the names of all handlers")
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
// Notice that we use panic mechanism to check the name.
// This is a more convenient way to use handlers (we think).
func HandlerOf(name string, params map[string]interface{}) Handler {
	mutexOfHandlers.RLock()
	defer mutexOfHandlers.RUnlock()
	newHandler, ok := handlers[name]
	if !ok {
		panic(HandlerIsNotExistedError)
	}
	return newHandler(params)
}

// ================================= StandardHandler =================================

// StandardHandler is a standard handler for use.
// Generally speaking, encoding a log to bytes then written by writer is the most of
// handlers do. So we provide a standard handler, which only need a writer and an encoder.
// Notice that this handler is not for config file but use in code, so we don't register it.
type StandardHandler struct {
	writer     io.Writer
	encoder    Encoder
	timeFormat string
}

// NewStandardHandler returns a StandardHandler holder with given writer and encoder.
func NewStandardHandler(writer io.Writer, encoder Encoder, timeFormat string) Handler {
	return &StandardHandler{
		writer:     writer,
		encoder:    encoder,
		timeFormat: timeFormat,
	}
}

// Handle will encode log and write log by internal writer.
// Return true so that handlers after it will be used.
func (sh *StandardHandler) Handle(log *Log) bool {
	sh.writer.Write(sh.encoder.Encode(log, sh.timeFormat))
	return true
}
