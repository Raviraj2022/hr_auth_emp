package shift

import "time"


type CreateShiftRequest struct{
	Name                 string         `json: "name" binding: "required"`
	StartTime            time.Time      `json: "start_time binding: "required"`
	EndTime              time.Time      `json: "end_time" binding: required`
	GraceMinutes          int           `json: "grace_minutes"`
	WorkingHours         float64        `json: "working_hours" binding: required`
	Description          string         `json: "description"`	
}

type UpdateShiftRequest struct {
	Name *string `json:"name"`

	StartTime *time.Time `json:"start_time"`

	EndTime *time.Time `json:"end_time"`

	GraceMinutes *int `json:"grace_minutes"`

	WorkingHours *float64 `json:"working_hours"`

	Description *string `json:"description"`
}