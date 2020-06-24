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
// Created at 2020/03/03 16:01:38

package files

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"
)

// 测试时间间隔滚动文件
func TestNewDurationRollingFile(t *testing.T) {

	defer func() {
		if err := recover(); err == nil {
			t.Errorf("时间间隔大小限制测试出现问题！")
		}
	}()

	// 创建临时测试文件夹
	root, err := ioutil.TempDir("", "TestNewDurationRollingFile_*")
	if err != nil {
		t.Fatal(err)
	}

	file := NewDurationRollingFile(root, time.Second)
	defer file.Close()

	for i := 0; i < 5; i++ {
		file.Write([]byte("测试"))
		time.Sleep(666 * time.Millisecond)
	}

	dir, err := os.Open(root)
	if err != nil {
		t.Fatal("获取测试文件夹失败！")
	}

	fileInfos, err := dir.Readdir(0)
	if err != nil {
		t.Fatal("获取测试文件夹信息失败！")
	}

	// 如果创建的文件数不符合，直接报错
	if len(fileInfos) != 3 {
		t.Fatal("文件滚动出现问题！")
	}

	file.Close()
	file = NewDurationRollingFile("", 999*time.Millisecond)
	file.SetNameGenerator(func(directory string, now time.Time) string {
		return ""
	})
}

// 测试名字生成器的设置方法
func TestDurationRollingFileSetNameGenerator(t *testing.T) {

	dir, err := ioutil.TempDir("", "TestDurationRollingFileSetNameGenerator_*")
	if err != nil {
		t.Fatal(err)
	}

	// 创建文件，并写入内容
	file := NewDurationRollingFile(dir, 2*time.Second)
	defer file.Close()
	file.Write([]byte("hello!"))

	// 更换命名器，等待滚动时间到了之后，再次写入内容
	file.SetNameGenerator(func(directory string, now time.Time) string {
		return filepath.Join(directory, now.Format("2006年01月02日的15点04分05秒产生的文件.log"))
	})
	time.Sleep(2 * time.Second)
	file.Write([]byte("hi!"))
}
