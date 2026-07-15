package shift

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

func (r *repository) Create(shift *Shift) error {
	return r.db.Create(shift).Error
}

func (r *repository) FindAll() ([]Shift, error) {

	var shifts []Shift

	err := r.db.
		Order("created_at DESC").
		Find(&shifts).Error

	if err != nil {
		return nil, err
	}

	return shifts, nil
}

func (r *repository) FindByID(id uuid.UUID) (*Shift, error) {

	var shift Shift

	err := r.db.
		First(&shift, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &shift, nil
}

func (r *repository) Update(shift *Shift) error {
	return r.db.Save(shift).Error
}

func (r *repository) Delete(id uuid.UUID) error {
	return r.db.Delete(&Shift{}, "id = ?", id).Error
}