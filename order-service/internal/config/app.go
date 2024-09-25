package config

import (
	"github.com/adafatya/micro-services/order-service/internal/repository"
	"github.com/adafatya/micro-services/order-service/internal/server"
	"github.com/adafatya/micro-services/order-service/internal/services"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	"github.com/adafatya/micro-services/order-service/internal/proto/inventoryservice"
	pb "github.com/adafatya/micro-services/order-service/internal/proto/orderservice"
)

type BootstrapConfig struct {
	DB                     *gorm.DB
	GRPCServer             *grpc.Server
	InventoryServiceClient inventoryservice.InventoryServiceClient
}

func Bootstrap(config *BootstrapConfig) {
	userAddressRepository := repository.NewUserAddressRepository(config.DB)
	orderRepository := repository.NewOrderRepository(config.DB)

	userAddressService := services.NewUserAddressService(userAddressRepository)
	orderService := services.NewOrderService(config.InventoryServiceClient, orderRepository)

	orderServiceServer := server.NewOrderServiceServer(userAddressService, orderService)

	pb.RegisterOrderServiceServer(config.GRPCServer, orderServiceServer)
}
