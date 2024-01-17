package main

import (
	"net/http"

	"github.com/ffelipelimao/rate-limit-gin-middleware/pkg/ratelimit"
	"github.com/gin-gonic/gin"
)

func main() {

	app := gin.Default()

	rateLimit := ratelimit.NewRateLimit(5, 5)

	app.Use(rateLimit.Apply())

	app.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	app.Run()
}
