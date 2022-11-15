package api

import (
	"github.com/Patrick564/url-shortener-backend/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type tem struct {
	books models.UrlModel
}

func SetupRouter(db *pgxpool.Pool) *gin.Engine {
	r := gin.Default()
	b := tem{books: models.UrlModel{DB: db}}

	r.GET("/api/all", b.urlsIndex)

	return r
}

func (t *tem) urlsIndex(ctx *gin.Context) {}
