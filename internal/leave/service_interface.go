package leave

import "github.com/google/uuid"

type Service interface {
	Create(req CreateLeaveRequest) (*LeaveResponse, error)

	GetAll() ([]LeaveResponse, error)

	GetByID(id uuid.UUID) (*LeaveResponse, error)

	GetByEmployee(employeeID uuid.UUID) ([]LeaveResponse, error)

	Approve(id uuid.UUID, approvedBy uuid.UUID, req UpdateLeaveStatusRequest) error

	Reject(id uuid.UUID, approvedBy uuid.UUID, req UpdateLeaveStatusRequest) error

	Delete(id uuid.UUID) error
}
