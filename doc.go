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
    logit.Warn("I am a warn message!")
    logit.Error("I am an error message!")

    // Also, you can create a new independent Logger to use. See logit.NewLogger.

    // If you want to output log with file info, try this:
    logit.EnableFileInfo()
    logit.Info("Show file info!")

2. logger:

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

4. change logger level:

    logit.Debug("Default logger level is info, so debug message will not be logged!")

    // Change logger level to debug level.
    logit.ChangeLevelTo(logit.DebugLevel)

    logit.Debug("Now debug message will be logged!")

5. log to file:

    // NewFileLogger creates a new logger which logs to file.
    // It just need a file path like "D:/test.log" and a logger level.
    logger := logit.NewFileLogger("D:/test.log", logit.DebugLevel)
    logger.Info("我是 info 日志！")

    // NewDurationRollingLogger creates a duration rolling logger with given duration.
    // You should appoint a directory to store all log files generated in this time.
    // Notice that duration must not less than minDuration (generally time.Second), see wrapper.minDuration.
    // Also, default filename of log file is like "20200304-145246-45.log", see wrapper.NewFilename.
    // If you want to appoint another filename, check this and do it by this way.
    // See wrapper.NewDurationRollingFile (it is an implement of io.writer).
    logger = logit.NewDurationRollingLogger("D:/", time.Second, logit.DebugLevel)
    logger.Info("Rolling!!!")

    // NewDayRollingLogger creates a day rolling logger.
    // You should appoint a directory to store all log files generated in this time.
    // See logit.NewDurationRollingLogger.
    logger = logit.NewDayRollingLogger("D:/", logit.DebugLevel)
    logger.Info("Today is Friday!!!")

    // NewSizeRollingLogger creates a file size rolling logger with given limitedSize.
    // You should appoint a directory to store all log files generated in this time.
    // Notice that limitedSize must not less than minLimitedSize (generally 64 KB), see wrapper.minLimitedSize.
    // Check wrapper.KB, wrapper.MB, wrapper.GB to know what unit you gonna to use.
    // Also, default filename of log file is like "20200304-145246-45.log", see nextFilename.
    // If you want to appoint another filename, check this and do it by this way.
    // See wrapper.NewSizeRollingFile (it is an implement of io.writer).
    logger = logit.NewSizeRollingLogger("D:/", 64*wrapper.KB, logit.DebugLevel)
    logger.Info("file size???")

    // NewDayRollingLogger creates a file size rolling logger.
    // You should appoint a directory to store all log files generated in this time.
    // Default means limitedSize is 64 MB. See NewSizeRollingLogger.
    logger = logit.NewDefaultSizeRollingLogger("D:/", logit.DebugLevel)
    logger.Info("64 MB rolling!!!")

6. logger handler:

    // Create a logger holder.
    // Default handler is logit.DefaultLoggerHandler.
    logger := logit.NewLogger(os.Stdout, logit.InfoLevel)
    logger.Info("before logging...")

    // Customize your own handler.
    handlers1 := func(logger *logit.Logger, level logit.LoggerLevel, now time.Time, msg string) bool {
        logger.Writer().Write([]byte("handlers1: " + msg + "\n"))
        return true
    }

    handlers2 := func(logger *logit.Logger, level logit.LoggerLevel, now time.Time, msg string) bool {
        logger.Writer().Write([]byte("handlers2: " + msg + "\n"))
        return true
    }

    // Add handlers to logger.
    // There are three handlers in logger because logger has a default handler inside after creating.
    // See logit.DefaultLoggerHandler.
    logger.AddHandlers(handlers1, handlers2)
    fmt.Println("fmt =========================================")
    logger.Info("after adding handlers...")

    // Set handlers to logger.
    // There are two handlers in logger because the default handler inside was removed.
    logger.SetHandlers(handlers1, handlers2)
    fmt.Println("fmt =========================================")
    logger.Info("after setting handlers...")

*/
package logit // import "github.com/FishGoddess/logit"

// Version is the version string representation of the "logit" package.
const Version = "0.0.9"
