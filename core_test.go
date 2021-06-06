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
// Created at 2021/06/06 15:51:29

package logit

import (
	"os"
	"testing"
)

// go test -v -cover -run=^TestCore$
func TestCore(t *testing.T) {

	c := newCore(NewTextEncoder(TimeFormat), os.Stdout)
	c.SetLevel(WarnLevel)
	if c.Level() != WarnLevel {
		t.Fatalf("level %+v of core is wrong", c.Level())
	}

	c.SetNeedCaller(true)
	if c.NeedCaller() != true {
		t.Fatalf("needCaller %+v of core is wrong", c.NeedCaller())
	}
}
