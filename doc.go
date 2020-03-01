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
// Created at 2020/02/29 15:41:09

/*
Package logit provides an easy way to use foundation for your logging operations.

1. The basic usage:

    // Log messages with four levels.
    // Notice that the default level is info, so first line of debug message
    // will not be logged! If you want to change level, see logit.ChangeLevelTo
    logit.Debug("I am a debug message! But I will not be logged in default level!")
    logit.Info("I am an info message!")
    logit.Warning("I am a warning message!")
    logit.Error("I am an error message!")

    // Also, you can create a new independent Logger to use. See logit.NewLogger.

2. logger:

    // NewStdoutLogger creates a new Logger holder to standard output, generally a terminal or a console.
    logger := logit.NewStdoutLogger(logit.DebugLevel)

    // Then you will be easy to log!
    logger.Debug("this is a debug message!")
    logger.Info("this is a info message!")
    logger.Warning("this is a warning message!")
    logger.Error("this is a error message!")

    // NewLogger creates a new Logger holder.
    // The first parameter "os.Stdout" is a writer for logging.
    // As you know, file also can be written, just replace "os.Stdout" with your file!
    // The second parameter "logit.DebugLevel" is the level of this Logger.
    logger = logit.NewLogger(os.Stdout, logit.DebugLevel)

3. enable or disable:

    // Every new Logger is running.
    logger := logit.NewLogger(os.Stdout, logit.DebugLevel)
    logger.Info("I am running!")

    // Shutdown the Logger.
    // So the info message next line will not be logged!
    logger.Disable()
    logger.Info("I will not be logged!")

    // Enable the Logger.
    // The info message next line will be logged again!
    logger.Enable()
    logger.Info("I am running again!")

4. change log level:

    logit.Debug("Default log level is info, so debug message will not be logged!")

    // Change log level to debug level.
    logit.ChangeLevelTo(logit.DebugLevel)

    logit.Debug("Now debug message will be logged!")


*/
package logit // import "github.com/FishGoddess/logit"

// Version is the version string representation of the "logit" package.
const Version = "0.0.2"
