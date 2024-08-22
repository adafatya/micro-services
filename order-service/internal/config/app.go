package config

import (
	"github.com/adafatya/micro-services/order-service/internal/repository"
	"github.com/adafatya/micro-services/order-service/internal/server"
	"github.com/adafatya/micro-services/order-service/internal/services"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	pb "github.com/adafatya/micro-services/order-service/internal/proto/orderservice"
)

type BootstrapConfig struct {
	DB         *gorm.DB
	GRPCServer *grpc.Server
}

func Bootstrap(config *BootstrapConfig) {
	userAddressRepository := repository.NewUserAddressRepository(config.DB)

	userAddressService := services.NewUserAddressService(userAddressRepository)

	orderServiceServer := server.NewOrderServiceServer(userAddressService)

	pb.RegisterOrderServiceServer(config.GRPCServer, orderServiceServer)
}
