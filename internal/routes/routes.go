package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/ravirajsahu/auth_app/config"
	"github.com/ravirajsahu/auth_app/internal/auth"
	"github.com/ravirajsahu/auth_app/internal/employee"
	"github.com/ravirajsahu/auth_app/internal/middleware"
	
)

func Setup(router *gin.Engine) {

	// Auth Module
	authRepo := auth.NewRepository(config.DB)
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService)

	// Refresh Module
	// refreshRepo := refresh.NewRepository(config.DB)
	// refreshService := refresh.NewService(refreshRepo, authRepo)
	// refreshHandler := refresh.NewHandler(refreshService)
employeeRepo := employee.NewRepository(config.DB)
employeeService := employee.NewService(employeeRepo)
employeeHandler := employee.NewHandler(employeeService)

	api := router.Group("/api")

	// Public Routes
	api.POST("/register", authHandler.Register)
	api.POST("/login", authHandler.Login)

	// api.POST("/refresh", refreshHandler.Refresh)
	// api.POST("/logout", refreshHandler.Logout)

	// Protected Routes
	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/profile", authHandler.Profile)
		protected.POST("/employees", employeeHandler.Create)
	protected.GET("/employees", employeeHandler.GetAll)
	protected.GET("/employees/:id", employeeHandler.GetByID)
	protected.PUT("/employees/:id", employeeHandler.Update)
	protected.DELETE("/employees/:id", employeeHandler.Delete)
	}


}