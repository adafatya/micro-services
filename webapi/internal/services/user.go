package services

import (
	"context"

	"github.com/adafatya/micro-services/webapi/internal/dto"
	"github.com/adafatya/micro-services/webapi/internal/models"
	pb "github.com/adafatya/micro-services/webapi/internal/proto/authservice"
)

type UserService struct {
	AuthServiceClient pb.AuthServiceClient
}

func NewUserService(authServiceClient pb.AuthServiceClient) *UserService {
	return &UserService{
		AuthServiceClient: authServiceClient,
	}
}

func (u *UserService) Register(ctx context.Context, data dto.UserRegisterRequest) (dto.MessageResponse, error) {
	user := models.User{
		Email:       data.Email,
		Pass:        data.Password,
		FullName:    data.FullName,
		PhoneNumber: data.PhoneNumber,
	}

	in := &pb.RegisterRequest{
		Email:       &user.Email,
		Password:    &user.Pass,
		FullName:    &user.FullName,
		PhoneNumber: &user.PhoneNumber,
	}

	resp, err := u.AuthServiceClient.Register(ctx, in)
	return dto.MessageResponse{Message: resp.GetMessage()}, err
}

func (u *UserService) Login(ctx context.Context, data dto.UserLoginRequest) (*dto.UserLoginResponse, error) {
	in := &pb.LoginRequest{
		Email:    &data.Email,
		Password: &data.Password,
	}

	resp, err := u.AuthServiceClient.Login(ctx, in)
	return &dto.UserLoginResponse{Message: resp.GetMessage(), Token: resp.GetToken()}, err
}
