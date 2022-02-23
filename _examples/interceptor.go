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

// serverInterceptor is the global interceptor applied to all logs.
func serverInterceptor(ctx context.Context, log *logit.Log) {
	log.String("server", "logit.interceptor")
}

// traceInterceptor is the global interceptor applied to all logs.
func traceInterceptor(ctx context.Context, log *logit.Log) {
	trace, ok := ctx.Value("trace").(string)
	if !ok {
		trace = "unknown trace"
	}

	log.String("trace", trace)
}

// userInterceptor is the global interceptor applied to all logs.
func userInterceptor(ctx context.Context, log *logit.Log) {
	user, ok := ctx.Value("user").(string)
	if !ok {
		user = "unknown user"
	}

	log.String("user", user)
}

// businessInterceptor is the log-level interceptor applied to one/some logs.
func businessInterceptor(ctx context.Context, log *logit.Log) {
	business, ok := ctx.Value("business").(string)
	if !ok {
		business = "unknown business"
	}

	log.String("business", business)
}

func main() {
	// Use logit.Options().WithInterceptors to append some interceptors.
	logger := logit.NewLogger(logit.Options().WithInterceptors(serverInterceptor, traceInterceptor, userInterceptor))
	defer logger.Close()

	// By default, context passed to interceptor is context.Background().
	logger.Info("try interceptor - round one").End()

	// You can use WithContext to change context passed to interceptor.
	ctx := context.WithValue(context.Background(), "trace", "666")
	ctx = context.WithValue(ctx, "user", "FishGoddess")
	logger.Info("try interceptor - round two").WithContext(ctx).End()

	// The interceptors appended to logger will apply to all logs.
	// You can use Intercept to intercept one log rather than all logs.
	logger.Info("try interceptor - round three").WithContext(ctx).Intercept(businessInterceptor).End()

	// Notice that WithContext should be called before Intercept if you want to pass this context to Intercept.
	ctx = context.WithValue(ctx, "business", "logger")
	logger.Info("try interceptor - round four").WithContext(ctx).Intercept(businessInterceptor).End()
}
