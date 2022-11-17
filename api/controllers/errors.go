package controllers

import "github.com/gin-gonic/gin"

func ErrorHandler(ctx *gin.Context) {
	ctx.Next()
}
