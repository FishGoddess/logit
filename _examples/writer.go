// Copyright 2021 Ye Zi Jie. All Rights Reserved.
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
// Created at 2021/07/11 23:33:04

package main

import (
	"os"

	"github.com/FishGoddess/logit"
	"github.com/FishGoddess/logit/core/writer"
)

func main() {
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
	writer.Wrapped(os.Stdout)  // Wrap io.Writer to writer.Writer.
	writer.Buffered(os.Stderr) // Wrap io.Writer to writer.Writer with buffer, which needs invoking Flush() or Close().

	// Use the writer without buffer.
	logger := logit.NewLogger(logit.Options().WithWriter(os.Stdout, false))
	logger.Info("WriterWithoutBuffer").End()

	// Use the writer with buffer, which is good for io.
	logger = logit.NewLogger(logit.Options().WithWriter(os.Stdout, true))
	defer logger.Close() // Flush data and close writer.

	logger.Info("WriterWithBuffer").End()
	logger.Flush() // Remember flushing data or flushing by Close().

	// Every level has its own appender so you can append logs in different level with different appender.
	logger = logit.NewLogger(
		logit.Options().WithDebugWriter(os.Stdout, true),
		logit.Options().WithInfoWriter(os.Stdout, true),
		logit.Options().WithWarnWriter(os.Stdout, false),
		logit.Options().WithErrorWriter(os.Stdout, false),
	)
}
