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
// Created at 2020/03/01 14:51:46

package main

import (
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/FishGoddess/logit"
)

func main() {

	// NewDevelopLogger creates a new Logger holder for developing, generally log to terminal or console.
	// You can switch to logit.NewProductionLogger for production environment.
	//logger := logit.NewProductionLogger(os.Stdout)
	logger := logit.NewDevelopLogger()

	// Then you will be easy to log!
	logger.Debug("this is a debug message!")
	logger.Info("this is an info message!")
	logger.Warn("this is a warn message!")
	logger.Error("this is an error message!")

	// NewLoggerWithoutEncoder creates a new Logger holder with given level and handlers.
	// As you know, file also can be written, just replace os.Stdout with your file!
	// A logger is made of level and handlers, so we provide some handlers for use, see logit.Handler.
	// This method is the most original way to create a logger for use.
	logger = logit.NewLogger(logit.DebugLevel, logit.NewStandardHandler(os.Stdout, logit.TextEncoder(), "2006/01/02 15:04:05"))
	logger.Info("What time is it now?")

	// For convenience, we provide a register mechanism and you can use handlers like this:
	logger = logit.NewLogger(logit.DebugLevel, logit.HandlerOf("console", map[string]interface{}{}))
	logger.Info("What handler is it now?")

	// NewLoggerFrom creates a new Logger holder with given config.
	// The config has all the things to create a logger, such as level.
	// We provide some encoders: default encoder and json encoder.
	// See logit.Encoder to check more information.
	logger = logit.NewLoggerFrom(logit.Config{
		Level:    logit.DebugLevel,
		Handlers: []logit.Handler{logit.NewStandardHandler(os.Stdout, logit.JsonEncoder(), "")},
	})
	logger.Info("I am a json log!")

	// If you want to output log with file info, try this:
	logger.EnableFileInfo()
	logger.Info("What file is it? Which line?")
	logger.DisableFileInfo()

	// If you have a long log and it is made of many variables, try this:
	// The msg is the return value of msgGenerator.
	logger.DebugFunc(func() string {
		// Use time as the source of random number generator.
		r := rand.New(rand.NewSource(time.Now().Unix()))
		return "debug rand int: " + strconv.Itoa(r.Intn(100))
	})
}
