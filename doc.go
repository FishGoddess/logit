// Copyright 2022 FishGoddess. All Rights Reserved.
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

	// Create a new logger for use.
	// Default level is debug, so all logs will be logged.
	// Invoke Close() isn't necessary in all situations.
	// If logger's writer has buffer or something like that, it's better to invoke Close() for syncing buffer or something else.
	logger := logit.NewLogger()
	//defer logger.Close()

	// Then, you can log anything you want.
	// Remember, logs will be ignored if their level is smaller than logger's level.
	// Log() will do some finishing work, so this invocation is necessary.
	logger.Debug("This is a debug message").Log()
	logger.Info("This is an info message").Log()
	logger.Warn("This is a warn message").Log()
	logger.Error(nil, "This is an error message").Log()
	logger.Error(nil, "This is a %s message, with format", "error").Log() // Format with params.

	// As you know, we provide some levels: debug, info, warn, error, off.
	// The lowest is debug and the highest is off.
	// If you want to change the level of your logger, do it at creating.
	logger = logit.NewLogger(logit.Options().WithWarnLevel())
	logger.Debug("This is a debug message, but ignored").Log()
	logger.Info("This is an info message, but ignored").Log()
	logger.Warn("This is a warn message, not ignored").Log()
	logger.Error(nil, "This is an error message, not ignored").Log()

	// Also, we provide some "old school" log method :)
	// (Don't mistake~ I love old school~)
	logger.Printf("This is a log %s, and it's for compatibility", "printed")
	logger.Print("This is a log printed, and it's for compatibility", 123)
	logger.Println("This is a log printed, and it's for compatibility", 666)

	// If you want to log with some fields, try this:
	user := struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		ID:   666,
		Name: "FishGoddess",
		Age:  3,
	}

	logger.Warn("This is a structured message").Any("user", user).Json("userJson", user).Log()
	logger.Error(io.EOF, "This is a structured message").Int("trace", 123).Log()

	// You may notice logit.Options() which returns an options list.
	// Here is some of them:
	options := logit.Options()
	options.WithCaller()                          // Let logs carry caller information.
	options.WithLevelKey("lvl")                   // Change logger's level key to "lvl".
	options.WithWriter(os.Stderr)                 // Change logger's writer to os.Stderr without buffer or batch.
	options.WithBufferWriter(os.Stderr)           // Change logger's writer to os.Stderr with buffer.
	options.WithBatchWriter(os.Stderr)            // Change logger's writer to os.Stderr with batch.
	options.WithErrorWriter(os.Stderr)            // Change logger's error writer to os.Stderr without buffer or batch.
	options.WithTimeFormat("2006-01-02 15:04:05") // Change the format of time (Only the log's time will apply it).

	// You can bind context with logger and use it as long as you can get the context.
	ctx := logit.NewContext(context.Background(), logger)
	logger = logit.FromContext(ctx)
	logger.Info("Logger from context").Log()

	// You can initialize the global logger if you don't want to use an independent logger.
	// WithCallerDepth will set the depth of caller, and default is core.CallerDepth.
	// Functions in global logger are wrapped so depth of caller should be increased 1.
	// You can specify your depth if you wrap again or have something else reasons.
	logger = logit.NewLogger(options.WithCallerDepth(global.CallerDepth + 1))
	logit.SetGlobal(logger)
	logit.Info("Info from logit").Log()

	// We don't recommend you to call logit.SetGlobal unless you really need to call.
	// Instead, we recommend you to call logger.SetToGlobal to set one logger to global if you need.
	logger.SetToGlobal()
	logit.Println("Println from logit")

