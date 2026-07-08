package leave

import (
	"time"

	"github.com/google/uuid"
	"github.com/ravirajsahu/auth_app/internal/employee"
	"gorm.io/gorm"
)

const (
	StatusPending  = "Pending"
	StatusApproved = "Approved"
	StatusRejected = "Rejected"
	StatusCancelled = "Cancelled"
)

const (
	CasualLeave    = "Casual"
	SickLeave      = "Sick"
	EarnedLeave    = "Earned"
	MaternityLeave = "Maternity"
	PaternityLeave = "Paternity"
	LossOfPay      = "Loss Of Pay"
)

type Leave struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	EmployeeID uuid.UUID `gorm:"type:uuid;not null;index"`

	Employee employee.Employee `gorm:"foreignKey:EmployeeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	LeaveType string `gorm:"size:50;not null"`

	FromDate time.Time `gorm:"type:date;not null"`

	ToDate time.Time `gorm:"type:date;not null"`

	TotalDays int `gorm:"default:1"`

	Reason string `gorm:"type:text"`

	Status string `gorm:"size:20;default:'Pending'"`

	ApprovedBy *uuid.UUID `gorm:"type:uuid"`

	ApprovedAt *time.Time

	Remarks string `gorm:"type:text"`

	CreatedAt time.Time

	UpdatedAt time.Time

	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (l *Leave) BeforeCreate(tx *gorm.DB) error {
	l.ID = uuid.New()
	return nil
}