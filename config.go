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

    // Encoder is how to encode a log to be writable.
    // There are two encoders for you:
    //     1. DefaultEncoder
    //     2. JsonEncoder
    //
    // DefaultEncoder encodes a log to a plain string like "[Info] [2020-03-06 16:10:44] msg" in bytes.
    // JsonEncoder encodes a log to a Json string like `{"level":"debug", "time":"2020-03-22 22:35:00", "msg":"log content..."}` in bytes.
    // Of cause, you can implement Encoder interface to do you encoding job in you own way.
    Encoder Encoder

    Handlers []Handler
}

type FileConfig struct {
    Config

    // Writer is where to write a log.
    Writer io.Writer
}
