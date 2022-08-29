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

package file

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/go-logit/logit/support/size"
)

var (
	Mode os.FileMode = 0644

	CreateFile = func(filename string) (*os.File, error) {
		return os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, Mode)
	}
)

type File struct {
	filename  string
	directory string

	maxSize     size.ByteSize
	currentSize size.ByteSize

	file    *os.File
	eventCh chan struct{}
	lock    sync.Mutex
}

func New(filename string, maxSize size.ByteSize) (*File, error) {
	// TODO 处理参数的合法性校验，考虑使用 options 机制创建
	directory := filepath.Dir(filename)
	err := os.MkdirAll(directory, Mode)
	if err != nil {
		return nil, err
	}

	file, err := CreateFile(filename)
	if err != nil {
		return nil, err
	}

	f := &File{
		filename:    filename,
		directory:   directory,
		maxSize:     maxSize,
		currentSize: 0,
		file:        file,
		eventCh:     make(chan struct{}, 1),
	}

	go f.handleEvents()
	return f, nil
}

func (f *File) handleEvents() {
	for range f.eventCh {
		// TODO 处理过期文件的清理和压缩等工作
	}
}

func (f *File) roll() error {
	f.lock.Lock()
	defer f.lock.Unlock()

	err := f.close()
	if err != nil {
		return err
	}

	err = os.Remove(f.filename)
	if err != nil {
		return err
	}

	// TODO ...
	return nil
}

// beforeWrite do some checks before writing.
func (f *File) beforeWrite(p []byte) (err error) {
	if f.file == nil {
		_, err := os.Stat(f.filename)
		if os.IsNotExist(err) {
			f.file, err = CreateFile(f.filename)
			if err != nil {
				return err
			}
		}

		if err != nil {
			// TODO ...
		}
	}

	writeSize := size.ByteSize(len(p))
	if writeSize > f.maxSize-f.currentSize {
		f.roll() // ignore rolling error so this p won't be discarded.
	}
	return nil
}

func (f *File) Write(p []byte) (n int, err error) {
	f.lock.Lock()
	defer f.lock.Unlock()

	err = f.beforeWrite(p)
	if err != nil {
		return 0, err
	}

	n, err = f.file.Write(p)
	if err != nil {
		return 0, err
	}

	f.currentSize += size.ByteSize(n)
	return n, nil
}

func (f *File) close() (err error) {
	if f.file == nil {
		return nil
	}

	file := f.file
	f.file = nil
	return file.Close()
}

func (f *File) Close() (err error) {
	f.lock.Lock()
	defer f.lock.Unlock()
	return f.close()
}
