package attendance

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/ravirajsahu/auth_app/internal/employee"
)

type service struct {
	repo         Repository
	employeeRepo employee.Repository
}

func NewService(
	repo Repository,
	employeeRepo employee.Repository,
) Service {
	return &service{
		repo:         repo,
		employeeRepo: employeeRepo,
	}
}

func (s *service) CheckIn(req CheckInRequest) (*AttendanceResponse, error) {

	// Employee must exist
	emp, err := s.employeeRepo.FindByID(req.EmployeeID)
	if err != nil || emp == nil {
		return nil, errors.New("employee not found")
	}

	now := time.Now()

	today := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		0, 0, 0, 0,
		now.Location(),
	)

	// Already checked in?
	existing, err := s.repo.FindTodayAttendance(req.EmployeeID, today)
	if err != nil {
		return nil, err
	}

	if existing != nil {
		return nil, errors.New("employee already checked in today")
	}

	attendance := &Attendance{
		EmployeeID: req.EmployeeID,
		Date:       today,
		CheckIn:    &now,

		CheckInLatitude:  &req.Latitude,
		CheckInLongitude: &req.Longitude,

		Status: "Present",
	}

	if err := s.repo.Create(attendance); err != nil {
		return nil, err
	}

	return s.mapAttendanceResponse(attendance, emp.User.Name), nil
}


func (s *service) CheckOut(req CheckOutRequest) (*AttendanceResponse, error) {

	now := time.Now()

	today := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		0, 0, 0, 0,
		now.Location(),
	)

	attendance, err := s.repo.FindTodayAttendance(req.EmployeeID, today)
	if err != nil {
		return nil, err
	}

	if attendance == nil {
		return nil, errors.New("check-in not found")
	}

	if attendance.CheckOut != nil {
		return nil, errors.New("already checked out")
	}

	attendance.CheckOut = &now

	attendance.CheckOutLatitude = &req.Latitude
	attendance.CheckOutLongitude = &req.Longitude

	attendance.WorkingHours =
		now.Sub(*attendance.CheckIn).Hours()

	switch {
	case attendance.WorkingHours >= 8:
		attendance.Status = "Present"

	case attendance.WorkingHours >= 4:
		attendance.Status = "Half Day"

	default:
		attendance.Status = "Absent"
	}

	if err := s.repo.Update(attendance); err != nil {
		return nil, err
	}

	emp, _ := s.employeeRepo.FindByID(req.EmployeeID)

	return s.mapAttendanceResponse(attendance, emp.User.Name), nil
}


func (s *service) GetAll() ([]AttendanceResponse, error) {

	records, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	response := make([]AttendanceResponse, 0, len(records))

	for _, item := range records {
		response = append(
			response,
			*s.mapAttendanceResponse(&item, item.Employee.User.Name),
		)
	}

	return response, nil
}


func (s *service) GetByID(id uuid.UUID) (*AttendanceResponse, error) {

	record, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return s.mapAttendanceResponse(record, record.Employee.User.Name), nil
}


func (s *service) GetByEmployee(employeeID uuid.UUID) ([]AttendanceResponse, error) {

	records, err := s.repo.FindByEmployee(employeeID)
	if err != nil {
		return nil, err
	}

	response := make([]AttendanceResponse, 0, len(records))

	for _, item := range records {
		response = append(
			response,
			*s.mapAttendanceResponse(&item, item.Employee.User.Name),
		)
	}

	return response, nil
}


func (s *service) Delete(id uuid.UUID) error {

	_, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}

func (s *service) mapAttendanceResponse(
	a *Attendance,
	name string,
) *AttendanceResponse {

	return &AttendanceResponse{
		ID: a.ID,

		EmployeeID: a.EmployeeID,

		EmployeeName: name,

		Date: a.Date,

		CheckIn: a.CheckIn,

		CheckOut: a.CheckOut,

		CheckInLatitude: a.CheckInLatitude,
		CheckInLongitude: a.CheckInLongitude,

		CheckOutLatitude: a.CheckOutLatitude,
		CheckOutLongitude: a.CheckOutLongitude,

		WorkingHours: a.WorkingHours,

		Status: a.Status,
	}
}