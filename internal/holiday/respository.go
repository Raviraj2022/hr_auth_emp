package holiday

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

func (r *repository) Create(holiday *Holiday) error {
	return r.db.Create(holiday).Error
}

func (r *repository) FindAll() ([]Holiday, error) {

	var holidays []Holiday

	err := r.db.
		Order("date ASC").
		Find(&holidays).Error

	if err != nil {
		return nil, err
	}

	return holidays, nil
}

func (r *repository) FindByID(id uuid.UUID) (*Holiday, error) {

	var holiday Holiday

	err := r.db.
		First(&holiday, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &holiday, nil
}

func (r *repository) FindByYear(year int) ([]Holiday, error) {

	var holidays []Holiday

	err := r.db.
		Where("EXTRACT(YEAR FROM date) = ?", year).
		Order("date ASC").
		Find(&holidays).Error

	if err != nil {
		return nil, err
	}

	return holidays, nil
}

func (r *repository) Update(holiday *Holiday) error {

	return r.db.
		Save(holiday).
		Error
}

func (r *repository) Delete(id uuid.UUID) error {

	return r.db.
		Delete(&Holiday{}, "id = ?", id).
		Error
}