package leave

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

func (r *repository) Create(leave *Leave) error {
	return r.db.Create(leave).Error
}

func (r *repository) FindAll() ([]Leave, error) {

	var leaves []Leave

	err := r.db.
		Preload("Employee").
		Preload("Employee.User").
		Order("created_at DESC").
		Find(&leaves).Error

	if err != nil {
		return nil, err
	}

	return leaves, nil
}

func (r *repository) FindByID(id uuid.UUID) (*Leave, error) {

	var leave Leave

	err := r.db.
		Preload("Employee").
		Preload("Employee.User").
		First(&leave, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &leave, nil
}


func (r *repository) FindByEmployee(employeeID uuid.UUID) ([]Leave, error) {

	var leaves []Leave

	err := r.db.
		Preload("Employee").
		Preload("Employee.User").
		Where("employee_id = ?", employeeID).
		Order("created_at DESC").
		Find(&leaves).Error

	if err != nil {
		return nil, err
	}

	return leaves, nil
}

func (r *repository) Update(leave *Leave) error {
	return r.db.Model(leave).Omit("Employee", "Employee.User").Updates(leave).Error
}

func (r *repository) Delete(id uuid.UUID) error {
	return r.db.Delete(&Leave{}, "id = ?", id).Error
}