package department

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(department *Department) error {
	return r.db.Create(department).Error
}

func (r *repository) FindAll() ([]Department, error) {

	var departments []Department

	err := r.db.Find(&departments).Error
	if err != nil {
		return nil, err
	}

	return departments, nil
}

func (r *repository) FindByID(id uuid.UUID) (*Department, error) {

	var department Department

	err := r.db.First(&department, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &department, nil
}

func (r *repository) FindByName(name string) (*Department, error) {

	var department Department

	err := r.db.Where("LOWER(name) = LOWER(?)", name).
		First(&department).Error

	if err != nil {
		return nil, err
	}

	return &department, nil
}

func (r *repository) Update(department *Department) error {

	return r.db.Model(department).Updates(map[string]interface{}{
		"name":        department.Name,
		"description": department.Description,
		"status":      department.Status,
	}).Error
}

func (r *repository) Delete(id uuid.UUID) error {

	return r.db.Delete(&Department{}, "id = ?", id).Error
}
