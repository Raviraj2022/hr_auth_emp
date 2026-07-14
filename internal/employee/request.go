package employee

import (
	"time"

	"github.com/google/uuid"
)

type CreateEmployeeRequest struct {
	UserID uuid.UUID `json:"user_id" binding:"required"`

	Designation string `json:"designation" binding:"required"`

	DepartmentID uuid.UUID `json:"department_id" binding:"required"`

	Salary float64 `json:"salary"`

	JoiningDate time.Time `json:"joining_date" binding:"required"`

	Status string `json:"status"`
}

type UpdateEmployeeRequest struct {
	UserID       uuid.UUID `json:"user_id"` // Added
	Designation  string    `json:"designation"`
	DepartmentID uuid.UUID `json:"department_id"`
	Salary       float64   `json:"salary"`
	JoiningDate  string    `json:"joining_date"` // Added (or use time.Time)
	Status       string    `json:"status"`
}
