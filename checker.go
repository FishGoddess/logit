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
// Created at 2020/12/01 00:18:24

package logit

// Checker is an interface of checking if a file writer need to roll.
// File writer will call Check() to know if time to roll before writing.
type Checker interface {

	// Check returns true if a file writer need to roll.
	// Although file is a pointer, you shouldn't modify it in this method.
	// Remember, fw in this method should be read only.
	// n is the length of slice to be written, and we provide it for some purposes.
	Check(fw *FileWriter, n int) bool
}
