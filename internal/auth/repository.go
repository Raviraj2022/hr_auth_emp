package auth

import (
	"errors"
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

func (r *repository) Create(user *User) error {
	return r.db.Create(user).Error
}

func (r *repository) FindByEmail(email string) (*User, error) {
	var user User

	err := r.db.Where("email = ?", email).Limit(1).Find(&user).Error
	if err != nil {
        return nil, err
    }
    
    // Check if a record was actually populated in memory
    if user.ID == uuid.Nil { 
        return nil, nil // User not found, clean return
    }

	return &user, nil
}

func (r *repository) FindByID(id uuid.UUID) (*User, error) {
	var user User

	err := r.db.First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repository) Update(user *User) error {
	return r.db.Save(user).Error
}

func (r *repository) Delete(id uuid.UUID) error {
	return r.db.Delete(&User{}, "id = ?", id).Error
}