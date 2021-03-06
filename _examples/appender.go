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
// Created at 2021/07/11 23:32:35

package main

import (
	"github.com/FishGoddess/logit"
	"github.com/FishGoddess/logit/core/appender"
)

func main() {

	// We provide some ways to change the form of logs
	// Actually, appender is an interface with some common methods, see appender.Appender
	appender.Text()
	appender.Json()

	// Set appender to the one you want to use when creating a logger
	// Default appender is appender.Text()
	logger := logit.NewLogger()
	logger.Info("appender.Text()").End()

	// You can switch appender to the other one, such appender.Json()
	logger = logit.NewLogger(logit.Options().WithAppender(appender.Json()))
	logger.Info("appender.Json()").End()

	// Appender is an interface so you can implement your own appender
	// However, we don't recommend you to do that
	// This interface may change in every version, so you will pay lots of extra attention to it
	// So you should implement it only if you really need to do
}
