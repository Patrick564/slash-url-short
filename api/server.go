package api

import (
	"github.com/Patrick564/url-shortener-backend/api/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(env *controllers.Env) *gin.Engine {
	r := gin.Default()
	// r.Use(controllers.ErrorHandler)

	r.GET("/api/all", env.UrlsIndex)
	r.POST("/api/add", env.UrlsAdd)

	return r
}
