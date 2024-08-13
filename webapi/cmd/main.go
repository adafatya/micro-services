package main

import (
	"github.com/adafatya/microservices/webapi/internal/config"
)

func main() {
	app := config.NewGinApp()
	config.Bootstrap(&config.BootstrapConfig{
		App: app,
	})

	app.Run()
}
