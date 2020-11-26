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
// Created at 2020/11/26 23:26:22

package logit

import "testing"

// go test -v -cover -run=^TestLevelString$
func TestLevelString(t *testing.T) {

	if DebugLevel.String() != "debug" {
		t.Fatalf("debug level %s is wrong", DebugLevel.String())
	}

	if InfoLevel.String() != "info" {
		t.Fatalf("info level %s is wrong", InfoLevel.String())
	}

	if WarnLevel.String() != "warn" {
		t.Fatalf("warn level %s is wrong", WarnLevel.String())
	}

	if ErrorLevel.String() != "error" {
		t.Fatalf("error level %s is wrong", ErrorLevel.String())
	}
}
