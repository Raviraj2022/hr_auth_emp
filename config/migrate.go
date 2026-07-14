package config

import (
	"log"

	"github.com/ravirajsahu/auth_app/internal/attendance"
	"github.com/ravirajsahu/auth_app/internal/auth"
	"github.com/ravirajsahu/auth_app/internal/department"
	"github.com/ravirajsahu/auth_app/internal/employee"
	"github.com/ravirajsahu/auth_app/internal/leave"
	"github.com/ravirajsahu/auth_app/internal/payroll"
	"github.com/ravirajsahu/auth_app/internal/holiday"
)

func AutoMigrate() {
	err := DB.AutoMigrate(
		&auth.User{},
		&employee.Employee{},
		&department.Department{},
		&attendance.Attendance{},
		&leave.Leave{},
		&payroll.Payroll{},
		&holiday.Holiday{},
	)

	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("✅ Database migrated successfully")
}
