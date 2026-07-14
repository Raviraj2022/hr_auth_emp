package attendance

import (
	"time"

	"github.com/google/uuid"
	"github.com/ravirajsahu/auth_app/internal/employee"
	"gorm.io/gorm"
)

type Attendance struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	EmployeeID uuid.UUID `gorm:"type:uuid;not null;index"`

	Employee employee.Employee `gorm:"foreignKey:EmployeeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	Date time.Time `gorm:"type:date;not null"`

	CheckIn *time.Time

	CheckOut *time.Time

	WorkingHours float64 `gorm:"type:numeric(5,2);default:0"`

	Status           string `gorm:"size:20;default:'Present'"`
	CheckInLatitude  *float64
	CheckInLongitude *float64

	CheckOutLatitude  *float64
	CheckOutLongitude *float64

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (a *Attendance) BeforeCreate(tx *gorm.DB) error {
	a.ID = uuid.New()
	return nil
}
