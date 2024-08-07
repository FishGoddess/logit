// Copyright 2024 FishGoddess. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
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
)

const (
	// minBatchSize is the min size of batch.
	// A panic will happen if batch size is smaller than it.
	minBatchSize = 1
)

// BatchWriter is a writer having a buffer inside to reduce times of writing underlying writer.
type BatchWriter struct {
	// writer is the underlying writer to write data.
	writer io.Writer

	// maxBatches is the max size of batch.
	maxBatches uint64

	// currentBatches is the current size of batch.
	currentBatches uint64

	// buffer is for keeping data together and writing them one time.
	// Data won't be written to underlying writer if batch size is less than max batch size,
	// so you can pre-write them by Sync() if you want.
	buffer *bytes.Buffer

	lock sync.Mutex
}

// Batch returns a new batch writer of writer with specified batchSize.
// Notice that batchSize must be larger than minBatchSize or a panic will happen.
// See minBatchSize.
func Batch(writer io.Writer, batchSize uint64) *BatchWriter {
	if batchSize < minBatchSize {
		panic(fmt.Errorf("logit: batchSize %d < minBatchSize %d", batchSize, minBatchSize))
	}

	if bw, ok := writer.(*BatchWriter); ok {
		return bw
	}

	bw := &BatchWriter{
		writer:         writer,
		maxBatches:     batchSize,
		currentBatches: 0,
		buffer:         bytes.NewBuffer(make([]byte, 0, defaultBufferSize)),
	}

	return bw
}

// Write writes p to buffer and syncs data to underlying writer first if it needs.
func (bw *BatchWriter) Write(p []byte) (n int, err error) {
	bw.lock.Lock()
	defer bw.lock.Unlock()

	if bw.currentBatches >= bw.maxBatches {
		bw.sync()
		bw.currentBatches = 0
	}

	bw.currentBatches++
	return bw.buffer.Write(p)
}

func (bw *BatchWriter) sync() error {
	_, err := bw.buffer.WriteTo(bw.writer)
	return err
}

// Sync writes data in buffer to underlying writer if buffer has data.
// It's safe in concurrency.
func (bw *BatchWriter) Sync() error {
	bw.lock.Lock()
	defer bw.lock.Unlock()

	if bw.buffer.Len() > 0 {
		return bw.sync()
	}

	return nil
}

func (bw *BatchWriter) close() error {
	if closer, ok := bw.writer.(io.Closer); ok && notStdoutAndStderr(bw.writer) {
		return closer.Close()
	}

	return nil
}

// Close syncs data and closes underlying writer if writer implements io.Closer.
func (bw *BatchWriter) Close() error {
	bw.lock.Lock()
	defer bw.lock.Unlock()

	if err := bw.sync(); err != nil {
		return err
	}

	return bw.close()
}
