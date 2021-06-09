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
// Created at 2021/06/09 00:34:42

package main

import (
	"os"

	"github.com/FishGoddess/logit"
)

func main() {

	// Create a new logger with extra values
	// Every log will carry these values.
	// logit.M is a type alia of map[string]interface{}, so you can use it like map
	logger := logit.NewLogger(logit.KV{
		"trace": 123,
		"xxx":   "abc",
	})

	// There are four levels can be logged
	logger.Debug("Hello, I am debug!") // Ignore because default level is info
	logger.Info("Hello, I am info!")
	logger.Warn("Hello, I am warn!")
	logger.Error("Hello, I am error!", logit.KV{"err": "xxx"})

	// You can format log with some parameters if you want
	logger.DebugF("Hello, I am debug %d!", 2) // Ignore because default level is info
	logger.InfoF("Hello, I am info %d!", 2)
	logger.WarnF("Hello, I am warn %d!", 2)
	logger.ErrorF("Hello, I am error %d!", 2)

	// Set level to debug
	logger.SetLevel(logit.DebugLevel)
	logger.Debug("Now debug logs will come up!")

	// Log won't carry caller information in default
	// So, try SetNeedCaller if you need
	logger.SetNeedCaller(true)
	logger.Info("I need caller!")

	// Set encoder and writer
	// Actually, every level has own encoder and writer
	// This way will set encoder and writer of all levels to the same one
	logger.Encoders().SetEncoder(logit.NewJsonEncoder("2006-01-02 15:04:05"))
	logger.Writers().SetWriter(os.Stdout)

	// We also provide some functions to set encoder and writer of each level
	logger.Encoders().SetDebugEncoder(logit.NewJsonEncoder("2006-01-02 15:04:05"))
	logger.Encoders().SetInfoEncoder(logit.NewJsonEncoder("2006-01-02 15:04:05"))
	logger.Encoders().SetWarnEncoder(logit.NewJsonEncoder("2006-01-02 15:04:05"))
	logger.Encoders().SetErrorEncoder(logit.NewJsonEncoder("2006-01-02 15:04:05"))
	logger.Writers().SetDebugWriter(os.Stdout)
	logger.Writers().SetInfoWriter(os.Stdout)
	logger.Writers().SetWarnWriter(os.Stdout)
	logger.Writers().SetErrorWriter(os.Stdout)
}
