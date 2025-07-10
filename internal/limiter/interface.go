package limiter

type Limiter interface {
	Allow(key string) bool
	GetRemaining(key string) int
	Reset(key string)
}
