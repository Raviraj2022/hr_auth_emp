package employee

import (
	"errors"
    "fmt"
	"github.com/google/uuid"
)

type service struct {
	repo Repository
	departmentRepo department.Repository
}

func NewService(repo Repository, departmentRepo department.Repository,) Service {
	return &service{
		repo: repo,
		departmentRepo: departmentRepo,
	}
}

// Create Employee
func (s *service) Create(req CreateEmployeeRequest) (*EmployeeResponse, error) {

	dept, err := s.departmentRepo.FindByID(req.DepartmentID)
if err != nil {
	return nil, errors.New("department not found")
}

if dept == nil {
	return nil, errors.New("department not found")
}
	// Check if employee already exists for this user
	existing, _ := s.repo.FindByUserID(req.UserID)
	if existing != nil {
		return nil, errors.New("employee already exists for this user")
	}

	employee := Employee{
		UserID:       req.UserID,
		Designation:  req.Designation,
		DepartmentID: req.DepartmentID,
		Salary:       req.Salary,
		JoiningDate:  req.JoiningDate,
		Status:       req.Status,
	}

	if employee.Status == "" {
		employee.Status = "Active"
	}

	if err := s.repo.Create(&employee); err != nil {
		return nil, err
	}

	created, err := s.repo.FindByID(employee.ID)
	if err != nil {
		return nil, err
	}

	return mapEmployeeResponse(created), nil
}

// Get All Employees
func (s *service) GetAll() ([]EmployeeResponse, error) {

	employees, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	response := make([]EmployeeResponse, 0, len(employees))

	for _, emp := range employees {
		response = append(response, *mapEmployeeResponse(&emp))
	}

	return response, nil
}

// Get Employee By ID
func (s *service) GetByID(id uuid.UUID) (*EmployeeResponse, error) {

	employee, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return mapEmployeeResponse(employee), nil
}

// Update Employee
func (s *service) Update(id uuid.UUID, req UpdateEmployeeRequest) (*EmployeeResponse, error) {

	employee, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Request: %+v\n", employee)

	if req.Designation != "" {
		employee.Designation = req.Designation
	}

	if req.DepartmentID != nil {
    employee.DepartmentID = req.DepartmentID
}

	if req.Salary != 0 {
		employee.Salary = req.Salary
	}

	if req.Status != "" {
		employee.Status = req.Status
	}

	if err := s.repo.Update(employee); err != nil {
		return nil, err
	}

	updated, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return mapEmployeeResponse(updated), nil
}

// Delete Employee
func (s *service) Delete(id uuid.UUID) error {

	_, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}

// Mapper
func mapEmployeeResponse(emp *Employee) *EmployeeResponse {

	return &EmployeeResponse{
		ID:           emp.ID,
		UserID:       emp.UserID,
		Name:         emp.User.Name,
		Email:        emp.User.Email,
		Designation:  emp.Designation,
		DepartmentID: emp.DepartmentID,
		Department:   emp.Department.Name,
		Salary:       emp.Salary,
		JoiningDate:  emp.JoiningDate,
		Status:       emp.Status,
		CreatedAt:    emp.CreatedAt,
	}
}