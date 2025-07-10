package test

import (
	"testing"
	"time"

	"github.com/balraj111/ratelimiter/internal/limiter"
)

func BenchMarkFixedWindowLimiter(b *testing.B) {
	r := limiter.NewFixedWindowLimiter(1000, time.Second)
	key := "benchmark-test-key"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.Allow(key)
	}

}

func BenchMarkSlidingWindowLimiter(b *testing.B) {
	r := limiter.NewSlidingWindowLimiter(1000, time.Second)
	key := "benchmark-test-key"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.Allow(key)
	}

}