2. option:

	// We provide some options for you.
	options := logit.Options()
	options.WithDebugLevel()
	options.WithInfoLevel()
	options.WithWarnLevel()
	options.WithErrorLevel()
	options.WithAppender(appender.Text())
	options.WithDebugAppender(appender.Text())
	options.WithInfoAppender(appender.Text())
	options.WithWarnAppender(appender.Text())
	options.WithErrorAppender(appender.Text())
	options.WithPrintAppender(appender.Text())
	options.WithWriter(os.Stderr)
	options.WithBufferWriter(os.Stdout)
	options.WithBatchWriter(os.Stdout)
	options.WithDebugWriter(os.Stderr)
	options.WithInfoWriter(os.Stderr)
	options.WithWarnWriter(os.Stderr)
	options.WithErrorWriter(os.Stderr)
	options.WithPrintWriter(os.Stderr)
	options.WithPID()
	options.WithCaller()
	options.WithMsgKey("msg")
	options.WithTimeKey("time")
	options.WithLevelKey("level")
	options.WithPIDKey("pid")
	options.WithFileKey("file")
	options.WithLineKey("line")
	options.WithFuncKey("func")
	options.WithErrorKey("err")
	options.WithTimeFormat(global.UnixTimeFormat) // UnixTimeFormat means time will be logged as unix time, an int64 number.
	options.WithCallerDepth(3)                    // Set caller depth to 3 so the log will get the third depth caller.
	options.WithInterceptors()

	// Remember, these options is only used for creating a logger.
	logger := logit.NewLogger(
		options.WithPID(),
		options.WithWriter(os.Stdout),
		options.WithTimeFormat("2006/01/02 15:04:05"),
		options.WithCaller(),
		options.WithCallerDepth(4),
		// ...
	)

	defer logger.Close()
	logger.Info("check options").Log()

	// You can use many options at the same time, but some of them is exclusive.
	// So only the last one in order will take effect if you use them at the same time.
	logit.NewLogger(
		options.WithDebugLevel(),
		options.WithInfoLevel(),
		options.WithWarnLevel(),
		options.WithErrorLevel(), // The level of logger is error.
	)

	// You can customize an option for your logger.
	// Actually, Option is just a function like func(logger *Logger).
	// So you can do what you want in creating a logger.
	syncOption := func(logger *logit.Logger) {
		go func() {
			select {
			case <-time.Tick(time.Second):
				logger.Sync()
			}
		}()
	}

	logit.NewLogger(syncOption)

3. appender:

	// We provide some ways to change the form of logs.
	// Actually, appender is an interface with some common methods, see appender.Appender.
	appender.Text()
	appender.Json()

	// Set appender to the one you want to use when creating a logger.
	// Default appender is appender.Text().
	logger := logit.NewLogger()
	logger.Info("appender.Text()").Log()

	// You can switch appender to the other one, such appender.Json().
	logger = logit.NewLogger(logit.Options().WithAppender(appender.Json()))
	logger.Info("appender.Json()").Log()

	// Every level has its own appender, so you can append logs in different level with different appender.
	logger = logit.NewLogger(
		logit.Options().WithDebugAppender(appender.Text()),
		logit.Options().WithInfoAppender(appender.Text()),
		logit.Options().WithWarnAppender(appender.Json()),
		logit.Options().WithErrorAppender(appender.Json()),
	)

	// Appender is an interface, so you can implement your own appender.
	// However, we don't recommend you to do that.
	// This interface may change in every version, so you will pay lots of extra attention to it.
	// So you should implement it only if you really need to do.

	// appender.TextWith can let you configure the escaping flags of text appender.
	// Default will escape keys and values.
	logger = logit.NewLogger(logit.Options().WithAppender(appender.Text()))
	logger.Info("appender.Text() try \t \b \n and see?").Byte("byte\n", '\n').Rune("rune\n", '\n').String("1\t2\b3\n4", "1\t2\b3\n4").Log()
	logger.Info("appender.Text() try \t \b \n and see?").Bytes("bytes\n", []byte{'\t', '\b', '\n'}).Runes("runes\n", []rune{'\t', '\b', '\n'}).Strings("1\t2\b3\n4", []string{"1\t2\b3\n4"}).Log()

	logger = logit.NewLogger(logit.Options().WithAppender(appender.TextWith(true, false)))
	logger.Info("appender.TextWith(true, false) try \t \b \n and see?").Byte("byte\n", '\n').Rune("rune\n", '\n').String("1\t2\b3\n4", "1\t2\b3\n4").Log()
	logger.Info("appender.TextWith(true, false) try \t \b \n and see?").Bytes("bytes\n", []byte{'\t', '\b', '\n'}).Runes("runes\n", []rune{'\t', '\b', '\n'}).Strings("1\t2\b3\n4", []string{"1\t2\b3\n4"}).Log()

	logger = logit.NewLogger(logit.Options().WithAppender(appender.TextWith(false, true)))
	logger.Info("appender.TextWith(false, true) try \t \b \n and see?").Byte("byte\n", '\n').Rune("rune\n", '\n').String("1\t2\b3\n4", "1\t2\b3\n4").Log()
	logger.Info("appender.TextWith(false, true) try \t \b \n and see?").Bytes("bytes\n", []byte{'\t', '\b', '\n'}).Runes("runes\n", []rune{'\t', '\b', '\n'}).Strings("1\t2\b3\n4", []string{"1\t2\b3\n4"}).Log()

	logger = logit.NewLogger(logit.Options().WithAppender(appender.TextWith(false, false)))
	logger.Info("appender.TextWith(true, true) try \t \b \n and see?").Byte("byte\n", '\n').Rune("rune\n", '\n').String("1\t2\b3\n4", "1\t2\b3\n4").Log()
	logger.Info("appender.TextWith(true, true) try \t \b \n and see?").Bytes("bytes\n", []byte{'\t', '\b', '\n'}).Runes("runes\n", []rune{'\t', '\b', '\n'}).Strings("1\t2\b3\n4", []string{"1\t2\b3\n4"}).Log()

