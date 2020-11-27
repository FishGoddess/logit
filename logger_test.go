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

// go test -v -cover -run=^TestLoggerTimeFormat$
func TestLoggerTimeFormat(t *testing.T) {

	logger := NewLogger()

	logger.TimeFormat("2006/01/02 15:04:05")
	if logger.timeFormat != "2006/01/02 15:04:05" {
		t.Fatal("timeFormat set is wrong")
	}
}

// go test -v -cover -run=^TestLoggerSetEncoder$
func TestLoggerSetEncoder(t *testing.T) {

	logger := NewLogger()

	logger.SetEncoder(func(log *Log, timeFormat string) []byte {
		return []byte("TestLoggerSetEncoder")
	})

	encoders := logger.encoders
	for level, encoder := range encoders {
		encoded := string(encoder.Encode(nil, ""))
		if encoded != "TestLoggerSetEncoder" {
			t.Fatalf("encoded %s of level %v is wrong", encoded, level)
		}
	}
}

// go test -v -cover -run=^TestLoggerSetEncoderByLevel$
func TestLoggerSetEncoderByLevel(t *testing.T) {

	logger := NewLogger()

	logger.SetDebugEncoder(func(log *Log, timeFormat string) []byte {
		return []byte(DebugLevel.String())
	})

	logger.SetInfoEncoder(func(log *Log, timeFormat string) []byte {
		return []byte(InfoLevel.String())
	})

	logger.SetWarnEncoder(func(log *Log, timeFormat string) []byte {
		return []byte(WarnLevel.String())
	})

	logger.SetErrorEncoder(func(log *Log, timeFormat string) []byte {
		return []byte(ErrorLevel.String())
	})

	encoders := logger.encoders
	for level, encoder := range encoders {
		encoded := string(encoder.Encode(nil, ""))
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

// go test -v -cover -run=^TestLoggerCore$
func TestLoggerCore(t *testing.T) {

	logger := NewLogger()
	logger.SetLevel(DebugLevel)

	buffer := bytes.NewBuffer(make([]byte, 0))
	logger.SetEncoder(func(log *Log, timeFormat string) []byte {
		return []byte(log.msg)
	})
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

// go test -v -cover -run=^TestLoggerCoreF$
func TestLoggerCoreF(t *testing.T) {

	logger := NewLogger()
	logger.SetLevel(DebugLevel)

	buffer := bytes.NewBuffer(make([]byte, 0))
	logger.SetEncoder(func(log *Log, timeFormat string) []byte {
		return []byte(log.msg)
	})
	logger.SetWriter(buffer)

	logger.DebugF("%sF", DebugLevel.String())
	logger.InfoF("%sF", InfoLevel.String())
	logger.WarnF("%sF", WarnLevel.String())
	logger.ErrorF("%sF", ErrorLevel.String())

	s := buffer.String()
	if s != fmt.Sprintf("%sF%sF%sF%sF", DebugLevel.String(), InfoLevel.String(), WarnLevel.String(), ErrorLevel.String()) {
		t.Fatalf("coreF log %s is wrong", s)
	}
}
