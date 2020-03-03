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
// Created at 2020/03/03 14:59:08

/*
Package wrapper provides some writers to extend your Logger.

1. DurationRollingFile:

    // DurationRollingFile is a time sensitive file.
    file := NewDurationRollingFile(time.Second, func(lastTime, currentTime time.Time) string {
        return "D:/" + currentTime.Format("20060102-150405") + ".txt"
    })
    defer file.Close()

    // You can use it like using os.File!
    file.Write([]byte("Hello!")

*/
package wrapper // import "github.com/FishGoddess/logit/wrapper"
