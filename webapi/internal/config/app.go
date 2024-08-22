package config

import (
	"cloud.google.com/go/storage"
	"github.com/adafatya/micro-services/webapi/internal/http"
	"github.com/adafatya/micro-services/webapi/internal/http/handlers"
	"github.com/adafatya/micro-services/webapi/internal/repository"
	"github.com/adafatya/micro-services/webapi/internal/services"
	"github.com/gin-gonic/gin"

	"github.com/adafatya/micro-services/webapi/internal/proto/authservice"
	"github.com/adafatya/micro-services/webapi/internal/proto/inventoryservice"
	"github.com/adafatya/micro-services/webapi/internal/proto/orderservice"
)

type BootstrapConfig struct {
	App                    *gin.Engine
	AuthServiceClient      authservice.AuthServiceClient
	InventoryServiceClient inventoryservice.InventoryServiceClient
	OrderServiceClient     orderservice.OrderServiceClient
	GCSBucket              *storage.BucketHandle
}

func Bootstrap(config *BootstrapConfig) {
	gcsRepository := repository.NewGCSRepository(config.GCSBucket)

	userService := services.NewUserService(config.AuthServiceClient)
	productService := services.NewProductService(config.InventoryServiceClient, gcsRepository)
	userAddressService := services.NewUserAddressService(config.OrderServiceClient)

	basicHandler := handlers.NewBasicHandler()
	userHandler := handlers.NewUserHandler(userService)
	productHandler := handlers.NewProductHandler(productService)
	userAddressHandler := handlers.NewUserAddressHandler(userAddressService)

	routeConfig := http.RouteConfig{
		App:                config.App,
		BasicHandler:       basicHandler,
		UserHandler:        userHandler,
		ProductHandler:     productHandler,
		UserAddressHandler: userAddressHandler,
	}
	routeConfig.Setup()
}
