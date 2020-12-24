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
// Created at 2020/03/25 22:01:26

package logit

import "time"

// caller stores some calling information.
type caller struct {

	// File is the file path of this log.
	File string

	// Line is the line number in file.
	Line int
}

// Log is representation of a logging message, including all information about this message.
type Log struct {

	// msg is the message of this log.
	msg string

	// level is the level of this log.
	level Level

	// time is the publishing time of this log.
	time time.Time

	// caller stores some calling information, such as file path and line number.
	caller *caller
}

// Msg returns the message of this log.
func (l *Log) Msg() string {
	return l.msg
}

// Level returns the level of this log.
func (l *Log) Level() Level {
	return l.level
}

// Time returns the publishing time of this log.
func (l *Log) Time() time.Time {
	return l.time
}

// Caller returns the caller information of this log.
// Notice that ok will be false if this log doesn't have caller information.
func (l *Log) Caller() (caller *caller, ok bool) {
	return l.caller, l.caller != nil
}
