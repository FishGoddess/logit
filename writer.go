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
	"fmt"
	"os"
	"sync"
)

const (
	// maxRetriedTimes is the max times which will retry after error.
	maxRetriedTimes = 10
)

// FileWriter writes logs to one or more files.
// We provide a powerful writer of file, which can roll to several files
// automatically in some conditions.
type FileWriter struct {

	// name is the name of log file.
	name string

	// file is a pointer to the real file in os.
	file *os.File

	// seq is the serial number of rolling file.
	// If name is "test.log" and seq is 1, then the file rolled will be "test.log.0000000001"
	seq int

	// checkers stores all checkers will be used.
	// If one of them say: "Time to roll!", then this file writer will start to roll.
	// After rolling, the checkers after it will be skipped in this loop.
	checkers []Checker

	// lock is for safe-concurrency.
	lock *sync.Mutex
}

// NewFileWriter returns a new file writer with given name and checkers.
func NewFileWriter(name string, checkers ...Checker) (*FileWriter, error) {

	file, err := os.OpenFile(name, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	return &FileWriter{
		name:     name,
		file:     file,
		seq:      0,
		checkers: checkers,
		lock:     &sync.Mutex{},
	}, nil
}

// roll changes fw.file to a new file and nothing happens if failed.
func (fw *FileWriter) roll() {

	fw.seq++
	fw.file.Close()
	for i := 0; i < maxRetriedTimes; i++ {
		if os.Rename(fw.name, fmt.Sprintf("%s.%.10d", fw.name, fw.seq)) != nil {
			continue
		}
		for j := 0; j < maxRetriedTimes; j++ {
			if newFile, err := os.OpenFile(fw.name, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644); err == nil {
				fw.file = newFile
				return
			}
		}
	}
}

// Write writes len(p) bytes from p to file and returns an error if failed.
// The precise count of written bytes is n.
func (fw *FileWriter) Write(p []byte) (n int, err error) {

	fw.lock.Lock()
	defer fw.lock.Unlock()

	// Check rolling condition first and replace to newFile only if roll returns nil error
	for _, checker := range fw.checkers {
		if checker.Check(fw, len(p)) {
			fw.roll()
		}
	}
	return fw.file.Write(p)
}

// Close closes this file writer and returns an error if failed.
func (fw *FileWriter) Close() error {
	fw.lock.Lock()
	defer fw.lock.Unlock()
	fw.checkers = nil
	return fw.file.Close()
}
