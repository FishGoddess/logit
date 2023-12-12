// Copyright 2023 FishGoddess. All Rights Reserved.
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

/*
Package logit provides an easy way to use foundation for your logging operations.

1. basic:

	// Use default logger to log.
	// By default, logs will be output to stdout.
	logit.Default().Info("hello from logit", "key", 123)

	// Use a new logger to log.
	// By default, logs will be output to stdout.
	logger := logit.NewLogger()

	logger.Debug("new version of logit", "version", "1.5.0-alpha", "date", 20231122)
	logger.Error("new version of logit", "version", "1.5.0-alpha", "date", 20231122)

	type user struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}

	u := user{123456, "fishgoddess"}
	logger.Info("user information", "user", u, "pi", 3.14)

	// Yep, I know you want to output logs to a file, try WithFile option.
	// The path in WithFile is where the log file will be stored.
	// Also, it's a good choice to call logger.Close() when program shutdown.
	logger = logit.NewLogger(logit.WithFile("./logit.log"))
	defer logger.Close()

	logger.Info("check where I'm logged", "file", "logit.log")

	// What if I want to use default logger and output logs to a file? Try SetDefault.
	// It sets a logger to default and you can use it by package function or Default().
	logit.SetDefault(logger)
	logit.Default().Warn("this is from default logger", "pi", 3.14, "default", true)

	// If you want to change level of logger to info, try WithInfoLevel.
	// Other levels is similar to info level.
	logger = logit.NewLogger(logit.WithInfoLevel())

	logger.Debug("debug logs will be ignored")
	logger.Info("info logs can be logged")

	// If you want to pass logger by context, use NewContext and FromContext.
	ctx := logit.NewContext(context.Background(), logger)

	logger = logit.FromContext(ctx)
	logger.Info("logger from context", "from", "context")

	// Don't want to panic when new a logger? Try NewLoggerGracefully.
	logger, err := logit.NewLoggerGracefully(logit.WithFile(""))
	if err != nil {
		fmt.Println("new logger gracefully failed:", err)
	}

2. logger:

	// Default() will return the default logger.
	// You can new a logger or just use the default logger.
	logger := logit.Default()
	logger.Info("nothing carried")

	// Use With() to carry some args in logger.
	// All logs output by this logger will carry these args.
	logger = logger.With("carry", 666, "who", "me")

	logger.Info("see what are carried")
	logger.Error("error carried", "err", io.EOF)

	// Use WithGroup() to group args in logger.
	// All logs output by this logger will group args.
	logger = logger.WithGroup("xxx")

	logger.Info("what group")
	logger.Error("error group", "err", io.EOF)

	// If you want to check if one level can be logged, try this:
	if logger.DebugEnabled() {
		logger.Debug("debug enabled")
	}

	// We provide some old-school logging methods.
	// They are using info level by default.
	// If you want to change the level, see defaults.LevelPrint.
	logger.Printf("printf %s log", "formatted")
	logger.Print("print log")
	logger.Println("println log")

	// Some useful method:
	logger.Sync()
	logger.Close()

3. handler:

	// By default, logit uses text handler to output logs.
	logger := logit.NewLogger()
	logger.Info("default handler is text")

	// You can change it to other handlers by options.
	// For example, use json handler:
	logger = logit.NewLogger(logit.WithJsonHandler())
	logger.Info("using json handler")

	// Or you want to use customized handlers, try RegisterHandler.
	newHandler := func(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
		return slog.NewTextHandler(w, opts)
	}

	if err := logit.RegisterHandler("demo", newHandler); err != nil {
		panic(err)
	}

	logger = logit.NewLogger(logit.WithHandler("demo"))
	logger.Info("using demo handler")

	// As you can see, our handler is slog's handler, so you can use any handlers implement this interface.
	newHandler = func(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
		return slog.NewJSONHandler(w, opts)
	}

4. writer:

	// A new logger outputs logs to stdout.
	logger := logit.NewLogger()
	logger.Debug("log to stdout")

	// What if I want to output logs to stderr? Try WithStderr.
	logger = logit.NewLogger(logit.WithStderr())
	logger.Debug("log to stderr")

	// Also, you can use WithWriter to specify your own writer.
	logger = logit.NewLogger(logit.WithWriter(os.Stdout))
	logger.Debug("log to writer")

	// How to output logs to a file? Try WithFile and WithRotateFile.
	// Rotate file is useful in production, see _examples/file.go.
	logger = logit.NewLogger(logit.WithFile("logit.log"))
	logger.Debug("log to file")

	logger = logit.NewLogger(logit.WithRotateFile("logit.log"))
	logger.Debug("log to rotate file")

5. file:

	// AS we know, you can use WithFile to output logs to a file.
	logger := logit.NewLogger(logit.WithFile("logit.log"))
	logger.Debug("debug to file")

	// However, a single file stored all logs isn't enough in production.
	// Sometimes we want a log file has a limit size and count of files not greater than a number.
	// So we provide a rotate file to do this thing.
	logger = logit.NewLogger(logit.WithRotateFile("logit.log"))
	defer logger.Close()

	logger.Debug("debug to rotate file")

	// Maybe you have noticed that WithRotateFile can pass some rotate.Option.
	// These options are used to setup the rotate file.
	opts := []rotate.Option{
		rotate.WithMaxSize(128 * rotate.MB),
		rotate.WithMaxAge(30 * rotate.Day),
		rotate.WithMaxBackups(60),
	}

	logger = logit.NewLogger(logit.WithRotateFile("logit.log", opts...))
	defer logger.Close()

	logger.Debug("debug to rotate file with rotate options")

	// See rotate.File if you want to use this magic in other scenes.
	file, err := rotate.New("logit.log")
	if err != nil {
		panic(err)
	}

	defer file.Close()

6. option:

	// As you can see, NewLogger can use some options to create a logger.
	logger := logit.NewLogger(logit.WithDebugLevel())
	logger.Debug("debug log")

	// We provide some options for different scenes and all options have prefix "With".
	// Change logger level:
	logit.WithDebugLevel()
	logit.WithInfoLevel()
	logit.WithWarnLevel()
	logit.WithDebugLevel()

	// Change logger handler:
	logit.WithHandler("xxx")
	logit.WithTextHandler()
	logit.WithJsonHandler()

	// Change handler writer:
	logit.WithWriter(os.Stdout)
	logit.WithStdout()
	logit.WithStderr()
	logit.WithFile("")
	logit.WithRotateFile("")

	// Some useful flags:
	logit.WithSource()
	logit.WithPID()

	// More options can be found in logit package which have prefix "With".
	// What's more? We provide a options pack that we think it's useful in production.
	// It outputs logs to a rotate file using batch write, so you should call Sync() or Close() when shutdown.
	opts := logit.ProductionOptions()

	logger = logit.NewLogger(opts...)
	defer logger.Close()

	logger.Info("log from production options")
	logger.Error("error log from production options")

7. context:

	// We provide a way for getting logger from a context.
	// By default, the default logger will be returned if there is no logit.Logger in context.
	ctx := context.Background()

	logger := logit.FromContext(ctx)
	logger.Debug("logger from context debug")

	if logger == logit.Default() {
		logger.Info("logger from context is default logger")
	}

	// Use NewContext to set a logger to context.
	// We use WithGroup here to make a difference to default logger.
	logger = logit.NewLogger().WithGroup("context").With("user_id", 123456)
	ctx = logit.NewContext(ctx, logger)

	// Then you can get the logger from context.
	logger = logit.FromContext(ctx)
	logger.Debug("logger from context debug", "key", "value")

8. default:

	// We set a defaults package that setups all shared fields.
	// For example, if you want to customize the time getter:
	defaults.CurrentTime = func() time.Time {
		// Return a fixed time for example.
		return time.Unix(666, 0).In(time.Local)
	}

	logit.Default().Print("println log is info level")

	// If you want change the level of old-school logging methods:
	defaults.LevelPrint = slog.LevelDebug

	logit.Default().Print("println log is debug level now")

	// More fields see defaults package.
	defaults.HandleError = func(label string, err error) {
		fmt.Printf("%s: %+n\n", label, err)
	}
*/
package logit // import "github.com/FishGoddess/logit"

const (
	// Version is the version string representation of logit.
	Version = "v1.5.3-alpha"
)
