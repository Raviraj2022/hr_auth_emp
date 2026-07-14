package holiday

import "github.com/google/uuid"

type Service interface {
	Create(req CreateHolidayRequest) (*HolidayResponse, error)

	GetAll() ([]HolidayResponse, error)

	GetByID(id uuid.UUID) (*HolidayResponse, error)

	GetByYear(year int) ([]HolidayResponse, error)

	Update(id uuid.UUID, req UpdateHolidayRequest) (*HolidayResponse, error)

	Delete(id uuid.UUID) error
}