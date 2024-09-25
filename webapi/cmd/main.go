package main

import (
	"github.com/adafatya/micro-services/webapi/internal/config"
)

func main() {
	authServiceClient := config.NewAuthServiceClient()
	inventoryServiceClient := config.NewInventoryServiceClient()
	orderServiceClient := config.NewOrderServiceClient()

	app := config.NewGinApp()

	gcsBucket := config.NewGCSBucket()

	config.Bootstrap(&config.BootstrapConfig{
		App:                    app,
		AuthServiceClient:      authServiceClient,
		InventoryServiceClient: inventoryServiceClient,
		OrderServiceClient:     orderServiceClient,
		GCSBucket:              gcsBucket,
	})

	app.Run()
}
