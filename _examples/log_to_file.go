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
// Created at 2020/03/03 23:39:39

package main

import (
	"time"

	"github.com/FishGoddess/logit"
	"github.com/FishGoddess/logit/writer"
)

func main() {

	// NewFileLogger creates a new logger which logs to file.
	// It just need a file path like "D:/test.log" and a logger level.
	logger := logit.NewLogger(logit.DebugLevel, logit.NewFileHandler("D:/test.log", logit.TextEncoder(), logit.DefaultTimeFormat))
	logger.Info("I am info message！")

	// NewDurationRollingLogger creates a duration rolling logger with given duration.
	// You should appoint a directory to store all log files generated in this time.
	// Notice that duration must not less than minDuration (generally time.Second), see writer.minDuration.
	// Also, default filename of log file is like "20200304-145246-45.log", see writer.NewFilename.
	// If you want to appoint another filename, check this and do it by this way.
	// See writer.NewDurationRollingFile (it is an implement of io.writer).
	logger = logit.NewLogger(logit.DebugLevel, logit.NewDurationRollingHandler(24*time.Hour, "D:/", logit.TextEncoder(), logit.DefaultTimeFormat))
	logger.Info("Rolling!!!")

	// NewSizeRollingLogger creates a file size rolling logger with given limitedSize.
	// You should appoint a directory to store all log files generated in this time.
	// Notice that limitedSize must not less than minLimitedSize (generally 64 KB), see writer.minLimitedSize.
	// Check writer.KB, writer.MB, writer.GB to know what unit you gonna to use.
	// Also, default filename of log file is like "20200304-145246-45.log", see nextFilename.
	// If you want to appoint another filename, check this and do it by this way.
	// See writer.NewSizeRollingFile (it is an implement of io.writer).
	logger = logit.NewLogger(logit.DebugLevel, logit.NewSizeRollingHandler(64*writer.KB, "D:/", logit.TextEncoder(), logit.DefaultTimeFormat))
	logger.Info("file size???")
}
