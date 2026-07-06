package auth

import (
	"errors"
	// "github.com/ravirajsahu/auth_app/config"
	"github.com/ravirajsahu/auth_app/pkg/hash"
	"github.com/ravirajsahu/auth_app/pkg/jwt"
)

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}
func (s *service) Register(req RegisterRequest) error {

	// Check if email already exists
	existingUser, _ := s.repo.FindByEmail(req.Email)
	if existingUser != nil {
		return errors.New("email already exists")
	}

	// Hash password
	hashedPassword, err := hash.HashPassword(req.Password)
	if err != nil {
		return err
	}
	user := User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}

	return s.repo.Create(&user)
}

func (s *service) Login(req LoginRequest) (*LoginResponse, error) {

	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("invalid email or password")
	}

	// Compare password
	err = hash.ComparePassword(user.Password, req.Password)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	token, err := jwt.GenerateToken(
	user.ID.String(),
	user.Email,
)

if err != nil {
	return nil, err
}
	// JWT will be added in next step
	response := &LoginResponse{
		AccessToken:  token,
		RefreshToken: "",
		User: UserResponse{
			ID:    user.ID.String(),
			Name:  user.Name,
			Email: user.Email,
		},
	}

	return response, nil
}