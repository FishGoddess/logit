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
    // 由于每一次写操作都需要获取文件大小，这样会发起一次系统调用，比较耗时
    // 这里将文件大小记录起来，每一次写操作直接使用这个大小进行判断
    // 这在没有外界修改过文件的情况下一般是没有问题的，而且能提高性能
    // 只有当判断到 currentSize >= limitedSize 的时候，才进行一次系统调用查询真实的文件大小
    // 如果真实的文件大小小于 currentSize，重新设置 currentSize，此时不划分文件
    // 如果真实的文件大小大于 currentSize，此时需要划分文件，并设置 currentSize 为 0
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
