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
	"io"
)

// wrapWriter is a Writer implement for wrapping normal writer.
type wrapWriter struct {
	writer io.Writer
}

// newWrapWriter returns an wrapWriter of writer.
func newWrapWriter(writer io.Writer) *wrapWriter {
	return &wrapWriter{writer: writer}
}

// Sync syncs data to underlying writer.
func (ww *wrapWriter) Sync() error {
	if syncer, ok := ww.writer.(Syncer); ok && notStdoutAndStderr(ww.writer) {
		return syncer.Sync()
	}
	return nil
}

// Write writes data to underlying writer.
func (ww *wrapWriter) Write(p []byte) (n int, err error) {
	return ww.writer.Write(p)
}

// Close closes the underlying writer.
func (ww *wrapWriter) Close() error {
	if closer, ok := ww.writer.(io.Closer); ok && notStdoutAndStderr(ww.writer) {
		return closer.Close()
	}
	return nil
}
