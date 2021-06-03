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
// Created at 2020/11/27 23:40:29

package main

import (
	"os"

	"github.com/FishGoddess/logit"
)

func main() {

	// Create a new logger and set its level to debug
	logger := logit.NewLogger()
	logger.SetLevel(logit.DebugLevel)

	// Then, just use it like normal logger
	logger.Debug("Hello, I am debug!")
	logger.Info("Hello, I am info!")
	logger.Warn("Hello, I am warn!")
	logger.Error("Hello, I am error!")

	// Log won't carry caller information in default
	// So, try NeedCaller if you need
	logger.NeedCaller(true)

	// Set encoder and writer
	// Actually, every level has own encoder and writer
	// This way will set encoder and writer of all levels to the same one
	logger.SetEncoder(logit.NewJsonEncoder(logit.TimeFormat))
	logger.SetWriter(os.Stdout)

	// More features can be discovered by API
}
