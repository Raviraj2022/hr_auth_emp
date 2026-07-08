package attendance

import "github.com/google/uuid"

type Service interface {
	CheckIn(req CheckInRequest) (*AttendanceResponse, error)

	CheckOut(req CheckOutRequest) (*AttendanceResponse, error)

	GetAll() ([]AttendanceResponse, error)

	GetByID(id uuid.UUID) (*AttendanceResponse, error)

	GetByEmployee(employeeID uuid.UUID) ([]AttendanceResponse, error)

	Delete(id uuid.UUID) error
}