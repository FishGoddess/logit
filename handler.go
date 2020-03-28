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
    "io"
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
    Handle(log []byte, raw *Log) bool
}

// DefaultHandler is a default handler for use.
// Generally speaking, encoding a log to bytes then written by writer is the most of
// handlers do. So we provide a default handler, which only need a writer and an encoder.
type DefaultHandler struct {
    writer io.Writer
}

// NewDefaultHandler returns a DefaultHandler holder with given writer.
func NewDefaultHandler(writer io.Writer) Handler {
    return &DefaultHandler{
        writer: writer,
    }
}

// Handle will encode log to bytes with internal encoder and written by internal writer.
// Return true so that handlers after it will be used.
func (dh *DefaultHandler) Handle(log []byte, raw *Log) bool {
    dh.writer.Write(log)
    return true
}
