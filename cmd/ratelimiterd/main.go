package main

import (
	"fmt"
	"ratelimiter/internal/limiter"
	"time"
)

func main() {
	r := limiter.NewFixedWindowLimiter(5, 10*time.Second)
	key := "user123"

	for i := 0; i < 10; i++ {
		if r.Allow(key) {
			fmt.Println("Request allowed")
		} else {
			fmt.Println("Rate limit exceeded")
		}
		fmt.Println("Request left", r.GetRemaining(key))
		time.Sleep(1 * time.Second)
		if r.GetRemaining(key) == 0 {
			r.Reset(key)
		}
	}
}
