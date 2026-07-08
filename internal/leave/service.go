package leave

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
		repo: repo,
		employeeRepo: employeeRepo,
	}
}

func (s *service) Create(req CreateLeaveRequest) (*LeaveResponse, error) {

	emp, err := s.employeeRepo.FindByID(req.EmployeeID)
	if err != nil || emp == nil {
		return nil, errors.New("employee not found")
	}

	if req.FromDate.After(req.ToDate) {
		return nil, errors.New("from_date cannot be after to_date")
	}

	totalDays := int(req.ToDate.Sub(req.FromDate).Hours()/24) + 1

	leave := &Leave{
		EmployeeID: req.EmployeeID,
		LeaveType:  req.LeaveType,
		FromDate:   req.FromDate,
		ToDate:     req.ToDate,
		TotalDays:  totalDays,
		Reason:     req.Reason,
		Status:     StatusPending,
	}

	if err := s.repo.Create(leave); err != nil {
		return nil, err
	}

	return s.mapResponse(leave, emp.User.Name), nil
}



func (s *service) GetAll() ([]LeaveResponse, error) {

	records, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	response := make([]LeaveResponse, 0, len(records))

	for _, item := range records {
		response = append(response,
			*s.mapResponse(&item, item.Employee.User.Name))
	}

	return response, nil
}

func (s *service) GetByID(id uuid.UUID) (*LeaveResponse, error) {

	record, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return s.mapResponse(record, record.Employee.User.Name), nil
}

func (s *service) GetByEmployee(employeeID uuid.UUID) ([]LeaveResponse, error) {

	records, err := s.repo.FindByEmployee(employeeID)
	if err != nil {
		return nil, err
	}

	response := make([]LeaveResponse, 0, len(records))

	for _, item := range records {
		response = append(response,
			*s.mapResponse(&item, item.Employee.User.Name))
	}

	return response, nil
}

func (s *service) Approve(
	id uuid.UUID,
	approvedBy uuid.UUID,
	req UpdateLeaveStatusRequest,
) error {

	leave, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if leave.Status != StatusPending {
		return errors.New("leave is already processed")
	}

	now := time.Now()

	leave.Status = StatusApproved
	leave.ApprovedBy = &approvedBy
	leave.ApprovedAt = &now
	leave.Remarks = req.Remarks

	return s.repo.Update(leave)
}

func (s *service) Reject(
	id uuid.UUID,
	approvedBy uuid.UUID,
	req UpdateLeaveStatusRequest,
) error {

	leave, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if leave.Status != StatusPending {
		return errors.New("leave is already processed")
	}

	now := time.Now()

	leave.Status = StatusRejected
	leave.ApprovedBy = &approvedBy
	leave.ApprovedAt = &now
	leave.Remarks = req.Remarks

	return s.repo.Update(leave)
}


func (s *service) Delete(id uuid.UUID) error {

	leave, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if leave.Status != StatusPending {
		return errors.New("only pending leave can be cancelled")
	}

	return s.repo.Delete(id)
}

func (s *service) mapResponse(
	leave *Leave,
	name string,
) *LeaveResponse {

	return &LeaveResponse{
		ID: leave.ID,

		EmployeeID: leave.EmployeeID,

		EmployeeName: name,

		LeaveType: leave.LeaveType,

		FromDate: leave.FromDate,

		ToDate: leave.ToDate,

		TotalDays: leave.TotalDays,

		Reason: leave.Reason,

		Status: leave.Status,

		ApprovedBy: leave.ApprovedBy,

		ApprovedAt: leave.ApprovedAt,

		Remarks: leave.Remarks,

		CreatedAt: leave.CreatedAt,
	}
}