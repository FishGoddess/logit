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

    // NewStdoutLogger creates a new Logger holder to standard output, generally a terminal or a console.
    logger := logit.NewStdoutLogger(logit.DebugLevel)

    // Then you will be easy to log!
    logger.Debug("this is a debug message!")
    logger.Info("this is an info message!")
    logger.Warn("this is a warn message!")
    logger.Error("this is an error message!")

    // NewLogger creates a new Logger holder.
    // The first parameter "os.Stdout" is a writer for logging.
    // As you know, file also can be written, just replace "os.Stdout" with your file!
    // The second parameter "logit.DebugLevel" is the level of this Logger.
    logger = logit.NewLogger(os.Stdout, logit.DebugLevel)

    // If you want format your time, try this:
    logger.SetFormatOfTime("2006/01/02 15:04:05")
    logger.Info("What time is it now?")

    // If you want to output log with file info, try this:
    logger.EnableFileInfo()
    logger.Info("What file is it? Which line?")
    logger.DisableFileInfo()

    // If you have a long log and it is made of many variables, try this:
    // The msg is the return value of msgGenerator.
    logger.DebugFunction(func() string {
        // Use time as the source of random number generator.
        r := rand.New(rand.NewSource(time.Now().Unix()))
        return "debug rand int: " + strconv.Itoa(r.Intn(100))
    })

    // If you want to change logger's writer, try this:
    logger.ChangeWriterTo(os.Stdout)
}
