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

package wrapper

import (
    "math/rand"
    "os"
    "strconv"
    "time"
)

// PrefixOfLogFile is the prefix of log file.
const PrefixOfLogFile = ".log"

// FormatOfTime is the format of time.
const FormatOfTime = "20060102-150405"

// NewFile creates a new file with given filePath.
// Return a new File or an error if failed.
// Notice that the permission of new file is 0644, which means rw-rw-r-- in unix-like os.
func NewFile(filePath string) (*os.File, error) {
    return os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
}

// NewFilename creates a time-relative filename with given now time.
// Also, it uses random number to ensure this filename is available.
// The filename will be like "20200304-145246-45.log".
func NewFilename(now time.Time) string {
    return now.Format(FormatOfTime) + "-" + strconv.Itoa(rand.Intn(100)) + PrefixOfLogFile
}
