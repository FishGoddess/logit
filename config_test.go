// Copyright 2021 Ye Zi Jie. All Rights Reserved.
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
// Created at 2021/07/11 14:02:38

package logit

import "testing"

// go test -v -cover -run=^TestNewDefaultConfig$
func TestNewDefaultConfig(t *testing.T) {
	defaultConfig := newDefaultConfig()

	cfg := &config{
		level:      debugLevel,
		needPid:    false,
		needCaller: false,
		msgKey:     "log.msg",
		timeKey:    "log.time",
		levelKey:   "log.level",
		pidKey:     "log.pid",
		fileKey:    "log.file",
		lineKey:    "log.line",
		timeFormat: "2006-01-02 15:04:05",
	}

	if *defaultConfig != *cfg {
		t.Errorf("default config %+v is wrong", defaultConfig)
	}
}
