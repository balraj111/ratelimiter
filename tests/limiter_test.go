package test

import (
	"testing"
	"time"

	"github.com/balraj111/ratelimiter/internal/limiter"
)

func TestFixedWindowLimiter(t *testing.T) {
	limit := 3
	interval := 2 * time.Second

	r := limiter.NewFixedWindowLimiter(limit, interval)
	key := "test-user"
	for i := 0; i < limit; i++ {
		allowed := r.Allow(key)
		t.Logf("Attempt %d: allowed = %v, remaining = %d", i+1, allowed, r.GetRemaining(key))
		if !allowed {
			t.Fatalf("expected request %d should be allowed", i+1)
		}
	}

	if r.Allow(key) {
		t.Fatalf("expected request should be denied after reaching the limit")
	}

	time.Sleep(6 * time.Second)

	if !r.Allow(key) {
		t.Fatalf("expected request to be allowed after window reset")
	}
}

func TestSlidingWindowLimiter(t *testing.T) {
	limit := 3
	interval := 5 * time.Second
	key := "test-user"
	r := limiter.NewSlidingWindowLimiter(limit, interval)
	var allowed bool
	for i := 0; i < limit; i++ {
		allowed = r.Allow(key)
		t.Logf("Attempt %d: allowed = %v, remaining = %d", i+1, allowed, r.GetRemaining(key))
		if !allowed {
			t.Fatalf("expected request %d should be allowed", i+1)
		}
		time.Sleep(1 * time.Second)
	}

	allowed = r.Allow(key)
	t.Logf("allowed = %v, remaining = %d", allowed, r.GetRemaining(key))
	if allowed {
		t.Fatalf("expected request should be denied after reaching the limit")
	}

	time.Sleep(3 * time.Second)
	allowed = r.Allow(key)
	t.Logf(" allowed = %v, remaining = %d", allowed, r.GetRemaining(key))
	if !allowed {
		t.Fatalf("expected result should be allowed")
	}

}
