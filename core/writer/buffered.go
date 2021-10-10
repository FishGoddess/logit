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
// Created at 2020/11/30 22:17:07

package writer

import (
	"bytes"
	"fmt"
	"io"
	"sync"
	"time"
)

const (
	// minBufferSize is the min size of buffer.
	// A panic will happen if buffer size is smaller than it.
	minBufferSize = 4
)

// bufferedWriter is a writer having a buffer inside to reduce times of writing underlying writer.
// You can set buffer size or use it with default buffer size. Any writer implemented io.Writer can be used by it.
type bufferedWriter struct {
	// writer is the underlying writer to write data.
	writer io.Writer

	// buffer is for keeping data together and writing them one time.
	// Data won't be wrote to underlying writer if buffer doesn't full, so you can pre-write them by Flush() if you need.
	// Also, we provide a way to flush data automatically, see bufferedWriter.AutoFlush.
	buffer *bytes.Buffer

	// bufferSize is the size of buffer.
	bufferSize int

	// lock is for safe concurrency.
	lock sync.Mutex
}

// newBufferedWriter returns a new buffered writer of this writer with specified bufferSize.
// Notice that bufferSize must be larger than minBufferSize or a panic will happen. See minBufferSize.
// The size we want to use is bufferSize, but we add more bytes to it for avoiding buffer growing up.
func newBufferedWriter(writer io.Writer, bufferSize int) *bufferedWriter {
	if bufferSize <= minBufferSize {
		panic(fmt.Errorf("logit.NewBufferedWriterWithSize got a bufferSize %d smaller than minBufferSize %d", bufferSize, minBufferSize))
	}

	return &bufferedWriter{
		writer:     writer,
		buffer:     bytes.NewBuffer(make([]byte, 0, bufferSize+minBufferSize)),
		bufferSize: bufferSize,
		lock:       sync.Mutex{},
	}
}

// needFlush returns if bw need flush or not.
func (bw *bufferedWriter) needFlush() bool {
	return bw.buffer.Len() > bw.buffer.Cap()-minBufferSize // remain some bytes for avoiding buffer growing...
}

// flush writes data in buffer to underlying writer.
func (bw *bufferedWriter) flush() (n int, err error) {
	writen, err := bw.buffer.WriteTo(bw.writer)
	return int(writen), err
}

// Flush writes data in buffer to underlying writer if buffer has data.
// It's safe in concurrency.
func (bw *bufferedWriter) Flush() (n int, err error) {
	bw.lock.Lock()
	defer bw.lock.Unlock()

	if bw.buffer.Len() > 0 {
		return bw.flush()
	}

	return 0, nil
}

// AutoFlush starts a goroutine to flush data automatically.
// It returns a channel for stopping this goroutine.
func (bw *bufferedWriter) AutoFlush(frequency time.Duration) chan<- struct{} {
	stopChan := make(chan struct{}, 1)
	go func() {
		ticker := time.NewTicker(frequency)
		for {
			select {
			case <-ticker.C:
				bw.Flush()
			case <-stopChan:
				ticker.Stop()
				return
			}
		}
	}()

	return stopChan
}

// Write writes p to buffer and flushes data to underlying writer first if need.
func (bw *bufferedWriter) Write(p []byte) (n int, err error) {
	bw.lock.Lock()
	defer bw.lock.Unlock()

	if bw.needFlush() {
		bw.flush() // ignore error so this p can be written to buffer which may cause buffer grows up.
	}

	return bw.buffer.Write(p)
}

// Close flushes data and closes underlying writer if writer implements io.Closer.
func (bw *bufferedWriter) Close() error {
	bw.lock.Lock()
	defer bw.lock.Unlock()

	_, err := bw.flush()
	if err != nil {
		return err
	}

	if wCloser, ok := bw.writer.(io.Closer); ok && notStdoutAndStderr(bw.writer) {
		return wCloser.Close()
	}

	return nil
}
