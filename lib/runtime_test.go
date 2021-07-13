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
// Created at 2021/07/13 23:35:26

package lib

import (
	"os"
	"runtime"
	"testing"
)

// go test -v -cover -run=^TestPid$
func TestPid(t *testing.T) {

	pid := Pid()
	osPid := os.Getpid()
	if pid != osPid {
		t.Fatalf("pid %d is wrong with %d", pid, osPid)
	}
}

// go test -v -cover -run=^TestCaller$
func TestCaller(t *testing.T) {

	_, f, l, ok := runtime.Caller(0)
	if !ok {
		t.Fatalf("runtime.Caller failed")
	}

	file, line := Caller(1)
	if file != f {
		t.Fatalf("Caller returns wrong file %s", file)
	}

	if line != l+5 {
		t.Fatalf("Caller returns wrong line %d", line)
	}
}
