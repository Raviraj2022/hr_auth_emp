package payroll

import (
	"time"

	"github.com/google/uuid"
	"github.com/ravirajsahu/auth_app/internal/employee"
	"gorm.io/gorm"
)

const (
	StatusPending = "Pending"
	StatusPaid    = "Paid"
)

type Payroll struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	EmployeeID uuid.UUID `gorm:"type:uuid;not null;index"`

	Employee employee.Employee `gorm:"foreignKey:EmployeeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	Month int `gorm:"not null"`

	Year int `gorm:"not null"`

	BasicSalary float64 `gorm:"type:numeric(12,2);not null"`

	Allowances float64 `gorm:"type:numeric(12,2);default:0"`

	Bonus float64 `gorm:"type:numeric(12,2);default:0"`

	Deductions float64 `gorm:"type:numeric(12,2);default:0"`

	Tax float64 `gorm:"type:numeric(12,2);default:0"`

	WorkingDays int `gorm:"default:0"`

	PresentDays int `gorm:"default:0"`

	LeaveDays int `gorm:"default:0"`

	AbsentDays int `gorm:"default:0"`

	NetSalary float64 `gorm:"type:numeric(12,2);default:0"`

	Status string `gorm:"size:20;default:'Pending'"`

	PaidAt *time.Time

	CreatedAt time.Time

	UpdatedAt time.Time

	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (p *Payroll) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.New()
	return nil
}