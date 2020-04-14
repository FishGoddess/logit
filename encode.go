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
// Created at 2020/04/14 21:06:56

package logit

import (
    "bytes"
    "strconv"
    "strings"
)

// EncodeToText encodes a log to a plain string like "[Info] [2020-03-06 16:10:44] msg" in bytes.
func EncodeToText(log *Log, timeFormat string) []byte {

    // 组装 log
    buffer := bytes.NewBuffer(make([]byte, 0, 64))
    buffer.WriteString("[")
    buffer.WriteString(log.Level().String())
    buffer.WriteString("] [")
    buffer.WriteString(log.Now().Format(timeFormat))
    buffer.WriteString("] ")

    // 如果有文件信息，就把文件信息也加进去
    if log.file != "" && log.Line() != 0 {
        buffer.WriteString("[")
        buffer.WriteString(log.File() + ":" + strconv.Itoa(log.Line()))
        buffer.WriteString("] ")
    }

    buffer.WriteString(log.Msg())
    buffer.WriteString("\n")
    return buffer.Bytes()
}

// EncodeToJson encodes a log to a Json string like `{"level":"debug", "time":"2020-03-22 22:35:00", "msg":"log content..."}` in bytes.
// If timeFormat == "", then it will not format time and keep time in unix form.
func EncodeToJson(log *Log, timeFormat string) []byte {

    // 组装 log
    buffer := bytes.NewBuffer(make([]byte, 0, 64))
    buffer.WriteString(`{"level":"`)
    buffer.WriteString(log.Level().String())
    buffer.WriteString(`", "time":`)

    // 判断是否需要格式化时间
    if timeFormat != "" {
        buffer.WriteString(strconv.Quote(log.Now().Format(timeFormat)))
    } else {
        buffer.WriteString(strconv.FormatInt(log.Now().Unix(), 10))
    }

    // 如果有文件信息，就把文件信息也加进去
    if log.file != "" && log.Line() != 0 {
        buffer.WriteString(`, "file":"` + log.File())
        buffer.WriteString(`", "line":` + strconv.Itoa(log.Line()))
    }

    buffer.WriteString(`, "msg":"`)
    buffer.WriteString(escapeString(log.Msg()))
    buffer.WriteString("\"}\n")
    return buffer.Bytes()
}

// escapeString is for escaping string from special character, such as double quotes.
// See issue: https://github.com/FishGoddess/logit/issues/1
func escapeString(s string) string {

    builder := strings.Builder{}
    runes := []rune(s)
    for _, r := range runes {

        // Json 中需要进行转义的字符主要是 \ 和 "，还有控制字符（\u0020 以内的 ascii 字符）
        switch r {
        case '"', '\\':
            builder.WriteRune('\\')
            builder.WriteRune(r)
        default:
            // ascii 小于 16 需要在前面补 \u000，介于 16 和 32 之间的需要补 \u00
            if r < 16 {
                builder.WriteString("\\u000" + strconv.FormatInt(int64(r), 16))
            } else if r < 32 {
                builder.WriteString("\\u00" + strconv.FormatInt(int64(r), 16))
            } else {
                builder.WriteRune(r)
            }
        }
    }

    return builder.String()
}
