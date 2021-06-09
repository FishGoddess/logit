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
// Created at 2020/11/27 23:41:13

package main

import (
	"fmt"

	"github.com/FishGoddess/logit"
)

type MyEncoder struct {
	name string
}

func (me *MyEncoder) Encode(log *logit.Log) []byte {
	return []byte(fmt.Sprintf("%s: %s ==> %+v\n", me.name, log.Msg(), log.Values()))
}

func main() {

	// Use default encoder
	logger := logit.NewLogger(logit.KV{"trace": 123})
	logger.Info("Default encoder is like this...")

	// We provide some encoders, such as text and json
	// Try TextEncoder and JsonEncoder
	logger.Encoders().SetEncoder(logit.NewTextEncoder("2006-01-02 15:04:05"))
	logger.Encoders().SetEncoder(logit.NewJsonEncoder("2006-01-02 15:04:05"))

	// In fact, encoder is an interface like "func(log *logit.Log) []byte"
	// So you can implement your own encoder as you want
	// All information of log is stored in log
	// No matter what you do, return a byte slice
	// The returned slice will be written by logger
	logger.Encoders().SetEncoder(&MyEncoder{name: "MyEncoder"})
	logger.Info("see what I got!", logit.KV{
		"id": "xxx",
	})

	// You can set encoder of each level, for example:
	logger.Encoders().SetWarnEncoder(logit.NewJsonEncoder("2006/01/02 15:04:05"))
	logger.Encoders().SetErrorEncoder(logit.NewJsonEncoder(logit.TimeFormat))
	logger.Warn("warn...")
	logger.Error("error...")
}
