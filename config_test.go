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

// go test -v -cover -count=1 -test.cpu=1 -run=^TestConfigNewHandlerOptions$
func TestConfigNewHandlerOptions(t *testing.T) {
	replaceAttr := func(groups []string, attr slog.Attr) slog.Attr { return attr }

	conf := &config{
		level:       slog.LevelWarn,
		withSource:  true,
		replaceAttr: replaceAttr,
	}

	opts := conf.newHandlerOptions()

	if opts.Level != conf.level {
		t.Fatalf("opts.Level %v != conf.level %v", opts.Level, conf.level)
	}

	if opts.AddSource != conf.withSource {
		t.Fatalf("opts.AddSource %v != conf.withSource %v", opts.AddSource, conf.withSource)
	}

	if fmt.Sprintf("%p", opts.ReplaceAttr) != fmt.Sprintf("%p", conf.replaceAttr) {
		t.Fatalf("opts.ReplaceAttr %p != conf.replaceAttr %p", opts.ReplaceAttr, conf.replaceAttr)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestConfigNewHandler$
func TestConfigNewHandler(t *testing.T) {
	RegisterHandler(t.Name(), func(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
		return &testConfigHandler{
			w:    w,
			opts: *opts,
		}
	})

	newWriter := func() (io.Writer, error) { return os.Stderr, nil }
	replaceAttr := func(groups []string, attr slog.Attr) slog.Attr { return attr }

	conf := &config{
		level:       slog.LevelWarn,
		handler:     t.Name(),
		newWriter:   newWriter,
		replaceAttr: replaceAttr,
		withSource:  true,
	}

	handler, syncer, closer, err := conf.newHandler()
	if err != nil {
		t.Fatal(err)
	}

	if syncer == nil {
		t.Fatal("syncer is nil")
	}

	if closer == nil {
		t.Fatal("closer is nil")
	}

	tcHandler, ok := handler.(*testConfigHandler)
	if !ok {
		t.Fatalf("handler type %T is wrong", handler)
	}

	if tcHandler.w != os.Stderr {
		t.Fatalf("tcHandler.w %p != os.Stderr %p", tcHandler.w, os.Stderr)
	}

	if tcHandler.opts.Level != conf.level {
		t.Fatalf("tcHandler.opts.Level %v != conf.level %v", tcHandler.opts.Level, conf.level)
	}

	if tcHandler.opts.AddSource != conf.withSource {
		t.Fatalf("tcHandler.opts.AddSource %v != conf.withSource %v", tcHandler.opts.AddSource, conf.withSource)
	}

	if fmt.Sprintf("%p", tcHandler.opts.ReplaceAttr) != fmt.Sprintf("%p", conf.replaceAttr) {
		t.Fatalf("tcHandler.opts.ReplaceAttr %p != conf.replaceAttr %p", tcHandler.opts.ReplaceAttr, conf.replaceAttr)
	}
}
