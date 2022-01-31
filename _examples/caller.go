// Copyright 2022 Ye Zi Jie. All Rights Reserved.
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
// Created at 2022/01/29 20:51:29

package main

import (
	"github.com/FishGoddess/logit"
	"time"
)

func main() {
	// Let's create a logger without caller information.
	logger := logit.NewLogger()
	logger.Info("I am without caller").End()

	// We provide a way to add caller information to log even logger doesn't carry caller.
	logger.Info("Invoke log.WithCaller()").WithCaller().End()
	logger.Close()

	time.Sleep(time.Second)

	// Now, let's create a logger with caller information.
	logger = logit.NewLogger(logit.Options().WithCaller())
	logger.Info("I am with caller").End()

	// We won't carry caller information twice or more if logger carries caller information originally.
	logger.Info("Invoke log.WithCaller() again").WithCaller().End()
	logger.Close()
}
