// Copyright 2022 FishGoddess. All Rights Reserved.
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

package logit

import (
	"context"
	"os"
	"testing"

	"github.com/go-logit/logit/core/appender"
)

// go test -v -cover -run=^TestOptions$
func TestOptions(t *testing.T) {
	opts := Options()
	if opts != nil {
		t.Errorf("Options returns wrong options %+v", opts)
	}
}

// go test -v -cover -run=^TestOptionsWithDebugLevel$
func TestOptionsWithDebugLevel(t *testing.T) {
	opts := Options()

	logger := NewLogger()
	logger.level = offLevel

	opts.WithDebugLevel()(logger)
	if logger.level != debugLevel {
		t.Errorf("logger's level %s is wrong", logger.level)
	}
}

// go test -v -cover -run=^TestOptionsWithInfoLevel$
func TestOptionsWithInfoLevel(t *testing.T) {
	opts := Options()

	logger := NewLogger()
	logger.level = offLevel

	opts.WithInfoLevel()(logger)
	if logger.level != infoLevel {
		t.Errorf("logger's level %s is wrong", logger.level)
	}
}

// go test -v -cover -run=^TestOptionsWithWarnLevel$
func TestOptionsWithWarnLevel(t *testing.T) {
	opts := Options()

	logger := NewLogger()
	logger.level = offLevel

	opts.WithWarnLevel()(logger)
	if logger.level != warnLevel {
		t.Errorf("logger's level %s is wrong", logger.level)
	}
}

// go test -v -cover -run=^TestOptionsWithErrorLevel$
func TestOptionsWithErrorLevel(t *testing.T) {
	opts := Options()

	logger := NewLogger()
	logger.level = offLevel

	opts.WithErrorLevel()(logger)
	if logger.level != errorLevel {
		t.Errorf("logger's level %s is wrong", logger.level)
	}
}

// go test -v -cover -run=^TestOptionsWithOffLevel$
func TestOptionsWithOffLevel(t *testing.T) {
	opts := Options()

	logger := NewLogger()
	logger.level = debugLevel

	opts.WithOffLevel()(logger)
	if logger.level != offLevel {
		t.Errorf("logger's level %s is wrong", logger.level)
	}
}

// go test -v -cover -run=^TestOptionsWithAppender$
func TestOptionsWithAppender(t *testing.T) {
	opts := Options()

	logger := NewLogger()
	logger.debugAppender = nil
	logger.infoAppender = nil
	logger.warnAppender = nil
	logger.errorAppender = nil
	logger.printAppender = nil
	opts.WithAppender(appender.Text())(logger)

	if logger.debugAppender != appender.Text() {
		t.Errorf("logger's debugAppender %s is wrong", logger.debugAppender)
	}

	if logger.infoAppender != appender.Text() {
		t.Errorf("logger's infoAppender %s is wrong", logger.infoAppender)
	}

	if logger.warnAppender != appender.Text() {
		t.Errorf("logger's warnAppender %s is wrong", logger.warnAppender)
	}

	if logger.errorAppender != appender.Text() {
		t.Errorf("logger's errorAppender %s is wrong", logger.errorAppender)
	}

	if logger.printAppender != appender.Text() {
		t.Errorf("logger's printAppender %s is wrong", logger.printAppender)
	}

	logger.debugAppender = nil
	logger.infoAppender = nil
	logger.warnAppender = nil
	logger.errorAppender = nil
	logger.printAppender = nil
	opts.WithAppender(appender.Json())(logger)

	if logger.debugAppender != appender.Json() {
		t.Errorf("logger's debugAppender %s is wrong", logger.debugAppender)
	}

	if logger.infoAppender != appender.Json() {
		t.Errorf("logger's infoAppender %s is wrong", logger.infoAppender)
	}

	if logger.warnAppender != appender.Json() {
		t.Errorf("logger's warnAppender %s is wrong", logger.warnAppender)
	}

	if logger.errorAppender != appender.Json() {
		t.Errorf("logger's errorAppender %s is wrong", logger.errorAppender)
	}

	if logger.printAppender != appender.Json() {
		t.Errorf("logger's printAppender %s is wrong", logger.printAppender)
	}
}

