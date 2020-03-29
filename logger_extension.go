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
    "io"
    "math/rand"
    "os"
    "path"
    "strconv"
    "time"

    "github.com/FishGoddess/logit/wrapper"
)

const (
    // SuffixOfLogFile is the suffix of log file.
    SuffixOfLogFile = ".log"
)

// nextFilename creates a time-relative filename with given now time.
// Also, it uses random number to ensure this filename is available.
// The filename will be like "20200304-145246-45.log".
// Notice that directory stores all log files generated in this time.
func nextFilename(directory string) func(now time.Time) string {
    rand.Seed(time.Now().UnixNano())
    return func(now time.Time) string {
        name := now.Format("20060102-150405") + "-" + strconv.Itoa(rand.Intn(1000)) + SuffixOfLogFile
        return path.Join(directory, name)
    }
}

// NewLoggerFrom returns a logger with given config.
// See logit.Config.
func NewLoggerFrom(config Config) *Logger {
    return NewLogger(config.Level, config.Handlers...)
}

// NewLoggerFromConfigFile returns a logger with config file.
// A config file is a file like "xxx.cfg", and its content looks like:
//
//     "handlers":[
//         {
//             "name":"json"
//         }
//     ]
//
// Check examples to get more information. See logit.ParseConfigFile.
func NewLoggerFromConfigFile(configFile string) *Logger {
    config, err := ParseConfigFile(configFile)
    if err != nil {
        panic(err)
    }
    return NewLoggerFrom(config)
}

// NewDevelopLogger returns a logger for developing.
// A logger for developing should be easy-to-read and output to console.
// Also, the level should be DebugLevel.
func NewDevelopLogger() *Logger {
    return NewLoggerFrom(Config{
        Level:    DebugLevel,
        Handlers: []Handler{NewDefaultHandler(os.Stdout, DefaultTimeFormat)},
    })
}

// NewProductionLogger returns a logger for production.
// A logger for production should be easy-to-resolve and output to somewhere not only console.
// Also, the level should be WarnLevel.
func NewProductionLogger(writer io.Writer) *Logger {
    return NewLoggerFrom(Config{
        Level:    WarnLevel,
        Handlers: []Handler{NewJsonHandler(writer, "")},
    })
}

// NewFileLogger returns a Logger holder which log to a file with given logFile.
func NewFileLogger(logFile string) *Logger {
    file, err := wrapper.NewFile(logFile)
    if err != nil {
        panic(err)
    }
    return NewLoggerFrom(Config{
        Level:    InfoLevel,
        Handlers: []Handler{NewDefaultHandler(file, DefaultTimeFormat)},
    })
}

// NewDurationRollingLogger creates a duration rolling logger with given duration.
// You should appoint a directory to store all log files generated in this time.
// Notice that duration must not less than minDuration (generally one second), see wrapper.minDuration.
// Also, default filename of log file is like "20200304-145246-45.log", see nextFilename.
// If you want to appoint another filename, check this and do it by this way.
// See wrapper.NewDurationRollingFile (it is an implement of io.writer).
func NewDurationRollingLogger(directory string, duration time.Duration) *Logger {
    file := wrapper.NewDurationRollingFile(duration, nextFilename(directory))
    return NewLoggerFrom(Config{
        Level:    InfoLevel,
        Handlers: []Handler{NewDefaultHandler(file, DefaultTimeFormat)},
    })
}

// NewDayRollingLogger creates a day rolling logger.
// You should appoint a directory to store all log files generated in this time.
// See logit.NewDurationRollingLogger.
func NewDayRollingLogger(directory string) *Logger {
    return NewDurationRollingLogger(directory, 24*time.Hour)
}

// NewSizeRollingLogger creates a file size rolling logger with given limitedSize.
// You should appoint a directory to store all log files generated in this time.
// Notice that limitedSize must not less than minLimitedSize (generally 64 KB), see wrapper.minLimitedSize.
// Check wrapper.KB, wrapper.MB, wrapper.GB to know what unit you gonna to use.
// Also, default filename of log file is like "20200304-145246-45.log", see nextFilename.
// If you want to appoint another filename, check this and do it by this way.
// See wrapper.NewSizeRollingFile (it is an implement of io.writer).
func NewSizeRollingLogger(directory string, limitedSize int64) *Logger {
    file := wrapper.NewSizeRollingFile(limitedSize, nextFilename(directory))
    return NewLoggerFrom(Config{
        Level:    InfoLevel,
        Handlers: []Handler{NewDefaultHandler(file, DefaultTimeFormat)},
    })
}

// NewDayRollingLogger creates a file size rolling logger.
// You should appoint a directory to store all log files generated in this time.
// Default means limitedSize is 64 MB. See logit.NewSizeRollingLogger.
func NewDefaultSizeRollingLogger(directory string) *Logger {
    return NewSizeRollingLogger(directory, 64*wrapper.MB)
}

// DebugFunc will output msg as a debug message.
// The msg is the return value of msgGenerator.
// This is a better way to output a long log made of many variables.
func (l *Logger) DebugFunc(msgGenerator func() string) {
    l.Debug(msgGenerator())
}

// InfoFunc will output msg as an info message.
// The msg is the return value of msgGenerator.
// This is a better way to output a long log made of many variables.
func (l *Logger) InfoFunc(msgGenerator func() string) {
    l.Info(msgGenerator())
}

// WarnFunc will output msg as a warn message.
// The msg is the return value of msgGenerator.
// This is a better way to output a long log made of many variables.
func (l *Logger) WarnFunc(msgGenerator func() string) {
    l.Warn(msgGenerator())
}

// ErrorFunc will output msg as an error message.
// The msg is the return value of msgGenerator.
// This is a better way to output a long log made of many variables.
func (l *Logger) ErrorFunc(msgGenerator func() string) {
    l.Error(msgGenerator())
}
