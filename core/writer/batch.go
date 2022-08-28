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

package writer

import (
	"bytes"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/go-logit/logit/support/global"
)

const (
	// minBatchCount is the min count of batch.
	// A panic will happen if batch count is smaller than it.
	minBatchCount = 1
)

// batchWriter is a writer having a buffer inside to reduce times of writing underlying writer.
// You can set batch count or use it with default batch count. Any writer implemented io.Writer can be used by it.
type batchWriter struct {
	// writer is the underlying writer to write data.
	writer io.Writer

	// maxBatchCount is the max count of batch.
	maxBatchCount uint

	// currentBatchCount is the current count of batch.
	currentBatchCount uint

	// buffer is for keeping data together and writing them one time.
	// Data won't be written to underlying writer if batch count is less than max batch count, so you can pre-write them by Flush() if you need.
	// Also, we provide a way to flush data automatically, see batchWriter.AutoFlush.
	buffer *bytes.Buffer

	// lock is for safe concurrency.
	lock sync.Mutex
}

// newBatchWriter returns a new batch writer of this writer with specified batchCount.
// Notice that batchCount must be larger than minBatchCount or a panic will happen. See minBatchCount.
func newBatchWriter(writer io.Writer, batchCount uint) *batchWriter {
	if batchCount < minBatchCount {
		panic(fmt.Errorf("batchCount %d < minBatchCount %d", batchCount, minBatchCount))
	}

	return &batchWriter{
		writer:            writer,
		maxBatchCount:     batchCount,
		currentBatchCount: 0,
		buffer:            bytes.NewBuffer(make([]byte, 0, global.WriterBufferSize)),
		lock:              sync.Mutex{},
	}
}

// flush writes data in buffer to underlying writer.
func (bw *batchWriter) flush() (n int, err error) {
	writen, err := bw.buffer.WriteTo(bw.writer)
	return int(writen), err
}

// Flush writes data in buffer to underlying writer if buffer has data.
// It's safe in concurrency.
func (bw *batchWriter) Flush() (n int, err error) {
	bw.lock.Lock()
	defer bw.lock.Unlock()

	if bw.buffer.Len() > 0 {
		return bw.flush()
	}
	return 0, nil
}

// AutoFlush starts a goroutine to flush data automatically.
// It returns a channel for stopping this goroutine.
func (bw *batchWriter) AutoFlush(frequency time.Duration) chan<- struct{} {
	stopCh := make(chan struct{}, 1)

	go func() {
		ticker := time.NewTicker(frequency)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				bw.Flush()
			case <-stopCh:
				return
			}
		}
	}()

	return stopCh
}

// Write writes p to buffer and flushes data to underlying writer first if it needs.
func (bw *batchWriter) Write(p []byte) (n int, err error) {
	bw.lock.Lock()
	defer bw.lock.Unlock()

	needFlush := bw.currentBatchCount >= bw.maxBatchCount
	if needFlush {
		bw.flush()
		bw.currentBatchCount = 0
	}

	bw.currentBatchCount++
	return bw.buffer.Write(p)
}

// Close flushes data and closes underlying writer if writer implements io.Closer.
func (bw *batchWriter) Close() error {
	bw.lock.Lock()
	defer bw.lock.Unlock()

	_, err := bw.flush()
	if err != nil {
		return err
	}

	if closer, ok := bw.writer.(io.Closer); ok && notStdoutAndStderr(bw.writer) {
		return closer.Close()
	}
	return nil
}
