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

package fastclock

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

// go test -v -run=^$ -bench=^BenchmarkTimeNow$ -benchtime=1s
func BenchmarkTimeNow(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		time.Now()
	}
}

// go test -v -run=^$ -bench=^BenchmarkFastClockNow$ -benchtime=1s
func BenchmarkFastClockNow(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Now()
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestNow$
func TestNow(t *testing.T) {
	duration := 100 * time.Millisecond

	for i := 0; i < 100; i++ {
		got := Now()
		gap := time.Since(got)
		t.Logf("got: %v, gap: %v", got, gap)

		if math.Abs(float64(gap.Nanoseconds())) > float64(duration)*1.1 {
			t.Errorf("now %v is wrong", got)
		}

		time.Sleep(time.Duration(rand.Int63n(int64(duration))))
	}
}
