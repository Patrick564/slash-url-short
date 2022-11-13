package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/all", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"ggl": "https://google.com",
		})
	})

	// redirect url
	// r.GET("/api/:url")

	// short url
	// r.POST("/long-url", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"https://google.com": "ggl",
	// 	})
	// })

	return r
}
