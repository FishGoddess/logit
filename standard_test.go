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

import (
	"log/slog"
	"testing"
	"time"
)

// go test -v -cover -count=1 -test.cpu=1 -run=^TestStandardHandler$
func TestStandardHandler(t *testing.T) {
	logger1 := NewLogger(WithHandler(handlerStandard)).With("id", 123456)
	logger1.Info("using console handler 1", slog.Group("log_group1", "k1", 666))

	logger2 := logger1.WithGroup("group2").With("name", "fishgoddess")
	logger2.Info("using console handler 2", slog.Group("log_group2", "k2", 888), "t", time.Now())

	logger1.Info("using console handler 1", slog.Group("log_group1", "k1", 666))
}
