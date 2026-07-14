package holiday

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	TypeNational = "National"
	TypeFestival = "Festival"
	TypeCompany  = "Company"
)

type Holiday struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	Name string `gorm:"size:150;not null"`

	Description string `gorm:"type:text"`

	Date time.Time `gorm:"type:date;not null"`

	Type string `gorm:"size:30;not null"`

	IsOptional bool `gorm:"default:false"`

	CreatedAt time.Time

	UpdatedAt time.Time

	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (h *Holiday) BeforeCreate(tx *gorm.DB) error {
	h.ID = uuid.New()
	return nil
}