// go test -v -cover -run=^TestOptionsWithDebugAppender$
func TestOptionsWithDebugAppender(t *testing.T) {
	opts := Options()

	logger := NewLogger()
	logger.debugAppender = nil
	logger.infoAppender = nil
	logger.warnAppender = nil
	logger.errorAppender = nil
	logger.printAppender = nil
	opts.WithDebugAppender(appender.Text())(logger)

	if logger.debugAppender != appender.Text() {
		t.Errorf("logger's debugAppender %s is wrong", logger.debugAppender)
	}

	if logger.infoAppender != nil {
		t.Errorf("logger's infoAppender %s is wrong", logger.infoAppender)
	}

	if logger.warnAppender != nil {
		t.Errorf("logger's warnAppender %s is wrong", logger.warnAppender)
	}

	if logger.errorAppender != nil {
		t.Errorf("logger's errorAppender %s is wrong", logger.errorAppender)
	}

	if logger.printAppender != nil {
		t.Errorf("logger's printAppender %s is wrong", logger.printAppender)
	}
}

// go test -v -cover -run=^TestOptionsWithInfoAppender$
func TestOptionsWithInfoAppender(t *testing.T) {
	opts := Options()

	logger := NewLogger()
	logger.debugAppender = nil
	logger.infoAppender = nil
	logger.warnAppender = nil
	logger.errorAppender = nil
	logger.printAppender = nil
	opts.WithInfoAppender(appender.Text())(logger)

	if logger.debugAppender != nil {
		t.Errorf("logger's debugAppender %s is wrong", logger.debugAppender)
	}

	if logger.infoAppender != appender.Text() {
		t.Errorf("logger's infoAppender %s is wrong", logger.infoAppender)
	}

	if logger.warnAppender != nil {
		t.Errorf("logger's warnAppender %s is wrong", logger.warnAppender)
	}

	if logger.errorAppender != nil {
		t.Errorf("logger's errorAppender %s is wrong", logger.errorAppender)
	}

	if logger.printAppender != nil {
		t.Errorf("logger's printAppender %s is wrong", logger.printAppender)
	}
}

// go test -v -cover -run=^TestOptionsWithWarnAppender$
func TestOptionsWithWarnAppender(t *testing.T) {
	opts := Options()

	logger := NewLogger()
	logger.debugAppender = nil
	logger.infoAppender = nil
	logger.warnAppender = nil
	logger.errorAppender = nil
	logger.printAppender = nil
	opts.WithWarnAppender(appender.Text())(logger)

	if logger.debugAppender != nil {
		t.Errorf("logger's debugAppender %s is wrong", logger.debugAppender)
	}

	if logger.infoAppender != nil {
		t.Errorf("logger's infoAppender %s is wrong", logger.infoAppender)
	}

	if logger.warnAppender != appender.Text() {
		t.Errorf("logger's warnAppender %s is wrong", logger.warnAppender)
	}

	if logger.errorAppender != nil {
		t.Errorf("logger's errorAppender %s is wrong", logger.errorAppender)
	}

	if logger.printAppender != nil {
		t.Errorf("logger's printAppender %s is wrong", logger.printAppender)
	}
}

// go test -v -cover -run=^TestOptionsWithErrorAppender$
func TestOptionsWithErrorAppender(t *testing.T) {
	opts := Options()

	logger := NewLogger()
	logger.debugAppender = nil
	logger.infoAppender = nil
	logger.warnAppender = nil
	logger.errorAppender = nil
	logger.printAppender = nil
	opts.WithErrorAppender(appender.Text())(logger)

	if logger.debugAppender != nil {
		t.Errorf("logger's debugAppender %s is wrong", logger.debugAppender)
	}

	if logger.infoAppender != nil {
		t.Errorf("logger's infoAppender %s is wrong", logger.infoAppender)
	}

	if logger.warnAppender != nil {
		t.Errorf("logger's warnAppender %s is wrong", logger.warnAppender)
	}

	if logger.errorAppender != appender.Text() {
		t.Errorf("logger's errorAppender %s is wrong", logger.errorAppender)
	}

	if logger.printAppender != nil {
		t.Errorf("logger's printAppender %s is wrong", logger.printAppender)
	}
}

