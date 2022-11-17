package api

import (
	"github.com/Patrick564/url-shortener-backend/api/controllers"
	"github.com/Patrick564/url-shortener-backend/internal/models"
	"github.com/gin-gonic/gin"
)

func SetupRouter(u *models.UrlModel) *gin.Engine {
	r := gin.Default()
	// r.Use(controllers.ErrorHandler)

	c := &controllers.Env{
		Urls: &models.UrlModel{DB: u.DB, Ctx: u.Ctx},
	}

	r.GET("/api/all", c.UrlsIndex)

	return r
}
