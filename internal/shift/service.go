package shift

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

func (s *service) Create(req CreateShiftRequest) (*ShiftResponse, error) {

	shift := &Shift{
		Name:          req.Name,
		StartTime:     req.StartTime,
		EndTime:       req.EndTime,
		GraceMinutes:  req.GraceMinutes,
		WorkingHours:  req.WorkingHours,
		Description:   req.Description,
	}

	if err := s.repo.Create(shift); err != nil {
		return nil, err
	}

	return &ShiftResponse{
		ID:             shift.ID,
		Name:           shift.Name,
		StartTime:      shift.StartTime,
		EndTime:        shift.EndTime,
		GraceMinutes:   shift.GraceMinutes,
		WorkingHours:   shift.WorkingHours,
		Description:    shift.Description,
		CreatedAt:      shift.CreatedAt,
	}, nil
}

func (s *service) GetAll() ([]ShiftResponse, error) {

	shifts, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	response := make([]ShiftResponse, 0, len(shifts))

	for _, shift := range shifts {
		response = append(response, ShiftResponse{
			ID:             shift.ID,
			Name:           shift.Name,
			StartTime:      shift.StartTime,
			EndTime:        shift.EndTime,
			GraceMinutes:   shift.GraceMinutes,
			WorkingHours:   shift.WorkingHours,
			Description:    shift.Description,
			CreatedAt:      shift.CreatedAt,
		})
	}

	return response, nil
}

func (s *service) GetByID(id uuid.UUID) (*ShiftResponse, error) {

	shift, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return &ShiftResponse{
		ID:             shift.ID,
		Name:           shift.Name,
		StartTime:      shift.StartTime,
		EndTime:        shift.EndTime,
		GraceMinutes:   shift.GraceMinutes,
		WorkingHours:   shift.WorkingHours,
		Description:    shift.Description,
		CreatedAt:      shift.CreatedAt,
	}, nil
}

func (s *service) Update(id uuid.UUID, req UpdateShiftRequest) (*ShiftResponse, error) {

	shift, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if shift == nil {
		return nil, errors.New("shift not found")
	}

	if req.Name != nil {
		shift.Name = *req.Name
	}

	if req.StartTime != nil {
		shift.StartTime = *req.StartTime
	}

	if req.EndTime != nil {
		shift.EndTime = *req.EndTime
	}

	if req.GraceMinutes != nil {
		shift.GraceMinutes = *req.GraceMinutes
	}

	if req.WorkingHours != nil {
		shift.WorkingHours = *req.WorkingHours
	}

	if req.Description != nil {
		shift.Description = *req.Description
	}

	if err := s.repo.Update(shift); err != nil {
		return nil, err
	}

	return &ShiftResponse{
		ID:             shift.ID,
		Name:           shift.Name,
		StartTime:      shift.StartTime,
		EndTime:        shift.EndTime,
		GraceMinutes:   shift.GraceMinutes,
		WorkingHours:   shift.WorkingHours,
		Description:    shift.Description,
		CreatedAt:      shift.CreatedAt,
	}, nil
}

func (s *service) Delete(id uuid.UUID) error {

	shift, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if shift == nil {
		return errors.New("shift not found")
	}

	return s.repo.Delete(id)
}