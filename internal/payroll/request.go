package payroll

import "github.com/google/uuid"

type GeneratePayrollRequest struct {
	EmployeeID uuid.UUID `json:"employee_id" binding:"required"`

	Month int `json:"month" binding:"required"`

	Year int `json:"year" binding:"required"`

	Allowances float64 `json:"allowances"`

	Bonus float64 `json:"bonus"`

	Deductions float64 `json:"deductions"`

	Tax float64 `json:"tax"`
}