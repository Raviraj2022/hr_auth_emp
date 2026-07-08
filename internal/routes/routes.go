package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/ravirajsahu/auth_app/config"
	"github.com/ravirajsahu/auth_app/internal/auth"
	"github.com/ravirajsahu/auth_app/internal/department"
	"github.com/ravirajsahu/auth_app/internal/employee"
	 "github.com/ravirajsahu/auth_app/internal/attendance"
	"github.com/ravirajsahu/auth_app/internal/middleware"
)

func Setup(router *gin.Engine) {

	// =========================
	// Auth Module
	// =========================
	authRepo := auth.NewRepository(config.DB)
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService)

	// =========================
	// Department Module
	// =========================
	departmentRepo := department.NewRepository(config.DB)
	departmentService := department.NewService(departmentRepo)
	departmentHandler := department.NewHandler(departmentService)

	// =========================
	// Employee Module
	// =========================
	employeeRepo := employee.NewRepository(config.DB)

	// Employee Service depends on Department Repository
	employeeService := employee.NewService(employeeRepo, departmentRepo)

	employeeHandler := employee.NewHandler(employeeService)


	attendanceRepo := attendance.NewRepository(config.DB)

attendanceService := attendance.NewService(
	attendanceRepo,
	employeeRepo,
)

attendanceHandler := attendance.NewHandler(attendanceService)
	// =========================
	// API Group
	// =========================
	api := router.Group("/api")

	// ---------- Public Routes ----------
	api.POST("/register", authHandler.Register)
	api.POST("/login", authHandler.Login)

	// ---------- Protected Routes ----------
	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware())

	// Profile
	protected.GET("/profile", authHandler.Profile)

	// Employee Routes
	protected.POST("/employees", employeeHandler.Create)
	protected.GET("/employees", employeeHandler.GetAll)
	protected.GET("/employees/:id", employeeHandler.GetByID)
	protected.PUT("/employees/:id", employeeHandler.Update)
	protected.DELETE("/employees/:id", employeeHandler.Delete)

	// Department Routes
	protected.POST("/departments", departmentHandler.Create)
	protected.GET("/departments", departmentHandler.GetAll)
	protected.GET("/departments/:id", departmentHandler.GetByID)
	protected.PUT("/departments/:id", departmentHandler.Update)
	protected.DELETE("/departments/:id", departmentHandler.Delete)

	// Attendance
protected.POST("/attendance/check-in", attendanceHandler.CheckIn)
protected.POST("/attendance/check-out", attendanceHandler.CheckOut)

protected.GET("/attendance", attendanceHandler.GetAll)
protected.GET("/attendance/:id", attendanceHandler.GetByID)

protected.GET("/attendance/employee/:employee_id", attendanceHandler.GetByEmployee)

protected.DELETE("/attendance/:id", attendanceHandler.Delete)
}