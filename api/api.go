package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// get all
	r.GET("/list", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"ggl": "https://google.com",
		})
	})

	// redirect url
	r.GET("/api/:url", redirectUrl)

	// short url
	r.POST("/long-url", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"https://google.com": "ggl",
		})
	})

	return r
}
