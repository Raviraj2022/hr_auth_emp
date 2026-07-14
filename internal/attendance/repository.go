package attendance

import (
	"time"

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

// Create Attendance
func (r *repository) Create(attendance *Attendance) error {
	return r.db.Create(attendance).Error
}

// Find Attendance By ID
func (r *repository) FindByID(id uuid.UUID) (*Attendance, error) {

	var attendance Attendance

	err := r.db.
		Preload("Employee").
		Preload("Employee.User").
		First(&attendance, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &attendance, nil
}

// Find Today's Attendance
func (r *repository) FindTodayAttendance(employeeID uuid.UUID, date time.Time) (*Attendance, error) {

	var attendance Attendance

	err := r.db.
		Where("employee_id = ? AND date = ?", employeeID, date).
		First(&attendance).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &attendance, nil
}

// Find Attendance By Employee
func (r *repository) FindByEmployee(employeeID uuid.UUID) ([]Attendance, error) {

	var attendances []Attendance

	err := r.db.
		Preload("Employee").
		Preload("Employee.User").
		Where("employee_id = ?", employeeID).
		Order("date DESC").
		Find(&attendances).Error

	if err != nil {
		return nil, err
	}

	return attendances, nil
}

// Find All Attendance
func (r *repository) FindAll() ([]Attendance, error) {

	var attendances []Attendance

	err := r.db.
		Preload("Employee").
		Preload("Employee.User").
		Order("date DESC").
		Find(&attendances).Error

	if err != nil {
		return nil, err
	}

	return attendances, nil
}

// Update Attendance
func (r *repository) Update(attendance *Attendance) error {
	return r.db.Save(attendance).Error
}

// Delete Attendance
func (r *repository) Delete(id uuid.UUID) error {
	return r.db.Delete(&Attendance{}, "id = ?", id).Error
}
