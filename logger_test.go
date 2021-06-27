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
// Created at 2021/06/27 16:41:20

package logit

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/FishGoddess/logit/appender"
)

// go test -v -bench=^BenchmarkLogitLogger$ -benchtime=3s
func BenchmarkLogitLogger(b *testing.B) {

	logger := NewLogger(DebugLevel, appender.Json(), ioutil.Discard)

	task := func() {
		logger.Debug().Str("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Msg("debug...")
		logger.Info().Str("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Msg("info...")
		logger.Warn().Str("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Msg("warning...")
		logger.Error().Str("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Msg("error...")
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		task()
	}
}

// go test -v -cover -run=^TestNewLogger$
func TestNewLogger(t *testing.T) {

	logger := NewLogger(DebugLevel, appender.Json(), os.Stdout)
	logger.Info().Str("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Msg("info...")
	logger.Error().Str("trace", "xxx").Int("id", 123).Float64("pi", 3.14).MsgF("error...")
	logger.Error().Str("trace", "xxx").Int("id", 123).Float64("pi", 3.14).MsgF("error with %d...", 666)
}
