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
	Designation string `json:"designation"`

	DepartmentID uuid.UUID `json:"department_id"`

	Salary float64 `json:"salary"`

	Status string `json:"status"`
}