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
	"fmt"
	"testing"
)

// go test -v -cover -run=^TestLoggerLevel$
func TestLoggerLevel(t *testing.T) {

	logger := NewLogger()

	if logger.Level() != logger.level {
		t.Fatalf("level returned %v is wrong", logger.Level())
	}

	realOldLevel := logger.level
	if oldLevel := logger.SetLevel(InfoLevel); oldLevel != realOldLevel || logger.level != InfoLevel {
		t.Fatalf("level set or old level %v is wrong", oldLevel)
	}
}

// go test -v -cover -run=^TestLoggerNeedCaller$
func TestLoggerNeedCaller(t *testing.T) {

	logger := NewLogger()

	logger.NeedCaller(true)
	if logger.needCaller != true {
		t.Fatal("needCaller set is wrong")
	}
}

type testLoggerSetEncoder struct{}

func (testLoggerSetEncoder) Encode(log *Log) []byte { return []byte("TestLoggerSetEncoder") }

// go test -v -cover -run=^TestLoggerSetEncoder$
func TestLoggerSetEncoder(t *testing.T) {

	logger := NewLogger()
	logger.SetEncoder(&testLoggerSetEncoder{})

	encoders := logger.encoders
	for level, encoder := range encoders {
		encoded := string(encoder.Encode(nil))
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
	logger.SetDebugEncoder(&testLoggerSetByLevelEncoder{})
	logger.SetInfoEncoder(&testLoggerSetByLevelEncoder{})
	logger.SetWarnEncoder(&testLoggerSetByLevelEncoder{})
	logger.SetErrorEncoder(&testLoggerSetByLevelEncoder{})

	encoders := logger.encoders
	for level, encoder := range encoders {
		encoded := string(encoder.Encode(&Log{level: level}))
		if encoded != level.String() {
			t.Fatalf("encoded %s of level %v is wrong", encoded, level)
		}
	}
}

// go test -v -cover -run=^TestLoggerSetWriter$
func TestLoggerSetWriter(t *testing.T) {

	logger := NewLogger()

	buffer := bytes.NewBuffer(make([]byte, 0))
	logger.SetWriter(buffer)

	writers := logger.writers
	for _, writer := range writers {
		_, err := writer.Write([]byte("1"))
		if err != nil {
			t.Fatal(err)
		}
	}

	s := buffer.String()
	if s != "1111" {
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

	logger.SetDebugWriter(writers[DebugLevel])
	logger.SetInfoWriter(writers[InfoLevel])
	logger.SetWarnWriter(writers[WarnLevel])
	logger.SetErrorWriter(writers[ErrorLevel])

	for level, writer := range logger.writers {
		_, err := writer.Write([]byte(level.String()))
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
	logger.SetEncoder(&testLoggerCoreEncoder{})
	logger.SetWriter(buffer)

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
	logger.SetEncoder(&testLoggerCoreParamsEncoder{})
	logger.SetWriter(buffer)

	logger.Debug("%sF", DebugLevel.String())
	logger.Info("%sF", InfoLevel.String())
	logger.Warn("%sF", WarnLevel.String())
	logger.Error("%sF", ErrorLevel.String())

	s := buffer.String()
	if s != fmt.Sprintf("%sF%sF%sF%sF", DebugLevel.String(), InfoLevel.String(), WarnLevel.String(), ErrorLevel.String()) {
		t.Fatalf("coreParams log %s is wrong", s)
	}
}
