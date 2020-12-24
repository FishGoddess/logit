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
// Created at 2020/11/27 23:41:13

package main

import "github.com/FishGoddess/logit"

func main() {

	// Use default encoder
	logit.Info("Default encoder is like this...")

	// We provide some encoders, such as text and json
	// Try TextEncoder() and JsonEncoder()
	logit.Me().SetEncoder(logit.TextEncoder())
	logit.Me().SetEncoder(logit.JsonEncoder())

	// In fact, encoder is a function like "func(log *logit.Log, timeFormat string) []byte"
	// All information of log is stored in log
	// timeFormat is a layout for formatting time
	// No matter what you do, return a byte slice
	// The returned slice will be written by logger
	logit.Me().SetEncoder(func(log *logit.Log, timeFormat string) []byte {
		logTime := log.Time().Format(timeFormat)
		return []byte(logTime + " => " + log.Msg() + "\r\n")
	})
	logit.Info("My encoder...")

	// You can set encoder of each level, for example:
	logit.Me().SetErrorEncoder(func(log *logit.Log, timeFormat string) []byte {
		logTime := log.Time().Format(timeFormat)
		return []byte("[Error] " + logTime + " => " + log.Msg() + "\n")
	})
	logit.Error("Panic...")

	// If you have a logger, just use it as logit.Me()
	logger := logit.NewLogger()
	logger.SetEncoder(logit.TextEncoder())
	logger.SetWarnEncoder(logit.JsonEncoder())
}
