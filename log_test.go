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
	"context"
	"encoding/json"
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/go-logit/logit/core/appender"
	"github.com/go-logit/logit/support/global"

	"github.com/go-logit/logit/support/runtime"
)

// go test -v -cover -run=^TestNewLog$
func TestNewLog(t *testing.T) {
	buffer := bytes.NewBuffer(make([]byte, 0, 1024))

	logger := NewLogger(Options().WithAppender(appender.Json()), Options().WithWriter(buffer))
	defer logger.Close()

	entries := map[string]interface{}{
		"Any":       map[int]string{1: "a", 2: "b", 3: "c"},
		"Bool":      true,
		"Byte":      'x',
		"Rune":      '国',
		"Int":       -123456,
		"Int8":      int8(-123),
		"Int16":     int16(-12345),
		"Int32":     int32(-123456),
		"Int64":     int64(-123456),
		"Uint":      uint(123456),
		"Uint8":     uint8(123),
		"Uint16":    uint16(12345),
		"Uint32":    uint32(123456),
		"Uint64":    uint64(123456),
		"String":    "abc",
		"Time":      time.Unix(12580, 0).Format("2006-01-02 15:04:05"),
		"Error":     io.EOF,
		"Stringer":  fmt.Stringer(time.Second),
		"Bools":     []bool{true},
		"Bytes":     []byte{'x'},
		"Runes":     []rune{'国'},
		"Ints":      []int{-123456},
		"Int8s":     []int8{-123},
		"Int16s":    []int16{-12345},
		"Int32s":    []int32{-123456},
		"Int64s":    []int64{-123456},
		"Uints":     []uint{123456},
		"Uint8s":    []uint8{123},
		"Uint16s":   []uint16{12345},
		"Uint32s":   []uint32{123456},
		"Uint64s":   []uint64{123456},
		"Strings":   []string{"abc"},
		"Times":     []string{time.Unix(12580, 0).Format("2006-01-02 15:04:05")},
		"Errors":    []error{io.EOF},
		"Stringers": []fmt.Stringer{time.Second},
		"WithTime":  time.Unix(12580, 0).Unix(),
	}

	//t.Logf("entries: %+v", entries)

	log := newLog()
	log.logger = logger
	log.appender = logger.debugAppender
	log.writer = logger.debugWriter
	log.begin()
	log.Any("Any", map[int]string{1: "a", 2: "b", 3: "c"})
	log.Bool("Bool", true)
	log.Byte("Byte", 'x')
	log.Rune("Rune", '国')
	log.Int("Int", -123456)
	log.Int8("Int8", int8(-123))
	log.Int16("Int16", int16(-12345))
	log.Int32("Int32", int32(-123456))
	log.Int64("Int64", int64(-123456))
	log.Uint("Uint", uint(123456))
	log.Uint8("Uint8", uint8(123))
	log.Uint16("Uint16", uint16(12345))
	log.Uint32("Uint32", uint32(123456))
	log.Uint64("Uint64", uint64(123456))
	log.String("String", "abc")
	log.Time("Time", time.Unix(12580, 0))
	log.Error("Error", io.EOF)
	log.Stringer("Stringer", fmt.Stringer(time.Second))
	log.Bools("Bools", []bool{true})
	log.Bytes("Bytes", []byte{'x'})
	log.Runes("Runes", []rune("国"))
	log.Ints("Ints", []int{-123456})
	log.Int8s("Int8s", []int8{-123})
	log.Int16s("Int16s", []int16{-12345})
	log.Int32s("Int32s", []int32{-123456})
	log.Int64s("Int64s", []int64{-123456})
	log.Uints("Uints", []uint{123456})
	log.Uint8s("Uint8s", []uint8{123})
	log.Uint16s("Uint16s", []uint16{12345})
	log.Uint32s("Uint32s", []uint32{123456})
	log.Uint64s("Uint64s", []uint64{123456})
	log.Strings("Strings", []string{"abc"})
	log.Times("Times", []time.Time{time.Unix(12580, 0)})
	log.Errors("Errors", []error{io.EOF})
	log.Stringers("Stringers", []fmt.Stringer{time.Second})
	log.WithTime("WithTime", time.Unix(12580, 0), global.UnixTimeFormat)
	log.Log()

	outputMap := map[string]interface{}{}
	output := buffer.String()

	if err := json.Unmarshal(buffer.Bytes(), &outputMap); err != nil {
		t.Errorf("unmarshal output %+v from Json failed", output)
	}

	//t.Logf("outputMap: %+v", outputMap)
	for k, v := range entries {
		outputValue, ok := outputMap[k]
		if !ok {
			t.Errorf("outputMap missed key %s", k)
		}

		switch ov := outputValue.(type) {
		case byte:
			if ov != v.(byte) {
				t.Errorf("key %s: outputValue %v is wrong with %v", k, outputValue, v)
			}
		case rune:
			if ov != v.(rune) {
				t.Errorf("key %s: outputValue %v is wrong with %v", k, outputValue, v)
			}
		case []byte:
			for i, c := range ov {
				if c != v.([]byte)[i] {
					t.Errorf("key %s: outputValue %v is wrong with %v", k, outputValue, v)
				}
			}
		case []rune:
			for i, r := range ov {
				if r != v.([]rune)[i] {
					t.Errorf("key %s: outputValue %v is wrong with %v", k, outputValue, v)
				}
			}
		}

	}
}

