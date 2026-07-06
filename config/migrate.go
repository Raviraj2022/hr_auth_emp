package config

import (
	"log"

	"github.com/ravirajsahu/auth_app/internal/auth"
)

func AutoMigrate() {
	err := DB.AutoMigrate(
		&auth.User{},
	)

	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("✅ Database migrated successfully")
}