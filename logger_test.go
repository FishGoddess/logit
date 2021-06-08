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
// Email: fishgoddess@qq.com
// Created at 2020/02/29 16:41:41

package logit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"runtime"
	"testing"
)

// go test -v -cover -run=^TestLoggerLevel$
func TestLoggerLevel(t *testing.T) {

	logger := NewLogger()

	if logger.Level() != logger.level.Load().(Level) {
		t.Fatalf("level returned %v is wrong", logger.Level())
	}

	logger.SetLevel(InfoLevel)
	if logger.level.Load().(Level) != InfoLevel {
		t.Fatalf("level set %+v is wrong", logger.level.Load().(Level))
	}
}

// go test -v -cover -run=^TestLoggerNeedCaller$
func TestLoggerNeedCaller(t *testing.T) {

	logger := NewLogger()

	logger.SetNeedCaller(true)
	if logger.needCaller.Load().(bool) != true {
		t.Fatalf("needCaller set %+v is wrong", logger.needCaller.Load().(bool))
	}
}

type testLoggerSetEncoder struct{}

func (testLoggerSetEncoder) Encode(log *Log) []byte { return []byte("TestLoggerSetEncoder") }

// go test -v -cover -run=^TestLoggerSetEncoder$
func TestLoggerSetEncoder(t *testing.T) {

	logger := NewLogger()
	logger.Encoders().SetEncoder(&testLoggerSetEncoder{})

	encoders := logger.Encoders()
	for level, _ := range levels {
		encoded := string(encoders.of(level).Encode(nil))
		if encoded != "TestLoggerSetEncoder" {
			t.Fatalf("encoded %s of level %v is wrong", encoded, level)
		}
	}
}

type testLoggerSetByLevelEncoder struct{}

func (testLoggerSetByLevelEncoder) Encode(log *Log) []byte { return []byte(log.level.String()) }

// go test -v -cover -run=^TestLoggerSetEncoderByLevel$
func TestLoggerSetEncoderByLevel(t *testing.T) {

	logger := NewLogger()
	logger.Encoders().SetDebugEncoder(&testLoggerSetByLevelEncoder{})
	logger.Encoders().SetInfoEncoder(&testLoggerSetByLevelEncoder{})
	logger.Encoders().SetWarnEncoder(&testLoggerSetByLevelEncoder{})
	logger.Encoders().SetErrorEncoder(&testLoggerSetByLevelEncoder{})

	encoders := logger.encoders
	for level, _ := range levels {
		encoded := string(encoders.of(level).Encode(&Log{level: level}))
		if encoded != level.String() {
			t.Fatalf("encoded %s of level %v is wrong", encoded, level)
		}
	}
}

// go test -v -cover -run=^TestLoggerSetWriter$
func TestLoggerSetWriter(t *testing.T) {

	logger := NewLogger()

	buffer := bytes.NewBuffer(make([]byte, 0))
	logger.Writers().SetWriter(buffer)

	writers := logger.Writers()
	for level, _ := range levels {
		_, err := writers.of(level).Write([]byte("1"))
		if err != nil {
			t.Fatal(err)
		}
	}

	s := buffer.String()
	if s != "11111" {
		t.Fatalf("write %s to buffer is wrong", s)
	}
}

// go test -v -cover -run=^TestLoggerSetWriterByLevel$
func TestLoggerSetWriterByLevel(t *testing.T) {

	logger := NewLogger()

	writers := map[Level]*bytes.Buffer{
		DebugLevel: bytes.NewBuffer(make([]byte, 0)),
		InfoLevel:  bytes.NewBuffer(make([]byte, 0)),
		WarnLevel:  bytes.NewBuffer(make([]byte, 0)),
		ErrorLevel: bytes.NewBuffer(make([]byte, 0)),
	}

	logger.Writers().SetDebugWriter(writers[DebugLevel])
	logger.Writers().SetInfoWriter(writers[InfoLevel])
	logger.Writers().SetWarnWriter(writers[WarnLevel])
	logger.Writers().SetErrorWriter(writers[ErrorLevel])

	for level, _ := range writers {
		_, err := logger.Writers().of(level).Write([]byte(level.String()))
		if err != nil {
			t.Fatal(err)
		}
	}

	for level, writer := range writers {
		s := writer.String()
		if s != level.String() {
			t.Fatalf("write %s to buffer of level %v is wrong", s, level)
		}
	}
}

type testLoggerCoreEncoder struct{}

func (testLoggerCoreEncoder) Encode(log *Log) []byte { return []byte(log.msg) }

// go test -v -cover -run=^TestLoggerCore$
func TestLoggerCore(t *testing.T) {

	logger := NewLogger()
	logger.SetLevel(DebugLevel)

	buffer := bytes.NewBuffer(make([]byte, 0))
	logger.Encoders().SetEncoder(&testLoggerCoreEncoder{})
	logger.Writers().SetWriter(buffer)

	logger.Debug(DebugLevel.String())
	logger.Info(InfoLevel.String())
	logger.Warn(WarnLevel.String())
	logger.Error(ErrorLevel.String())

	s := buffer.String()
	if s != DebugLevel.String()+InfoLevel.String()+WarnLevel.String()+ErrorLevel.String() {
		t.Fatalf("core log %s is wrong", s)
	}
}

type testLoggerCoreParamsEncoder struct{}

func (testLoggerCoreParamsEncoder) Encode(log *Log) []byte { return []byte(log.msg) }

// go test -v -cover -run=^TestLoggerCoreParams$
func TestLoggerCoreParams(t *testing.T) {

	logger := NewLogger()
	logger.SetLevel(DebugLevel)

	buffer := bytes.NewBuffer(make([]byte, 0))
	logger.Encoders().SetEncoder(&testLoggerCoreParamsEncoder{})
	logger.Writers().SetWriter(buffer)

	logger.Debug("%sF", DebugLevel.String())
	logger.Info("%sF", InfoLevel.String())
	logger.Warn("%sF", WarnLevel.String())
	logger.Error("%sF", ErrorLevel.String())

	s := buffer.String()
	if s != fmt.Sprintf("%sF%sF%sF%sF", DebugLevel.String(), InfoLevel.String(), WarnLevel.String(), ErrorLevel.String()) {
		t.Fatalf("coreParams log %s is wrong", s)
	}
}

// go test -v -cover -run=^TestLoggerWithCaller$
func TestLoggerWithCaller(t *testing.T) {

	logger := NewLogger()
	logger.SetNeedCaller(true)

	buffer := bytes.NewBuffer(make([]byte, 0))
	logger.Encoders().SetEncoder(NewJsonEncoder(""))
	logger.Writers().SetWriter(buffer)
	logger.Info("xxx")

	_, file, line, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("get caller failed")
	}
	t.Logf("file: %s, ling: %d", file, line)

	logStr := buffer.String()
	log := map[string]interface{}{}
	err := json.Unmarshal([]byte(logStr), &log)
	if err != nil {
		t.Fatalf("unmarshal json %s failed", logStr)
	}
	t.Logf("log: %+v", log)

	if log["file"] != file {
		t.Fatalf("file %s in caller is wrong, correct: %s", log["file"], file)
	}

	if int(log["line"].(float64)) != line - 2 {
		t.Fatalf("line %d in caller is wrong, correct: %d", log["line"], line-2)
	}
}