// go test -v -cover -run=^TestLogWithPID$
func TestLogWithPID(t *testing.T) {
	buffer := bytes.NewBuffer(make([]byte, 0, 1024))

	logger := NewLogger(Options().WithWriter(buffer))
	logger.withPID = true

	log := newLog()
	log.logger = logger
	log.appender = logger.debugAppender
	log.writer = logger.debugWriter
	log.begin()
	log.WithPID()
	log.Log()

	str := buffer.String()
	if str != "\n" {
		t.Errorf("str %q != '\n'", str)
	}

	buffer.Reset()
	logger.withPID = false
	log.begin()
	log.WithPID()
	log.Log()

	pid := runtime.PID()
	right := fmt.Sprintf("%s=%d\n", logger.pidKey, pid)

	str = buffer.String()
	if str != right {
		t.Errorf("str %s != right %s", str, right)
	}
}

// go test -v -cover -run=^TestLogWithCaller$
func TestLogWithCaller(t *testing.T) {
	buffer := bytes.NewBuffer(make([]byte, 0, 1024))

	logger := NewLogger(Options().WithWriter(buffer))
	logger.withCaller = true

	log := newLog()
	log.logger = logger
	log.appender = logger.debugAppender
	log.writer = logger.debugWriter

	log.begin()
	log.WithCaller()
	log.Log()

	str := buffer.String()
	if str != "\n" {
		t.Errorf("str %q != '\n'", str)
	}

	buffer.Reset()
	logger.withCaller = false

	log.begin()
	log.WithCaller()
	log.Log()

	file, line, function := runtime.Caller(1)
	line -= 3 // Between log.WithCaller() and pkg.Caller(1) is 3
	right := fmt.Sprintf("%s=%s|%s=%d|%s=%s\n", logger.fileKey, file, logger.lineKey, line, logger.funcKey, function)

	str = buffer.String()
	if str != right {
		t.Errorf("str %s != right %s", str, right)
	}
}

// go test -v -cover -run=^TestLogWithContext$
func TestLogWithContext(t *testing.T) {
	log := newLog()
	log.WithContext(context.WithValue(context.Background(), "key", "value"))

	value, ok := log.ctx.Value("key").(string)
	if !ok || value != "value" {
		t.Errorf("!ok %v || value %s != 'value'", ok, value)
	}
}

// go test -v -cover -run=^TestLogIntercept$
func TestLogIntercept(t *testing.T) {
	buffer := bytes.NewBuffer(make([]byte, 0, 1024))

	logger := NewLogger(Options().WithWriter(buffer), Options().WithInterceptors(func(ctx context.Context, log *Log) {
		log.Int("xxx", 123)
	}))
	logger.withCaller = true

	log := newLog()
	log.logger = logger
	log.appender = logger.debugAppender
	log.writer = logger.debugWriter

	log.begin()
	log.Log()

	str := buffer.String()
	if str != "xxx=123\n" {
		t.Errorf("str %q != '\n'", str)
	}

	buffer.Reset()
	log = newLog()
	log.logger = logger
	log.appender = logger.debugAppender
	log.writer = logger.debugWriter

	log.begin()
	log.Intercept(func(ctx context.Context, log *Log) {
		log.String("abc", "666")
	})
	log.Log()

	str = buffer.String()
	if str != "abc=666|xxx=123\n" {
		t.Errorf("str %q != '\n'", str)
	}
}
