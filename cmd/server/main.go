//	@title			HRMS API
//	@version		1.0
//	@description	HRMS Backend API using Go, Gin, GORM & PostgreSQL.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Raviraj Sahu
//	@contact.email	ravi@example.com

//	@license.name	MIT

//	@host		localhost:8080
//	@BasePath	/api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
package main

import (
	"log"
	"time"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"

	"github.com/ravirajsahu/auth_app/config"
	"github.com/ravirajsahu/auth_app/internal/routes"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/ravirajsahu/auth_app/docs"
)

func main() {

	// Load .env
	config.LoadEnv()

	// Connect Database
	config.ConnectDB()

	config.AutoMigrate()

	// Create Gin server
	router := gin.Default()

	router.Use(cors.New(cors.Config{
    AllowOrigins: []string{
        "http://localhost:3000", // Next.js
        "http://localhost:5173", // Vite
    },

// 	AllowOrigins: []string{
//     "https://hrms.yourcompany.com",
//     "https://admin.yourcompany.com",
// }

    AllowMethods: []string{
        "GET",
        "POST",
        "PUT",
        "PATCH",
        "DELETE",
        "OPTIONS",
    },

    AllowHeaders: []string{
        "Origin",
        "Content-Type",
        "Authorization",
    },

    ExposeHeaders: []string{
        "Content-Length",
    },

    AllowCredentials: true,

    MaxAge: 12 * time.Hour,
}))

	// router.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "HR Auth API is running 🚀",
	// 	})
	// })

	routes.Setup(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Printf("Server running on port %s", config.App.Port)

	if err := router.Run(":" + config.App.Port); err != nil {
		log.Fatal(err)
	}
}
