package ratelimit

import (
	"time"

	"github.com/gin-gonic/gin"
)

type RateLimit struct {
	Requests int
	Interval time.Duration
	Storage  map[string]int
}

func NewRateLimit(req int, interval time.Duration) *RateLimit {
	return &RateLimit{
		Requests: req,
		Interval: interval,
		Storage:  make(map[string]int),
	}
}

func (r *RateLimit) Apply() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (r *RateLimit) purge() {}
