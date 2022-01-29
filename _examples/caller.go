package main

import (
	"github.com/FishGoddess/logit"
	"time"
)

func main() {
	// Let's create a logger without caller information.
	logger := logit.NewLogger()
	logger.Info("I am without caller").End()

	// We provide a way to add caller information to log even logger doesn't carry caller.
	logger.Info("Invoke log.WithCaller()").WithCaller().End()
	logger.Close()

	time.Sleep(time.Second)

	// Now, let's create a logger with caller information.
	logger = logit.NewLogger(logit.Options().WithCaller())
	logger.Info("I am with caller").End()

	// We won't carry caller information twice or more if logger carries caller information originally.
	logger.Info("Invoke log.WithCaller() again").WithCaller().End()
	logger.Close()
}
