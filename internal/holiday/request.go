package holiday

import "time"

type CreateHolidayRequest struct {
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	Date        time.Time `json:"date" binding:"required"`
	Type        string    `json:"type" binding:"required,oneof=National Festival Company"`
	IsOptional  bool      `json:"is_optional"`
}

type UpdateHolidayRequest struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Type        string    `json:"type" binding:"omitempty,oneof=National Festival Company"`
	IsOptional  *bool     `json:"is_optional"`
}