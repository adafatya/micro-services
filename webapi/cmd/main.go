package main

import (
	"github.com/adafatya/micro-services/webapi/internal/config"
)

func main() {
	authServiceClient := config.NewAuthServiceClient()
	inventoryServiceClient := config.NewInventoryServiceClient()

	app := config.NewGinApp()

	gcsBucket := config.NewGCSBucket()

	config.Bootstrap(&config.BootstrapConfig{
		App:                    app,
		AuthServiceClient:      authServiceClient,
		InventoryServiceClient: inventoryServiceClient,
		GCSBucket:              gcsBucket,
	})

	app.Run()
}
