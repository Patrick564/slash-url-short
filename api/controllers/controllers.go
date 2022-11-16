package controllers

import (
	"net/http"

	"github.com/Patrick564/url-shortener-backend/internal/models"
	"github.com/gin-gonic/gin"
)

type Controllers interface {
	All() (models.Urls, error)
}

type Env struct {
	Urls Controllers
}

func (e *Env) All(ctx *gin.Context) {
	m, _ := e.Urls.All()

	ctx.JSON(http.StatusOK, gin.H{"data": m})
}
