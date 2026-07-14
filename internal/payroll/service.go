package payroll

import (
	"errors"
	"time"

	//  "fmt"
	"github.com/google/uuid"

	"github.com/ravirajsahu/auth_app/internal/attendance"
	"github.com/ravirajsahu/auth_app/internal/employee"
	"github.com/ravirajsahu/auth_app/internal/leave"
)

type service struct {
	repo Repository

	employeeRepo employee.Repository

	attendanceRepo attendance.Repository

	leaveRepo leave.Repository
}

func NewService(
	repo Repository,
	employeeRepo employee.Repository,
	attendanceRepo attendance.Repository,
	leaveRepo leave.Repository,
) Service {

	return &service{
		repo: repo,

		employeeRepo: employeeRepo,

		attendanceRepo: attendanceRepo,

		leaveRepo: leaveRepo,
	}
}

func (s *service) Generate(req GeneratePayrollRequest) (*PayrollResponse, error) {

	// Employee exists?
	emp, err := s.employeeRepo.FindByID(req.EmployeeID)
	if err != nil || emp == nil {
		return nil, errors.New("employee not found")
	}

	// Already generated?
	existing, err := s.repo.FindByEmployeeMonthYear(
		req.EmployeeID,
		req.Month,
		req.Year,
	)

	// fmt.Println(existing)

	if err != nil {
		return nil, err
	}

	if existing != nil {
		return nil, errors.New("payroll already generated")
	}

	// ---------------------------------------------------
	// Temporary values.
	// Later we'll calculate these from Attendance & Leave.
	// ---------------------------------------------------

	workingDays := 26

	presentDays := 26

	leaveDays := 0

	absentDays := workingDays - presentDays - leaveDays

	basicSalary := emp.Salary

	perDaySalary := basicSalary / float64(workingDays)

	lop := float64(absentDays) * perDaySalary

	netSalary := basicSalary +
		req.Allowances +
		req.Bonus -
		req.Deductions -
		req.Tax -
		lop

	payroll := &Payroll{

		EmployeeID: req.EmployeeID,

		Month: req.Month,

		Year: req.Year,

		BasicSalary: basicSalary,

		Allowances: req.Allowances,

		Bonus: req.Bonus,

		Deductions: req.Deductions,

		Tax: req.Tax,

		WorkingDays: workingDays,

		PresentDays: presentDays,

		LeaveDays: leaveDays,

		AbsentDays: absentDays,

		NetSalary: netSalary,

		Status: StatusPending,
	}

	if err := s.repo.Create(payroll); err != nil {
		return nil, err
	}

	return s.mapResponse(payroll, emp.User.Name), nil
}

func (s *service) GetAll() ([]PayrollResponse, error) {

	items, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	response := make([]PayrollResponse, 0, len(items))

	for _, item := range items {

		response = append(
			response,
			*s.mapResponse(&item, item.Employee.User.Name),
		)
	}

	return response, nil
}

func (s *service) GetByID(id uuid.UUID) (*PayrollResponse, error) {

	item, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return s.mapResponse(item, item.Employee.User.Name), nil
}

func (s *service) GetByEmployee(employeeID uuid.UUID) ([]PayrollResponse, error) {

	items, err := s.repo.FindByEmployee(employeeID)
	if err != nil {
		return nil, err
	}

	response := make([]PayrollResponse, 0, len(items))

	for _, item := range items {

		response = append(
			response,
			*s.mapResponse(&item, item.Employee.User.Name),
		)
	}

	return response, nil
}

func (s *service) MarkPaid(id uuid.UUID) error {
	//  fmt.Println(id)
	payroll, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	// fmt.Println(payroll)

	if payroll.Status == StatusPaid {
		return errors.New("salary already paid")
	}

	now := time.Now()

	payroll.Status = StatusPaid

	payroll.PaidAt = &now

	return s.repo.Update(payroll)
}

func (s *service) Delete(id uuid.UUID) error {

	return s.repo.Delete(id)
}

func (s *service) mapResponse(
	p *Payroll,
	name string,
) *PayrollResponse {

	return &PayrollResponse{

		ID: p.ID,

		EmployeeID: p.EmployeeID,

		EmployeeName: name,

		Month: p.Month,

		Year: p.Year,

		BasicSalary: p.BasicSalary,

		Allowances: p.Allowances,

		Bonus: p.Bonus,

		Deductions: p.Deductions,

		Tax: p.Tax,

		WorkingDays: p.WorkingDays,

		PresentDays: p.PresentDays,

		LeaveDays: p.LeaveDays,

		AbsentDays: p.AbsentDays,

		NetSalary: p.NetSalary,

		Status: p.Status,

		PaidAt: p.PaidAt,

		CreatedAt: p.CreatedAt,
	}
}
