package leave

import (
	"time"

	"github.com/google/uuid"
)

type LeaveResponse struct {
	ID uuid.UUID `json:"id"`

	EmployeeID uuid.UUID `json:"employee_id"`

	EmployeeName string `json:"employee_name"`

	LeaveType string `json:"leave_type"`

	FromDate time.Time `json:"from_date"`

	ToDate time.Time `json:"to_date"`

	TotalDays int `json:"total_days"`

	Reason string `json:"reason"`

	Status string `json:"status"`

	ApprovedBy *uuid.UUID `json:"approved_by,omitempty"`

	ApprovedAt *time.Time `json:"approved_at,omitempty"`

	Remarks string `json:"remarks"`

	CreatedAt time.Time `json:"created_at"`
}
