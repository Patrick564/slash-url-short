package api

import (
	"github.com/Patrick564/url-shortener-backend/api/controllers"
	"github.com/Patrick564/url-shortener-backend/internal/models"
	"github.com/gin-gonic/gin"
)

func SetupRouter(db *models.Database) *gin.Engine {
	r := gin.Default()
	c := &controllers.Env{
		Urls: &models.UrlModel{DB: db.DB, Ctx: db.Ctx},
	}

	r.GET("/api/all", c.All)

	return r
}
