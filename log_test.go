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
// Created at 2020/11/26 23:19:26

package logit

import (
	"testing"
	"time"
)

// go test -v -cover -run=^TestLog$
func TestLog(t *testing.T) {

	log := &Log{
		msg:    "test",
		level:  DebugLevel,
		time:   time.Now(),
		caller: nil,
	}

	if log.Msg() != log.msg {
		t.Fatalf("msg returned %s is wrong!", log.Msg())
	}

	if log.Level() != log.level {
		t.Fatalf("level returned %d is wrong!", log.Level())
	}

	if log.Time().Unix() != log.time.Unix() {
		t.Fatalf("time returned %v is wrong!", log.Time())
	}

	if caller, ok := log.Caller(); ok {
		t.Fatalf("caller %v is wrong", caller)
	}

	log.caller = &caller{
		File: "file.go",
		Line: 123,
	}

	caller, ok := log.Caller()
	if !ok || caller.File != "file.go" || caller.Line != 123 {
		t.Fatalf("caller %v is wrong", caller)
	}
}