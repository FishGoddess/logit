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
// Created at 2020/03/03 15:10:45

package writer

import (
	"math/rand"
	"os"
	"path"
	"strconv"
	"time"
)

// These are units representation of file size.
//
// KB = 1024 bytes.
// MB = 1024 * 1024 bytes.
// GB = 1024 * 1024 * 1024 bytes.
const (
	KB int64 = 1 << (10 * (iota + 1))
	MB
	GB
)

const (
	// SuffixOfLogFile is the suffix of log file.
	SuffixOfLogFile = ".log"
)

// nextFilename creates a time-relative filename with given now time.
// Also, it uses random number to ensure this filename is available.
// The filename will be like "20200304-145246-45.log".
// Notice that directory stores all log files generated in this time.
func NextFilename(directory string) func(now time.Time) string {
	rand.Seed(time.Now().UnixNano())
	return func(now time.Time) string {
		name := now.Format("20060102-150405") + "-" + strconv.Itoa(rand.Intn(1000)) + SuffixOfLogFile
		return path.Join(directory, name)
	}
}

// NewFile creates a new file with given filePath.
// Return a new File or an error if failed.
// Notice that the permission of new file is 0644, which means rw-rw-r-- in unix-like os.
func NewFile(filePath string) (*os.File, error) {
	return os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
}

// generateFirstFile creates the first file with given nextFilename function.
func generateFirstFile(nextFilename func(now time.Time) string) (*os.File, time.Time) {
	now := time.Now()
	file, err := NewFile(nextFilename(now))
	if err != nil {
		panic(err)
	}
	return file, now
}
