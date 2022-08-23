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
	// If logger's writer has buffer or something like that, it's better to invoke Close() for flushing buffer or something else.
	logger := logit.NewLogger()
	//defer logger.Close()

	// Then, you can log anything you want.
	// Remember, logs will be ignored if their level is smaller than logger's level.
	// Log() will do some finishing work, so this invocation is necessary.
	logger.Debug("This is a debug message").Log()
	logger.Info("This is a info message").Log()
	logger.Warn("This is a warn message").Log()
	logger.Error("This is a error message").Log()
	logger.Error("This is a %s message, with format", "error").Log() // Format with params.

	// As you know, we provide some levels: debug, info, warn, error, off.
	// The lowest is debug and the highest is off.
	// If you want to change the level of your logger, do it at creating.
	logger = logit.NewLogger(logit.Options().WithWarnLevel())
	logger.Debug("This is a debug message, but ignored").Log()
	logger.Info("This is a info message, but ignored").Log()
	logger.Warn("This is a warn message, not ignored").Log()
	logger.Error("This is a error message, not ignored").Log()

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
	logger.Error("This is a structured message").Error("err", io.EOF).Int("trace", 123).Log()

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
	logger = logit.NewLogger(options.WithCallerDepth(core.CallerDepth + 1))
	logit.SetGlobal(logger)
	logit.Info("Info from logit").Log()

	// We don't recommend you to call logit.SetGlobal unless you really need to call.
	// Instead, we recommend you to call logger.SetToGlobal to set one logger to global if you need.
	logger.SetToGlobal()
	logit.Println("Println from logit")

2. options:

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
	options.WithTimeFormat(appender.UnixTime) // UnixTime means time will be logged as unix time, an int64 number.
	options.WithCallerDepth(3)                // Set caller depth to 3 so the log will get the third depth caller.
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
	autoFlushOption := func(logger *logit.Logger) {
		go func() {
			select {
			case <-time.Tick(time.Second):
				logger.Flush()
			}
		}()
	}

	logit.NewLogger(autoFlushOption)

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

	// Every level has its own appender so you can append logs in different level with different appender.
	logger = logit.NewLogger(
		logit.Options().WithDebugAppender(appender.Text()),
		logit.Options().WithInfoAppender(appender.Text()),
		logit.Options().WithWarnAppender(appender.Json()),
		logit.Options().WithErrorAppender(appender.Json()),
	)

	// Appender is an interface so you can implement your own appender.
	// However, we don't recommend you to do that.
	// This interface may change in every version, so you will pay lots of extra attention to it.
	// So you should implement it only if you really need to do.

4. writer:

	// As you know, writer in logit is customized, not io.Writer.
	// The reason why we create a new Writer interface is we want a flushable writer.
	// Then, we notice a flushable writer also need a close method to flush all data in buffer when closing.
	// So, a new Writer is born:
	//
	//     type Writer interface {
	//	       Flusher
	//	       io.WriteCloser
	//     }
	//
	// In package writer, we provide some writers for you.
	writer.Wrap(os.Stdout)   // Wrap io.Writer to writer.Writer.
	writer.Buffer(os.Stderr) // Wrap io.Writer to writer.Writer with buffer, which needs invoking Flush() or Close().

	// Use the writer without buffer.
	logger := logit.NewLogger(logit.Options().WithWriter(os.Stdout))
	logger.Info("WriterWithoutBuffer").Log()

	// Use the writer with buffer, which is good for io.
	logger = logit.NewLogger(logit.Options().WithBufferWriter(os.Stdout))
	logger.Info("WriterWithBuffer").Log()
	logger.Flush() // Remember flushing data or flushing by Close().
	logger.Close()

	// Use the writer with batch, which is also good for io.
	logger = logit.NewLogger(logit.Options().WithBatchWriter(os.Stdout))
	logger.Info("WriterWithBatch").Log()
	logger.Flush() // Remember flushing data or flushing by Close().
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
	writer.BufferWithSize(os.Stdout, 16*core.KB)
	writer.BatchWithCount(os.Stdout, 64)

