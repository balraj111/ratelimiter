package limiter

import (
	"sync"
	"time"
)

type SlidingWindowLimiter struct {
	mu       sync.Mutex
	limit    int
	interval time.Duration
	store    map[string][]time.Time
}

func NewSlidingWindowLimiter(limit int, interval time.Duration) *SlidingWindowLimiter {
	return &SlidingWindowLimiter{
		limit:    limit,
		interval: interval,
		store:    make(map[string][]time.Time),
	}
}

func (s *SlidingWindowLimiter) Allow(key string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now()
	windowStart := time.Now().Add(-s.interval)
	timeStamps := s.store[key]

	filtered := timeStamps[:0]

	for _, ts := range timeStamps {
		if ts.After(windowStart) {
			filtered = append(filtered, ts)
		}
	}

	if len(filtered) < s.limit {
		filtered = append(filtered, now)
		s.store[key] = filtered
		return true
	}

	s.store[key] = filtered
	return false
}

func (s *SlidingWindowLimiter) GetRemaining(key string) int {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now()
	startWindow := now.Add(-s.interval)
	timeStamps := s.store[key]
	count := 0

	for _, ts := range timeStamps {
		if ts.After(startWindow) {
			count++
		}
	}

	return s.limit - count
}

func (s *SlidingWindowLimiter) Reset(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.store, key)
}
