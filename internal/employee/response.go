package employee

import (
	"time"

	"github.com/google/uuid"
)

type EmployeeResponse struct {
	ID uuid.UUID `json:"id"`

	UserID uuid.UUID `json:"user_id"`

	Name string `json:"name"`

	Email string `json:"email"`

	Designation string `json:"designation"`

	DepartmentID uuid.UUID `json:"department_id"`
	Department   string    `json:"department"`

	Salary float64 `json:"salary"`

	JoiningDate time.Time `json:"joining_date"`

	Status string `json:"status"`

	CreatedAt time.Time `json:"created_at"`
}
