package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/Patrick564/url-shortener-backend/internal/models"
	"github.com/Patrick564/url-shortener-backend/utils"
	"github.com/gin-gonic/gin"
)

type Controllers interface {
	All() ([]models.Url, error)
	Add(url string) (models.Url, error)
}

type Env struct {
	Urls Controllers
}

func (e *Env) UrlsIndex(ctx *gin.Context) {
	u, err := e.Urls.All()
	if err != nil {
		log.Println(err)
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err, "urls": []models.Url{}},
		)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"error": nil, "urls": u})
}

func (e *Env) UrlsAdd(ctx *gin.Context) {
	var body struct {
		Url string `json:"url"`
	}

	err := ctx.BindJSON(&body)
	if err != nil {
		if err.Error() == "EOF" {
			log.Println(err)
			ctx.JSON(
				http.StatusBadRequest,
				gin.H{"error": utils.ErrEmptyBody, "url": models.Url{}},
			)
			return
		}

		log.Println(err)
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err, "url": models.Url{}},
		)
		return
	}

	u, err := e.Urls.Add(body.Url)
	if err != nil {
		if errors.Is(err, utils.ErrEmptyBody) {
			log.Println(err)
			ctx.JSON(
				http.StatusBadRequest,
				gin.H{"error": err, "url": models.Url{}},
			)
			return
		}
		log.Println(err)
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err, "url": models.Url{}},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{"error": nil, "url": u},
	)
}
