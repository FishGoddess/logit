// Copyright 2025 FishGoddess. All Rights Reserved.
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

package main

import (
	"context"

	"github.com/FishGoddess/logit"
)

func main() {
	// We provide a way for getting logger from a context.
	// By default, the default logger will be returned if there is no logit.Logger in context.
	ctx := context.Background()

	logger := logit.FromContext(ctx)
	logger.Debug("logger from context debug")

	if logger == logit.Default() {
		logger.Info("logger from context is default logger")
	}

	// Use NewContext to set a logger to context.
	// We use WithGroup here to make a difference to default logger.
	logger = logit.NewLogger().WithGroup("context").With("user_id", 123456)
	ctx = logit.NewContext(ctx, logger)

	// Then you can get the logger from context.
	logger = logit.FromContext(ctx)
	logger.Debug("logger from context debug", "key", "value")
}
