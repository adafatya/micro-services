package main

import (
	"github.com/adafatya/micro-services/webapi/internal/config"
)

func main() {
	authServiceClient := config.NewAuthServiceClient()

	app := config.NewGinApp()

	config.Bootstrap(&config.BootstrapConfig{
		App:               app,
		AuthServiceClient: authServiceClient,
	})

	app.Run()
}
