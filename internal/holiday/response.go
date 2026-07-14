package holiday

import (
	"time"

	"github.com/google/uuid"
)

type HolidayResponse struct {
	ID uuid.UUID `json:"id"`

	Name string `json:"name"`

	Description string `json:"description"`

	Date time.Time `json:"date"`

	Type string `json:"type"`

	IsOptional bool `json:"is_optional"`

	CreatedAt time.Time `json:"created_at"`

	UpdatedAt time.Time `json:"updated_at"`
}