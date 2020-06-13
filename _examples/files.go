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
// Created at 2020/03/05 17:30:21

package main

import (
	"path/filepath"
	"time"

	"github.com/FishGoddess/logit/files"
)

func main() {

	// 1. DurationRollingFile is a time sensitive file.
	durationRollingFile := files.NewDurationRollingFile("D:/", 24*time.Hour)
	defer durationRollingFile.Close()

	// You can use it like using io.Writer!
	durationRollingFile.Write([]byte("durationRollingFile!"))

	// If you want to change the way generating file name, try this:
	durationRollingFile.SetNameGenerator(func(directory string, now time.Time) string {
		// directory is the directory stores all file rolled before.
		// now is the time calling this method.
		return filepath.Join(directory, now.Format("2006-01-02-15-04-05.log"))
	})

	// What's more, you can add a hook in rolling process, see RollingHook.
	//durationRollingFile.SetRollingHook(xxx)

	// =================================================================================

	// 2. SizeRollingFile is a file size sensitive file.
	sizeRollingFile := files.NewSizeRollingFile("D:/", 64*files.KB)
	defer sizeRollingFile.Close()

	// You can use it like using io.Writer!
	sizeRollingFile.Write([]byte("sizeRollingFile!"))

	// If you want to change the way generating file name, try this:
	sizeRollingFile.SetNameGenerator(func(directory string, now time.Time) string {
		// directory is the directory stores all file rolled before.
		// now is the time calling this method.
		return filepath.Join(directory, now.Format("2006-01-02-15-04-05.log"))
	})

	// What's more, you can add a hook in rolling process, see RollingHook.
	//sizeRollingFile.SetRollingHook(xxx)
}
