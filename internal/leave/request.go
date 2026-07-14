package leave

import (
	"time"

	"github.com/google/uuid"
)

type CreateLeaveRequest struct {
	EmployeeID uuid.UUID `json:"employee_id" binding:"required"`

	LeaveType string `json:"leave_type" binding:"required"`

	FromDate time.Time `json:"from_date" binding:"required"`

	ToDate time.Time `json:"to_date" binding:"required"`

	Reason string `json:"reason" binding:"required"`
}

type UpdateLeaveStatusRequest struct {
	Remarks string `json:"remarks"`
}
