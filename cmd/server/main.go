package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/ravirajsahu/auth_app/config"
	"github.com/ravirajsahu/auth_app/internal/routes"
)

func main() {

	// Load .env
	config.LoadEnv()

	// Connect Database
	config.ConnectDB()

	config.AutoMigrate()

	// Create Gin server
	router := gin.Default()

	// router.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "HR Auth API is running 🚀",
	// 	})
	// })

routes.Setup(router)

router.Run(":" + config.App.Port)

	log.Printf("Server running on port %s", config.App.Port)

	err := router.Run(":" + config.App.Port)
	if err != nil {
		log.Fatal(err)
	}
}