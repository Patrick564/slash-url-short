package list

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListRoute(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message":  "pong",
		"response": "ping",
	})
}
