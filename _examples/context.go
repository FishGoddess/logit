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

package main

import (
	"context"
	"log/slog"

	"github.com/FishGoddess/logit"
)

func resolveUser(ctx context.Context) []slog.Attr {
	userID, ok := ctx.Value("user_id").(int)
	if !ok {
		return nil
	}

	username, ok := ctx.Value("username").(string)
	if !ok {
		return nil
	}

	return []slog.Attr{
		slog.Int("user_id", userID), slog.String("username", username),
	}
}

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
	logger = logit.NewLogger().WithGroup("context")
	ctx = logit.NewContext(ctx, logger)

	// Then you can get the logger from context.
	logger = logit.FromContext(ctx)
	logger.Debug("logger from context debug", "key", "value")

	// Maybe you have noticed logger has some methods with context.
	// These methods will pass the context to underlying handler so that we can use it to process some logics.
	logger.DebugContext(ctx, "debug context")
	logger.InfoContext(ctx, "info context")
	logger.WarnContext(ctx, "warn context")
	logger.ErrorContext(ctx, "error context")

	// You can carry some attributes through context.
	ctx = context.WithValue(ctx, "user_id", 123456)
	ctx = context.WithValue(ctx, "username", "fishgoddess")

	// Then use AttrResolver resolves attributes from context, see WithAttrResolvers.
	logger = logit.NewLogger(logit.WithAttrResolvers(resolveUser))
	logger.InfoContext(ctx, "see what attributes in this log")

	// However, attributes are gone if log without context.
	logger.Info("see what attributes in this log")
}
