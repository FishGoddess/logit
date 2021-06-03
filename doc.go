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
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2020/02/29 15:41:09

/*
Package logit provides an easy way to use foundation for your logging operations.

1. basic

	// There are four levels can be logged
	logit.Debug("Hello, I am debug!") // Ignore because default level is info
	logit.Info("Hello, I am info!")
	logit.Warn("Hello, I am warn!")
	logit.Error("Hello, I am error!")

	// You can format log with some parameters if you want
	logit.Debug("Hello, I am debug %d!", 2) // Ignore because default level is info
	logit.Info("Hello, I am info %d!", 2)
	logit.Warn("Hello, I am warn %d!", 2)
	logit.Error("Hello, I am error %d!", 2)

	// logit.Me() returns a completed logger for use

	// Set level to debug
	logit.Me().SetLevel(logit.DebugLevel)

	// Log won't carry caller information in default
	// So, try NeedCaller if you need
	logit.Me().NeedCaller(true)
	logit.Info("I need caller!")

	// Set encoder and writer
	// Actually, every level has own encoder and writer
	// This way will set encoder and writer of all levels to the same one
	logit.Me().SetEncoder(logit.NewJsonEncoder("2006-01-02 15:04:05"))
	logit.Me().SetWriter(os.Stdout)

	// We also provide some functions to set encoder and writer of each level
	logit.Me().SetDebugEncoder(logit.NewJsonEncoder("2006-01-02 15:04:05"))
	logit.Me().SetInfoEncoder(logit.NewJsonEncoder("2006-01-02 15:04:05"))
	logit.Me().SetWarnEncoder(logit.NewJsonEncoder("2006-01-02 15:04:05"))
	logit.Me().SetErrorEncoder(logit.NewJsonEncoder("2006-01-02 15:04:05"))
	logit.Me().SetDebugWriter(os.Stdout)
	logit.Me().SetInfoWriter(os.Stdout)
	logit.Me().SetWarnWriter(os.Stdout)
	logit.Me().SetErrorWriter(os.Stdout)

2. logger

	// Create a new logger and set its level to debug
	logger := logit.NewLogger()
	logger.SetLevel(logit.DebugLevel)

	// Then, just use it like normal logger
	logger.Debug("Hello, I am debug!")
	logger.Info("Hello, I am info!")
	logger.Warn("Hello, I am warn!")
	logger.Error("Hello, I am error!")

	// Log won't carry caller information in default
	// So, try NeedCaller if you need
	logger.NeedCaller(true)

	// Set encoder and writer
	// Actually, every level has own encoder and writer
	// This way will set encoder and writer of all levels to the same one
	logger.SetEncoder(logit.NewJsonEncoder(logit.TimeFormat))
	logger.SetWriter(os.Stdout)

	// More features can be discovered by API

3. encoder

	// Use default encoder
	logit.Info("Default encoder is like this...")

	// We provide some encoders, such as text and json
	// Try TextEncoder and JsonEncoder
	logit.Me().SetEncoder(logit.NewTextEncoder("2006-01-02 15:04:05"))
	logit.Me().SetEncoder(logit.NewJsonEncoder("2006-01-02 15:04:05"))

	// In fact, encoder is an interface like "func(log *logit.Log) []byte"
	// So you can implement your own encoder as you want
	// All information of log is stored in log
	// No matter what you do, return a byte slice
	// The returned slice will be written by logger
	logit.Me().SetEncoder(logit.NewTextEncoder("2006-01-02 15:04:05"))
	logit.Info("My encoder...")

	// You can set encoder of each level, for example:
	logit.Me().SetErrorEncoder(logit.NewJsonEncoder(logit.TimeFormat))
	logit.Error("Panic...")

	// If you have a logger, just use it as logit.Me()
	logger := logit.NewLogger()
	logger.SetEncoder(logit.NewTextEncoder("2006-01-02 15:04:05"))
	logger.SetWarnEncoder(logit.NewJsonEncoder("2006-01-02 15:04:05"))

4. writer

	// We provide a writer writing to file
	// You should use a config instance to initialize this writer
	// logit.DefaultConfig() returns a default config
	// config.LogFileName is the path of this file
	// However, it isn't the final name of log file
	// We will add a date suffix to the name, so the final name will look like "test.log.20210328"
	// This means every day has its own log file, and it's an automatic task
	config := logit.DefaultConfig()
	config.LogFileName = "Z:/test.log"

	writer, err := logit.NewFileWriter(config)
	if err != nil {
		panic(err)
	}
	defer writer.Close()

	// Use Write() to write something to file underlying
	writer.Write([]byte("Something new...\n"))

	// Use file writer in logger
	logger := logit.NewLogger()
	logger.SetWriter(writer)
	logger.Info("log by file writer")

*/
package logit // import "github.com/FishGoddess/logit"

const (
	// Version is the version string representation of logit.
	Version = "v0.4.0-alpha"
)
