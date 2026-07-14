package auth

type Service interface {
	Register(req RegisterRequest) error
	Login(req LoginRequest) (*LoginResponse, error)
}
