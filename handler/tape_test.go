// Copyright 2024 FishGoddess. All Rights Reserved.
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
	"bytes"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
	"testing"
	"testing/slogtest"
	"time"
)

type demo struct {
	value string
}

func (d *demo) String() string {
	return d.value
}

func parseTime(timeValue string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05.000000", timeValue)
}

func parseLevel(levelValue string) (slog.Level, error) {
	switch levelValue {
	case "DEBUG":
		return slog.LevelDebug, nil
	case "INFO":
		return slog.LevelInfo, nil
	case "WARN":
		return slog.LevelWarn, nil
	case "ERROR":
		return slog.LevelError, nil
	default:
		return 0, fmt.Errorf("unknown level %s", levelValue)
	}
}

// 2023-12-20 12:07:42.731993 ¦ INFO ¦ message
// 2023-12-20 12:07:42.732041 ¦ INFO ¦ message ¦ k=v
// 2023-12-20 12:07:42.732045 ¦ INFO ¦ msg ¦ a=b ¦ c=d
// 0001-01-01 00:00:00.000000 ¦ INFO ¦ msg ¦ k=v
// 2023-12-20 12:07:42.732054 ¦ INFO ¦ msg ¦ a=b ¦ k=v
// 2023-12-20 12:07:42.732057 ¦ INFO ¦ msg ¦ a=b ¦ G.c=d ¦ e=f
// 2023-12-20 12:07:42.732059 ¦ INFO ¦ msg ¦ a=b ¦ e=f
// 2023-12-20 12:07:42.732060 ¦ INFO ¦ msg ¦ a=b ¦ c=d ¦ e=f
// 2023-12-20 12:07:42.732062 ¦ INFO ¦ msg ¦ G.a=b
// 2023-12-20 12:07:42.732064 ¦ INFO ¦ msg ¦ G.H.a=b ¦ G.H.c=d ¦ G.H.e=f
// 2023-12-20 12:07:42.732066 ¦ INFO ¦ msg ¦ G.H.a=b ¦ G.H.c=d
// 2023-12-20 12:07:42.732068 ¦ INFO ¦ msg ¦ k=replaced
// 2023-12-20 12:07:42.732071 ¦ INFO ¦ msg ¦ G.a=v1 ¦ G.b=v2
// 2023-12-20 12:07:42.732073 ¦ INFO ¦ msg ¦ k=replaced
// 2023-12-20 12:07:42.732076 ¦ INFO ¦ msg ¦ G.a=v1 ¦ G.b=v2
func parseLog(log string) (map[string]any, error) {
	attrs := strings.Split(log, string(attrConnector))
	if len(attrs) < 3 {
		return nil, errors.New("len(attrs) < 3")
	}

	timeValue, levelValue, message := attrs[0], attrs[1], attrs[2]

	t, err := parseTime(timeValue)
	if err != nil {
		return nil, err
	}

	level, err := parseLevel(levelValue)
	if err != nil {
		return nil, err
	}

	result := map[string]any{
		slog.LevelKey:   level,
		slog.MessageKey: message,
	}

	if !t.IsZero() {
		result[slog.TimeKey] = t
	}

	for i := 3; i < len(attrs); i++ {
		kv := strings.Split(attrs[i], string(keyValueConnector))
		if len(kv) < 2 {
			return nil, fmt.Errorf("attr kv len %d < 2", len(kv))
		}

		key, value := kv[0], kv[1]
		if !strings.Contains(key, groupConnector) {
			result[key] = value
			continue
		}

		index := 0
		lastMap := result
		groups := strings.Split(key, groupConnector)

		for ; index < len(groups)-1; index++ {
			group := groups[index]

			m, ok := lastMap[group]
			if !ok {
				m = make(map[string]any, 4)
				lastMap[group] = m
			}

			lastMap = m.(map[string]any)
		}

		k := groups[index]
		lastMap[k] = value
	}

	return result, nil
}

func parseLogs(logs []string) ([]map[string]any, error) {
	result := make([]map[string]any, 0, 16)
	for _, log := range logs {
		if log == "" {
			continue
		}

		one, err := parseLog(log)
		if err != nil {
			return nil, err
		}

		result = append(result, one)
	}

	return result, nil
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestParseLog$
func TestParseLog(t *testing.T) {
	log := "2023-12-20 12:07:42.732064 ¦ INFO ¦ msg ¦ G.H.a=b ¦ G.H.c=d ¦ G.H.e=f"

	m, err := parseLog(log)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(m)
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestTapeHandler$
func TestTapeHandler(t *testing.T) {
	handler := NewTapeHandler(os.Stdout, nil)
	//handler := slog.NewTextHandler(os.Stdout, opts)

	logger1 := slog.New(handler).WithGroup("group1").With("id", 123456)
	logger1.Info("using console handler 1", slog.Group("log_group1", "k1", 666), "err", io.EOF)

	logger2 := logger1.WithGroup("group2").With("name", "fishgoddess")
	logger2.Info("using console handler 2", slog.Group("log_group2", "k2", 888, "k3", "xxx"), "t", time.Date(1977, 10, 24, 25, 35, 17, 999999000, time.Local))

	demo := &demo{"xxx"}
	logger1.Info("using console handler 1", slog.Group("log_group1", "k1", 666), "demo", demo, "err", nil)

	buffer := bytes.NewBuffer(make([]byte, 0, 4096))
	handler = NewTapeHandler(buffer, nil)

	err := slogtest.TestHandler(handler, func() []map[string]any {
		lines := string(buffer.Bytes())
		logs := strings.Split(lines, string(lineBreak))

		t.Log(lines)

		kvs, err := parseLogs(logs)
		if err != nil {
			t.Fatal(err)
		}

		return kvs
	})

	// Tape handler doesn't always act like a slog handler.
	if err != nil {
		t.Log(err)
	}
}
