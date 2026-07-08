package config

import (
	"log"

	"github.com/ravirajsahu/auth_app/internal/auth"
	"github.com/ravirajsahu/auth_app/internal/employee"
	"github.com/ravirajsahu/auth_app/internal/department"
)

func AutoMigrate() {
	err := DB.AutoMigrate(
		&auth.User{},
		&employee.Employee{},
		&department.Department{},
		
	)

	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("✅ Database migrated successfully")
}