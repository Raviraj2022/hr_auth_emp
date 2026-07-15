package shift

import (
	"time"

	"github.com/google/uuid"
)

type ShiftResponse struct {
	ID uuid.UUID `json:"id"`

	Name string `json:"name"`

	StartTime time.Time `json:"start_time"`

	EndTime time.Time `json:"end_time"`

	GraceMinutes int `json:"grace_minutes"`

	WorkingHours float64 `json:"working_hours"`

	Description string `json:"description"`

	CreatedAt time.Time `json:"created_at"`
}