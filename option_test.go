// Copyright 2021 Ye Zi Jie. All Rights Reserved.
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
// Created at 2021/07/11 14:00:53

package logit

import (
	"os"
	"testing"

	"github.com/FishGoddess/logit/core/appender"
)

// go test -v -cover -run=^TestOptions$
func TestOptions(t *testing.T) {

	opts := Options()
	if opts != nil {
		t.Fatalf("Options returns wrong options %+v", opts)
	}
}

// go test -v -cover -run=^TestOptionsWithDebugLevel$
func TestOptionsWithDebugLevel(t *testing.T) {

	opts := Options()

	logger := NewLogger()
	logger.level = offLevel
	opts.WithDebugLevel()(logger)
	if logger.level != debugLevel {
		t.Fatalf("logger's level %s is wrong", logger.level)
	}
}

// go test -v -cover -run=^TestOptionsWithInfoLevel$
func TestOptionsWithInfoLevel(t *testing.T) {

	opts := Options()

	logger := NewLogger()
	logger.level = offLevel
	opts.WithInfoLevel()(logger)
	if logger.level != infoLevel {
		t.Fatalf("logger's level %s is wrong", logger.level)
	}
}

// go test -v -cover -run=^TestOptionsWithWarnLevel$
func TestOptionsWithWarnLevel(t *testing.T) {

	opts := Options()

	logger := NewLogger()
	logger.level = offLevel
	opts.WithWarnLevel()(logger)
	if logger.level != warnLevel {
		t.Fatalf("logger's level %s is wrong", logger.level)
	}
}

// go test -v -cover -run=^TestOptionsWithErrorLevel$
func TestOptionsWithErrorLevel(t *testing.T) {

	opts := Options()

	logger := NewLogger()
	logger.level = offLevel
	opts.WithErrorLevel()(logger)
	if logger.level != errorLevel {
		t.Fatalf("logger's level %s is wrong", logger.level)
	}
}

// go test -v -cover -run=^TestOptionsWithAppender$
func TestOptionsWithAppender(t *testing.T) {

	opts := Options()

	logger := NewLogger()
	logger.appender = nil
	opts.WithAppender(appender.Text())(logger)
	if logger.appender != appender.Text() {
		t.Fatalf("logger's appender %s is wrong", logger.appender)
	}

	logger.appender = nil
	opts.WithAppender(appender.Json())(logger)
	if logger.appender != appender.Json() {
		t.Fatalf("logger's appender %s is wrong", logger.appender)
	}
}

// go test -v -cover -run=^TestOptionsWithWriter$
func TestOptionsWithWriter(t *testing.T) {

	opts := Options()

	logger := NewLogger()
	logger.writer = nil
	opts.WithWriter(os.Stdout)(logger)
	if logger.writer == nil {
		t.Fatalf("logger's writer %s is wrong", logger.writer)
	}
}

// go test -v -cover -run=^TestOptionsWithBuffered$
func TestOptionsWithBuffered(t *testing.T) {

	opts := Options()

	logger := NewLogger()
	logger.writer = nil
	opts.WithBuffered(os.Stdout)(logger)
	if logger.writer == nil {
		t.Fatalf("logger's writer %s is wrong", logger.writer)
	}
}

// go test -v -cover -run=^TestOptionsWithPid$
func TestOptionsWithPid(t *testing.T) {

	opts := Options()

	logger := NewLogger()
	logger.needPid = false
	opts.WithPid()(logger)
	if logger.needPid != true {
		t.Fatalf("logger's needPid %+v is wrong", logger.needPid)
	}
}

// go test -v -cover -run=^TestOptionsWithCaller$
func TestOptionsWithCaller(t *testing.T) {

	opts := Options()

	logger := NewLogger()
	logger.needCaller = false
	opts.WithCaller()(logger)
	if logger.needCaller != true {
		t.Fatalf("logger's needCaller %+v is wrong", logger.needCaller)
	}
}

// go test -v -cover -run=^TestOptionsWithMsgKey$
func TestOptionsWithMsgKey(t *testing.T) {

	opts := Options()

	logger := NewLogger()
	logger.msgKey = ""
	opts.WithMsgKey("msg")(logger)
	if logger.msgKey != "msg" {
		t.Fatalf("logger's msgKey %+v is wrong", logger.msgKey)
	}
}

// go test -v -cover -run=^TestOptionsWithTimeKey$
func TestOptionsWithTimeKey(t *testing.T) {

	opts := Options()

	logger := NewLogger()
	logger.timeKey = ""
	opts.WithTimeKey("time")(logger)
	if logger.timeKey != "time" {
		t.Fatalf("logger's timeKey %+v is wrong", logger.timeKey)
	}
}

// go test -v -cover -run=^TestOptionsWithLevelKey$
func TestOptionsWithLevelKey(t *testing.T) {

	opts := Options()

	logger := NewLogger()
	logger.levelKey = ""
	opts.WithLevelKey("level")(logger)
	if logger.levelKey != "level" {
		t.Fatalf("logger's levelKey %+v is wrong", logger.levelKey)
	}
}

// go test -v -cover -run=^TestOptionsWithPidKey$
func TestOptionsWithPidKey(t *testing.T) {

	opts := Options()

	logger := NewLogger()
	logger.pidKey = ""
	opts.WithPidKey("pid")(logger)
	if logger.pidKey != "pid" {
		t.Fatalf("logger's pidKey %+v is wrong", logger.pidKey)
	}
}

// go test -v -cover -run=^TestOptionsWithFileKey$
func TestOptionsWithFileKey(t *testing.T) {

	opts := Options()

	logger := NewLogger()
	logger.fileKey = ""
	opts.WithFileKey("file")(logger)
	if logger.fileKey != "file" {
		t.Fatalf("logger's fileKey %+v is wrong", logger.fileKey)
	}
}

// go test -v -cover -run=^TestOptionsWithLineKey$
func TestOptionsWithLineKey(t *testing.T) {

	opts := Options()

	logger := NewLogger()
	logger.lineKey = ""
	opts.WithLineKey("line")(logger)
	if logger.lineKey != "line" {
		t.Fatalf("logger's lineKey %+v is wrong", logger.lineKey)
	}
}

// go test -v -cover -run=^TestOptionsWithTimeFormat$
func TestOptionsWithTimeFormat(t *testing.T) {

	opts := Options()

	logger := NewLogger()
	logger.timeFormat = ""
	opts.WithTimeFormat("20060102150405")(logger)
	if logger.timeFormat != "20060102150405" {
		t.Fatalf("logger's timeFormat %+v is wrong", logger.timeFormat)
	}
}
