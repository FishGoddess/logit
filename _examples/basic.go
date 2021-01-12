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

	// There are four levels can be logged
	logit.Debug("Hello, I am debug!") // Ignore because default level is info
	logit.Info("Hello, I am info!")
	logit.Warn("Hello, I am warn!")
	logit.Error("Hello, I am error!")

	// You can format log with some parameters if you want
	logit.Debug("Hello, I am debug %d!", 2) // Ignore because default level is info
	logit.Info("Hello, I am info %d!", 2)
	logit.Warn("Hello, I am warn %d!", 2)
	logit.Error("Hello, I am error %d!", 2)

	// logit.Me() returns a completed logger for use

	// Set level to debug
	logit.Me().SetLevel(logit.DebugLevel)

	// Log won't carry caller information in default
	// So, try NeedCaller if you need
	logit.Me().NeedCaller(true)
	logit.Info("I need caller!")

	// Set format of time in log
	logit.Me().TimeFormat("2006/01/02 15:04:05")

	// Set encoder and writer
	// Actually, every level has own encoder and writer
	// This way will set encoder and writer of all levels to the same one
	logit.Me().SetEncoder(logit.JsonEncoder())
	logit.Me().SetWriter(os.Stdout)

	// We also provide some functions to set encoder and writer of each level
	logit.Me().SetDebugEncoder(logit.JsonEncoder())
	logit.Me().SetInfoEncoder(logit.JsonEncoder())
	logit.Me().SetWarnEncoder(logit.JsonEncoder())
	logit.Me().SetErrorEncoder(logit.JsonEncoder())
	logit.Me().SetDebugWriter(os.Stdout)
	logit.Me().SetInfoWriter(os.Stdout)
	logit.Me().SetWarnWriter(os.Stdout)
	logit.Me().SetErrorWriter(os.Stdout)
}
