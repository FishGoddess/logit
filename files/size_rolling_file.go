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
// Created at 2020/03/05 00:11:50

package files

import (
	"errors"
	"os"
	"strconv"
	"sync"
	"time"
)

// SizeRollingFile is a file size sensitive file.
//
//  file := NewSizeRollingFile(64*KB, func (now time.Time) string {
//      return "D:/" + now.Format("20060102150405.000") + ".log"
//  })
//  defer file.Close()
//  file.Write([]byte("Hello!"))
//
// You can use it like using os.File!
type SizeRollingFile struct {

	// file points the writer which will be used this moment.
	file *os.File

	// directory is the target storing all created files.
	directory string

	// limitedSize is the limited size of this file.
	// File will roll to next file if its size has reached to limitedSize.
	// This field should be always larger than minLimitedSize for some safe considerations.
	// Notice that it doesn't mean every file must be this size due to our performance optimization
	// scheme. Generally it equals to file size, however, it will not equal to file size
	// if someone modified this file. See currentSize.
	limitedSize int64

	// currentSize equals to the size of current file.
	// The currentSize will reset to 0 when rolling to next file.
	// The reason why we set this field is file.Stat() is too expensive!
	// Every writing operations will fetch file size, that means each operation
	// will call file.Stat() for size. It's not a good way to fetch file size.
	// So we keep one field inside, and record current size of current file with it.
	// Each time fetching file size, the only thing wo do is checking this field.
	// This way is cheaper, even cheapest. Of course, we should maintain this field
	// inside for precision, so it doesn't mean we won't call file.Stat() anymore.
	// If currentSize >= limitedSize, we will still call file.Stat() for precision.
	// Certainly, we will set currentSize to the real size of file. Hey, you know we
	// won't waste the time we spent on file.Stat() ^_^.
	currentSize int64

	// nextFilename is a function for generating next file name.
	// Every times rolling to next file will call it first.
	// now is the time of calling this function, also the
	// created time of next file.
	nextFilename func(now time.Time) string

	// nameGenerator is for generating the name of every created file.
	// You can customize your format of filename by implementing this function.
	// Default is DefaultNameGenerator().
	nameGenerator NameGenerator

	// rollingHook is a hook that will be invoked in rolling process.
	// This interface has two method: BeforeRolling and AfterRolling.
	// BeforeRolling will be called before rolling to next file.
	// AfterRolling will be called after rolling to next file.
	// Default is DefaultRollingHook, and it will do nothing when rolling to next file.
	rollingHook RollingHook

	// mu is a lock for safe concurrency.
	mu *sync.Mutex
}

const (
	// minLimitedSize prevents io system from creating file too fast.
	// Default is 64 KB (64 * 1024 bytes).
	minLimitedSize = 64 * KB
)

// NewSizeRollingFile creates a new size rolling file.
// limitedSize is how big did it roll to next file.
// nextFilename is a function for generating next file name.
// Every times rolling to next file will call nextFilename first.
// now is the created time of next file. Notice that limitedSize's min value
// is 64 KB (64 * 1024 bytes). See minLimitedSize.
func NewSizeRollingFile(directory string, limitedSize int64) *SizeRollingFile {

	// 防止文件限制尺寸太小导致滚动文件时 IO 的疯狂蠕动
	if limitedSize < minLimitedSize {
		panic(errors.New("LimitedSize is smaller than " + strconv.FormatUint(uint64(minLimitedSize)>>10, 10) + " KB!\n"))
	}

	return &SizeRollingFile{
		directory:     directory,
		limitedSize:   limitedSize,
		currentSize:   0,
		nameGenerator: DefaultNameGenerator(),
		rollingHook:   NewDefaultRollingHook(),
		mu:            &sync.Mutex{},
	}
}

// doRollingTask does the whole task of rolling process.
func (srf *SizeRollingFile) doRollingTask() {
	srf.rollingHook.BeforeRolling()
	srf.rollingToNextFile(time.Now())
	srf.rollingHook.AfterRolling()
}

// rollingToNextFile will roll to next file for srf.
func (srf *SizeRollingFile) rollingToNextFile(now time.Time) {

	// 如果创建新文件发生错误，就继续使用当前的文件，等到下一次时间间隔再重试
	newFile, err := CreateFileOf(srf.nameGenerator.NextName(srf.directory, now))
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

	// file 为 nil，进行初始化
	if srf.file == nil {
		srf.doRollingTask()
		return
	}

	// 判断文件大小是否超过限制值
	if srf.currentSize >= srf.limitedSize {

		// 这时候还不能确定 currentSize 是正确的，需要通过系统调用查询文件真实大小
		fileInfo, err := srf.file.Stat()

		// 需要划分文件的两种情况：
		// 1. err != nil，获取文件真实大小失败，选择相信 currentSize
		// 2. 真实文件大小确实大于 limitedSize
		if err != nil || fileInfo.Size() >= srf.limitedSize {
			srf.doRollingTask()
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

// SetNameGenerator replaces srf.nameGenerator to newNameGenerator.
func (srf *SizeRollingFile) SetNameGenerator(nameGenerator NameGenerator) {
	srf.mu.Lock()
	defer srf.mu.Unlock()
	srf.nameGenerator = nameGenerator
}

// SetRollingHook replaces srf.rollingHook to newRollingHook.
func (srf *SizeRollingFile) SetRollingHook(rollingHook RollingHook) {
	srf.mu.Lock()
	defer srf.mu.Unlock()
	srf.rollingHook = rollingHook
}
