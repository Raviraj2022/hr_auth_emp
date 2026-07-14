package payroll

import "github.com/google/uuid"

type GeneratePayrollRequest struct {
	EmployeeID uuid.UUID `json:"employee_id" binding:"required"`

	Month int `json:"month" binding:"gte=1,lte=12"`

	Year int `json:"year" binding:"gte=2000,lte=2100"`

	Allowances float64 `json:"allowances"`

	Bonus float64 `json:"bonus"`

	Deductions float64 `json:"deductions"`

	Tax float64 `json:"tax"`
}
