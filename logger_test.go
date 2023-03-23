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
	"bytes"
	"errors"
	"testing"

	"github.com/FishGoddess/logit/core/appender"
)

// go test -v -cover -run=^TestNewLogger$
func TestNewLogger(t *testing.T) {
	buffer := bytes.NewBuffer(make([]byte, 0, 1024))

	options := []Option{
		Options().WithDebugLevel(),
		Options().WithAppender(appender.Json()),
		Options().WithWriter(buffer),
		Options().WithCaller(),
		Options().WithMsgKey("msg"),
		Options().WithTimeKey(""),
		Options().WithLevelKey("level"),
		Options().WithPIDKey("pid"),
		Options().WithFileKey("file"),
		Options().WithLineKey("line"),
		Options().WithFuncKey("func"),
		Options().WithErrorKey("err"),
		Options().WithTimeFormat("060102"),
	}

	logger := NewLogger(options...)
	defer logger.Close()

	logger.Debug("debug...").String("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Any("any", map[string]interface{}{"a": 1, "b": "bbb"}).Log()
	logger.Error(errors.New("我是错误"), "error...").Byte("b", 'a').Byte("es", '\n').Runes("words", []rune("我是中国人")).Log()
	logger.Error(nil, "error with %d...", 666).String("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Log()
	logger.Warn("\"warn\"...\r\b\t\n").Strings("s\tb\nd\b", []string{"abc\r", "efg\n"}).Log()
	logger.Info("info...").Bools("bools", []bool{true, false}).Bytes("bytes", []byte{'\b', '\t', 'a', 'b', 'c', '"', '\n'}).Int16s("int16s", []int16{123, 4567, 8901}).Float32s("float32s", []float32{3.14, 6.18}).Log()

	logs := `{"level":"debug","file":"D:/GoProject/go-logit/logit/logger_test.go","line":48,"func":"github.com/FishGoddess/logit.TestNewLogger","msg":"debug...","trace":"xxx","id":123,"pi":3.14,"any":{"a":1,"b":"bbb"}}
{"level":"error","file":"D:/GoProject/go-logit/logit/logger_test.go","line":49,"func":"github.com/FishGoddess/logit.TestNewLogger","msg":"error...","err":"我是错误","b":"a","es":"\n","words":["我","是","中","国","人"]}
{"level":"error","file":"D:/GoProject/go-logit/logit/logger_test.go","line":50,"func":"github.com/FishGoddess/logit.TestNewLogger","msg":"error with 666...","err":null,"trace":"xxx","id":123,"pi":3.14}
{"level":"warn","file":"D:/GoProject/go-logit/logit/logger_test.go","line":51,"func":"github.com/FishGoddess/logit.TestNewLogger","msg":"\"warn\"...\r\b\t\n","s\tb\nd\b":["abc\r","efg\n"]}
{"level":"info","file":"D:/GoProject/go-logit/logit/logger_test.go","line":52,"func":"github.com/FishGoddess/logit.TestNewLogger","msg":"info...","bools":[true,false],"bytes":["\b","\t","a","b","c","\"","\n"],"int16s":[123,4567,8901],"float32s":[3.140000104904175,6.179999828338623]}
`

	output := buffer.String()
	if output != logs {
		t.Errorf("logs %s is wrong with %s", output, logs)
	}
}

// go test -v -cover -run=^TestLoggerSetToGlobal$
func TestLoggerSetToGlobal(t *testing.T) {
	logger := NewLogger().SetToGlobal()

	if logger != globalLogger {
		t.Errorf("logger %p != globalLogger %p", logger, globalLogger)
	}

	if logger.callerDepth != globalLogger.callerDepth {
		t.Errorf("logger.callerDepth %d != globalLogger.callerDepth %d", logger.callerDepth, globalLogger.callerDepth)
	}
}

// go test -v -cover -run=^TestLoggerPrintf$
func TestLoggerPrintf(t *testing.T) {
	buffer := bytes.NewBuffer(make([]byte, 0, 1024))

	options := Options()
	logger := NewLogger(
		options.WithErrorLevel(),
		options.WithAppender(appender.Json()),
		options.WithWriter(buffer),
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

// go test -v -cover -run=^TestLoggerSync$
func TestLoggerSync(t *testing.T) {
	logger := NewLogger()

	if err := logger.Sync(); err != nil {
		t.Error(err)
	}
}

// go test -v -cover -run=^TestLoggerClose$
func TestLoggerClose(t *testing.T) {
	logger := NewLogger()

	if err := logger.Close(); err != nil {
		t.Error(err)
	}

	if logger.level != offLevel {
		t.Errorf("level of logger %+v is wrong", logger.level)
	}
}
