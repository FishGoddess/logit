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
// Created at 2021/07/11 14:03:47

package logit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/FishGoddess/logit/core/appender"
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
		"Rune":      rune('国'),
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
		"Time":      time.Unix(12580, 0).Unix(),
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
		"Times":     []int64{time.Unix(12580, 0).Unix()},
		"Errors":    []error{io.EOF},
		"Stringers": []fmt.Stringer{time.Second},
	}
	//t.Logf("entries: %+v", entries)

	log := newLog(logger)
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
	log.Time("Time", time.Unix(12580, 0), appender.UnixTime)
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
	log.Times("Times", []time.Time{time.Unix(12580, 0)}, appender.UnixTime)
	log.Errors("Errors", []error{io.EOF})
	log.Stringers("Stringers", []fmt.Stringer{time.Second})
	log.End()

	outputMap := map[string]interface{}{}
	output := buffer.String()
	err := json.Unmarshal(buffer.Bytes(), &outputMap)
	if err != nil {
		t.Fatalf("unmarshal output %+v from Json failed", output)
	}

	//t.Logf("outputMap: %+v", outputMap)
	for k, v := range entries {

		outputValue, ok := outputMap[k]
		if !ok {
			t.Fatalf("outputMap missed key %s", k)
		}

		switch ov := outputValue.(type) {
		case byte:
			if ov != v.(byte) {
				t.Fatalf("key %s: outputValue %v is wrong with %v", k, outputValue, v)
			}
		case rune:
			if ov != v.(rune) {
				t.Fatalf("key %s: outputValue %v is wrong with %v", k, outputValue, v)
			}
		case []byte:
			for i, c := range ov {
				if c != v.([]byte)[i] {
					t.Fatalf("key %s: outputValue %v is wrong with %v", k, outputValue, v)
				}
			}
		case []rune:
			for i, r := range ov {
				if r != v.([]rune)[i] {
					t.Fatalf("key %s: outputValue %v is wrong with %v", k, outputValue, v)
				}
			}
		}

	}
}
