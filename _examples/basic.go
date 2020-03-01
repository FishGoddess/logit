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
// Created at 2020/02/29 21:59:13
package main

import (
    "os"

    "github.com/FishGoddess/logit"
)

func main() {

    // NewLogger creates a new Logger holder.
    // The first parameter "os.Stdout" is a writer for logging.
    // The second parameter "logit.DebugLevel" is the level of this Logger
    logger := logit.NewLogger(os.Stdout, logit.DebugLevel)

    // Then you will be easy to log!
    logger.Debug("this is a debug message!")
    logger.Info("this is a info message!")
    logger.Warning("this is a warning message!")
    logger.Error("this is a error message!")
}
