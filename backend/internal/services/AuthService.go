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
