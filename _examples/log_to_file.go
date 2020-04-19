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
	"github.com/FishGoddess/logit/wrapper"
)

func main() {

	// NewFileLogger creates a new logger which logs to file.
	// It just need a file path like "D:/test.log" and a logger level.
	logger := logit.NewFileLogger("D:/test.log")
	logger.Info("I am info message！")

	// NewDurationRollingLogger creates a duration rolling logger with given duration.
	// You should appoint a directory to store all log files generated in this time.
	// Notice that duration must not less than minDuration (generally time.Second), see wrapper.minDuration.
	// Also, default filename of log file is like "20200304-145246-45.log", see wrapper.NewFilename.
	// If you want to appoint another filename, check this and do it by this way.
	// See wrapper.NewDurationRollingFile (it is an implement of io.writer).
	logger = logit.NewDurationRollingLogger("D:/", 24*time.Hour)
	logger.Info("Rolling!!!")

	// NewDayRollingLogger creates a day rolling logger.
	// You should appoint a directory to store all log files generated in this time.
	// See logit.NewDurationRollingLogger.
	logger = logit.NewDayRollingLogger("D:/")
	logger.Info("Today is Friday!!!")

	// NewSizeRollingLogger creates a file size rolling logger with given limitedSize.
	// You should appoint a directory to store all log files generated in this time.
	// Notice that limitedSize must not less than minLimitedSize (generally 64 KB), see wrapper.minLimitedSize.
	// Check wrapper.KB, wrapper.MB, wrapper.GB to know what unit you gonna to use.
	// Also, default filename of log file is like "20200304-145246-45.log", see nextFilename.
	// If you want to appoint another filename, check this and do it by this way.
	// See wrapper.NewSizeRollingFile (it is an implement of io.writer).
	logger = logit.NewSizeRollingLogger("D:/", 64*wrapper.KB)
	logger.Info("file size???")

	// NewDayRollingLogger creates a file size rolling logger.
	// You should appoint a directory to store all log files generated in this time.
	// Default means limitedSize is 64 MB. See NewSizeRollingLogger.
	logger = logit.NewDefaultSizeRollingLogger("D:/")
	logger.Info("64 MB rolling!!!")
}