5. global:

	// There are some global settings for optimizations, and you can set all of them in need.
	//
	//     import "github.com/go-logit/logit/core"
	//
	// All global settings are stored in package core.

	// 1. LogMallocSize (The pre-malloc size of a new Log data)
	// If your logs are extremely long, such as 4000 bytes, you can set it to 4096 to avoid re-malloc.
	core.LogMallocSize = 4 * core.MB // 4096 Bytes

	// 2. WriterBufferSize (The default size of buffer writer)
	// If your logs are extremely long, such as 16 KB, you can set it to 2048 to avoid re-malloc.
	core.WriterBufferSize = 32 * core.KB

	// 3. MarshalToJson (The marshal function which marshal interface{} to json data)
	// Use std by default, and you can customize your marshal function.
	core.MarshalToJson = json.Marshal

	// After setting global settings, just use Logger as normal.
	logger := logit.NewLogger()
	defer logger.Close()

	logger.Info("set global settings").Uint64("LogMallocSize", core.LogMallocSize).Uint64("WriterBufferSize", core.WriterBufferSize).Log()

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

7. creator:

	type testLoggerCreator struct{}

	func (tlm *testLoggerCreator) CreateLogger(params ...interface{}) (*logit.Logger, error) {
		if len(params) < 1 {
			return nil, errors.New("testLoggerCreator: len(params) < 1")
		}

		if params[0].(string) == "error" {
			return nil, errors.New("testLoggerCreator: params[0] isn't a string")
		}

		// Customize your creation of logger here.
		return logit.NewLogger(), nil
	}

	name := "testLoggerCreator"

	// RegisterLoggerCreator registers creator to logit with given name.
	err := logit.RegisterLoggerCreator(name, new(testLoggerCreator))
	if err != nil {
		panic(err)
	}

	// NewLoggerFromCreator creates logger from creator with given params.
	// Panic will be invoked if params is "error" because CreateLogger in testLoggerCreator has this logic.
	logger, err := logit.NewLoggerFromCreator(name, "xxx")
	if err != nil {
		panic(err)
	}

	logger.Info("I am made of logger creator!").Log()

8. caller:

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

9. interceptor:

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

10. file:

	// Logger will log everything to console by default.
	logger := logit.NewLogger()
	logger.Info("I log everything to console.").Log()

	// You can use WithWriter to change writer in logger.
	logger = logit.NewLogger(logit.Options().WithWriter(os.Stdout))
	logger.Info("I also log everything to console.").Log()

	// As we know, we always log everything to file in production.
	// So we provide a convenient way to create a file.
	logFile := filepath.Join(os.TempDir(), "test.log")
	fmt.Println(logFile)
	logger = logit.NewLogger(logit.Options().WithWriter(file.MustNewFile(logFile)))
	logger.Info("I log everything to file.").String("logFile", logFile).Log()
	logger.Close()

	// Also, as you can see, there is a parameter called withBuffer in WithWriter option.
	// It will use a buffer writer to write logs if withBuffer is true which will bring a huge performance improvement.
	logFile = filepath.Join(os.TempDir(), "test_buffer.log")
	fmt.Println(logFile)
	logger = logit.NewLogger(logit.Options().WithWriter(file.MustNewFile(logFile)))
	logger.Info("I log everything to file with buffer.").String("logFile", logFile).Log()
	logger.Close()

	// We provide some high-performance file for you. Try these:
	writer.BufferWithSize(os.Stdout, 128*core.KB)
	writer.BatchWithCount(os.Stdout, 256)
	logit.Options().WithBufferWriter(os.Stdout)
	logit.Options().WithBatchWriter(os.Stdout)
*/
package logit // import "github.com/go-logit/logit"

const (
	// Version is the version string representation of logit.
	Version = "v0.5.0-alpha"
)
