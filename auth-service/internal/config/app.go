package config

import (
	"github.com/adafatya/micro-services/auth-service/internal/repositories"
	"github.com/adafatya/micro-services/auth-service/internal/server"
	"github.com/adafatya/micro-services/auth-service/internal/services"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	pb "github.com/adafatya/micro-services/auth-service/internal/proto/authservice"
)

type BootstrapConfig struct {
	DB         *gorm.DB
	GRPCServer *grpc.Server
}

func Bootstrap(config *BootstrapConfig) {
	userRepository := repositories.NewUserRepository(config.DB)

	userService := services.NewUserService(userRepository)

	authServiceServer := server.NewAuthServiceServer(userService)

	pb.RegisterAuthServiceServer(config.GRPCServer, authServiceServer)
}
