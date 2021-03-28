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
	"github.com/FishGoddess/logit"
)

func main() {

	// We provide a writer writing to file
	// You should use a config instance to initialize this writer
	// logit.DefaultConfig() returns a default config
	// config.LogFileName is the path of this file
	// However, it isn't the final name of log file
	// We will add a date suffix to the name, so the final name will look like "test.log.20210328"
	// This means every day has its own log file, and it's an automatic task
	config := logit.DefaultConfig()
	config.LogFileName = "Z:/test.log"

	writer, err := logit.NewFileWriter(config)
	if err != nil {
		panic(err)
	}
	defer writer.Close()

	// Use Write() to write something to file underlying
	writer.Write([]byte("Something new...\n"))

	// Use file writer in logger
	logger := logit.NewLogger()
	logger.SetWriter(writer)
	logger.Info("log by file writer")
}
