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
// Author: fish
// Email: fishinlove@163.com
// Created at 2020/03/05 16:10:31

package writer

import (
	"testing"
	"time"
)

// 测试创建根据文件大小滚动的文件类型
func TestNewSizeRollingFile(t *testing.T) {

	file := NewSizeRollingFile(64*KB, func(now time.Time) string {
		return "Z:/" + now.Format("20060102150405.000") + ".log"
	})
	defer file.Close()

	b := make([]byte, 1024)
	for i := 0; i < 1024; i++ {
		file.Write(b)
	}
}
