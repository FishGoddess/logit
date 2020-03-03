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
// Author: fish
// Email: fishinlove@163.com
// Created at 2020/03/03 14:58:21

package wrapper

import (
    "fmt"
    "os"
    "sync"
    "time"
)

// DurationRollingFile is a time sensitive file.
//
//  file := NewDurationRollingFile(time.Second, func(lastTime, currentTime time.Time) string {
//      return root + currentTime.Format(formatOfTime) + PrefixOfLogFile
//  })
//  defer file.Close()
//  file.Write([]byte("Hello!"))
//
// You can use it like using os.File!
type DurationRollingFile struct {

    // file points the writer which will be used this moment.
    file *os.File

    // lastTime is the created time of current file above.
    lastTime time.Time

    // duration is the core property of this struct.
    // Every times currentTime - lastTime >= duration, the file will
    // roll to an entire new file for writing.
    // Notice that its min value is time.Millisecond. See minDuration.
    duration time.Duration

    // nextFilename is a function for generating a new file name.
    // Every times rolling to next file will call it first.
    // lastTime is the time of previous file using.
    // currentTime is the time of calling this function, also the
    // created time of next file.
    nextFilename func(lastTime, currentTime time.Time) string

    // mu is a lock for safe concurrency.
    mu sync.Mutex
}

// minDuration prevents io system from creating file too fast.
const minDuration = time.Millisecond

// NewDurationRollingFile creates a new duration rolling file.
// duration is how long did it roll to next file.
// nextFilename is a function for generating a new file name.
// Every times rolling to next file will call nextFilename first.
// lastTime is the time of previous file using.
// currentTime is the time of calling this function, also the
// created time of next file.
// Notice that its min value is time.Millisecond. See minDuration.
func NewDurationRollingFile(duration time.Duration, nextFilename func(lastTime, currentTime time.Time) string) *DurationRollingFile {

    // 防止时间间隔太小导致滚动文件时 IO 的疯狂蠕动
    if duration < minDuration {
        panic(fmt.Errorf("Duration is smaller than %v!\n", minDuration))
    }

    // 获取当前时间，并生成第一个文件
    now := time.Now()
    file, err := NewFile(nextFilename(now, now))
    if err != nil {
        panic(err)
    }

    return &DurationRollingFile{
        file:         file,
        lastTime:     now,
        duration:     duration,
        nextFilename: nextFilename,
        mu:           sync.Mutex{},
    }
}

// rollingToNextFile will roll to next file for drf.
func (drf *DurationRollingFile) rollingToNextFile(now time.Time) {

    // 如果创建新文件发生错误，就继续使用当前的文件，等到下一次时间间隔再重试
    newFile, err := NewFile(drf.nextFilename(drf.lastTime, now))
    if err != nil {
        return
    }

    // 关闭当前使用的文件，初始化新文件
    drf.file.Close()
    drf.file = newFile
    drf.lastTime = now
}

// ensureFileIsCorrect ensures drf is writing to a correct file this moment.
func (drf *DurationRollingFile) ensureFileIsCorrect() {
    now := time.Now()
    if now.Sub(drf.lastTime) >= drf.duration {
        drf.rollingToNextFile(now)
    }
}

// Write writes len(p) bytes from p to the underlying data stream.
// It returns the number of bytes written from p (0 <= n <= len(p))
// and any error encountered that caused the write to stop early.
func (drf *DurationRollingFile) Write(p []byte) (n int, err error) {
    drf.mu.Lock()
    defer drf.mu.Unlock()

    // 确保当前文件对于当前时间点来说是正确的
    drf.ensureFileIsCorrect()
    return drf.file.Write(p)
}

// Close releases any resources using just moment.
// It returns error when closing.
func (drf *DurationRollingFile) Close() error {
    drf.mu.Lock()
    defer drf.mu.Unlock()
    return drf.file.Close()
}
