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
// Created at 2020/04/24 16:43:49

package logit

import (
	"testing"
	"time"
)

// 测试获取编码器
func TestEncoderOf(t *testing.T) {

	defer func() {
		err := recover()
		if err == nil {
			t.Fatal("获取编码器测试出现问题！")
		}
	}()

	// 这个不存在的编码器会引起 panic，然后被上面的 recover 捕获
	encoder := EncoderOf("nil")
	if encoder != nil {
		t.Fatal("EncoderOf 中 encoder 应该为 nil")
	}

	log := &Log{
		level: DebugLevel,
		now:   time.Now(),
		msg:   "xxx",
	}

	// 判断获取的编码器是否正确
	if string(EncoderOf("text").Encode(log, DefaultTimeFormat)) != string(EncodeToText(log, DefaultTimeFormat)) {
		t.Fatal("EncoderOf(\"text\") 出现问题！")
	}
	if string(EncoderOf("text").Encode(log, "")) != string(EncodeToJson(log, "")) {
		t.Fatal("EncoderOf(\"json\") 出现问题！")
	}
}
