package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/souravdev-eng/toktok-api/pkg/config"
	"github.com/souravdev-eng/toktok-api/pkg/routes"
	"github.com/souravdev-eng/toktok-api/utils"
)

func main() {
	utils.LoadEnv()
	port := utils.GetEnvValue("PORT")

	server := gin.Default()
	routes.RegisterRoutes(server)
	client, err := config.ConnectDB()

	if err != nil {
		fmt.Println("MongoDB connection error:", err)
		return
	}

	defer client.Disconnect(nil)

	fmt.Println("Connected to MongoDB successfully!")
	server.Run(":" + port)
}
