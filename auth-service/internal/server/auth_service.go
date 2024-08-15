package server

import (
	"context"
	"errors"

	"github.com/adafatya/micro-services/auth-service/internal/dto"
	"github.com/adafatya/micro-services/auth-service/internal/models"
	"github.com/adafatya/micro-services/auth-service/internal/services"

	pb "github.com/adafatya/micro-services/auth-service/internal/proto/authservice"
)

type AuthServiceServer struct {
	pb.UnimplementedAuthServiceServer

	UserService *services.UserService
}

func NewAuthServiceServer(userService *services.UserService) *AuthServiceServer {
	return &AuthServiceServer{
		UserService: userService,
	}
}

func (s *AuthServiceServer) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.MessageResponse, error) {
	message := "Pendaftaran berhasil! Silahkan log in"

	data := &models.User{
		Email:       in.GetEmail(),
		Pass:        in.GetPassword(),
		FullName:    in.GetFullName(),
		PhoneNumber: in.GetPhoneNumber(),
	}

	if data.Email == "" || data.Pass == "" || data.FullName == "" || data.PhoneNumber == "" {
		message = "Pendaftaran gagal! Harap isi data dengan lengkap"
		return &pb.MessageResponse{Message: &message}, errors.New("data tidak lengkap")
	}

	err := s.UserService.Register(ctx, data)
	if err != nil {
		message = "Pendaftaran gagal! " + err.Error()
		return &pb.MessageResponse{Message: &message}, err
	}

	return &pb.MessageResponse{Message: &message}, nil
}

func (s *AuthServiceServer) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	data := dto.UserLoginRequest{
		Email:    in.GetEmail(),
		Password: in.GetPassword(),
	}

	token, err := s.UserService.Login(ctx, data)
	if err != nil {
		switch token {
		case "1":
			msg := "Gagal login! Email tidak terdaftar"
			return &pb.LoginResponse{Message: &msg, Token: nil}, err
		case "2":
			msg := "Gagal login! Password salah"
			return &pb.LoginResponse{Message: &msg, Token: nil}, err
		default:
			msg := "Gagal login! Terjadi kesalahan"
			return &pb.LoginResponse{Message: &msg, Token: nil}, err
		}
	}

	msg := "Berhasil login!"
	return &pb.LoginResponse{Message: &msg, Token: &token}, nil
}