// go test -v -cover -run=^TestOptionsWithPrintAppender$
func TestOptionsWithPrintAppender(t *testing.T) {
	opts := Options()

	logger := NewLogger()
	logger.debugAppender = nil
	logger.infoAppender = nil
	logger.warnAppender = nil
	logger.errorAppender = nil
	logger.printAppender = nil
	opts.WithPrintAppender(appender.Text())(logger)

	if logger.debugAppender != nil {
		t.Errorf("logger's debugAppender %s is wrong", logger.debugAppender)
	}

	if logger.infoAppender != nil {
		t.Errorf("logger's infoAppender %s is wrong", logger.infoAppender)
	}

	if logger.warnAppender != nil {
		t.Errorf("logger's warnAppender %s is wrong", logger.warnAppender)
	}

	if logger.errorAppender != nil {
		t.Errorf("logger's errorAppender %s is wrong", logger.errorAppender)
	}

	if logger.printAppender != appender.Text() {
		t.Errorf("logger's printAppender %s is wrong", logger.printAppender)
	}
}

// go test -v -cover -run=^TestOptionsWithWriter$
func TestOptionsWithWriter(t *testing.T) {
	opts := Options()

	logger := NewLogger()
	logger.debugWriter = nil
	logger.infoWriter = nil
	logger.warnWriter = nil
	logger.errorWriter = nil
	logger.printWriter = nil
	opts.WithWriter(os.Stdout, false)(logger)

	if logger.debugWriter == nil {
		t.Errorf("logger's debugWriter %s is wrong", logger.debugWriter)
	}

	if logger.infoWriter == nil {
		t.Errorf("logger's infoWriter %s is wrong", logger.infoWriter)
	}

	if logger.warnWriter == nil {
		t.Errorf("logger's warnWriter %s is wrong", logger.warnWriter)
	}

	if logger.errorWriter == nil {
		t.Errorf("logger's errorWriter %s is wrong", logger.errorWriter)
	}

	if logger.printWriter == nil {
		t.Errorf("logger's printWriter %s is wrong", logger.printWriter)
	}
}

// go test -v -cover -run=^TestOptionsWithDebugWriter$
func TestOptionsWithDebugWriter(t *testing.T) {
	opts := Options()

	logger := NewLogger()
	logger.debugWriter = nil
	logger.infoWriter = nil
	logger.warnWriter = nil
	logger.errorWriter = nil
	logger.printWriter = nil
	opts.WithDebugWriter(os.Stdout, false)(logger)

	if logger.debugWriter == nil {
		t.Errorf("logger's debugWriter %s is wrong", logger.debugWriter)
	}

	if logger.infoWriter != nil {
		t.Errorf("logger's infoWriter %s is wrong", logger.infoWriter)
	}

	if logger.warnWriter != nil {
		t.Errorf("logger's warnWriter %s is wrong", logger.warnWriter)
	}

	if logger.errorWriter != nil {
		t.Errorf("logger's errorWriter %s is wrong", logger.errorWriter)
	}

	if logger.printWriter != nil {
		t.Errorf("logger's printWriter %s is wrong", logger.printWriter)
	}
}

// go test -v -cover -run=^TestOptionsWithInfoWriter$
func TestOptionsWithInfoWriter(t *testing.T) {
	opts := Options()

	logger := NewLogger()
	logger.debugWriter = nil
	logger.infoWriter = nil
	logger.warnWriter = nil
	logger.errorWriter = nil
	logger.printWriter = nil
	opts.WithInfoWriter(os.Stdout, false)(logger)

	if logger.debugWriter != nil {
		t.Errorf("logger's debugWriter %s is wrong", logger.debugWriter)
	}

	if logger.infoWriter == nil {
		t.Errorf("logger's infoWriter %s is wrong", logger.infoWriter)
	}

	if logger.warnWriter != nil {
		t.Errorf("logger's warnWriter %s is wrong", logger.warnWriter)
	}

	if logger.errorWriter != nil {
		t.Errorf("logger's errorWriter %s is wrong", logger.errorWriter)
	}

	if logger.printWriter != nil {
		t.Errorf("logger's printWriter %s is wrong", logger.printWriter)
	}
}

