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
// Created at 2021/06/27 16:38:31

package logit

import (
	"math"
)

const (
	debugLevel level = iota
	infoLevel
	warnLevel
	errorLevel

	// offLevel is for turning off the logger.
	offLevel = math.MaxUint8
)

var (
	// levels store all names of levels provided.
	levels = map[level]string{
		debugLevel: "debug",
		infoLevel:  "info",
		warnLevel:  "warn",
		errorLevel: "error",
		offLevel:   "off",
	}
)

// level is position of one log.
type level uint8

// String returns the name of ll.
// This method will be called when using printing operations like fmt.Println.
func (ll level) String() string {
	return levels[ll]
}
