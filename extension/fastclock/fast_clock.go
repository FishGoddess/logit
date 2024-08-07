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
	"sync"
	"sync/atomic"
	"time"
)

// fastClock is a clock for getting current time faster.
// It caches current time in nanos and updates it in fixed duration, so it's not a precise way to get current time.
// In fact, we don't recommend you to use it unless you do need a fast way to get current time even the time is "incorrect".
// According to our benchmarks, it does run faster than time.Now:
//
// In my linux server with 2 cores:
// BenchmarkTimeNow-2               19150246                62.26 ns/op           0 B/op          0 allocs/op
// BenchmarkFastClockNow-2         357209233                 3.46 ns/op           0 B/op          0 allocs/op
// BenchmarkFastClockNowNanos-2    467461363                 2.55 ns/op           0 B/op          0 allocs/op
//
// However, the performance of time.Now is faster enough for 99.9% situations, so we hope you never use it :)
type fastClock struct {
	nanos int64
}

func newClock() *fastClock {
	clock := &fastClock{
		nanos: time.Now().UnixNano(),
	}

	go clock.start()
	return clock
}

func (fc *fastClock) start() {
	const duration = 100 * time.Millisecond

	for {
		for i := 0; i < 9; i++ {
			time.Sleep(duration)
			atomic.AddInt64(&fc.nanos, int64(duration))
		}

		time.Sleep(duration)
		atomic.StoreInt64(&fc.nanos, time.Now().UnixNano())
	}
}

func (fc *fastClock) Nanos() int64 {
	return atomic.LoadInt64(&fc.nanos)
}

var (
	clock     *fastClock
	clockOnce sync.Once
)

// Now returns the current time from fast clock.
func Now() time.Time {
	clockOnce.Do(func() {
		clock = newClock()
	})

	nanos := NowNanos()
	return time.Unix(0, nanos)
}

// NowNanos returns the current time in nanos from fast clock.
func NowNanos() int64 {
	clockOnce.Do(func() {
		clock = newClock()
	})

	return clock.Nanos()
}
