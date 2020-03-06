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
    "github.com/FishGoddess/logit"
)

func main() {

    // Log messages with four levels.
    // Notice that the default level is info, so first line of debug message
    // will not be logged! If you want to change level, see logit.ChangeLevelTo
    logit.Debug("I am a debug message! But I will not be logged in default level!")
    logit.Info("I am an info message!")
    logit.Warn("I am a warn message!")
    logit.Error("I am an error message!")

    // Also, you can create a new independent Logger to use. See logit.NewLogger.

    // If you want format your message, just add arguments!
    logit.Info("format info message! id = %d, content = %s", 1, "info!")

    // If you want to output log with file info, try this:
    logit.EnableFileInfo()
    logit.Info("Show file info!")
}
