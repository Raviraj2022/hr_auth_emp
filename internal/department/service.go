package department

import (
	"errors"

	"github.com/google/uuid"
)

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

// Create Department
func (s *service) Create(req CreateDepartmentRequest) (*DepartmentResponse, error) {

	existing, _ := s.repo.FindByName(req.Name)
	if existing != nil {
		return nil, errors.New("department already exists")
	}

	department := &Department{
		Name:        req.Name,
		Description: req.Description,
		Status:      req.Status,
	}

	if department.Status == "" {
		department.Status = "Active"
	}

	if err := s.repo.Create(department); err != nil {
		return nil, err
	}

	return mapDepartmentResponse(department), nil
}

// Get All Departments
func (s *service) GetAll() ([]DepartmentResponse, error) {

	departments, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	response := make([]DepartmentResponse, 0, len(departments))

	for _, dept := range departments {
		response = append(response, *mapDepartmentResponse(&dept))
	}

	return response, nil
}

// Get Department By ID
func (s *service) GetByID(id uuid.UUID) (*DepartmentResponse, error) {

	department, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return mapDepartmentResponse(department), nil
}

// Update Department
func (s *service) Update(id uuid.UUID, req UpdateDepartmentRequest) (*DepartmentResponse, error) {

	department, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" && req.Name != department.Name {
		existing, _ := s.repo.FindByName(req.Name)
		if existing != nil {
			return nil, errors.New("department name already exists")
		}
		department.Name = req.Name
	}

	if req.Description != "" {
		department.Description = req.Description
	}

	if req.Status != "" {
		department.Status = req.Status
	}

	if err := s.repo.Update(department); err != nil {
		return nil, err
	}

	updated, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return mapDepartmentResponse(updated), nil
}

// Delete Department
func (s *service) Delete(id uuid.UUID) error {

	_, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}

// Mapper
func mapDepartmentResponse(dept *Department) *DepartmentResponse {
	return &DepartmentResponse{
		ID:          dept.ID,
		Name:        dept.Name,
		Description: dept.Description,
		Status:      dept.Status,
		CreatedAt:   dept.CreatedAt,
	}
}