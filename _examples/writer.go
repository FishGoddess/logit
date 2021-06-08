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
	"os"
	"time"

	"github.com/FishGoddess/logit"
)

func main() {

	// If you want to set output to another one, try SetWriter
	// You should use Writers() to get all writers in logger and invoke SetWriter on it
	// Any writer implemented io.Writer can be used here
	logger := logit.NewLogger()
	logger.Writers().SetWriter(os.Stdout)
	logger.InfoF("SetWriter...")

	// Also, all levels have its own writer
	logger.Writers().SetDebugWriter(os.Stdout)
	logger.Writers().SetInfoWriter(os.Stdout)
	logger.Writers().SetWarnWriter(os.Stderr)
	logger.Writers().SetErrorWriter(os.Stderr)

	// In fact, write logs to disk is expensive in time, so we provide a special writer for you
	// This writer uses a buffer to reduce times of writing to disk, so it has a extremely-high performance
	// Write logs to disk is just like write logs to memory after using this writer in our benchmark
	// Amazing, right? Try logit.NewBufferedWriter immediately!
	writer := logit.NewBufferedWriter(os.Stdout)
	logger.Writers().SetWriter(writer)
	logger.InfoF("NewBufferedWriter...")
	writer.Flush() // Notice that Flush() should be invoked after finishing writing or you may miss some data

	// Of cause we provide a way to change the buffer size of it
	writer = logit.NewBufferedWriter(os.Stdout)
	logger.Writers().SetWriter(writer)
	logger.InfoF("Oh! Faster! Faster!!! Yeah~~")
	writer.Flush() // Notice that Flush() should be invoked after finishing writing or you may miss some data

	// The buffered writer won't flush data automatically in default
	// Does it puzzle you? Try AutoFlush() to get it if you want!
	writer = logit.NewBufferedWriter(os.Stdout)
	writer.AutoFlush(time.Second)
	logger.Writers().SetWriter(writer)
	logger.InfoF("AutoFlush...")
	time.Sleep(2 * time.Second)
}
