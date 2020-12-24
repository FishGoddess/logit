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
// Created at 2020/03/29 22:59:14

package logit

import "testing"

// go test -v -cover -run=^TestLevelOf$
func TestLevelOf(t *testing.T) {

	if level, err := levelOf("notExistedLevel"); err == nil {
		t.Fatalf("level (%s) should not be returned", level)
	}

	if level, err := levelOf("info"); err != nil || level != InfoLevel {
		t.Fatalf("returned level (%s) is wrong", level)
	}
}

// go test -v -cover -run=^TestEncodeOf$
func TestEncodeOf(t *testing.T) {

	if _, err := encoderOf("notExistedEncoder"); err == nil {
		t.Fatal("encoder should not be returned")
	}

	if _, err := encoderOf("text"); err != nil {
		t.Fatal("failed to get encoder")
	}
}
