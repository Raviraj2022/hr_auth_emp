package shift

import "github.com/google/uuid"

type Service interface {
	Create(req CreateShiftRequest) (*ShiftResponse, error)

	GetAll() ([]ShiftResponse, error)

	GetByID(id uuid.UUID) (*ShiftResponse, error)

	Update(id uuid.UUID, req UpdateShiftRequest) (*ShiftResponse, error)

	Delete(id uuid.UUID) error
}