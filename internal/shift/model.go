package shift

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Shift struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	Name string `gorm:"size:100;not null"`

	StartTime time.Time `gorm:"type:time;not null"`
	EndTime   time.Time `gorm:"type:time;not null"`

	GraceMinutes int     `gorm:"default:0"`
	WorkingHours float64 `gorm:"type:numeric(4,2);not null"`

	Description string `gorm:"size:255"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}