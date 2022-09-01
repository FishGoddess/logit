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

package main

import (
	"os"

	"github.com/go-logit/logit"
	"github.com/go-logit/logit/core/writer"
	"github.com/go-logit/logit/support/size"
)

func main() {
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
}
