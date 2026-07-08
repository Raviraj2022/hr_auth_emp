package employee

import "github.com/google/uuid"

type Repository interface {
	Create(employee *Employee) error

	FindAll() ([]Employee, error)

	FindByID(id uuid.UUID) (*Employee, error)

	FindByUserID(userID uuid.UUID) (*Employee, error)

	Update(employee *Employee) error

	Delete(id uuid.UUID) error
}