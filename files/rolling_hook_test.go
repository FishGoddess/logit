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
// Created at 2020/06/13 00:07:00

package files

import (
	"fmt"
	"io/ioutil"
	"testing"
	"time"
)

// 测试 LifeRollingHook 的 AfterRolling 方法
func TestLifeRollingHookAfterRolling(t *testing.T) {

	// 创建测试文件夹
	directory, err := ioutil.TempDir("", "TestLifeRollingHookAfterRolling_*")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(directory)

	// 创建测试文件
	for i := 0; i < 3; i++ {
		file, err := ioutil.TempFile(directory, "test_file_*.log")
		if err != nil {
			t.Fatal(err)
		}
		file.Close()

		_, err = ioutil.TempDir(directory, "test_directory_*")
		if err != nil {
			t.Fatal(err)
		}

		time.Sleep(2 * time.Second)
	}

	// 判断测试文件是否准备成功
	checkFileCountInDirectory := func(count int) {
		fileInfos, err := ioutil.ReadDir(directory)
		t.Log("count:", count)
		t.Log("len:", len(fileInfos))
		if err != nil || len(fileInfos) != count {
			t.Fatal(err, len(fileInfos))
		}
	}

	// 创建基于生命周期的滚动钩子器
	rollingHook := NewLifeBasedRollingHook(directory, 4*time.Second)

	// 开始测试
	// 判断滚动之前的文件数量是否正确
	checkFileCountInDirectory(6)
	for i := 1; i <= 3; i++ {
		rollingHook.AfterRolling()
		checkFileCountInDirectory(6 - i)
		time.Sleep(2 * time.Second)
	}

	rollingHook.AfterRolling()
	checkFileCountInDirectory(3)
}

// 测试用的 rollingHook
type testRollingHook struct{}

func (trh *testRollingHook) BeforeRolling() {
	fmt.Println("testRollingHook.BeforeRolling()...")
}

func (trh *testRollingHook) AfterRolling() {
	fmt.Println("testRollingHook.AfterRolling()...")
}

// 测试注册 rollingHook 的方法
func TestRegisterRollingHook(t *testing.T) {

	// default 已经存在，测试是否报错
	err := RegisterRollingHook("default", func(params map[string]interface{}) RollingHook {
		return &testRollingHook{}
	})
	if err == nil {
		t.Fatal("name 为 default 的 rollingHook 已经存在，本来要报错的，但是没有报错")
	}

	// life 已经存在，测试是否报错
	err = RegisterRollingHook("life", func(params map[string]interface{}) RollingHook {
		return &testRollingHook{}
	})
	if err == nil {
		t.Fatal("name 为 life 的 rollingHook 已经存在，本来要报错的，但是没有报错")
	}

	// test 不存在，测试是否报错
	err = RegisterRollingHook("TestRegisterRollingHook", func(params map[string]interface{}) RollingHook {
		return &testRollingHook{}
	})
	if err != nil {
		t.Fatal("name 为 test 的 rollingHook 不存在，不应该报错的，但是报错了")
	}

	rollingHook := rollingHookOf("TestRegisterRollingHook", map[string]interface{}{})
	rollingHook.BeforeRolling()
	rollingHook.AfterRolling()

	//rollingHookOf("noExist", map[string]interface{}{})
}

// 测试 rollingHookOf 方法
func TestRollingHookOf(t *testing.T) {

	// 先注册，再获取
	err := RegisterRollingHook("TestRollingHookOf", func(params map[string]interface{}) RollingHook {
		return &testRollingHook{}
	})
	if err != nil {
		t.Fatal(err)
	}

	rollingHook := rollingHookOf("TestRollingHookOf", map[string]interface{}{})
	rollingHook.BeforeRolling()
	rollingHook.AfterRolling()
}
