package limiter

import (
	"sync"
	"time"
)

type FixedWindowLimiter struct {
	mu       sync.Mutex
	limit    int
	interval time.Duration
	store    map[string]*FixedWindowEntry
}

type FixedWindowEntry struct {
	count     int
	startTime time.Time
}

func NewFixedWindowLimiter(limit int, interval time.Duration) *FixedWindowLimiter {
	return &FixedWindowLimiter{
		limit:    limit,
		interval: interval,
		store:    make(map[string]*FixedWindowEntry),
	}
}

func (f *FixedWindowLimiter) Allow(key string) bool {
	f.mu.Lock()
	defer f.mu.Unlock()

	now := time.Now()
	entry, exist := f.store[key]

	if !exist || now.Sub(entry.startTime) > f.interval {
		f.store[key] = &FixedWindowEntry{count: 1, startTime: time.Now()}
		return true
	}

	if entry.count < f.limit {
		entry.count++
		return true
	}

	return false
}

func (f *FixedWindowLimiter) GetRemaining(key string) int {
	f.mu.Lock()
	defer f.mu.Unlock()

	entry, exist := f.store[key]
	if !exist {
		return f.limit
	}
	return f.limit - entry.count
}

func (f *FixedWindowLimiter) Reset(key string) {
	f.mu.Lock()
	defer f.mu.Unlock()
	delete(f.store, key)
}
