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
// Author: fish
// Email: fishinlove@163.com
// Created at 2020/02/29 22:20:35
package main

import (
	"os"

	"github.com/FishGoddess/logit"
)

func main() {

	// Every new Logger is running.
	logger := logit.NewLogger(os.Stdout, logit.DebugLevel)
	logger.Info("I am running!")

	// Shutdown the Logger.
	// So the info message next line will be not logged!
	logger.Disable()
	logger.Info("I will be not logged!")

	// Enable the Logger.
	// The info message next line will be logged again!
	logger.Enable()
	logger.Info("I am running again!")
}
