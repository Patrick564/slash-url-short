package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// change name to controller

func All(ctx *gin.Context) {
	// db.AllUrls()
	// u, _ := models.All()

	ctx.JSON(http.StatusOK, gin.H{
		"ggl": "https://google.com",
		"go":  "https://go.dev",
		"tw":  "https://twitch.tv",
	})
}
