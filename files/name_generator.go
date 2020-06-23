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
// Created at 2020/06/11 21:15:37

package files

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

var (
	// nameGenerators stores all nameGenerators registered.
	// mutexOfNameGenerators is for concurrency.
	nameGenerators = map[string]NameGenerator{
		"default": DefaultNameGenerator(),
	}
	mutexOfNameGenerators = &sync.RWMutex{}

	// NameGeneratorIsExistedError is an error happening on repeating nameGenerator name.
	NameGeneratorIsExistedError = errors.New("the name of nameGenerator you want to register already exists! May be you should give it an another name")
)

// NameGenerator is the type for generating the name of every created file.
// You can customize your format of filename by implementing this function.
// The two parameters string and time.Time is useful. The string parameter is the directory
// stores all files created in this time and the time.Time parameter is the time of current moment.
type NameGenerator func(string, time.Time) string

// NextName is for code-readable.
// Directory stores all files created in this time and now is the time of current moment.
func (ng NameGenerator) NextName(directory string, now time.Time) string {
	return ng(directory, now)
}

// RegisterNameGenerator registers your nameGenerator to logit so that you can use them in config file.
// Return an error if the name is existed, and you should change another name for your nameGenerator.
func RegisterNameGenerator(name string, nameGenerator NameGenerator) error {
	mutexOfNameGenerators.Lock()
	defer mutexOfNameGenerators.Unlock()

	if _, ok := nameGenerators[name]; ok {
		return NameGeneratorIsExistedError
	}
	nameGenerators[name] = nameGenerator
	return nil
}

// NameGeneratorOf returns nameGenerator whose name is given name.
// Notice that we use tips+exit mechanism to check the name.
// This is a more convenient way to use nameGenerator (we think).
// so if the nameGenerator doesn't exist, a tip will be printed and
// the program will exit with status code 10.
func NameGeneratorOf(name string) NameGenerator {
	mutexOfNameGenerators.RLock()
	defer mutexOfNameGenerators.RUnlock()
	nameGenerator, ok := nameGenerators[name]
	if !ok {
		fmt.Fprintf(os.Stderr, "Error: The nameGenerator \"%s\" doesn't exist! Please change it to another nameGenerator.\n", name)
		os.Exit(10)
	}
	return nameGenerator
}

// ================================== default name generator ==================================

// DefaultNameGenerator returns a name generator that creates a time-relative filename
// with given now time. Also, it uses random number to ensure this filename is available.
// The filename will be like "20200304-145246-45.log".
// Notice that directory stores all files created in this time and now is the time of current moment.
func DefaultNameGenerator() func(directory string, now time.Time) string {
	// 考虑使用原子计数器替换随机数
	// 这个 Seed 方法最好不要并发执行，但是这个方法有可能会被并发执行，这是个隐患
	// 在测试阶段就已经出现了随机数重复的情况，导致一个文件被写入多个文件的内容
	// issue: https://github.com/FishGoddess/logit/issues/7
	rand.Seed(time.Now().UnixNano())
	return func(directory string, now time.Time) string {
		name := now.Format("20060102-150405") + "-" + strconv.Itoa(rand.Intn(1000000)) + SuffixOfLogFile
		return filepath.Join(directory, name)
	}
}
