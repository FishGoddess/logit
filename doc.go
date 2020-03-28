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

1. the basic usage:

    // Log messages with four levels.
    logit.Debug("I am a debug message!")
    logit.Info("I am an info message!")
    logit.Warn("I am a warn message!")
    logit.Error("I am an error message!")

    // If you want to output log with file info, try this:
    logit.EnableFileInfo()
    logit.Info("Show file info!")

    // If you have a long log and it is made of many variables, try this:
    // The msg is the return value of msgGenerator.
    logit.DebugFunc(func() string {
        // Use time as the source of random number generator.
        r := rand.New(rand.NewSource(time.Now().Unix()))
        return "debug rand int: " + strconv.Itoa(r.Intn(100))
    })

2. logger:

    // NewDevelopLogger creates a new Logger holder for developing, generally log to terminal or console.
    // You can switch to logit.NewProductionLogger for production environment.
    //logger := logit.NewProductionLogger(os.Stdout)
    logger := logit.NewDevelopLogger()

    // Then you will be easy to log!
    logger.Debug("this is a debug message!")
    logger.Info("this is an info message!")
    logger.Warn("this is a warn message!")
    logger.Error("this is an error message!")

    // NewLoggerWithoutEncoder creates a new Logger holder with given level and handlers.
    // As you know, file also can be written, just replace os.Stdout with your file!
    // A logger is made of level and handlers, so we provide some handlers for use, see logit.Handler.
    // This method is the most original way to create a logger for use.
    logger = logit.NewLogger(logit.DebugLevel, logit.NewDefaultEncoder("2006/01/02 15:04:05"), logit.NewDefaultHandler(os.Stdout))
    logger.Info("What time is it now?")

    // NewLoggerFrom creates a new Logger holder with given config.
    // The config has all the things to create a logger, such as level.
    // We provide some encoders: default encoder and json encoder.
    // See logit.Encoder to check more information.
    logger = logit.NewLoggerFrom(logit.Config{
        Level:    logit.DebugLevel,
        Encoder:  logit.NewJsonEncoder("2006/01/02 15:04:05", false),
        Handlers: []logit.Handler{logit.NewDefaultHandler(os.Stdout)},
    })
    logger.Info("I am a json log!")

    // If you want to output log with file info, try this:
    logger.EnableFileInfo()
    logger.Info("What file is it? Which line?")
    logger.DisableFileInfo()

    // If you have a long log and it is made of many variables, try this:
    // The msg is the return value of msgGenerator.
    logger.DebugFunc(func() string {
        // Use time as the source of random number generator.
        r := rand.New(rand.NewSource(time.Now().Unix()))
        return "debug rand int: " + strconv.Itoa(r.Intn(100))
    })

3. level_and_disable:

    logit.Debug("Default logger level is debug.")

    // Change logger level to warn level.
    // So debug and info log will be ignored.
    logit.ChangeLevelTo(logit.WarnLevel)
    logit.Debug("You never see me!")

    // In particular, you can change level to OffLevel to disable the logger.
    // So the info message next line will not be logged!
    logit.ChangeLevelTo(logit.OffLevel)
    logit.Info("I will not be logged!")

    // Enable the Logger.
    // The info message next line will be logged again!
    logit.ChangeLevelTo(logit.InfoLevel)
    logit.Info("I am running again!")

4. log to file:

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

5. handler:

    type myHandler struct{}

    // Customize your own handler.
    func (mh *myHandler) Handle(log []byte, raw *logit.Log) bool {
        os.Stdout.Write([]byte("myHandler: "))
        os.Stdout.Write(log) // Try `os.Stdout.WriteString(raw.Msg())` ?
        return true
    }

    // Create a logger holder.
    // Default handler is logit.DefaultHandler.
    logger := logit.NewDevelopLogger()
    logger.Info("before adding handlers...")

    // Add handlers to logger.
    // There are two handlers in logger because logger has a default handler inside after creating.
    // See logit.DefaultHandler.
    logger.AddHandlers(&myHandler{})
    fmt.Println("fmt =========================================")
    logger.Info("after adding handlers...")

    // Set handlers to logger.
    // There are two handlers in logger because the default handler inside was removed.
    logger.SetHandlers(&myHandler{})
    fmt.Println("fmt =========================================")
    logger.Info("after setting handlers...")

6. encoder:

    type myEncoder struct{}

    // Customize you own encoder.
    func (me *myEncoder) Encode(log *logit.Log) []byte {
        if log.Level() == logit.DebugLevel {
            return []byte("I am debug log ==> " + log.Msg() + "\n")
        }
        return []byte("Normal log --> " + log.Msg() + "\n")
    }

    // logit.NewDefaultEncoder returns a default encoder with given time format.
    // ChangeEncoderTo will change current encoder to new encoder, and return old encoder.
    logger := logit.NewDevelopLogger()
    logger.ChangeEncoderTo(logit.NewDefaultEncoder("2006/01/02 15:04:05"))
    logger.Info("What encoder is it now?")

    // logit.NewJsonEncoder returns a json encoder with given time format and need formatting time.
    logger.ChangeEncoderTo(logit.NewJsonEncoder(logit.DefaultTimeFormat, false))
    logger.Info("I am json log!")

    // You can customize you own encoder, see myEncoder.
    logger.ChangeEncoderTo(&myEncoder{})
    logger.Debug("Ha ha!")
    logger.Info("Hello!")

*/
package logit // import "github.com/FishGoddess/logit"

// Version is the version string representation of logit.
const Version = "v0.1.0-alpha"
