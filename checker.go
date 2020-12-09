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
// Created at 2020/12/01 00:18:24

package logit

import (
	"time"
)

const (
	// KB = 1024 bytes.
	KB int64 = 1 << (10 * (iota + 1))

	// MB = 1024 * 1024 bytes.
	MB

	// GB = 1024 * 1024 * 1024 bytes.
	GB

	// minDuration prevents io system from creating file too fast.
	// Default is one second.
	minDuration = 1 * time.Second

	// minLimitedSize prevents io system from creating file too fast.
	// Default is 64 KB (64 * 1024 bytes).
	minLimitedSize = 64 * KB
)

// Checker is an interface of checking if a file writer need to roll.
// File writer will call Check() to know if time to roll before writing.
type Checker interface {

	// Check returns true if a file writer need to roll.
	// Although file is a pointer, you shouldn't modify it in this method.
	// Remember, fw in this method should be read only.
	// n is the length of slice to be written, and we provide it for some purposes.
	Check(fw *FileWriter, n int) bool
}

// ================================= time checker =================================

// TimeChecker is an implement of Checker of time dimension.
type TimeChecker struct {

	// lastTime is the last time of returning true.
	lastTime time.Time

	// duration is the core field of this struct.
	// Everytime currentTime - lastTime >= duration, Check() will return true.
	// This field should be always larger than minDuration for some safe considerations.
	// See minDuration.
	duration time.Duration
}

// NewTimeChecker returns a new TimeChecker with given duration.
// Everytime currentTime - lastTime >= duration, Check() will return true.
// If duration is less than minDuration, then duration will be set to minDuration.
// See minDuration.
func NewTimeChecker(duration time.Duration) *TimeChecker {

	if duration < minDuration {
		duration = minDuration
	}
	return &TimeChecker{
		lastTime: time.Now(),
		duration: duration,
	}
}

// Check checks fw with lastTime and duration.
func (tc *TimeChecker) Check(fw *FileWriter, n int) bool {
	return time.Now().Sub(tc.lastTime) >= tc.duration
}

// ================================= size checker =================================

type SizeChecker struct {

	// limitedSize is the limited size when Check() returns true.
	// Check() will return true if currentSize has reached to limitedSize.
	// This field should be always larger than minLimitedSize for some safe considerations.
	limitedSize int64

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
}

// NewSizeChecker returns a new SizeChecker with given limitedSize.
// Everytime size >= limitedSize, Check() will return true.
// If limitedSize is less than minLimitedSize, then limitedSize will be set to minLimitedSize.
// See minLimitedSize.
func NewSizeChecker(limitedSize int64) *SizeChecker {

	if limitedSize < minLimitedSize {
		limitedSize = minLimitedSize
	}
	return &SizeChecker{
		limitedSize: limitedSize,
		currentSize: 0,
	}
}

// Check checks fw with limitedSize and currentSize.
func (sc *SizeChecker) Check(fw *FileWriter, n int) bool {

	if sc.currentSize >= sc.limitedSize {

		// 1. err != nil: failed to fetch real size, choose to believe currentSize
		// 2. real size >= limitedSize
		fileInfo, err := fw.file.Stat()
		if err != nil || fileInfo.Size() >= sc.limitedSize {
			return true
		}

		// Correct currentSize
		sc.currentSize = fileInfo.Size()
	}
	return false
}

// ================================= count checker =================================

type CountChecker struct {

}
