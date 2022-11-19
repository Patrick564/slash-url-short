package controllers

import (
	"log"
	"net/http"

	"github.com/Patrick564/url-shortener-backend/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
)

type Controllers interface {
	All() ([]models.Url, error)
	Add(id string, url string) (models.Url, error)
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
			gin.H{"urls": "", "error": err},
		)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"urls": u, "error": ""})
}

type UrlResponse struct {
	Url string `json:"url"`
}

func (e *Env) UrlsAdd(ctx *gin.Context) {
	sid, _ := shortid.New(1, shortid.DefaultABC, 2342)
	id, _ := sid.Generate()

	var r UrlResponse

	err := ctx.BindJSON(&r)
	if err != nil {
		log.Println(err)
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error(), "url": ""},
		)
		return
	}

	mn, err := e.Urls.Add(id, r.Url)
	if err != nil {
		log.Println(err)
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error(), "urls": ""},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{"error": "", "url": mn},
	)
}
