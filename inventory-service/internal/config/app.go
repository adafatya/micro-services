package config

import (
	"github.com/adafatya/micro-services/inventory-service/internal/repository"
	"github.com/adafatya/micro-services/inventory-service/internal/server"
	"github.com/adafatya/micro-services/inventory-service/internal/services"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	pb "github.com/adafatya/micro-services/inventory-service/internal/proto/inventoryservice"
)

type BootstrapConfig struct {
	DB         *gorm.DB
	GRPCServer *grpc.Server
}

func Bootstrap(config *BootstrapConfig) {
	productRepository := repository.NewProductRepository(config.DB)

	productService := services.NewProductService(productRepository)

	inventoryServiceServer := server.NewInventoryServiceServer(productService)

	pb.RegisterInventoryServiceServer(config.GRPCServer, inventoryServiceServer)
}
