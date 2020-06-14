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
// Created at 2020/06/14 23:11:04

package files

import (
	"testing"
	"time"
)

// 测试注册 nameGenerator 的方法
func TestRegisterNameGenerator(t *testing.T) {

	// default 已经存在，测试是否报错
	err := RegisterNameGenerator("default", func(directory string, now time.Time) string {
		return ""
	})
	if err == nil {
		t.Fatal(err)
	}

	// test 不存在，测试是否报错
	err = RegisterNameGenerator("TestRegisterNameGenerator", func(directory string, now time.Time) string {
		return "test"
	})
	if err != nil {
		t.Fatal(err)
	}

	nameGenerator := nameGeneratorOf("TestRegisterNameGenerator")
	if nameGenerator.NextName("", time.Now()) != "test" {
		t.Fatal("注册可能失败了，获取也可能是失败了...")
	}
}

// 测试 nameGeneratorOf 方法
func TestNameGeneratorOf(t *testing.T) {

	// 先注册，再获取
	err := RegisterNameGenerator("TestNameGeneratorOf", func(directory string, now time.Time) string {
		return "test"
	})
	if err != nil {
		t.Fatal(err)
	}

	nameGenerator := nameGeneratorOf("TestNameGeneratorOf")
	if nameGenerator.NextName("", time.Now()) != "test" {
		t.Fatal("注册可能失败了，获取也可能是失败了...")
	}
}
