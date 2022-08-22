// Copyright 2022 FishGoddess. All Rights Reserved.
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

	"github.com/go-logit/logit"
)

func main() {
	// By NewContext, you can bind a context with a logger and get it from context again.
	// So you can use this logger from everywhere as long as you can get this context.
	ctx := logit.NewContext(context.Background(), logit.NewLogger())

	// FromContext returns the logger in context.
	logger := logit.FromContext(ctx)
	logger.Info("This is a message logged by logger from context").Log()

	// Actually, you also have a chance to specify the key of logger in context.
	// It gives you a way to discriminate different businesses in using logger.
	// For example, you can create two loggers for your two different usages and
	// set them to a context with different key, so you can get each logger from context with each key.
	businessOneKey := "businessOne"
	logger = logit.NewLogger(logit.Options().WithMsgKey("businessOneMsg"))
	ctx = logit.NewContextWithKey(context.Background(), businessOneKey, logger)

	businessTwoKey := "businessTwo"
	logger = logit.NewLogger(logit.Options().WithMsgKey("businessTwoMsg"))
	ctx = logit.NewContextWithKey(ctx, businessTwoKey, logger)

	// Get different logger from the same context with different key.
	logger = logit.FromContextWithKey(ctx, businessOneKey)
	logger.Info("This is a message logged by logger from context with businessOneKey").Log()

	logger = logit.FromContextWithKey(ctx, businessTwoKey)
	logger.Info("This is a message logged by logger from context with businessTwoKey").Log()
}
