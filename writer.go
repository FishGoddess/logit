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
// Created at 2020/11/30 22:17:07

package logit

import (
	"bytes"
	"fmt"
	"io"
	"sync"
	"time"
)

const (
	KB = 1024

	minBufferSize     = 4
	defaultBufferSize = 16*KB + minBufferSize
)

type BufferedWriter struct {
	writer     io.Writer
	buffer     *bytes.Buffer
	bufferSize int
	lock       *sync.Mutex
}

func NewBufferedWriter(writer io.Writer) *BufferedWriter {
	return NewBufferedWriterWithSize(writer, defaultBufferSize)
}

func NewBufferedWriterWithSize(writer io.Writer, bufferSize int) *BufferedWriter {

	if bufferSize <= minBufferSize {
		panic(fmt.Errorf("logit.NewBufferedWriterWithSize got a bufferSize %d smaller than minBufferSize %d", bufferSize, minBufferSize))
	}

	result := &BufferedWriter{
		writer:     writer,
		buffer:     bytes.NewBuffer(make([]byte, 0, bufferSize)),
		bufferSize: bufferSize,
		lock:       &sync.Mutex{},
	}
	result.AutoFlush()
	return result
}

func (bf *BufferedWriter) needFlush() bool {
	return bf.buffer.Len() > bf.buffer.Cap()-minBufferSize // remain some bytes for avoiding buffer growing...
}

func (bf *BufferedWriter) flush() (n int, err error) {
	writen, err := bf.buffer.WriteTo(bf.writer)
	return int(writen), err
}

func (bf *BufferedWriter) Flush() (n int, err error) {

	bf.lock.Lock()
	defer bf.lock.Unlock()

	if bf.buffer.Len() > 0 {
		return bf.flush()
	}
	return 0, nil
}

func (bf *BufferedWriter) AutoFlush() {
	go func() {
		for {
			bf.Flush()
			time.Sleep(time.Second)
		}
	}()
}

func (bf *BufferedWriter) Write(p []byte) (n int, err error) {

	bf.lock.Lock()
	defer bf.lock.Unlock()

	if bf.needFlush() {
		bf.flush() // ignore error so this p can be written to buffer which may grow up
	}
	return bf.buffer.Write(p)
}

func (bf *BufferedWriter) Close() error {

	bf.lock.Lock()
	defer bf.lock.Unlock()

	_, err := bf.flush()
	if err != nil {
		return err
	}

	if wCloser, ok := bf.writer.(io.Closer); ok {
		return wCloser.Close()
	}
	return nil
}
