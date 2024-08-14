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

func (u *UserService) Register(ctx context.Context, data dto.UserRegisterRequest) (string, error) {
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
	return resp.GetMessage(), err
}
