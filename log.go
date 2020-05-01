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
// Email: fishinlove@163.com
// Created at 2020/03/25 22:01:26

package logit

import "time"

// Log is representation of a logging message, including all information about this message.
type Log struct {

	// logger is the publisher of this log.
	logger *Logger

	// level is the level of this log.
	level Level

	// now is the publishing time of this log.
	now time.Time

	// file is the file path of this log.
	file string

	// line is the line number in file.
	line int

	// msg is the message of this log.
	msg string
}

// Logger returns the publisher of this log.
func (l *Log) Logger() *Logger {
	return l.logger
}

// Level returns the level of this log.
func (l *Log) Level() Level {
	return l.level
}

// Now returns the publishing time of this log.
func (l *Log) Now() time.Time {
	return l.now
}

// File returns the file path of this log.
func (l *Log) File() string {
	return l.file
}

// Line returns the line number in file of this log.
func (l *Log) Line() int {
	return l.line
}

// Msg returns the message of this log.
func (l *Log) Msg() string {
	return l.msg
}
