package payroll

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

func (r *repository) Create(payroll *Payroll) error {
	return r.db.Create(payroll).Error
}

func (r *repository) FindAll() ([]Payroll, error) {

	var payrolls []Payroll

	err := r.db.
		Preload("Employee").
		Preload("Employee.User").
		Order("year DESC, month DESC").
		Find(&payrolls).Error

	if err != nil {
		return nil, err
	}

	return payrolls, nil
}

func (r *repository) FindByID(id uuid.UUID) (*Payroll, error) {

	var payroll Payroll

	err := r.db.
		Preload("Employee").
		Preload("Employee.User").
		First(&payroll, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &payroll, nil
}

func (r *repository) FindByEmployee(employeeID uuid.UUID) ([]Payroll, error) {

	var payrolls []Payroll

	err := r.db.
		Preload("Employee").
		Preload("Employee.User").
		Where("employee_id = ?", employeeID).
		Order("year DESC, month DESC").
		Find(&payrolls).Error

	if err != nil {
		return nil, err
	}

	return payrolls, nil
}

func (r *repository) FindByEmployeeMonthYear(
	employeeID uuid.UUID,
	month int,
	year int,
) (*Payroll, error) {

	var payroll Payroll

	err := r.db.
		Where(
			"employee_id = ? AND month = ? AND year = ?",
			employeeID,
			month,
			year,
		).
		First(&payroll).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &payroll, nil
}

func (r *repository) Update(payroll *Payroll) error {
	return r.db.Save(payroll).Error
}

func (r *repository) Delete(id uuid.UUID) error {
	return r.db.Delete(&Payroll{}, "id = ?", id).Error
}