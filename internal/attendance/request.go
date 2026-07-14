package attendance

import "github.com/google/uuid"

type CheckInRequest struct {
	EmployeeID uuid.UUID `json:"employee_id" binding:"required"`
	Latitude   float64   `json:"latitude" binding:"required"`
	Longitude  float64   `json:"longitude" binding:"required"`
}

type CheckOutRequest struct {
	EmployeeID uuid.UUID `json:"employee_id" binding:"required"`
	Latitude   float64   `json:"latitude" binding:"required"`
	Longitude  float64   `json:"longitude" binding:"required"`
}
