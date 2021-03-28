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
	"container/list"
	"fmt"
	"os"
	"sync"
	"time"
)

const (
	// maxRetriedTimes is the max times which will retry after error.
	maxRetriedTimes = 10

	// timeFormat is the format of time.
	timeFormat = "20060102"
)

// ======================================== queue ========================================

// queue is a queue for limiting number of log files.
type queue struct {

	// l stores all file names.
	l *list.List

	// maxSize is the max number of file names in queue.
	maxNumber int
}

// newQueue returns a queue instance.
func newQueue(maxNumber int) *queue {
	return &queue{
		l:         list.New(),
		maxNumber: maxNumber,
	}
}

// push pushes a file name to queue and executes onRemove when removing an item in queue.
func (q *queue) push(fileName string, onRemove func(fileName string)) {
	if q.l.Len() >= q.maxNumber {
		onRemove(q.l.Remove(q.l.Front()).(string))
	}
	q.l.PushBack(fileName)
}

// ======================================== FileWriter ========================================

// FileWriter writes logs to one or more files.
// We provide a powerful writer of file, which can roll to several files
// automatically in some conditions.
type FileWriter struct {

	// config is the configuration of file writer.
	config Config

	// currentSize equals to the size of current file checked.
	// The currentSize will reset to 0 when checked file has rolled to the next file.
	// The reason why we set this field is file.Stat() is too expensive!
	// Every writing operations will fetch file size, that means each operation
	// will call file.Stat() for size. It's not a good way to fetch file size.
	// So we keep one field inside, and record size of current file by it.
	// Each time fetching file size, the only thing wo do is checking this field.
	// This way is cheaper, even cheapest. Of course, we should maintain this field
	// inside for precision, so it doesn't mean we won't call file.Stat() anymore.
	// If currentSize >= limitedSize, we will still call file.Stat() for precision.
	// Certainly, we will set currentSize to the real size of file. Hey, you know we
	// won't waste any time we spent on file.Stat() ^_^.
	currentSize int64

	// fileQueue is a queue limiting number of log files.
	fileQueue *queue

	// file is a pointer to the real file in os.
	file *os.File

	// lock is for safe-concurrency.
	lock *sync.RWMutex
}

// NewFileWriter returns a new file writer with given name and checkers.
func NewFileWriter(config Config) (*FileWriter, error) {

	checkConfig(config)

	w := &FileWriter{
		config:    config,
		fileQueue: newQueue(config.MaxLogFileNumber),
		lock:      &sync.RWMutex{},
	}

	w.prepareLogFile(time.Now())
	go w.startUpdateFileTask()
	return w, nil
}

// newLogFileName returns a new log file name of date.
func (fw *FileWriter) newLogFileName(date time.Time) string {
	return fmt.Sprintf("%s.%s", fw.config.LogFileName, date.Format(timeFormat))
}

// prepareLogFile ensures log file is ready and right at time of date.
func (fw *FileWriter) prepareLogFile(date time.Time) {

	for i := 0; i < maxRetriedTimes; i++ {
		newName := fw.newLogFileName(date)
		newFile, err := os.OpenFile(newName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err == nil && (fw.file == nil || fw.file.Close() == nil) {
			fw.file = newFile
			fw.fileQueue.push(newFile.Name(), func(fileName string) {
				os.Remove(fileName)
			})
			break
		}
	}
}

// startUpdateFileTask updates fw.file to a new file automatically and nothing happens if failed.
func (fw *FileWriter) startUpdateFileTask() {

	for {
		now := time.Now()
		next := now.Add(24 * time.Hour)
		timer := time.NewTimer(time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, time.Local).Sub(now))
		select {
		case t := <-timer.C:
			fw.lock.Lock()
			fw.prepareLogFile(t)
			fw.lock.Unlock()
		}
	}
}

// Write writes len(p) bytes from p to file and returns an error if failed.
// The precise count of written bytes is n.
func (fw *FileWriter) Write(p []byte) (n int, err error) {
	fw.lock.RLock()
	defer fw.lock.RUnlock()
	return fw.file.Write(p)
}

// Close closes this file writer and returns an error if failed.
func (fw *FileWriter) Close() error {
	fw.lock.Lock()
	defer fw.lock.Unlock()
	return fw.file.Close()
}
