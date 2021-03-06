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
// Created at 2021/07/11 23:31:57

package main

import (
	"os"
	"time"

	"github.com/FishGoddess/logit"
	"github.com/FishGoddess/logit/core/appender"
)

func main() {

	// We provide some options for you
	options := logit.Options()
	options.WithDebugLevel()
	options.WithInfoLevel()
	options.WithWarnLevel()
	options.WithErrorLevel()
	options.WithAppender(appender.Text())
	options.WithWriter(os.Stderr)
	options.WithBuffered(os.Stderr)
	options.WithPid()
	options.WithCaller()
	options.WithMsgKey("msg")
	options.WithTimeKey("time")
	options.WithLevelKey("level")
	options.WithPidKey("pid")
	options.WithFileKey("file")
	options.WithLineKey("line")
	options.WithTimeFormat(appender.UnixTime) // UnixTime means time will be logged as unix time, an int64 number

	// Remember, these options is only used for creating a logger
	logger := logit.NewLogger(
		options.WithPid(),
		options.WithBuffered(os.Stdout),
		options.WithTimeFormat("2006/01/02 15:04:05"),
		// ...
	)
	defer logger.Close()
	logger.Info("check options").End()

	// You can use many options at the same time, but some of them is exclusive
	// So only the last one in order will take effect if you use them at the same time
	logit.NewLogger(
		options.WithDebugLevel(),
		options.WithInfoLevel(),
		options.WithWarnLevel(),
		options.WithErrorLevel(), // The level of logger is error
	)

	// You can customize an option for your logger
	// Actually, Option is just a function like func(logger *Logger)
	// So you can do what you want in creating a logger
	autoFlushOption := func(logger *logit.Logger) {
		go func() {
			select {
			case <-time.Tick(time.Second):
				logger.Flush()
			}
		}()
	}
	logit.NewLogger(autoFlushOption)
}
