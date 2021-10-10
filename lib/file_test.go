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
// Created at 2021/08/10 02:41:27

package lib

import (
	"os"
	"testing"
)

// go test -v -cover -run=^TestNewFile$
func TestNewFile(t *testing.T) {
	filePath := "Z:/" + t.Name()
	os.Remove(filePath)

	file, err := NewFile(filePath)
	if err != nil {
		t.Error(err)
	}

	err = file.Close()
	if err != nil {
		t.Error(err)
	}

	stat, err := os.Stat(filePath)
	if err != nil {
		t.Error(err)
	}

	if stat.IsDir() {
		t.Error("file is a directory, not file")
	}
}
