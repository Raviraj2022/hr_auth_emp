// package routes

// import (
// 	"github.com/gin-gonic/gin"

// 	"github.com/ravirajsahu/auth_app/config"
// 	"github.com/ravirajsahu/auth_app/internal/auth"
// 	"github.com/ravirajsahu/auth_app/internal/department"
// 	"github.com/ravirajsahu/auth_app/internal/employee"
// 	 "github.com/ravirajsahu/auth_app/internal/attendance"
// 	 "github.com/ravirajsahu/auth_app/internal/leave"
// 	 "github.com/ravirajsahu/auth_app/internal/payroll"
// 	"github.com/ravirajsahu/auth_app/internal/middleware"
// )

// func Setup(router *gin.Engine) {

// 	// =========================
// 	// Auth Module
// 	// =========================
// 	authRepo := auth.NewRepository(config.DB)
// 	authService := auth.NewService(authRepo)
// 	authHandler := auth.NewHandler(authService)

// 	// =========================
// 	// Department Module
// 	// =========================
// 	departmentRepo := department.NewRepository(config.DB)
// 	departmentService := department.NewService(departmentRepo)
// 	departmentHandler := department.NewHandler(departmentService)

// 	// =========================
// 	// Employee Module
// 	// =========================
// 	employeeRepo := employee.NewRepository(config.DB)

// 	// Employee Service depends on Department Repository
// 	employeeService := employee.NewService(employeeRepo, departmentRepo)

// 	employeeHandler := employee.NewHandler(employeeService)

// 	attendanceRepo := attendance.NewRepository(config.DB)

// attendanceService := attendance.NewService(
// 	attendanceRepo,
// 	employeeRepo,
// )

// attendanceHandler := attendance.NewHandler(attendanceService)

// leaveRepo := leave.NewRepository(config.DB)

// leaveService := leave.NewService(
// 	leaveRepo,
// 	employeeRepo,
// )

// leaveHandler := leave.NewHandler(leaveService)

// payrollRepo := payroll.NewRepository(config.DB)

// payrollService := payroll.NewService(
// 	payrollRepo,
// 	employeeRepo,
// 	attendanceRepo,
// 	leaveRepo,
// )

// payrollHandler := payroll.NewHandler(payrollService)
// 	// =========================
// 	// API Group
// 	// =========================
// 	api := router.Group("/api")

// 	// ---------- Public Routes ----------
// 	api.POST("/register", authHandler.Register)
// 	api.POST("/login", authHandler.Login)

// 	// ---------- Protected Routes ----------
// 	protected := api.Group("/")
// 	protected.Use(middleware.AuthMiddleware())

// 	// Profile
// 	protected.GET("/profile", authHandler.Profile)

// 	// Employee Routes
// 	protected.POST("/employees", middleware.AdminOrHR(), employeeHandler.Create)
// 	protected.GET("/employees", employeeHandler.GetAll)
// 	protected.GET("/employees/:id", employeeHandler.GetByID)
// 	protected.PUT("/employees/:id", middleware.AdminOrHR(), employeeHandler.Update)
// 	protected.DELETE("/employees/:id", middleware.AdminOnly(), employeeHandler.Delete)

// 	// Department Routes
// 	protected.POST("/departments", middleware.AdminOrHR(), departmentHandler.Create)
// 	protected.GET("/departments", departmentHandler.GetAll)
// 	protected.GET("/departments/:id", departmentHandler.GetByID)
// 	protected.PUT("/departments/:id", departmentHandler.Update)
// 	protected.DELETE("/departments/:id", middleware.AdminOnly(), departmentHandler.Delete)

// 	// Attendance
// protected.POST("/attendance/check-in", attendanceHandler.CheckIn)
// protected.POST("/attendance/check-out", attendanceHandler.CheckOut)

// protected.GET("/attendance", attendanceHandler.GetAll)
// protected.GET("/attendance/:id", attendanceHandler.GetByID)

// protected.GET("/attendance/employee/:employee_id", attendanceHandler.GetByEmployee)

// protected.DELETE("/attendance/:id", attendanceHandler.Delete)

// // Leave
// protected.POST("/leaves", middleware.EmployeeOnly(), leaveHandler.Create)

// protected.GET("/leaves", leaveHandler.GetAll)

// protected.GET("/leaves/:id", leaveHandler.GetByID)

// protected.GET("/leaves/employee/:employee_id", leaveHandler.GetByEmployee)

// protected.PUT("/leaves/:id/approve", middleware.AdminHRManager(), leaveHandler.Approve)

// protected.PUT("/leaves/:id/reject", middleware.AdminHRManager(), leaveHandler.Reject)

// protected.DELETE("/leaves/:id", leaveHandler.Delete)

// // Payroll

// protected.POST("/payroll/generate", middleware.AdminOrHR(), payrollHandler.Generate)

// protected.GET("/payroll", payrollHandler.GetAll)

// protected.GET("/payroll/:id", payrollHandler.GetByID)

// protected.GET("/payroll/employee/:employee_id", payrollHandler.GetByEmployee)

// protected.PUT("/payroll/:id/pay", middleware.AdminOrHR(), payrollHandler.MarkPaid)

// protected.DELETE("/payroll/:id", payrollHandler.Delete)
// }

