package attendance

import (
	"time"

	"github.com/google/uuid"
)

type Repository interface {
	Create(attendance *Attendance) error

	FindByID(id uuid.UUID) (*Attendance, error)

	FindTodayAttendance(employeeID uuid.UUID, date time.Time) (*Attendance, error)

	FindByEmployee(employeeID uuid.UUID) ([]Attendance, error)

	FindAll() ([]Attendance, error)

	Update(attendance *Attendance) error

	Delete(id uuid.UUID) error
}
