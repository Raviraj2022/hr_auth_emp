package employee

import "github.com/google/uuid"

type Service interface {
	Create(req CreateEmployeeRequest) (*EmployeeResponse, error)

	GetAll() ([]EmployeeResponse, error)

	GetByID(id uuid.UUID) (*EmployeeResponse, error)

	Update(id uuid.UUID, req UpdateEmployeeRequest) (*EmployeeResponse, error)

	Delete(id uuid.UUID) error
}
