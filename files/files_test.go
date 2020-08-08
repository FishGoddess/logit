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
// Created at 2020/08/08 15:37:04

package files

import (
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"testing"
)

// 测试 CreateFileOf 方法
func TestCreateFileOf(t *testing.T) {

	// 先测试完全不存在父级目录的情况
	basedDirectory := strconv.Itoa(rand.Int())
	notExistedFile := filepath.Join(os.TempDir(), basedDirectory, "fishgoddess", "logit", strconv.Itoa(rand.Int())+".txt")
	t.Log(notExistedFile)
	_, err := CreateFileOf(notExistedFile)
	if err != nil {
		t.Fatal(err)
	}

	// 再测试存在父级目录的情况
	notExistedFile = filepath.Join(os.TempDir(), basedDirectory, "fishgoddess", "logit", strconv.Itoa(rand.Int())+".txt")
	t.Log(notExistedFile)
	_, err = CreateFileOf(notExistedFile)
	if err != nil {
		t.Fatal(err)
	}
}
