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
	logger := logit.NewLogger(logit.DebugLevel, logit.NewConsoleHandler(logit.TextEncoder(), logit.DefaultTimeFormat))

	// Then you will be easy to log!
	logger.Debug("this is a debug message!")
	logger.Info("this is an info message!")
	logger.Warn("this is a warn message!")
	logger.Error("this is an error message!")

	// NewLoggerWithoutEncoder creates a new Logger holder with given level and handlers.
	// As you know, file also can be written, just replace os.Stdout with your file!
	// A logger is made of level and handlers, so we provide some handlers for use, see logit.Handler.
	// This method is the most original way to create a logger for use.
	logger = logit.NewLogger(logit.DebugLevel, logit.NewStandardHandler(os.Stdout, logit.TextEncoder(), "2006/01/02 15:04:05"))
	logger.Info("What time is it now?")

	// For convenience, we provide a register mechanism and you can use handlers like this:
	logger = logit.NewLogger(logit.DebugLevel, logit.NewConsoleHandler(logit.TextEncoder(), logit.DefaultTimeFormat))
	logger.Info("What handler is it now?")

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

5. handler:

    type myHandler struct{}

    // Customize your own handler.
    func (mh *myHandler) Handle(log *logit.Log) bool {
        os.Stdout.Write([]byte("myHandler: "))
        os.Stdout.Write(logit.TextEncoder().Encode(log, "")) // Try `os.Stdout.WriteString(log.Msg())` ?
        return true
    }

    func init() {
        // We recommend you to register your handler to logit, so that
        // you can use your handler in config file.
        // See logit.RegisterHandler.
        logit.RegisterHandler("myHandler", func(params map[string]interface{}) logit.Handler {
            return &myHandler{}
        })
    }

	// Create a logger holder with a console handler.
	logger := logit.NewLogger(logit.DebugLevel, logit.NewConsoleHandler(logit.TextEncoder(), logit.DefaultTimeFormat))
	logger.Info("before adding handlers...")
	fmt.Println("fmt =========================================")

	// Add handlers to logger.
	// There are three handlers in logger because logger has one handler inside before adding.
	// See logit.NewConsoleHandler.
	logger.AddHandlers(&myHandler{}, logit.NewConsoleHandler(logit.JsonEncoder(), ""))
	logger.Info("after adding two handlers...")
	fmt.Println("fmt =========================================")

	// Set handlers to logger.
	// There are one handler in logger because all handlers inside was removed.
	// If you register your handler to logit by logit.RegisterHandler, then you can
	// use your handler everywhere like this:
	logger.SetHandlers(&myHandler{})
	logger.Info("after setting one handlers...")

6. config file:

	// Create a logger from config file.
	//
	// logger.conf:
	//
	//     "level": "debug",
	//
	//     "caller": false,
	//
	//     "handlers": {
	//         "console": {
	//             "timeFormat":"unix",
	//             "encoder":"json"
	//         },
	//         "file":{
	//             "path":"D:/logit.log"
	//         }
	//     }
	//
	logger := logit.NewLoggerFrom("./logger.conf")
	logger.Info("I am working!")
	logger.Info("My level is " + logger.Level().String())
	fmt.Println("fmt ==============================================")

	handlers := logger.Handlers()
	for i, handler := range handlers {
		logger.Info(fmt.Sprintf("No.%d hadler ==> %T", i+1, handler))
	}

*/
package logit // import "github.com/FishGoddess/logit"

// Version is the version string representation of logit.
const Version = "v0.2.1-alpha"
