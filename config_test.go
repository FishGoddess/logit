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
	"fmt"
	"io"
	"log/slog"
	"os"
	"testing"
)

type testConfigHandler struct {
	slog.TextHandler
	w    io.Writer
	opts slog.HandlerOptions
}

// go test -v -cover -run=^TestConfigHandlerOptions$
func TestConfigHandlerOptions(t *testing.T) {
	replaceAttr := func(groups []string, attr slog.Attr) slog.Attr { return attr }

	conf := &config{
		level:       LevelWarn,
		withSource:  true,
		replaceAttr: replaceAttr,
	}

	opts := conf.handlerOptions()

	if opts.Level != conf.level.Peel() {
		t.Errorf("opts.Level %v != conf.level %v", opts.Level, conf.level.Peel())
	}

	if opts.AddSource != conf.withSource {
		t.Errorf("opts.AddSource %v != conf.withSource %v", opts.AddSource, conf.withSource)
	}

	if fmt.Sprintf("%p", opts.ReplaceAttr) != fmt.Sprintf("%p", conf.replaceAttr) {
		t.Errorf("opts.ReplaceAttr %p != conf.replaceAttr %p", opts.ReplaceAttr, conf.replaceAttr)
	}
}

// go test -v -cover -run=^TestConfigHandler$
func TestConfigHandler(t *testing.T) {
	replaceAttr := func(groups []string, attr slog.Attr) slog.Attr { return attr }

	conf := &config{
		level:       LevelWarn,
		withSource:  true,
		replaceAttr: replaceAttr,

		newWriter: func() (io.Writer, error) {
			return os.Stderr, nil
		},
		newHandler: func(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
			return &testConfigHandler{
				w:    w,
				opts: *opts,
			}
		},
	}

	handler, err := conf.handler()
	if err != nil {
		t.Error(err)
	}

	tcHandler, ok := handler.(*testConfigHandler)
	if !ok {
		t.Errorf("handler type %T is wrong", handler)
	}

	if tcHandler.w != os.Stderr {
		t.Errorf("tcHandler.w %p != os.Stderr %p", tcHandler.w, os.Stderr)
	}

	if tcHandler.opts.Level != conf.level.Peel() {
		t.Errorf("tcHandler.opts.Level %v != conf.level %v", tcHandler.opts.Level, conf.level.Peel())
	}

	if tcHandler.opts.AddSource != conf.withSource {
		t.Errorf("tcHandler.opts.AddSource %v != conf.withSource %v", tcHandler.opts.AddSource, conf.withSource)
	}

	if fmt.Sprintf("%p", tcHandler.opts.ReplaceAttr) != fmt.Sprintf("%p", conf.replaceAttr) {
		t.Errorf("tcHandler.opts.ReplaceAttr %p != conf.replaceAttr %p", tcHandler.opts.ReplaceAttr, conf.replaceAttr)
	}
}
