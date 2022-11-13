package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AllRoute(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"asd":  "pong",
		"zxcs": "ping",
	})
}
