package services

import (
	"context"

	"github.com/adafatya/micro-services/auth-service/internal/models"
	"github.com/adafatya/micro-services/auth-service/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (u *UserService) Register(ctx context.Context, data *models.User) error {
	password, err := bcrypt.GenerateFromPassword([]byte(data.Pass), 10)
	if err != nil {
		return err
	}

	data.Pass = string(password)
	return u.UserRepository.Create(ctx, data)
}
