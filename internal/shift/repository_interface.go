package shift

import "github.com/google/uuid"

type Repository interface {
	Create(shift *Shift) error

	FindAll() ([]Shift, error)

	FindByID(id uuid.UUID) (*Shift, error)

	Update(shift *Shift) error

	Delete(id uuid.UUID) error
}