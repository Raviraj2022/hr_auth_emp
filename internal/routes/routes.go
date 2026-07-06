package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/ravirajsahu/auth_app/config"
	"github.com/ravirajsahu/auth_app/internal/auth"
)

func Setup(router *gin.Engine) {

	// Dependency Injection
	repository := auth.NewRepository(config.DB)
	service := auth.NewService(repository)
	handler := auth.NewHandler(service)

	api := router.Group("/api")

	{
		api.POST("/register", handler.Register)
		api.POST("/login", handler.Login)
	}
}