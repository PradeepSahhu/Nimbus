package services

import (
	"fmt"

	"github.com/yashpatil74/nimbus/internal/domain/entities"
	"github.com/yashpatil74/nimbus/internal/repository"
	"github.com/yashpatil74/nimbus/internal/utils"
)

type AuthService struct {
	UserRepository *repository.UserRepository
}

func NewAuthService(userRepository *repository.UserRepository) *AuthService {
	return &AuthService{
		UserRepository: userRepository,
	}
}

func (s *AuthService) Register(username string, email string, password string) error {
	isEmailValid := utils.CheckValidEmail(email)
	if !isEmailValid {
		return fmt.Errorf("invalid email")
	}

	hashedPassword, err := utils.Hash(password, 12)
	if err != nil {
		return err
	}

	User := &entities.User{
		Username: username,
		Email:    email,
		Password: hashedPassword,
	}

	err = s.UserRepository.CreateUser(User)
	if err != nil {
		return err
	}
	return nil
}

func (s *AuthService) Login(email string, password string) (string, error) {
	user, err := s.UserRepository.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	isValid := utils.CheckHash(password, user.Password)
	if !isValid {
		return "", fmt.Errorf("invalid credentials")
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