package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/ravirajsahu/auth_app/config"

	"github.com/ravirajsahu/auth_app/internal/attendance"
	"github.com/ravirajsahu/auth_app/internal/auth"
	"github.com/ravirajsahu/auth_app/internal/department"
	"github.com/ravirajsahu/auth_app/internal/employee"
	"github.com/ravirajsahu/auth_app/internal/leave"
	"github.com/ravirajsahu/auth_app/internal/payroll"
    "github.com/ravirajsahu/auth_app/internal/dashboard"
	"github.com/ravirajsahu/auth_app/internal/holiday"
	"github.com/ravirajsahu/auth_app/internal/middleware"
)

func Setup(router *gin.Engine) {

	// ========= Repositories =========

	authRepo := auth.NewRepository(config.DB)
	employeeRepo := employee.NewRepository(config.DB)
	departmentRepo := department.NewRepository(config.DB)
	attendanceRepo := attendance.NewRepository(config.DB)
	leaveRepo := leave.NewRepository(config.DB)
	payrollRepo := payroll.NewRepository(config.DB)
	dashboardRepo := dashboard.NewRepository(config.DB)
    holidayRepo := holiday.NewRepository(config.DB)
	// ========= Services =========

	authService := auth.NewService(authRepo)

	employeeService := employee.NewService(
		employeeRepo,
		departmentRepo,
	)

	departmentService := department.NewService(
		departmentRepo,
	)

	attendanceService := attendance.NewService(
		attendanceRepo,
		employeeRepo,
	)

	leaveService := leave.NewService(
		leaveRepo,
		employeeRepo,
	)

	payrollService := payroll.NewService(
		payrollRepo,
		employeeRepo,
		attendanceRepo,
		leaveRepo,
	)

	dashboardService := dashboard.NewService(dashboardRepo)

	holidayService := holiday.NewService(holidayRepo)
	// ========= Handlers =========

	authHandler := auth.NewHandler(authService)

	employeeHandler := employee.NewHandler(employeeService)

	departmentHandler := department.NewHandler(departmentService)

	attendanceHandler := attendance.NewHandler(attendanceService)

	leaveHandler := leave.NewHandler(leaveService)

	payrollHandler := payroll.NewHandler(payrollService)

	dashboardHandler := dashboard.NewHandler(dashboardService)

	holidayHandler := holiday.NewHandler(holidayService)
	api := router.Group("/api")

	// -------------------------
	// Public Routes
	// -------------------------

	api.POST("/register", authHandler.Register)
	api.POST("/login", authHandler.Login)

	// -------------------------
	// Authenticated Routes
	// -------------------------

	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware())

	{
		protected.GET("/profile", authHandler.Profile)
	}

	// -------------------------
	// Admin
	// -------------------------

	admin := protected.Group("/")
	admin.Use(middleware.AdminOnly())

	{
		admin.DELETE("/employees/:id", employeeHandler.Delete)

		admin.DELETE("/departments/:id", departmentHandler.Delete)

		admin.DELETE("/payroll/:id", payrollHandler.Delete)

		// Admin only
admin.DELETE("/holidays/:id", holidayHandler.Delete)
	}

	// -------------------------
	// Admin + HR
	// -------------------------

	hr := protected.Group("/")
	hr.Use(middleware.AdminOrHR())

	{
		// Employee

		hr.POST("/employees", employeeHandler.Create)
		hr.PUT("/employees/:id", employeeHandler.Update)

		// Department

		hr.POST("/departments", departmentHandler.Create)
		hr.PUT("/departments/:id", departmentHandler.Update)

		// Payroll

		hr.POST("/payroll/generate", payrollHandler.Generate)
		hr.PUT("/payroll/:id/pay", payrollHandler.MarkPaid)

		hr.POST("/holidays", holidayHandler.Create)
hr.PUT("/holidays/:id", holidayHandler.Update)
// hr.DELETE("/holidays/:id", holidayHandler.Delete)
	}

	// -------------------------
	// Admin + HR + Manager
	// -------------------------

	manager := protected.Group("/")
	manager.Use(middleware.AdminHRManager())

	{
		manager.PUT("/leaves/:id/approve", leaveHandler.Approve)

		manager.PUT("/leaves/:id/reject", leaveHandler.Reject)
	}

	// -------------------------
	// All Logged-in Users
	// -------------------------

	employeeRoutes := protected.Group("/")

	{

		employeeRoutes.GET("/dashboard", dashboardHandler.GetDashboard)

		// Employee

		employeeRoutes.GET("/employees", employeeHandler.GetAll)
		employeeRoutes.GET("/employees/:id", employeeHandler.GetByID)

		// Department

		employeeRoutes.GET("/departments", departmentHandler.GetAll)
		employeeRoutes.GET("/departments/:id", departmentHandler.GetByID)

		// Attendance

		employeeRoutes.POST("/attendance/check-in", attendanceHandler.CheckIn)
		employeeRoutes.PUT("/attendance/check-out", attendanceHandler.CheckOut)
		employeeRoutes.GET("/attendance", attendanceHandler.GetAll)

		// Leave

		employeeRoutes.POST("/leaves", leaveHandler.Create)
		employeeRoutes.GET("/leaves", leaveHandler.GetAll)

		// Payroll

		employeeRoutes.GET("/payroll", payrollHandler.GetAll)
		employeeRoutes.GET("/payroll/:id", payrollHandler.GetByID)
		employeeRoutes.GET("/payroll/employee/:employee_id", payrollHandler.GetByEmployee)

		// All authenticated users
employeeRoutes.GET("/holidays", holidayHandler.GetAll)
employeeRoutes.GET("/holidays/:id", holidayHandler.GetByID)
employeeRoutes.GET("/holidays/year/:year", holidayHandler.GetByYear)

	}
}
