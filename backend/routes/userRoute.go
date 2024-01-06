package routes

import (
	controller "github.com/souravdev-eng/toktok-api/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/api/users/signup", controller.SignUp())
}