// go test -v -cover -run=^TestOptionsWithWarnWriter$
func TestOptionsWithWarnWriter(t *testing.T) {
	opts := Options()

	logger := NewLogger()
	logger.debugWriter = nil
	logger.infoWriter = nil
	logger.warnWriter = nil
	logger.errorWriter = nil
	logger.printWriter = nil
	opts.WithWarnWriter(os.Stdout, false)(logger)

	if logger.debugWriter != nil {
		t.Errorf("logger's debugWriter %s is wrong", logger.debugWriter)
	}

	if logger.infoWriter != nil {
		t.Errorf("logger's infoWriter %s is wrong", logger.infoWriter)
	}

	if logger.warnWriter == nil {
		t.Errorf("logger's warnWriter %s is wrong", logger.warnWriter)
	}

	if logger.errorWriter != nil {
		t.Errorf("logger's errorWriter %s is wrong", logger.errorWriter)
	}

	if logger.printWriter != nil {
		t.Errorf("logger's printWriter %s is wrong", logger.printWriter)
	}
}

// go test -v -cover -run=^TestOptionsWithErrorWriter$
func TestOptionsWithErrorWriter(t *testing.T) {
	opts := Options()

	logger := NewLogger()
	logger.debugWriter = nil
	logger.infoWriter = nil
	logger.warnWriter = nil
	logger.errorWriter = nil
	logger.printWriter = nil
	opts.WithErrorWriter(os.Stdout, false)(logger)

	if logger.debugWriter != nil {
		t.Errorf("logger's debugWriter %s is wrong", logger.debugWriter)
	}

	if logger.infoWriter != nil {
		t.Errorf("logger's infoWriter %s is wrong", logger.infoWriter)
	}

	if logger.warnWriter != nil {
		t.Errorf("logger's warnWriter %s is wrong", logger.warnWriter)
	}

	if logger.errorWriter == nil {
		t.Errorf("logger's errorWriter %s is wrong", logger.errorWriter)
	}

	if logger.printWriter != nil {
		t.Errorf("logger's printWriter %s is wrong", logger.printWriter)
	}
}

// go test -v -cover -run=^TestOptionsWithPrintWriter$
func TestOptionsWithPrintWriter(t *testing.T) {
	opts := Options()

	logger := NewLogger()
	logger.debugWriter = nil
	logger.infoWriter = nil
	logger.warnWriter = nil
	logger.errorWriter = nil
	logger.printWriter = nil
	opts.WithPrintWriter(os.Stdout, false)(logger)

	if logger.debugWriter != nil {
		t.Errorf("logger's debugWriter %s is wrong", logger.debugWriter)
	}

	if logger.infoWriter != nil {
		t.Errorf("logger's infoWriter %s is wrong", logger.infoWriter)
	}

	if logger.warnWriter != nil {
		t.Errorf("logger's warnWriter %s is wrong", logger.warnWriter)
	}

	if logger.errorWriter != nil {
		t.Errorf("logger's errorWriter %s is wrong", logger.errorWriter)
	}

	if logger.printWriter == nil {
		t.Errorf("logger's printWriter %s is wrong", logger.printWriter)
	}
}

// go test -v -cover -run=^TestOptionsWithPid$
func TestOptionsWithPid(t *testing.T) {
	opts := Options()

	logger := NewLogger()
	logger.needPid = false

	opts.WithPid()(logger)
	if logger.needPid != true {
		t.Errorf("logger's needPid %+v is wrong", logger.needPid)
	}
}

// go test -v -cover -run=^TestOptionsWithCaller$
func TestOptionsWithCaller(t *testing.T) {
	opts := Options()

	logger := NewLogger()
	logger.needCaller = false

	opts.WithCaller()(logger)
	if logger.needCaller != true {
		t.Errorf("logger's needCaller %+v is wrong", logger.needCaller)
	}
}

// go test -v -cover -run=^TestOptionsWithMsgKey$
func TestOptionsWithMsgKey(t *testing.T) {
	opts := Options()

	logger := NewLogger()
	logger.msgKey = ""

	opts.WithMsgKey("msg")(logger)
	if logger.msgKey != "msg" {
		t.Errorf("logger's msgKey %+v is wrong", logger.msgKey)
	}
}

