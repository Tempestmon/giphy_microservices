package handlers

import (
	"github.com/gin-gonic/gin"
)

func GetMeme(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(200, id)
}

func SetMeme(ctx *gin.Context) {
	ctx.JSON(200, "kek")
}
