package controllers

import (
	"log"
	"net/http"

	"github.com/Patrick564/url-shortener-backend/internal/models"
	"github.com/gin-gonic/gin"
)

type Controllers interface {
	GetAll() ([]models.Url, error)
}

type Env struct {
	Urls Controllers
}

func (e *Env) UrlsIndex(ctx *gin.Context) {
	u, err := e.Urls.GetAll()
	if err != nil {
		log.Println(err)
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"urls": "", "error": err},
		)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"urls": u, "error": ""})
}
