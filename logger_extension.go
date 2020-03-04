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
// Created at 2020/03/02 15:36:49

package logit

import (
    "math/rand"
    "os"
    "path"
    "strconv"
    "time"

    "github.com/FishGoddess/logit/wrapper"
)

// PrefixOfLogFile is the prefix of log file.
const PrefixOfLogFile = ".log"

// nextFilename creates a time-relative filename with given now time.
// Also, it uses random number to ensure this filename is available.
// The filename will be like "20200304-145246-45.log".
// Notice that directory stores all log files generated in this time.
func nextFilename(directory string) func(now time.Time) string {
    return func(now time.Time) string {
        name := now.Format("20060102-150405") + "-" + strconv.Itoa(rand.Intn(100)) + PrefixOfLogFile
        return path.Join(directory, name)
    }
}

// NewStdoutLogger returns a Logger holder with given logger level.
func NewStdoutLogger(level LoggerLevel) *Logger {
    return NewLogger(os.Stdout, level)
}

// NewFileLogger returns a Logger holder which log to a file with given logFile and level.
func NewFileLogger(logFile string, level LoggerLevel) *Logger {
    file, err := wrapper.NewFile(logFile)
    if err != nil {
        panic(err)
    }
    return NewLogger(file, level)
}

// NewDurationRollingLogger creates a duration rolling logger with given duration.
// You should appoint a directory to store all log files generated in this time.
// Notice that duration must not less than minDuration (generally time.Second), see wrapper.minDuration.
// Also, default filename of log file is like "20200304-145246-45.log", see wrapper.NewFilename.
// If you want to appoint another filename, check this and do it by this way.
// See wrapper.NewDurationRollingFile (it is an implement of io.writer).
func NewDurationRollingLogger(directory string, duration time.Duration, level LoggerLevel) *Logger {
    file := wrapper.NewDurationRollingFile(duration, nextFilename(directory))
    return NewLogger(file, level)
}

// NewDayRollingLogger creates a day rolling logger.
// You should appoint a directory to store all log files generated in this time.
// See NewDurationRollingLogger.
func NewDayRollingLogger(directory string, level LoggerLevel) *Logger {
    return NewDurationRollingLogger(directory, 24*time.Hour, level)
}
