package services

import (
	"context"

	"github.com/adafatya/micro-services/auth-service/internal/dto"
	"github.com/adafatya/micro-services/auth-service/internal/models"
	"github.com/adafatya/micro-services/auth-service/internal/repositories"
	"github.com/adafatya/micro-services/auth-service/pkg/util"
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

func (u *UserService) Login(ctx context.Context, data dto.UserLoginRequest) (string, error) {
	user, err := u.UserRepository.FindOne(ctx, "id, email, pass", "email = ?", data.Email)
	if err != nil {
		return "1", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Pass), []byte(data.Password)); err != nil {
		return "2", err
	}

	token, err := util.CreateJWTToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
