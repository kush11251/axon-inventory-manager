package main

import (
	"axon-inventory-manager/src/controllers"
	"axon-inventory-manager/src/models"
	"axon-inventory-manager/src/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the Gin router
	r := gin.Default()

	// Initialize the controllers
	controllers.Init(r)

	// Start the server
	r.Run("localhost:8080")
}
