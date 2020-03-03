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
// Created at 2020/03/03 23:39:39

package main

import "github.com/FishGoddess/logit"

func main() {

    // NewFileLogger creates a new logger which logs to file.
    // It just need a file path like "D:/test.log" and a logger level.
    logger := logit.NewFileLogger("D:/test.log", logit.DebugLevel)
    logger.Info("我是 info 日志！")
}
