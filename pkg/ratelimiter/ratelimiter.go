package ratelimiter

import "github.com/balraj111/ratelimiter/internal/limiter"

type Limiter = limiter.Limiter
type Config = limiter.LimiterConfig
type LimiterType = limiter.LimiterType

const (
	FixedWindow  = limiter.FixedWindow
	SldingWindow = limiter.SldingWindow
)

func New(config Config) Limiter {
	return limiter.NewLimiterFactory(config)
}
