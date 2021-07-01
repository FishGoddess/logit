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
	"errors"
	"os"
	"testing"

	"github.com/FishGoddess/logit/core/appender"
)

// go test -v -cover -run=^TestNewLogger$
func TestNewLogger(t *testing.T) {

	options := Options()
	logger := NewLogger(
		options.WithDebug(),
		options.WithAppender(appender.Json()),
		options.WithWriter(os.Stdout),
		//options.WithCaller(),
		//options.WithMsgKey("message"),
		//options.WithTimeKey("when"),
		//options.WithLevelKey("level"),
		//options.WithFileKey("file"),
		//options.WithLineKey("line"),
		//options.WithTimeFormat("060102"),
	)

	logger.Debug().String("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Any("any", map[string]interface{}{"a": 1, "b": "bbb"}).Msg("info...")
	logger.Error().Byte("b", 'a').Byte("es", '\n').Runes("words", []rune("我是中国人")).Error("err", errors.New("我是错误")).Msg("error...")
	logger.Error().String("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Msg("error with %d...", 666)
	logger.Warn().Strings("s\tb\nd\b", []string{"abc\r", "efg\n"}).Msg("\"warn\"...\r\b\t\n")
	logger.Info().Bools("bools", []bool{true, false}).Bytes("bytes", []byte{'\b', '\t', 'a', 'b', 'c', '"', '\n'}).Int16s("int16s", []int16{123, 4567, 8901}).Float32s("float32s", []float32{3.14, 6.18}).Msg("warn...")
}
