// Copyright 2023 FishGoddess. All Rights Reserved.
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
)

const (
	// minBufferSize is the min size of buffer.
	// A panic will happen if buffer size is smaller than it.
	minBufferSize = 16
)

// BufferWriter is a writer having a buffer inside to reduce times of writing underlying writer.
// You can set buffer size or use it with default buffer size. Any writer implemented io.Writer can be used by it.
type BufferWriter struct {
	// writer is the underlying writer to write data.
	writer io.Writer

	// maxBufferSize is the max size of buffer.
	maxBufferSize uint64

	// buffer is for keeping data together and writing them one time.
	// Data won't be written to underlying writer if buffer doesn't full, so you can pre-write them by Sync() if you need.
	buffer *bytes.Buffer

	// lock is for safe concurrency.
	lock sync.Mutex
}

// Buffer returns a new buffer writer of writer with specified bufferSize.
// Notice that bufferSize must be larger than minBufferSize or a panic will happen. See minBufferSize.
// The size we want to use is bufferSize, but we add more bytes to it for avoiding buffer growing up.
func Buffer(writer io.Writer, bufferSize uint64) *BufferWriter {
	if bufferSize < minBufferSize {
		panic(fmt.Errorf("bufferSize %d < minBufferSize %d", bufferSize, minBufferSize))
	}

	if bw, ok := writer.(*BufferWriter); ok {
		return bw
	}

	bw := &BufferWriter{
		writer:        writer,
		maxBufferSize: bufferSize,
		buffer:        bytes.NewBuffer(make([]byte, 0, bufferSize+minBufferSize)),
	}

	return bw
}

// Write writes p to buffer and syncs data to underlying writer first if it needs.
func (bw *BufferWriter) Write(p []byte) (n int, err error) {
	bw.lock.Lock()
	defer bw.lock.Unlock()

	// This p is too large, so we write it directly to avoid copying.
	needBufferSize := len(p)
	tooLarge := uint64(needBufferSize) >= bw.maxBufferSize
	if tooLarge {
		// Sync before writing to keep the sequence between writes.
		bw.sync()
		return bw.writer.Write(p)
	}

	// The remaining buffer is not enough, sync data to write this p.
	needBufferSize = bw.buffer.Len() + len(p)
	notEnough := uint64(needBufferSize) >= bw.maxBufferSize
	if notEnough {
		bw.sync()
	}

	return bw.buffer.Write(p)
}

// sync writes data in buffer to underlying writer.
func (bw *BufferWriter) sync() error {
	if _, err := bw.buffer.WriteTo(bw.writer); err != nil {
		return err
	}

	syncer, ok := bw.writer.(interface{ Sync() error })
	if ok && notStdoutAndStderr(bw.writer) {
		return syncer.Sync()
	}

	return nil
}

// Sync writes data in buffer to underlying writer if buffer has data.
// It's safe in concurrency.
func (bw *BufferWriter) Sync() error {
	bw.lock.Lock()
	defer bw.lock.Unlock()

	if bw.buffer.Len() > 0 {
		return bw.sync()
	}

	return nil
}

// Close syncs data and closes underlying writer if writer implements io.Closer.
func (bw *BufferWriter) Close() error {
	bw.lock.Lock()
	defer bw.lock.Unlock()

	if err := bw.sync(); err != nil {
		return err
	}

	if closer, ok := bw.writer.(io.Closer); ok && notStdoutAndStderr(bw.writer) {
		return closer.Close()
	}

	return nil
}
