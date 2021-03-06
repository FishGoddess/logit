// Copyright 2021 Ye Zi Jie. All Rights Reserved.
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
// Created at 2021/07/02 02:03:45

package lib

import (
	"os"
	"runtime"
)

var (
	pid = os.Getpid()
)

func Pid() int {
	return pid
}

func Caller(depth int) (file string, line int) {

	var ok bool
	_, file, line, ok = runtime.Caller(depth)
	if ok {
		return file, line
	}
	return "unknown file", -1
}
