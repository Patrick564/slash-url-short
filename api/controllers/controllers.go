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

// It's ok use a pointer this
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

type UrlResponse struct {
	Url string `json:"url"`
}

func (e *Env) UrlsAdd(ctx *gin.Context) {
	var r UrlResponse

	err := ctx.BindJSON(&r)
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

	mn, err := e.Urls.Add(r.Url)
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
			gin.H{"error": err.Error(), "url": models.Url{}},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{"error": nil, "url": mn},
	)
}
