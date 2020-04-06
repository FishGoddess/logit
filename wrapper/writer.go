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
// Created at 2020/03/03 15:10:45

package wrapper

import (
    "errors"
    "io"
    "os"
    "sync"
    "time"
)

// These are units representation of file size.
//
// KB = 1024 bytes.
// MB = 1024 * 1024 bytes.
// GB = 1024 * 1024 * 1024 bytes.
const (
    KB int64 = 1 << (10 * (iota + 1))
    MB
    GB
)

var (
    // writers stores all registered writers.
    // mutexOfWriters is for concurrency.
    writers        = map[string]func(params map[string]string) io.Writer{}
    mutexOfWriters = &sync.RWMutex{}

    // WriterIsExistedError is an error happens on repeating writer name.
    WriterIsExistedError = errors.New("the name of writer you want to register already exists! May be you should give it an another name")
)

// RegisterWriter registers your writer to logit so that you can use them easily.
// Return an error if the name is existed, and you should change another name for your writer.
// Notice that newWriter has a parameter called params, which will inject into newWriter
// by logit automatically. Different writer may have different params, so what params should
// inject into newWriter is dependent to specific writer.
func RegisterWriter(name string, newWriter func(params map[string]string) io.Writer) error {
    mutexOfWriters.Lock()
    defer mutexOfWriters.Unlock()
    if _, ok := writers[name]; ok {
        return WriterIsExistedError
    }
    writers[name] = newWriter
    return nil
}

// WriterOf returns writer whose name is given name and params.
// Different writer may have different params, so what params should
// inject into newWriter is dependent to specific writer.
// Notice that we don't use an error mechanism or ok mechanism to check the name but
// a default writer returning mechanism. This is a more convenient way to use writers (we think).
func WriterOf(name string, params map[string]string) io.Writer {
    mutexOfWriters.RLock()
    defer mutexOfWriters.RUnlock()
    newWriter, ok := writers[name]
    if !ok {
        return os.Stdout
    }
    return newWriter(params)
}

// NewFile creates a new file with given filePath.
// Return a new File or an error if failed.
// Notice that the permission of new file is 0644, which means rw-rw-r-- in unix-like os.
func NewFile(filePath string) (*os.File, error) {
    return os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
}

// generateFirstFile creates the first file with given nextFilename function.
func generateFirstFile(nextFilename func(now time.Time) string) (*os.File, time.Time) {
    now := time.Now()
    file, err := NewFile(nextFilename(now))
    if err != nil {
        panic(err)
    }
    return file, now
}
