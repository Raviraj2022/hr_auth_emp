package auth

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	RoleAdmin    = "Admin"
	RoleHR       = "HR"
	RoleManager  = "Manager"
	RoleEmployee = "Employee"
)

type User struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	Name string `gorm:"size:100;not null"`

	Email string `gorm:"size:100;uniqueIndex;not null"`

	Password string `gorm:"not null"`
	Role     string `gorm:"size:20;default:'Employee'"`

	CreatedAt time.Time

	UpdatedAt time.Time
}

// BeforeCreate runs automatically before INSERT.
func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}
