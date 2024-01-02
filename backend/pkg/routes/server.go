package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/souravdev-eng/toktok-api/pkg/handlers"
)

func RegisterRoutes(r *gin.Engine) {
	userGroup := r.Group("/api/user")

	userGroup.POST("/signup", handlers.Signup)
	userGroup.POST("/login", handlers.Login)
}
