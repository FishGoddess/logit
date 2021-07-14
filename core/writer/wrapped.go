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
// Created at 2021/07/11 13:55:35

package writer

import "io"

type wrappedWriter struct {
	writer io.Writer
}

func newWrappedWriter(writer io.Writer) *wrappedWriter {
	return &wrappedWriter{writer: writer}
}

func (ww *wrappedWriter) Flush() (n int, err error) {
	if flusher, ok := ww.writer.(Flusher); ok {
		return flusher.Flush()
	}
	return 0, nil
}

func (ww *wrappedWriter) Write(p []byte) (n int, err error) {
	return ww.writer.Write(p)
}

func (ww *wrappedWriter) Close() error {
	if closer, ok := ww.writer.(io.Closer); ok {
		return closer.Close()
	}
	return nil
}
