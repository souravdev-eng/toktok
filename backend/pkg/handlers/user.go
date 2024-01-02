package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signup(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Signup successfully!",
	})
}

func Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Signup successfully!",
	})
}
