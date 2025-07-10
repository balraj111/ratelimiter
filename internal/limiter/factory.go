package limiter

import (
	"time"
)

type LimiterType string

const (
	FixedWindow  LimiterType = "fixed"
	SldingWindow LimiterType = "sliding"
)

type LimiterConfig struct {
	Type     LimiterType
	Limit    int
	Interval time.Duration
}

func NewLimiterFactory(config LimiterConfig) Limiter {
	switch config.Type {
	case FixedWindow:
		return NewFixedWindowLimiter(config.Limit, config.Interval)
	case SldingWindow:
		return NewSlidingWindowLimiter(config.Limit, config.Interval)
	default:
		panic("unsupported limiter type")
	}
}
