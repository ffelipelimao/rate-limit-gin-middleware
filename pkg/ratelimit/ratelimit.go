package ratelimit

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type RateLimit struct {
	Requests int
	Interval time.Duration
	Storage  map[string]storage
}

type storage struct {
	Value     int
	Timestamp time.Time
}

func NewRateLimit(req int, interval time.Duration) *RateLimit {
	rl := &RateLimit{
		Requests: req,
		Interval: interval,
		Storage:  make(map[string]storage),
	}

	go rl.purge()

	return rl
}

func (r *RateLimit) Apply() gin.HandlerFunc {
	return func(c *gin.Context) {
		IP := c.ClientIP()

		userRequests, ok := r.Storage[IP]
		if !ok {
			r.Storage[IP] = storage{
				Value:     1,
				Timestamp: time.Now(),
			}
			c.Next()
			return
		}
		userRequests.Value++
		r.Storage[IP] = userRequests

		if userRequests.Value >= r.Requests {
			c.JSON(http.StatusBadRequest, gin.H{"error": "rate limit quota"})
			c.Abort()
			return
		}
	}
}

func (r *RateLimit) purge() {
	for {
		time.Sleep(r.Interval)
		currentTime := time.Now()

		for key, s := range r.Storage {
			if currentTime.Sub(s.Timestamp) > r.Interval {
				delete(r.Storage, key)
			}
		}
	}
}
