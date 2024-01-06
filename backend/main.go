package main

import (
	"os"

	routes "github.com/souravdev-eng/toktok-api/routes"
	"github.com/souravdev-eng/toktok-api/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoadEnv()
	port := os.Getenv("PORT")

	if port == "" {
		port = "4000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)

	router.Run(":" + port)
}
