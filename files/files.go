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
// Created at 2020/03/03 15:10:45

package files

import (
	"os"
	"path/filepath"
)

const (
	// These are units representation of file size.
	//
	// KB = 1024 bytes.
	// MB = 1024 * 1024 bytes.
	// GB = 1024 * 1024 * 1024 bytes.
	KB int64 = 1 << (10 * (iota + 1))
	MB
	GB

	// SuffixOfLogFile is the suffix of log file.
	SuffixOfLogFile = ".log"
)

// CreateFileOf creates a new file with given filePath.
// Return a new File or an error if failed.
// Notice that the permission of new file is 0644, which means rw-rw-r-- in unix-like os.
func CreateFileOf(filePath string) (*os.File, error) {

	// 如果这个文件路径的父级目录不存在，全部自动创建
	// issue: https://github.com/FishGoddess/logit/issues/8
	err := os.MkdirAll(filepath.Dir(filePath), 0644)
	if err != nil {
		return nil, err
	}
	return os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
}
