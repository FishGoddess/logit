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
	logit.DebugF("Hello, I am debug!") // Ignore because default level is info
	logit.InfoF("Hello, I am info!")
	logit.WarnF("Hello, I am warn!")
	logit.ErrorF("Hello, I am error!")

	// You can format log with some parameters if you want
	logit.DebugF("Hello, I am debug %d!", 2) // Ignore because default level is info
	logit.InfoF("Hello, I am info %d!", 2)
	logit.WarnF("Hello, I am warn %d!", 2)
	logit.ErrorF("Hello, I am error %d!", 2)

	// logit.Me() returns a completed logger for use

	// Set level to debug
	logit.Me().SetLevel(logit.DebugLevel)

	// Log won't carry caller information in default
	// So, try SetNeedCaller if you need
	logit.Me().SetNeedCaller(true)
	logit.InfoF("I need caller!")

	// Set encoder and writer
	// Actually, every level has own encoder and writer
	// This way will set encoder and writer of all levels to the same one
	logit.Me().Encoders().SetEncoder(logit.NewJsonEncoder("2006-01-02 15:04:05"))
	logit.Me().Writers().SetWriter(os.Stdout)

	// We also provide some functions to set encoder and writer of each level
	logit.Me().Encoders().SetDebugEncoder(logit.NewJsonEncoder("2006-01-02 15:04:05"))
	logit.Me().Encoders().SetInfoEncoder(logit.NewJsonEncoder("2006-01-02 15:04:05"))
	logit.Me().Encoders().SetWarnEncoder(logit.NewJsonEncoder("2006-01-02 15:04:05"))
	logit.Me().Encoders().SetErrorEncoder(logit.NewJsonEncoder("2006-01-02 15:04:05"))
	logit.Me().Writers().SetDebugWriter(os.Stdout)
	logit.Me().Writers().SetInfoWriter(os.Stdout)
	logit.Me().Writers().SetWarnWriter(os.Stdout)
	logit.Me().Writers().SetErrorWriter(os.Stdout)

2. logger

	// Create a new logger and set its level to debug
	logger := logit.NewLogger()
	logger.SetLevel(logit.DebugLevel)

	// Then, just use it like normal logger
	logger.DebugF("Hello, I am debug!")
	logger.InfoF("Hello, I am info!")
	logger.WarnF("Hello, I am warn!")
	logger.ErrorF("Hello, I am error!")

	// Log won't carry caller information in default
	// So, try SetNeedCaller if you need
	logger.SetNeedCaller(true)
	logger.DebugF("Hello, I have caller information!")

	// Set encoder and writer
	// Actually, every level has own encoder and writer
	// This way will set encoder and writer of all levels to the same one
	logger.Encoders().SetEncoder(logit.NewJsonEncoder(logit.TimeFormat))
	logger.Writers().SetErrorWriter(os.Stderr)
	logger.ErrorF("Oh, I am error!")

	// More features can be discovered by API

3. encoder

	// Use default encoder
	logit.InfoF("Default encoder is like this...")

	// We provide some encoders, such as text and json
	// Try TextEncoder and JsonEncoder
	logit.Me().Encoders().SetEncoder(logit.NewTextEncoder("2006-01-02 15:04:05"))
	logit.Me().Encoders().SetEncoder(logit.NewJsonEncoder("2006-01-02 15:04:05"))

	// In fact, encoder is an interface like "func(log *logit.Log) []byte"
	// So you can implement your own encoder as you want
	// All information of log is stored in log
	// No matter what you do, return a byte slice
	// The returned slice will be written by logger
	logit.Me().Encoders().SetEncoder(&MyEncoder{name: "whatever"})
	logit.InfoF("My encoder...")

	// You can set encoder of each level, for example:
	logit.Me().Encoders().SetErrorEncoder(logit.NewJsonEncoder(logit.TimeFormat))
	logit.ErrorF("Panic...")

	// If you have a logger, just use it as logit.Me()
	logger := logit.NewLogger()
	logger.Encoders().SetEncoder(logit.NewTextEncoder("2006-01-02 15:04:05"))
	logger.Encoders().SetWarnEncoder(logit.NewJsonEncoder("2006-01-02 15:04:05"))
	logger.InfoF("info...")
	logger.WarnF("warn...")

4. writer

	// If you want to set output to another one, try SetWriter
	// You should use Writers() to get all writers in logger and invoke SetWriter on it
	// Any writer implemented io.Writer can be used here
	logger := logit.NewLogger()
	logger.Writers().SetWriter(os.Stdout)
	logger.InfoF("SetWriter...")

	// Also, all levels have its own writer
	logger.Writers().SetDebugWriter(os.Stdout)
	logger.Writers().SetInfoWriter(os.Stdout)
	logger.Writers().SetWarnWriter(os.Stderr)
	logger.Writers().SetErrorWriter(os.Stderr)

	// In fact, write logs to disk is expensive in time, so we provide a special writer for you
	// This writer uses a buffer to reduce times of writing to disk, so it has a extremely-high performance
	// Write logs to disk is just like write logs to memory after using this writer in our benchmark
	// Amazing, right? Try logit.NewBufferedWriter immediately!
	writer := logit.NewBufferedWriter(os.Stdout)
	logger.Writers().SetWriter(writer)
	logger.InfoF("NewBufferedWriter...")
	writer.Flush() // Notice that Flush() should be invoked after finishing writing or you may miss some data

	// Of cause we provide a way to change the buffer size of it
	writer = logit.NewBufferedWriter(os.Stdout)
	logger.Writers().SetWriter(writer)
	logger.InfoF("Oh! Faster! Faster!!! Yeah~~")
	writer.Flush() // Notice that Flush() should be invoked after finishing writing or you may miss some data

	// The buffered writer won't flush data automatically in default
	// Does it puzzle you? Try AutoFlush() to get it if you want!
	writer = logit.NewBufferedWriter(os.Stdout)
	writer.AutoFlush(time.Second)
	logger.Writers().SetWriter(writer)
	logger.InfoF("AutoFlush...")
	time.Sleep(2 * time.Second)

*/
package logit // import "github.com/FishGoddess/logit"

const (
	// Version is the version string representation of logit.
	Version = "v0.4.0-alpha"
)
