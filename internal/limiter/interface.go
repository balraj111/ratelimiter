package limiter

type limiter interface {
	Allow(key string) bool
	GetRemaining(key string) int
	reset(key string)
}

