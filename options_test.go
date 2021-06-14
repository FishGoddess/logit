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
// Created at 2021/06/13 16:10:29

package logit

import "testing"

// go test -v -cover -run=^TestOptionsApply$
func TestOptionsApply(t *testing.T) {

	logger := &Logger{}
	opt := Options(func(logger *Logger) {
		logger.kvs = M{"origin": 666}
	})

	logger.kvs = nil
	opt(logger)
	if logger.kvs == nil || logger.kvs["origin"] != 666 {
		t.Fatalf("kvs %+v after opt is wrong", logger.kvs)
	}

	logger.kvs = nil
	opt.Apply(logger)
	if logger.kvs == nil || logger.kvs["origin"] != 666 {
		t.Fatalf("kvs %+v after applying is wrong", logger.kvs)
	}
}
