package employee

import (
	"time"

	"github.com/google/uuid"
	"github.com/ravirajsahu/auth_app/internal/auth"
	"github.com/ravirajsahu/auth_app/internal/department"
	"gorm.io/gorm"
)

type Employee struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	UserID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex"`

	User auth.User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	Designation string `gorm:"size:100;not null"`

	DepartmentID *uuid.UUID
	Department department.Department `gorm:"foreignKey:DepartmentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
// Department department.Department `gorm:"foreignKey:DepartmentID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`


	Salary float64 `gorm:"type:numeric(12,2);default:0"`

	JoiningDate time.Time `gorm:"not null"`

	Status string `gorm:"size:20;default:'Active'"`

	CreatedAt time.Time

	UpdatedAt time.Time

	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (e *Employee) BeforeCreate(tx *gorm.DB) error {
	e.ID = uuid.New()
	return nil
}