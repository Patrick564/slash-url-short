package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/Patrick564/url-shortener-backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
)

type Controllers interface {
	All() ([]string, error)
	Add(sid string, url string) (string, error)
	GoTo(id string) (string, error)
}

type Env struct {
	Urls Controllers
}

// ShowAccount godoc
// @Summary      List all urls
// @Description  get all urls and short id
// @Tags         urls
// @Accept       json
// @Produce      json
// @Success      200  {array} string
// @Router       /api/all [get]
func (e *Env) UrlsIndex(ctx *gin.Context) {
	u, err := e.Urls.All()
	if err != nil {
		log.Println(err)
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"urls": u})
}

// ShowAccount godoc
// @Summary      Create a new short url
// @Description  post a new short id for url
// @Tags         urls
// @Accept       json
// @Produce      json
// @Param        id   body    string  true  "Short URL"
// @Success      200 {string} string
// @Router       /api/add [post]
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
				gin.H{"error": utils.ErrEmptyBody.Error()},
			)
			return
		}

		log.Println(err)
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	sid, err := shortid.GetDefault().Generate()
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
	}

	short, err := e.Urls.Add(sid, body.Url)
	if err != nil {
		if errors.Is(err, utils.ErrEmptyBody) {
			log.Println(err)
			ctx.JSON(
				http.StatusBadRequest,
				gin.H{"error": err.Error()},
			)
			return
		}

		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"short_url": short,
		},
	)
}

// ShowAccount godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Short URL"
// @Success      301
// @Router       /api/{id} [get]
func (e *Env) UrlsGoToID(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(
			http.StatusNotFound,
			gin.H{
				"status": http.StatusNotFound,
				"error":  utils.ErrEmptyID.Error(),
			},
		)
		return
	}

	url, err := e.Urls.GoTo(id)
	if err != nil {
		log.Println(err)
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status": http.StatusBadRequest,
				"error":  err.Error(),
			},
		)
		return
	}

	ctx.Redirect(http.StatusMovedPermanently, url)
}
