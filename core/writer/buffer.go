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

	"github.com/go-logit/logit/core"
)

const (
	// minBufferSize is the min size of buffer.
	// A panic will happen if buffer size is smaller than it.
	minBufferSize = 2 * core.B
)

// bufferWriter is a writer having a buffer inside to reduce times of writing underlying writer.
// You can set buffer size or use it with default buffer size. Any writer implemented io.Writer can be used by it.
type bufferWriter struct {
	// writer is the underlying writer to write data.
	writer io.Writer

	// maxBufferSize is the max size of buffer.
	maxBufferSize core.ByteSize

	// buffer is for keeping data together and writing them one time.
	// Data won't be written to underlying writer if buffer doesn't full, so you can pre-write them by Flush() if you need.
	// Also, we provide a way to flush data automatically, see bufferWriter.AutoFlush.
	buffer *bytes.Buffer

	// lock is for safe concurrency.
	lock sync.Mutex
}

// newBufferWriter returns a new buffer writer of this writer with specified bufferSize.
// Notice that bufferSize must be larger than minBufferSize or a panic will happen. See minBufferSize.
// The size we want to use is bufferSize, but we add more bytes to it for avoiding buffer growing up.
func newBufferWriter(writer io.Writer, bufferSize core.ByteSize) *bufferWriter {
	if bufferSize <= minBufferSize {
		panic(fmt.Errorf("bufferSize %d <= minBufferSize %d", bufferSize, minBufferSize))
	}

	return &bufferWriter{
		writer:        writer,
		maxBufferSize: bufferSize,
		buffer:        bytes.NewBuffer(make([]byte, 0, bufferSize+minBufferSize)),
		lock:          sync.Mutex{},
	}
}

// flush writes data in buffer to underlying writer.
func (bw *bufferWriter) flush() (n int, err error) {
	writen, err := bw.buffer.WriteTo(bw.writer)
	return int(writen), err
}

// Flush writes data in buffer to underlying writer if buffer has data.
// It's safe in concurrency.
func (bw *bufferWriter) Flush() (n int, err error) {
	bw.lock.Lock()
	defer bw.lock.Unlock()

	if bw.buffer.Len() > 0 {
		return bw.flush()
	}
	return 0, nil
}

// AutoFlush starts a goroutine to flush data automatically.
// It returns a channel for stopping this goroutine.
func (bw *bufferWriter) AutoFlush(frequency time.Duration) chan<- struct{} {
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
func (bw *bufferWriter) Write(p []byte) (n int, err error) {
	bw.lock.Lock()
	defer bw.lock.Unlock()

	// This p is too large, so we write it directly to avoid copying.
	tooLarge := core.ByteSize(len(p)) >= bw.maxBufferSize
	if tooLarge {
		bw.flush() // Flush before writing to keep the sequence between writes.
		return bw.writer.Write(p)
	}

	// The remaining buffer is not enough, flush data to write this p.
	notEnough := core.ByteSize(bw.buffer.Len()+len(p)) >= bw.maxBufferSize
	if notEnough {
		bw.flush()
	}
	return bw.buffer.Write(p)
}

// Close flushes data and closes underlying writer if writer implements io.Closer.
func (bw *bufferWriter) Close() error {
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
