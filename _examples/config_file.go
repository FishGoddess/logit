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
// Created at 2020/03/30 14:21:28

package main

import (
	"fmt"

	"github.com/FishGoddess/logit"
)

func main() {

	// Create a logger from config file.
	//
	// logger.conf:
	//
	//     "level": "debug",
	//
	//     "caller": false,
	//
	//     "handlers": {
	//         "console": {
	//             "timeFormat":"unix",
	//             "encoder":"json"
	//         },
	//         "file":{
	//             "path":"D:/logit.log"
	//         }
	//     }
	//
	logger := logit.NewLoggerFromPath("./logger.conf")
	logger.Info("I am working!")
	logger.Info("My level is " + logger.Level().String())
	fmt.Println("fmt ==============================================")

	handlers := logger.Handlers()
	for i, handler := range handlers {
		logger.Info(fmt.Sprintf("No.%d hadler ==> %T", i+1, handler))
	}
}
