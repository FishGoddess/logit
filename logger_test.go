// Copyright 2023 FishGoddess. All Rights Reserved.
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

package logit

import "testing"

func TestLogger(t *testing.T) {
	logger := NewDevelopLogger()
	defer logger.Close()

	logger.Debug("new version of logit", "version", "1.5.0-alpha", "date", 20231122)
	logger.Error("new version of logit", "version", "1.5.0-alpha", "date", 20231122)
}