// go test -v -cover -run=^TestOptionsWithTimeKey$
func TestOptionsWithTimeKey(t *testing.T) {
	opts := Options()

	logger := NewLogger()
	logger.timeKey = ""

	opts.WithTimeKey("time")(logger)
	if logger.timeKey != "time" {
		t.Errorf("logger's timeKey %+v is wrong", logger.timeKey)
	}
}

// go test -v -cover -run=^TestOptionsWithLevelKey$
func TestOptionsWithLevelKey(t *testing.T) {
	opts := Options()

	logger := NewLogger()
	logger.levelKey = ""

	opts.WithLevelKey("level")(logger)
	if logger.levelKey != "level" {
		t.Errorf("logger's levelKey %+v is wrong", logger.levelKey)
	}
}

// go test -v -cover -run=^TestOptionsWithPidKey$
func TestOptionsWithPidKey(t *testing.T) {
	opts := Options()

	logger := NewLogger()
	logger.pidKey = ""

	opts.WithPidKey("pid")(logger)
	if logger.pidKey != "pid" {
		t.Errorf("logger's pidKey %+v is wrong", logger.pidKey)
	}
}

// go test -v -cover -run=^TestOptionsWithFileKey$
func TestOptionsWithFileKey(t *testing.T) {
	opts := Options()

	logger := NewLogger()
	logger.fileKey = ""

	opts.WithFileKey("file")(logger)
	if logger.fileKey != "file" {
		t.Errorf("logger's fileKey %+v is wrong", logger.fileKey)
	}
}

// go test -v -cover -run=^TestOptionsWithLineKey$
func TestOptionsWithLineKey(t *testing.T) {
	opts := Options()

	logger := NewLogger()
	logger.lineKey = ""

	opts.WithLineKey("line")(logger)
	if logger.lineKey != "line" {
		t.Errorf("logger's lineKey %+v is wrong", logger.lineKey)
	}
}

// go test -v -cover -run=^TestOptionsWithFuncKey$
func TestOptionsWithFuncKey(t *testing.T) {
	opts := Options()

	logger := NewLogger()
	logger.funcKey = ""

	opts.WithFuncKey("func")(logger)
	if logger.funcKey != "func" {
		t.Errorf("logger's funcKey %+v is wrong", logger.funcKey)
	}
}

// go test -v -cover -run=^TestOptionsWithTimeFormat$
func TestOptionsWithTimeFormat(t *testing.T) {
	opts := Options()

	logger := NewLogger()
	logger.timeFormat = ""

	opts.WithTimeFormat("20060102150405")(logger)
	if logger.timeFormat != "20060102150405" {
		t.Errorf("logger's timeFormat %+v is wrong", logger.timeFormat)
	}
}

// go test -v -cover -run=^TestOptionsWithCallerDepth$
func TestOptionsWithCallerDepth(t *testing.T) {
	opts := Options()

	logger := NewLogger()
	logger.callerDepth = 0

	opts.WithCallerDepth(3)(logger)
	if logger.callerDepth != 3 {
		t.Errorf("logger's callerDepth %d is wrong", logger.callerDepth)
	}
}

// go test -v -cover -run=^TestOptionsWithInterceptors$
func TestOptionsWithInterceptors(t *testing.T) {
	opts := Options()

	logger := NewLogger()
	logger.interceptors = nil

	interceptors := []Interceptor{
		func(ctx context.Context, log *Log) {},
		func(ctx context.Context, log *Log) {},
		func(ctx context.Context, log *Log) {},
	}

	opts.WithInterceptors(interceptors...)(logger)
	if len(logger.interceptors) != len(interceptors) {
		t.Errorf("len(logger.interceptors) %d != len(interceptors) %d", len(logger.interceptors), len(interceptors))
	}

	logger.interceptors = nil

	opts.WithInterceptors()
	if len(logger.interceptors) != 0 {
		t.Errorf("len(logger.interceptors) %d != 0", len(logger.interceptors))
	}
}
