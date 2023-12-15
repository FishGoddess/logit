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

package handler

import (
	"io"
	"log/slog"
	"os"
	"testing"
	"time"
)

type demo struct {
	value string
}

func (d *demo) String() string {
	return d.value
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestMixHandler$
func TestMixHandler(t *testing.T) {
	opts := &slog.HandlerOptions{}

	//handler := NewMixHandler(os.Stdout, opts)
	handler := slog.NewTextHandler(os.Stdout, opts)

	logger1 := slog.New(handler).WithGroup("group1").With("id", 123456)
	logger1.Info("using console handler 1", slog.Group("log_group1", "k1", 666), "err", io.EOF)

	logger2 := logger1.WithGroup("group2").With("name", "fishgoddess")
	logger2.Info("using console handler 2", slog.Group("log_group2", "k2", 888, "k3", "xxx"), "t", time.Date(1977, 10, 24, 25, 35, 17, 222000000, time.Local))

	demo := &demo{"xxx"}
	logger1.Info("using console handler 1", slog.Group("log_group1", "k1", 666), "demo", demo, "err", nil)
}
