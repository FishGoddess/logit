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
// Created at 2021/06/13 15:51:29

package logit

type Options func(logger *Logger)

func (o Options) Apply(logger *Logger) {
	o(logger)
}

func WithLevel(level Level) Options {
	return func(logger *Logger) {
		logger.SetLevel(level)
	}
}

func WithCaller() Options {
	return func(logger *Logger) {
		logger.SetNeedCaller(true)
	}
}

func WithKVs(kvs M) Options {
	return func(logger *Logger) {
		logger.kvs = newM(kvs)
	}
}