4. writer:

	// As you know, writer in logit is customized, not io.Writer.
	// The reason why we create a new Writer interface is we want a sync-able writer.
	// Then, we notice a sync-able writer also need a close method to sync all data in buffer when closing.
	// So, a new Writer is born:
	//
	//     type Writer interface {
	//	       Syncer
	//	       io.WriteCloser
	//     }
	//
	// In package writer, we provide some writers for you.
	writer.Wrap(os.Stdout)   // Wrap io.Writer to writer.Writer.
	writer.Buffer(os.Stderr) // Wrap io.Writer to writer.Writer with buffer, which needs invoking Sync() or Close().

	// Use the writer without buffer.
	logger := logit.NewLogger(logit.Options().WithWriter(os.Stdout))
	logger.Info("WriterWithoutBuffer").Log()

	// Use the writer with buffer, which is good for io.
	logger = logit.NewLogger(logit.Options().WithBufferWriter(os.Stdout))
	logger.Info("WriterWithBuffer").Log()
	logger.Sync() // Remember syncing data or syncing by Close().
	logger.Close()

	// Use the writer with batch, which is also good for io.
	logger = logit.NewLogger(logit.Options().WithBatchWriter(os.Stdout))
	logger.Info("WriterWithBatch").Log()
	logger.Sync() // Remember syncing data or syncing by Close().
	logger.Close()

	// Every level has its own appender so you can append logs in different level with different appender.
	logger = logit.NewLogger(
		logit.Options().WithBufferWriter(os.Stdout),
		logit.Options().WithBatchWriter(os.Stdout),
		logit.Options().WithWarnWriter(os.Stdout),
		logit.Options().WithErrorWriter(os.Stdout),
	)

	// Let me explain buffer writer and batch writer.
	// Both of them are base on a byte buffer and merge some writes to one write.
	// Buffer writer will write data in buffer to underlying writer if bytes in buffer are too much.
	// Batch writer will write data in buffer to underlying writer if writes to buffer are too much.
	//
	// Let's see something more interesting:
	// A buffer writer with buffer size 16 KB and a batch writer with batch count 64, whose performance is better?
	//
	// 1. Assume one log is 512 Bytes and its size is fixed
	// In buffer writer, it will merge 32 writes to 1 writes (16KB / 512Bytes);
	// In batch writer, it will always merge 64 writes to 1 writes;
	// Batch writer wins the game! Less writes means it's better to IO.
	//
	// 2. Assume one log is 128 Bytes and its size is fixed
	// In buffer writer, it will merge 128 writes to 1 writes (16KB / 128Bytes);
	// In batch writer, it will always merge 64 writes to 1 writes;
	// Buffer writer wins the game! Less writes means it's better to IO.
	//
	// 3. How about one log is 256 Bytes and its size is fixed
	// In buffer writer, it will merge 64 writes to 1 writes (16KB / 256Bytes);
	// In batch writer, it will always merge 64 writes to 1 writes;
	// They are the same in writing times.
	//
	// Based on what we mentioned above, we can tell the performance of buffer writer is depends on the size of log, and the batch writer is more stable.
	// Actually, the size of logs in production isn't fixed-size, so batch writer may be a better choice.
	// However, the buffer in batch writer is out of our control, so it may grow too large if our logs are too large.
	writer.Buffer(os.Stdout)
	writer.Batch(os.Stdout)
	writer.BufferWithSize(os.Stdout, 16*size.KB)
	writer.BatchWithCount(os.Stdout, 64)

