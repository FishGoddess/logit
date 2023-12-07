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
	"io"
	"log/slog"

	"github.com/FishGoddess/logit"
)

func main() {
	// By default, logit uses text handler to output logs.
	logger := logit.NewLogger()
	logger.Info("default handler is text")

	// You can change it to other handlers by options.
	// For example, use json handler:
	logger = logit.NewLogger(logit.WithJsonHandler())
	logger.Info("using json handler")

	// Or you want to use slog's handlers in Go:
	newHandler := func(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
		return slog.NewTextHandler(w, opts)
	}

	logger = logit.NewLogger(logit.WithHandler(newHandler))
	logger.Info("using slog text handler")

	// As you can see, our handler is slog's handler, so you can use any handlers implement this interface.
	// Like slog's json handler, too:
	newHandler = func(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
		return slog.NewJSONHandler(w, opts)
	}
}
