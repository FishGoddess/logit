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
// Created at 2020/04/24 16:43:49

package logit

import (
	"fmt"
	"testing"
	"time"
)

// prepareTestLog prepares one log for testing.
func prepareTestLog() *Log {
	return &Log{
		msg:    "test",
		level:  InfoLevel,
		time:   time.Date(2020, time.November, 13, 19, 43, 23, 0, time.Local),
		caller: nil,
	}
}

// go test -v -cover -run=^TestTextEncoder$
func TestTextEncoder(t *testing.T) {

	encoder := NewTextEncoder("2006-01-02 15:04:05")

	log := prepareTestLog()
	logString := string(encoder.Encode(log))
	result := "[info] [2020-11-13 19:43:23] test\n"
	if logString != result {
		t.Fatalf("encoded log (%s) is wrong", logString)
	}

	log.level = ErrorLevel
	log.caller = &caller{
		File: "encoder_test.go",
		Line: 36,
	}

	encoder = NewTextEncoder("")
	logString = string(encoder.Encode(log))
	result = fmt.Sprintf("[error] [%d] [encoder_test.go:36] test\n", log.time.Unix())
	if logString != result {
		t.Fatalf("encoded log (%s) is wrong", logString)
	}
}

// go test -v -cover -run=^TestJsonEncoder$
func TestJsonEncoder(t *testing.T) {

	encoder := NewJsonEncoder("2006/01/02 15:04:05")

	log := prepareTestLog()
	logString := string(encoder.Encode(log))
	result := "{\"level\":\"info\",\"time\":\"2020/11/13 19:43:23\",\"msg\":\"test\"}\n"
	if logString != result {
		t.Fatalf("encoded log (%s) is wrong", logString)
	}

	log.level = ErrorLevel
	log.caller = &caller{
		File: "encoder_test.go",
		Line: 36,
	}

	encoder = NewJsonEncoder("")
	logString = string(encoder.Encode(log))
	result = fmt.Sprintf("{\"level\":\"error\",\"time\":%d,\"file\":\"encoder_test.go\",\"line\":36,\"msg\":\"test\"}\n", log.time.Unix())
	if logString != result {
		t.Fatalf("encoded log (%s) is wrong", logString)
	}
}
