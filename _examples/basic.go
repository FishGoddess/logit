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
// Email: fishinlove@163.com
// Created at 2020/02/29 21:59:13
package main

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/FishGoddess/logit"
)

func main() {

	// Log messages with four levels.
	logit.Debug("I am a debug message!")
	logit.Info("I am an info message!")
	logit.Warn("I am a warn message!")
	logit.Error("I am an error message!")

	// Notice that logit has blocked some methods for more refreshing method list.
	// If you want to use some higher level methods, you should call logit.Me() to
	// get the fully functional logger, then call what you want to call.
	// For example, if you want to output log with file info, try this:
	logit.Me().EnableFileInfo()
	logit.Info("Show file info!")

	// If you have a long log and it is made of many variables, try this:
	// The msg is the return value of msgGenerator.
	logit.DebugFunc(func() string {
		// Use time as the source of random number generator.
		r := rand.New(rand.NewSource(time.Now().Unix()))
		return "debug rand int: " + strconv.Itoa(r.Intn(100))
	})

	// If a config file "logit.conf" in "./", then logit will load it automatically.
	// This is more convenience to use config file and logger.
}
