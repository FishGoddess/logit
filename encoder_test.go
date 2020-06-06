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
// Created at 2020/04/24 16:43:49

package logit

import (
	"testing"
	"time"
)

// 测试获取编码器
func TestEncoderOf(t *testing.T) {

	// 这个不存在的编码器退出程序
	//encoder := encoderOf("fake-encoder")
	//if encoder != nil {
	//	t.Fatal("encoderOf 中 encoder 应该为 nil")
	//}

	log := &Log{
		level: DebugLevel,
		now:   time.Now(),
		msg:   "xxx",
	}

	// 判断获取的编码器是否正确
	if string(encoderOf("text").Encode(log, DefaultTimeFormat)) != string(TextEncoder().Encode(log, DefaultTimeFormat)) {
		t.Fatal("encoderOf(\"text\") 出现问题！")
	}
	if string(encoderOf("json").Encode(log, "")) != string(JsonEncoder().Encode(log, "")) {
		t.Fatal("encoderOf(\"json\") 出现问题！")
	}
}
