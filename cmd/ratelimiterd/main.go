package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/balraj111/ratelimiter/internal/limiter"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("running main")
	r := limiter.NewLimiterFactory(limiter.LimiterConfig{
		Limit:    5,
		Interval: 10 * time.Second,
		Type:     limiter.FixedWindow,
	})

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		key := c.ClientIP()
		if !r.Allow(key) {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "too many request",
			})
			return
		}
		c.Next()
	})

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Println("starting server at :8080")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed: %v", err)
	}

}
