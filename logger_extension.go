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
    "os"
    "path"
    "time"

    "github.com/FishGoddess/logit/wrapper"
)

func nextFilename(directory string) func(lastTime, currentTime time.Time) string {
    // TODO 这个文件函数可能需要调整参数为 currentTime 和 duration
    return func(lastTime, currentTime time.Time) string {
        return path.Join(directory)
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

func NewDurationRollingLogger(directory string, duration time.Duration, level LoggerLevel) *Logger {
    // TODO 创建一个文件名生成器，返回时间间隔滚动的日志记录器
    return nil
}
