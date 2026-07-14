package payroll

import "github.com/google/uuid"

type Repository interface {
	Create(payroll *Payroll) error

	FindAll() ([]Payroll, error)

	FindByID(id uuid.UUID) (*Payroll, error)

	FindByEmployee(employeeID uuid.UUID) ([]Payroll, error)

	FindByEmployeeMonthYear(employeeID uuid.UUID, month, year int) (*Payroll, error)

	Update(payroll *Payroll) error

	Delete(id uuid.UUID) error
}
