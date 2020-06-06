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
// Created at 2020/03/01 15:10:19

package main

import "github.com/FishGoddess/logit"

func main() {

	// Use logit.Debug method to output a debug level message.
	// Also, Info/Warn/Error method is available.
	logit.Debug("Default logger level is debug.")

	// Change logger level to info level, so logs in debug level will be ignored.
	// Notice that logit has blocked some methods for more refreshing method list.
	// If you want to use some higher level methods, you should call logit.Me() to
	// get the fully functional logger, then call what you want to call.
	logit.Me().ChangeLevelTo(logit.InfoLevel)
	logit.Debug("You never see me!")

	// In particular, you can change level to OffLevel to disable the logger.
	// So the info message next line will not be logged!
	level := logit.Me().ChangeLevelTo(logit.OffLevel)
	logit.Info("I will not be logged!")

	// Enable the Logger.
	// The info message next line will be logged again!
	logit.Me().ChangeLevelTo(level)
	logit.Info("I am running again!")
}
