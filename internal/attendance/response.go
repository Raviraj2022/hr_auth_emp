package attendance

import (
	"time"

	"github.com/google/uuid"
)

type AttendanceResponse struct {
	ID uuid.UUID `json:"id"`

	EmployeeID uuid.UUID `json:"employee_id"`

	EmployeeName string `json:"employee_name"`

	Date time.Time `json:"date"`

	CheckIn *time.Time `json:"check_in"`

	CheckOut *time.Time `json:"check_out"`

	// Check-in Location
	CheckInLatitude  *float64 `json:"check_in_latitude"`
	CheckInLongitude *float64 `json:"check_in_longitude"`

	// Check-out Location
	CheckOutLatitude  *float64 `json:"check_out_latitude"`
	CheckOutLongitude *float64 `json:"check_out_longitude"`

	WorkingHours float64 `json:"working_hours"`

	Status string `json:"status"`
}
