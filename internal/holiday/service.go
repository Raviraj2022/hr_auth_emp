package holiday

import (
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

func (s *service) Create(req CreateHolidayRequest) (*HolidayResponse, error) {

	holiday := Holiday{
		Name:        req.Name,
		Description: req.Description,
		Date:        req.Date,
		IsOptional:  req.IsOptional,
	}

	if err := s.repo.Create(&holiday); err != nil {
		return nil, err
	}

	return &HolidayResponse{
		ID:          holiday.ID,
		Name:        holiday.Name,
		Description: holiday.Description,
		Date:        holiday.Date,
		IsOptional:  holiday.IsOptional,
		CreatedAt:   holiday.CreatedAt,
	}, nil
}

func (s *service) GetAll() ([]HolidayResponse, error) {

	holidays, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	response := make([]HolidayResponse, 0, len(holidays))

	for _, holiday := range holidays {
		response = append(response, HolidayResponse{
			ID:          holiday.ID,
			Name:        holiday.Name,
			Description: holiday.Description,
			Date:        holiday.Date,
			IsOptional:  holiday.IsOptional,
			CreatedAt:   holiday.CreatedAt,
		})
	}

	return response, nil
}

func (s *service) GetByID(id uuid.UUID) (*HolidayResponse, error) {

	holiday, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return &HolidayResponse{
		ID:          holiday.ID,
		Name:        holiday.Name,
		Description: holiday.Description,
		Date:        holiday.Date,
		IsOptional:  holiday.IsOptional,
		CreatedAt:   holiday.CreatedAt,
	}, nil
}

func (s *service) GetByYear(year int) ([]HolidayResponse, error) {

	holidays, err := s.repo.FindByYear(year)
	if err != nil {
		return nil, err
	}

	response := make([]HolidayResponse, 0, len(holidays))

	for _, holiday := range holidays {
		response = append(response, HolidayResponse{
			ID:          holiday.ID,
			Name:        holiday.Name,
			Description: holiday.Description,
			Date:        holiday.Date,
			IsOptional:  holiday.IsOptional,
			CreatedAt:   holiday.CreatedAt,
		})
	}

	return response, nil
}

func (s *service) Update(id uuid.UUID, req UpdateHolidayRequest) (*HolidayResponse, error) {

	holiday, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		holiday.Name = req.Name
	}

	if req.Description != "" {
		holiday.Description = req.Description
	}

	if !req.Date.IsZero() {
		holiday.Date = req.Date
	}

	if req.IsOptional != nil {
    holiday.IsOptional = *req.IsOptional
}

	if err := s.repo.Update(holiday); err != nil {
		return nil, err
	}

	return &HolidayResponse{
		ID:          holiday.ID,
		Name:        holiday.Name,
		Description: holiday.Description,
		Date:        holiday.Date,
		IsOptional:  holiday.IsOptional,
		CreatedAt:   holiday.CreatedAt,
	}, nil
}

func (s *service) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}