5. global:

	// There are some global settings for optimizations, and you can set all of them in need.

	// 1. LogMallocSize (The pre-malloc size of a new Log data)
	// If your logs are extremely long, such as 4000 bytes, you can set it to 4096 to avoid re-malloc.
	global.LogMallocSize = 4 * size.MB

	// 2. WriterBufferSize (The default size of buffer writer)
	// If your logs are extremely long, such as 16 KB, you can set it to 2048 to avoid re-malloc.
	global.WriterBufferSize = 32 * size.KB

	// 3. MarshalToJson (The marshal function which marshal interface{} to json data)
	// Use std by default, and you can customize your marshal function.
	global.MarshalToJson = json.Marshal

	// After setting global settings, just use Logger as normal.
	logger := logit.NewLogger()
	defer logger.Close()

	logger.Info("set global settings").Uint64("LogMallocSize", global.LogMallocSize).Uint64("WriterBufferSize", global.WriterBufferSize).Log()

6. context:

	// By NewContext, you can bind a context with a logger and get it from context again.
	// So you can use this logger from everywhere as long as you can get this context.
	ctx := logit.NewContext(context.Background(), logit.NewLogger())

	// FromContext returns the logger in context.
	logger := logit.FromContext(ctx)
	logger.Info("This is a message logged by logger from context").Log()

	// Actually, you also have a chance to specify the key of logger in context.
	// It gives you a way to discriminate different businesses in using logger.
	// For example, you can create two loggers for your two different usages and
	// set them to a context with different key, so you can get each logger from context with each key.
	businessOneKey := "businessOne"
	logger = logit.NewLogger(logit.Options().WithMsgKey("businessOneMsg"))
	ctx = logit.NewContextWithKey(context.Background(), businessOneKey, logger)

	businessTwoKey := "businessTwo"
	logger = logit.NewLogger(logit.Options().WithMsgKey("businessTwoMsg"))
	ctx = logit.NewContextWithKey(ctx, businessTwoKey, logger)

	// Get different logger from the same context with different key.
	logger = logit.FromContextWithKey(ctx, businessOneKey)
	logger.Info("This is a message logged by logger from context with businessOneKey").Log()

	logger = logit.FromContextWithKey(ctx, businessTwoKey)
	logger.Info("This is a message logged by logger from context with businessTwoKey").Log()

7. caller:

	// Let's create a logger without caller information.
	logger := logit.NewLogger()
	logger.Info("I am without caller").Log()

	// We provide a way to add caller information to log even logger doesn't carry caller.
	logger.Info("Invoke log.WithCaller()").WithCaller().Log()
	logger.Close()

	time.Sleep(time.Second)

	// Now, let's create a logger with caller information.
	logger = logit.NewLogger(logit.Options().WithCaller())
	logger.Info("I am with caller").Log()

	// We won't carry caller information twice or more if logger carries caller information already.
	logger.Info("Invoke log.WithCaller() again").WithCaller().Log()
	logger.Close()

8. interceptor:

	// serverInterceptor is the global interceptor applied to all logs.
	func serverInterceptor(ctx context.Context, log *logit.Log) {
		log.String("server", "logit.interceptor")
	}

	// traceInterceptor is the global interceptor applied to all logs.
	func traceInterceptor(ctx context.Context, log *logit.Log) {
		trace, ok := ctx.Value("trace").(string)
		if !ok {
			trace = "unknown trace"
		}

		log.String("trace", trace)
	}

	// userInterceptor is the global interceptor applied to all logs.
	func userInterceptor(ctx context.Context, log *logit.Log) {
		user, ok := ctx.Value("user").(string)
		if !ok {
			user = "unknown user"
		}

		log.String("user", user)
	}

	// businessInterceptor is the log-level interceptor applied to one/some logs.
	func businessInterceptor(ctx context.Context, log *logit.Log) {
		business, ok := ctx.Value("business").(string)
		if !ok {
			business = "unknown business"
		}

		log.String("business", business)
	}

	// Use logit.Options().WithInterceptors to append some interceptors.
	logger := logit.NewLogger(logit.Options().WithInterceptors(serverInterceptor, traceInterceptor, userInterceptor))
	defer logger.Close()

	// By default, context passed to interceptor is context.Background().
	logger.Info("try interceptor - round one").Log()

	// You can use WithContext to change context passed to interceptor.
	ctx := context.WithValue(context.Background(), "trace", "666")
	ctx = context.WithValue(ctx, "user", "FishGoddess")
	logger.Info("try interceptor - round two").WithContext(ctx).Log()

	// The interceptors appended to logger will apply to all logs.
	// You can use Intercept to intercept one log rather than all logs.
	logger.Info("try interceptor - round three").WithContext(ctx).Intercept(businessInterceptor).Log()

	// Notice that WithContext should be called before Intercept if you want to pass this context to Intercept.
	ctx = context.WithValue(ctx, "business", "logger")
	logger.Info("try interceptor - round four").WithContext(ctx).Intercept(businessInterceptor).Log()

