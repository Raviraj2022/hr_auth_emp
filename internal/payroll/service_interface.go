package payroll

import "github.com/google/uuid"

type Service interface {
	Generate(req GeneratePayrollRequest) (*PayrollResponse, error)

	GetAll() ([]PayrollResponse, error)

	GetByID(id uuid.UUID) (*PayrollResponse, error)

	GetByEmployee(employeeID uuid.UUID) ([]PayrollResponse, error)

	MarkPaid(id uuid.UUID) error

	Delete(id uuid.UUID) error
}
