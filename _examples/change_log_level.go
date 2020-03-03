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
// Created at 2020/03/01 15:10:19

package main

import "github.com/FishGoddess/logit"

func main() {

    logit.Debug("Default logger level is info, so debug message will not be logged!")

    // Change logger level to debug level.
    logit.ChangeLevelTo(logit.DebugLevel)

    logit.Debug("Now debug message will be logged!")
}