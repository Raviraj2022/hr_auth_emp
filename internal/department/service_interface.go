package department

import "github.com/google/uuid"

type Service interface {
	Create(req CreateDepartmentRequest) (*DepartmentResponse, error)

	GetAll() ([]DepartmentResponse, error)

	GetByID(id uuid.UUID) (*DepartmentResponse, error)

	Update(id uuid.UUID, req UpdateDepartmentRequest) (*DepartmentResponse, error)

	Delete(id uuid.UUID) error
}
