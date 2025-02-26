package services

import "github.com/yashpatil74/nimbus/internal/repository"

type AuthService struct {
	UserRepository *repository.UserRepository
}

func NewAuthService(userRepository *repository.UserRepository) *AuthService {
	return &AuthService{
		UserRepository: userRepository,
	}
}

func (s *AuthService) Login(email, password string) (string, error) {
	return "", nil
}
