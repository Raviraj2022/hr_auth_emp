package leave

import "github.com/google/uuid"

type Repository interface {
	Create(leave *Leave) error

	FindAll() ([]Leave, error)

	FindByID(id uuid.UUID) (*Leave, error)

	FindByEmployee(employeeID uuid.UUID) ([]Leave, error)

	Update(leave *Leave) error

	Delete(id uuid.UUID) error
}