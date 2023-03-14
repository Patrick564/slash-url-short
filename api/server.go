package api

import (
	"github.com/Patrick564/url-shortener-backend/api/controllers"
	"github.com/Patrick564/url-shortener-backend/docs"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Url Shortener API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @host      localhost:8080
// @BasePath  /api/v1

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func SetupRouter(env *controllers.Env) *gin.Engine {
	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/api"

	// r.Use(controllers.ErrorHandler)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/list", env.UrlsIndex)
	r.GET("/:id", env.UrlsGoToID)

	r.POST("/create", env.UrlsAdd)

	return r
}
