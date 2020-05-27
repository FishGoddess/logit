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
// Email: fishinlove@163.com
// Created at 2020/03/03 16:01:38

package writer

import (
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

	root := filepath.Join(os.TempDir(), "TestNewDurationRollingFile")
	os.RemoveAll(root)   // 先删除现有文件夹
	os.Mkdir(root, 0666) // 再创建测试文件夹

	file := NewDurationRollingFile(time.Second, NextFilename(root))
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
	file = NewDurationRollingFile(999*time.Millisecond, func(now time.Time) string {
		return ""
	})
}
