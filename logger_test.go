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
// Created at 2021/06/27 16:41:20

package logit

import (
	"bytes"
	"errors"
	"testing"

	"github.com/FishGoddess/logit/core/appender"
)

// go test -v -cover -run=^TestNewLogger$
func TestNewLogger(t *testing.T) {
	buffer := bytes.NewBuffer(make([]byte, 0, 1024))

	options := Options()
	logger := NewLogger(
		options.WithDebugLevel(),
		options.WithAppender(appender.Json()),
		options.WithWriter(buffer, false),
		//options.WithPid(),
		//options.WithCaller(),
		//options.WithMsgKey("message"),
		options.WithTimeKey(""),
		//options.WithLevelKey("level"),
		//options.WithPidKey("pid"),
		//options.WithFileKey("file"),
		//options.WithLineKey("line"),
		//options.WithTimeFormat("060102"),
	)
	defer logger.Close()

	logger.Debug("debug...").String("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Any("any", map[string]interface{}{"a": 1, "b": "bbb"}).End()
	logger.Error("error...").Byte("b", 'a').Byte("es", '\n').Runes("words", []rune("我是中国人")).Error("err", errors.New("我是错误")).End()
	logger.Error("error with %d...", 666).String("trace", "xxx").Int("id", 123).Float64("pi", 3.14).End()
	logger.Warn("\"warn\"...\r\b\t\n").Strings("s\tb\nd\b", []string{"abc\r", "efg\n"}).End()
	logger.Info("info...").Bools("bools", []bool{true, false}).Bytes("bytes", []byte{'\b', '\t', 'a', 'b', 'c', '"', '\n'}).Int16s("int16s", []int16{123, 4567, 8901}).Float32s("float32s", []float32{3.14, 6.18}).End()

	logs := `{"log.level":"debug","log.msg":"debug...","trace":"xxx","id":123,"pi":3.14,"any":{"a":1,"b":"bbb"}}
{"log.level":"error","log.msg":"error...","b":"a","es":"\n","words":["我","是","中","国","人"],"err":"我是错误"}
{"log.level":"error","log.msg":"error with 666...","trace":"xxx","id":123,"pi":3.14}
{"log.level":"warn","log.msg":"\"warn\"...\r\b\t\n","s\tb\nd\b":["abc\r","efg\n"]}
{"log.level":"info","log.msg":"info...","bools":[true,false],"bytes":["\b","\t","a","b","c","\"","\n"],"int16s":[123,4567,8901],"float32s":[3.140000104904175,6.179999828338623]}
`

	output := buffer.String()
	if output != logs {
		t.Errorf("logs %s is wrong with %s", output, logs)
	}
}

// go test -v -cover -run=^TestLoggerPrintf$
func TestLoggerPrintf(t *testing.T) {
	buffer := bytes.NewBuffer(make([]byte, 0, 1024))

	options := Options()
	logger := NewLogger(
		options.WithErrorLevel(),
		options.WithAppender(appender.Json()),
		options.WithWriter(buffer, false),
		options.WithTimeKey(""),
	)

	logger.Printf("printf%d", 123)
	logger.Print("print", 666)
	logger.Println("println", 999)

	output := buffer.String()
	logs := `{"log.level":"print","log.msg":"printf123"}
{"log.level":"print","log.msg":"print666"}
{"log.level":"print","log.msg":"println 999\n"}
`
	if output != logs {
		t.Errorf("logs %s is wrong with %s", output, logs)
	}
}

// go test -v -cover -run=^TestLoggerFlush$
func TestLoggerFlush(t *testing.T) {
	logger := NewLogger()

	_, err := logger.Flush()
	if err != nil {
		t.Error(err)
	}
}

// go test -v -cover -run=^TestLoggerClose$
func TestLoggerClose(t *testing.T) {
	logger := NewLogger()

	err := logger.Close()
	if err != nil {
		t.Error(err)
	}

	if logger.level != offLevel {
		t.Errorf("level of logger %+v is wrong", logger.level)
	}
}
