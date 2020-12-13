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
// Created at 2020/12/13 21:08:30

package main

import (
	"fmt"
	"time"

	"github.com/FishGoddess/logit"
)

func main() {

	// We provide a writer writing to file
	// "./test.log" is the path of this file
	writer, err := logit.NewFileWriter("Z:/test.log")
	if err != nil {
		panic(err)
	}
	defer writer.Close()

	// Use Write() to write something to file underlying
	writer.Write([]byte("Something new..."))

	// Also, we provide some checkers for advanced features
	// Every writing operation will call Check() in checker
	// If one checker's Check() returns true, then this file will roll to a new file
	// The old file will be renamed to be like "xxx.log.0000000001"
	// These are all checkers we provide:
	logit.NewTimeChecker(24 * time.Hour)
	logit.NewSizeChecker(64 * logit.MB)
	logit.NewCountChecker(1000)

	// If you want to use one of them above, try this:
	newWriter, err := logit.NewFileWriter("Z:/test_checker.log", logit.NewSizeChecker(128*logit.KB))
	if err != nil {
		panic(err)
	}
	defer newWriter.Close()

	// Check how many files starting with "test_checker.log"?
	for i := 0; i < 10000; i++ {
		newWriter.Write([]byte(fmt.Sprintf("Something new with checker...%d\n", i)))
	}

	// Also, you can use more than one of them in a writer
	//logit.NewFileWriter("Z:/test_checker.log", logit.NewTimeChecker(24*time.Hour), logit.NewSizeChecker(128*logit.KB))
}
