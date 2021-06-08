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
// Created at 2020/11/27 23:39:30

package main

import (
	"os"

	"github.com/FishGoddess/logit"
)

func main() {

	// Create a new logger first
	logger := logit.NewLogger()

	// There are four levels can be logged, and you can format log with some parameters
	logger.DebugF("Hello, I am debug %d!", 2) // Ignore because default level is info
	logger.InfoF("Hello, I am info %d!", 2)
	logger.WarnF("Hello, I am warn %d!", 2)
	logger.ErrorF("Hello, I am error %d!", 2)

	// Set level to debug
	logger.SetLevel(logit.DebugLevel)
	logger.DebugF("Now debug logs will come up!")

	// Log won't carry caller information in default
	// So, try SetNeedCaller if you need
	logger.SetNeedCaller(true)
	logger.InfoF("I need caller!")

	// Set encoder and writer
	// Actually, every level has own encoder and writer
	// This way will set encoder and writer of all levels to the same one
	logger.Encoders().SetEncoder(logit.NewTextEncoder("2006-01-02 15:04:05"))
	logger.Encoders().SetErrorEncoder(logit.NewJsonEncoder("2006-01-02 15:04:05"))
	logger.Writers().SetWriter(os.Stdout)
	logger.Writers().SetErrorWriter(os.Stderr)
}
