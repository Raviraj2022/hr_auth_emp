package department

import "github.com/google/uuid"

type Repository interface {
	Create(department *Department) error
	FindAll() ([]Department, error)
	FindByID(id uuid.UUID) (*Department, error)
	FindByName(name string) (*Department, error)
	Update(department *Department) error
	Delete(id uuid.UUID) error 
}