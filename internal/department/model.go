package department

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Department struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	Name string `gorm:"size:100;uniqueIndex;not null"`

	Description string `gorm:"size:255"`

	Status string `gorm:"size:20;default:'Active'"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (d *Department) BeforeCreate(tx *gorm.DB) error {
	d.ID = uuid.New()
	return nil
}
