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
// Created at 2020/03/25 23:06:59

package logit

import "io"

// Config is for configuring a Logger.
// You can use a config to create a Logger to use.
// See NewLoggerFrom(config Config).
type Config struct {

    // Level is the level of Logger.
    // If the level of log is smaller than this Level, this log will be ignored.
    Level Level

    Handlers []Handler
}

// FileConfig is the config mapping a file.
type FileConfig struct {

    // Config means that a file config is also a config.
    Config

    // Writer is where to write a log.
    Writer io.Writer
}
