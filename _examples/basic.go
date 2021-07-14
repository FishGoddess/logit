// Copyright 2021 Ye Zi Jie. All Rights Reserved.
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
// Created at 2021/07/11 22:53:20

package main

import (
	"os"

	"github.com/FishGoddess/logit"
)

func main() {

	// Create a new logger for use
	// Default level is debug, so all logs will be logged
	// Invoke Close() isn't necessary in all situations
	// If logger's writer has buffer or something like that, it's better to invoke Close() for flushing buffer or something else
	logger := logit.NewLogger()
	//defer logger.Close()

	// Then, you can log anything you want
	// Remember, logs will be ignored if their level is smaller than logger's level
	// End() will do some finishing work, so this invocation is necessary
	logger.Debug("This is a debug message").End()
	logger.Info("This is a info message").End()
	logger.Warn("This is a warn message").End()
	logger.Error("This is a error message").End()
	logger.Error("This is a %s message, with format", "error").End() // Format with params

	// As you know, we provide some levels: debug, info, warn, error, off
	// The lowest is debug and the highest is off
	// If you want to change the level of your logger, do it at creating
	logger = logit.NewLogger(logit.Options().WithWarnLevel())
	logger.Debug("This is a debug message, but ignored").End()
	logger.Info("This is a info message, but ignored").End()
	logger.Warn("This is a warn message, not ignored").End()
	logger.Error("This is a error message, not ignored").End()

	// You may notice logit.Options() which returns an options list
	// Here is some of them:
	options := logit.Options()
	options.WithCaller()                          // Let logs carry caller information
	options.WithLevelKey("lvl")                   // Change logger's level key to "lvl"
	options.WithWriter(os.Stderr)                 // Change logger's writer to os.Stderr
	options.WithBuffered(os.Stderr)               // Change logger's writer to os.Stderr with buffer
	options.WithTimeFormat("2006-01-02 15:04:05") // Change the format of time (Only the log's time will apply it)
}