9. file:

	func createFile(filePath string) *os.File {
		f, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}

		return f
	}

	// Logger will log everything to console by default.
	logger := logit.NewLogger()
	logger.Info("I log everything to console.").Log()

	// You can use WithWriter to change writer in logger.
	logger = logit.NewLogger(logit.Options().WithWriter(os.Stdout))
	logger.Info("I also log everything to console.").Log()

	// As we know, we always log everything to file in production.
	logFile := filepath.Join(os.TempDir(), "test.log")
	fmt.Println(logFile)

	logger = logit.NewLogger(logit.Options().WithWriter(createFile(logFile)))
	logger.Info("I log everything to file.").String("logFile", logFile).Log()
	logger.Close()

	// We provide some high-performance file for you. Try these:
	logger = logit.NewLogger(logit.Options().WithBufferWriter(createFile(logFile)))
	logger = logit.NewLogger(logit.Options().WithBatchWriter(createFile(logFile)))

	// Or you can use the original writer package to create a writer configured by you.
	writer.BufferWithSize(os.Stdout, 128*size.KB)
	writer.BatchWithCount(os.Stdout, 256)

	// Wait a minute, we also provide a powerful file for you!
	// See extension/file/file.go.
	// It will rotate file and clean backups automatically.
	// You can set maxSize, maxAge and maxBackups by options.
	logFile = filepath.Join(os.TempDir(), "test_powerful.log")

	f, err := file.New(logFile)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	_, err = f.Write([]byte("xxx"))
	if err != nil {
		panic(err)
	}

10. error:

	// uselessWriter is a demo writer to demonstrate the error handling function.
	type uselessWriter struct{}

	func (uw uselessWriter) Write(p []byte) (n int, err error) {
		return 0, errors.New("always error in writing")
	}

	// You can specify a function to handle errors happens in logger.
	// For example, you can count these errors and report them to team members by email.
	global.HandleError = func(name string, err error) {
		fmt.Printf("%s received an error: %+v\n", name, err)
	}

	// Let's log something to see what happen.
	logger := logit.NewLogger(logit.Options().WithWriter(&uselessWriter{}))
	logger.Info("See what happen?").Log()

11. config:

	// We provide a config which can be converted to option in logit.
	// It has many tags in fields, such json, yaml, toml, which means you can use config file to create logger.
	// You just need to define your config file then unmarshal your config file to this config.
	// Of course, you can embed this struct to your application config struct!
	cfg := config.Config{
		Level:         config.LevelDebug,
		TimeKey:       "x.time",
		LevelKey:      "x.level",
		MsgKey:        "x.msg",
		PIDKey:        "x.pid",
		FileKey:       "x.file",
		LineKey:       "x.line",
		FuncKey:       "x.func",
		ErrorKey:      "x.err",
		TimeFormat:    config.UnixTimeFormat,
		WithPID:       true,
		WithCaller:    true,
		CallerDepth:   0,
		AutoSync:      "10s",
		Appender:      config.AppenderText,
		DebugAppender: config.AppenderText,
		InfoAppender:  config.AppenderText,
		WarnAppender:  config.AppenderText,
		ErrorAppender: config.AppenderText,
		PrintAppender: config.AppenderJson,
		Writer: config.WriterConfig{
			Target:     config.WriterTargetStdout,
			Mode:       config.WriterModeDirect,
			BufferSize: "4MB",
			BatchCount: 1024,
			Filename:   "test.log",
			DirMode:    0755,
			FileMode:   0644,
			TimeFormat: "20060102150405",
			MaxSize:    "128MB",
			MaxAge:     "30d",
			MaxBackups: 32,
		},
		DebugWriter: config.WriterConfig{},
		InfoWriter:  config.WriterConfig{},
		WarnWriter:  config.WriterConfig{},
		ErrorWriter: config.WriterConfig{},
		PrintWriter: config.WriterConfig{},
	}

	// Once you got a config, use Options() to convert to option in logger.
	opts, err := cfg.Options()
	if err != nil {
		panic(err)
	}

	fmt.Println(opts)

	// Then you can create your logger by options.
	// Amazing!
	logger := logit.NewLogger(opts...)
	defer logger.Close()

	logger.Info("My mother is a config").Any("config", cfg).Log()
	logger.Info("See logger").Any("logger", logger).Log()
	logger.Error(io.EOF, "error message").Log()
*/
package logit // import "github.com/FishGoddess/logit"

const (
	// Version is the version string representation of logit.
	Version = "v1.0.0"
)
