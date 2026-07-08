package employee

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

// Create Employee
func (r *repository) Create(employee *Employee) error {
	return r.db.Create(employee).Error
}

// Get All Employees
func (r *repository) FindAll() ([]Employee, error) {

	var employees []Employee

	err := r.db.
		Preload("User").
		Preload("Department").
		Find(&employees).Error

	if err != nil {
		return nil, err
	}

	return employees, nil
}

// Get Employee By ID
func (r *repository) FindByID(id uuid.UUID) (*Employee, error) {

	var employee Employee

	err := r.db.
		Preload("User").
		Preload("Department").
		First(&employee, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &employee, nil
}

// Get Employee By User ID
func (r *repository) FindByUserID(userID uuid.UUID) (*Employee, error) {

	var employee Employee

	err := r.db.
		Preload("User").
		Preload("Department").
		First(&employee, "user_id = ?", userID).Error

	if err != nil {
		return nil, err
	}

	return &employee, nil
}

// Update Employee
func (r *repository) Update(employee *Employee) error {

	// return r.db.Save(employee).Error
	// return r.db.Session(&gorm.Session{FullSaveAssociations: false}).Save(employee).Error
	return r.db.Model(employee).Omit("User").Updates(employee).Error
}

// Delete Employee (Soft Delete)
func (r *repository) Delete(id uuid.UUID) error {

	return r.db.Delete(&Employee{}, "id = ?", id).Error
}