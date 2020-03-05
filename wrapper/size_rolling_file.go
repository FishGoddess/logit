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
// Created at 2020/03/05 00:11:50

package wrapper

import (
    "fmt"
    "os"
    "sync"
    "time"
)

// SizeRollingFile 是按照文件大小自动划分的文件类型
type SizeRollingFile struct {

    // file points the writer which will be used this moment.
    file *os.File

    // limitedSize 是指这个文件限制的大小，达到这个大小之后就会自动划分为另一个文件
    // 注意：出于对写操作的性能优化，采取了自计数方式的设计方案，参考 currentSize
    limitedSize int64

    // currentSize 表示当前文件大小
    // 由于每一次写操作都需要获取文件大小，这样会发起一次系统调用，file.Stat() 非常耗时
    // 这里将文件大小记录起来，每一次写操作直接使用这个大小进行判断
    // 这在没有外界修改过文件的情况下一般是没有问题的，而且能提高性能
    // 只有当判断到 currentSize >= limitedSize 的时候，才进行一次系统调用查询真实的文件大小
    // 如果真实的文件大小大于 currentSize，此时需要划分文件，并设置 currentSize 为 0
    // 如果真实的文件大小小于 currentSize，重新设置 currentSize，此时不划分文件
    // 注意：这个设计方案可以保证文件不小于 limitedSize 的大小
    currentSize int64

    // nextFilename is a function for generating a new file name.
    // Every times rolling to next file will call it first.
    // now is the time of calling this function, also the
    // created time of next file.
    nextFilename func(now time.Time) string

    // mu is a lock for safe concurrency.
    mu sync.Mutex
}

// minLimitedSize prevents io system from creating file too fast.
// Default is 64 KB (64 * 1024 bytes).
const minLimitedSize = 64 * KB

// NewSizeRollingFile creates a new size rolling file.
// limitedSize is how big did it roll to next file.
// nextFilename is a function for generating a new file name.
// Every times rolling to next file will call nextFilename first.
// now is the time of calling this function, also the created time of next file.
// Notice that limitedSize's min value is 64 KB (64 * 1024 bytes). See minLimitedSize.
func NewSizeRollingFile(limitedSize int64, nextFilename func(now time.Time) string) *SizeRollingFile {

    // 防止文件限制尺寸太小导致滚动文件时 IO 的疯狂蠕动
    if limitedSize < minLimitedSize {
        panic(fmt.Errorf("LimitedSize is smaller than %v KB!\n", uint64(minLimitedSize)>>10))
    }

    // 获取当前时间，并生成第一个文件
    file, _ := generateFirstFile(nextFilename)
    return &SizeRollingFile{
        file:         file,
        limitedSize:  limitedSize,
        currentSize:  0,
        nextFilename: nextFilename,
        mu:           sync.Mutex{},
    }
}

// rollingToNextFile will roll to next file for srf.
func (srf *SizeRollingFile) rollingToNextFile(now time.Time) {

    // 如果创建新文件发生错误，就继续使用当前的文件，等到下一次时间间隔再重试
    newFile, err := NewFile(srf.nextFilename(now))
    if err != nil {
        return
    }

    // 关闭当前使用的文件，初始化新文件
    srf.file.Close()
    srf.file = newFile
    srf.currentSize = 0
}

// ensureFileIsCorrect ensures srf is writing to a correct file this moment.
func (srf *SizeRollingFile) ensureFileIsCorrect() {
    if srf.currentSize >= srf.limitedSize {

        // 这时候还不能确定 currentSize 是正确的，需要通过系统调用查询文件真实大小
        fileInfo, err := srf.file.Stat()

        // 需要划分文件的两种情况：
        // 1. err != nil，获取文件真实大小失败，选择相信 currentSize
        // 2. 真实文件大小确实大于 limitedSize
        if err != nil || fileInfo.Size() >= srf.limitedSize {
            srf.rollingToNextFile(time.Now())
            return
        }

        // 否则修正 currentSize 为真实文件大小，不能浪费这一次系统调用
        srf.currentSize = fileInfo.Size()
    }
}

// writeAndUpdateCurrentSize writes p to srf.file and updates srf.currentSize with n.
func (srf *SizeRollingFile) writeAndUpdateCurrentSize(p []byte) (int, error) {
    n, err := srf.file.Write(p)
    srf.currentSize += int64(n)
    return n, err
}

// Write writes len(p) bytes from p to the underlying data stream.
// It returns the number of bytes written from p (0 <= n <= len(p))
// and any error encountered that caused the write to stop early.
func (srf *SizeRollingFile) Write(p []byte) (n int, err error) {
    srf.mu.Lock()
    defer srf.mu.Unlock()

    // 确保当前文件对于当前时间点来说是正确的
    srf.ensureFileIsCorrect()
    return srf.writeAndUpdateCurrentSize(p)
}

// Close releases any resources using just moment.
// It returns error when closing.
func (srf *SizeRollingFile) Close() error {
    srf.mu.Lock()
    defer srf.mu.Unlock()
    return srf.file.Close()
}
