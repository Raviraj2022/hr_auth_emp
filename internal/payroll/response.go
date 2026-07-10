package payroll

import (
	"time"

	"github.com/google/uuid"
)

type PayrollResponse struct {
	ID uuid.UUID `json:"id"`

	EmployeeID uuid.UUID `json:"employee_id"`

	EmployeeName string `json:"employee_name"`

	Month int `json:"month"`

	Year int `json:"year"`

	BasicSalary float64 `json:"basic_salary"`

	Allowances float64 `json:"allowances"`

	Bonus float64 `json:"bonus"`

	Deductions float64 `json:"deductions"`

	Tax float64 `json:"tax"`

	WorkingDays int `json:"working_days"`

	PresentDays int `json:"present_days"`

	LeaveDays int `json:"leave_days"`

	AbsentDays int `json:"absent_days"`

	NetSalary float64 `json:"net_salary"`

	Status string `json:"status"`

	PaidAt *time.Time `json:"paid_at,omitempty"`

	CreatedAt time.Time `json:"created_at"`
}