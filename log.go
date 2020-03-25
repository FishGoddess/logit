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
// Created at 2020/03/25 22:01:26

package logit

import "time"

type Log struct {
    logger *Logger

    level Level

    now time.Time

    msg string

    extra map[string]interface{}
}

func NewLog(logger *Logger, level Level, now time.Time, msg string) *Log {
    return &Log{
        logger: logger,
        level:  level,
        now:    now,
        msg:    msg,
    }
}

func (l *Log) Logger() *Logger {
    return l.logger
}

func (l *Log) Level() Level {
    return l.level
}

func (l *Log) Now() time.Time {
    return l.now
}

func (l *Log) Msg() string {
    return l.msg
}

func (l *Log) Extra() map[string]interface{} {
    return l.extra
}
