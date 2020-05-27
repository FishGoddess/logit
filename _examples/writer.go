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
// Created at 2020/03/05 17:30:21

package main

import (
	"time"

	"github.com/FishGoddess/logit/writer"
)

func main() {

	// 1. DurationRollingFile is a time sensitive file.
	durationRollingFile := writer.NewDurationRollingFile(24*time.Hour, func(now time.Time) string {
		return "D:/" + now.Format("20060102-150405") + ".txt"
	})
	defer durationRollingFile.Close()

	// You can use it like using io.Writer!
	durationRollingFile.Write([]byte("durationRollingFile!"))

	// =================================================================================

	// 2. SizeRollingFile is a file size sensitive file.
	sizeRollingFile := writer.NewSizeRollingFile(64*writer.KB, func(now time.Time) string {
		return "D:/" + now.Format("20060102150405.000") + ".txt"
	})
	defer sizeRollingFile.Close()

	// You can use it like using io.Writer!
	sizeRollingFile.Write([]byte("sizeRollingFile!"))
}
