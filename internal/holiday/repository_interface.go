package holiday

import "github.com/google/uuid"

type Repository interface {
	Create(holiday *Holiday) error

	FindAll() ([]Holiday, error)

	FindByID(id uuid.UUID) (*Holiday, error)

	FindByYear(year int) ([]Holiday, error)

	Update(holiday *Holiday) error

	Delete(id uuid.UUID) error
}