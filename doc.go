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
    logger = logit.NewLogger(logit.DebugLevel, logit.NewDefaultHandler(os.Stdout, "2006/01/02 15:04:05"))
    logger.Info("What time is it now?")

    // For convenience, we provide a register mechanism and you can use handlers like this:
    logger = logit.NewLogger(logit.DebugLevel, logit.HandlerOf("default", map[string]string{}))
    logger.Info("What handler is it now?")

    // NewLoggerFrom creates a new Logger holder with given config.
    // The config has all the things to create a logger, such as level.
    // We provide some encoders: default encoder and json encoder.
    // See logit.Encoder to check more information.
    logger = logit.NewLoggerFrom(logit.Config{
        Level:    logit.DebugLevel,
        Handlers: []logit.Handler{logit.NewJsonHandler(os.Stdout, "")},
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

    // Change logger level to info level.
    // So debug log will be ignored.
    logit.ChangeLevelTo(logit.InfoLevel)
    logit.Debug("You never see me!")

    // In particular, you can change level to OffLevel to disable the logger.
    // So the info message next line will not be logged!
    level := logit.ChangeLevelTo(logit.OffLevel)
    logit.Info("I will not be logged!")

    // Enable the Logger.
    // The info message next line will be logged again!
    logit.ChangeLevelTo(level)
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
    func (mh *myHandler) Handle(log *logit.Log) bool {
        os.Stdout.Write([]byte("myHandler: "))
        os.Stdout.Write(logit.EncodeToJson(log, "")) // Try `os.Stdout.WriteString(log.Msg())` ?
        return true
    }

    // We recommend you to register your handler to logit, so that
    // you can use your handler by logit.HandlerOf.
    // See logit.HandlerOf.
    func init() {
        logit.RegisterHandler("myHandler", func(params map[string]string) logit.Handler {
            return &myHandler{}
        })
    }

    // Create a logger holder.
    // Default handler is logit.DefaultHandler.
    logger := logit.NewDevelopLogger()
    logger.Info("before adding handlers...")

    // Add handlers to logger.
    // There are two handlers in logger because logger has a default handler inside after creating.
    // See logit.DefaultHandler.
    logger.AddHandlers(&myHandler{}, logit.HandlerOf("json", map[string]string{}))
    fmt.Println("fmt =========================================")
    logger.Info("after adding handlers...")

    // Set handlers to logger.
    // There are two handlers in logger because the default handler inside was removed.
    // If you register your handler to logit by logit.RegisterHandler, then you can
    // use your handler everywhere like this:
    logger.SetHandlers(logit.HandlerOf("myHandler", map[string]string{}))
    fmt.Println("fmt =========================================")
    logger.Info("after setting handlers...")

6. config file:

    // Create a logger from config file.
    //
    // logger.cfg:
    //     "level":"debug",
    //     "handlers":{
    //        "default":{
    //            "timeFormat" : "2006/01/02 15:04"
    //        },
    //        "json":{
    //            # "timeFormat" : "2006-01-02"
    //        }
    //     }
    //
    logger := logit.NewLoggerFromConfigFile("./logger.cfg")
    logger.Info("I am working!")
    logger.Info("My level is " + logger.Level().String())
    fmt.Println("fmt ==============================================")

    handlers := logger.Handlers()
    for i, handler := range handlers {
        logger.Info(fmt.Sprintf("No.%d hadler ==> %T", i + 1, handler))
    }

*/
package logit // import "github.com/FishGoddess/logit"

// Version is the version string representation of logit.
const Version = "v0.1.2"
