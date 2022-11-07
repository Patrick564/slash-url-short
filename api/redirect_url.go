package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func redirectUrl(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message":  "pong",
		"response": "ping",
	})
}
