package main

import (
	"net/http"
	"time"

	"github.com/ffelipelimao/rate-limit-gin-middleware/pkg/ratelimit"
	"github.com/gin-gonic/gin"
)

func main() {

	app := gin.Default()

	rateLimit := ratelimit.NewRateLimit(100, 5*time.Second)

	app.Use(rateLimit.Apply())

	app.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
		return
	})

	app.Run()
}
