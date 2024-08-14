package config

import (
	"github.com/adafatya/micro-services/webapi/internal/http"
	"github.com/adafatya/micro-services/webapi/internal/http/handlers"
	"github.com/adafatya/micro-services/webapi/internal/services"
	"github.com/gin-gonic/gin"

	pb "github.com/adafatya/micro-services/webapi/internal/proto/authservice"
)

type BootstrapConfig struct {
	App               *gin.Engine
	AuthServiceClient pb.AuthServiceClient
}

func Bootstrap(config *BootstrapConfig) {
	userService := services.NewUserService(config.AuthServiceClient)

	basicHandler := handlers.NewBasicHandler()
	userHandler := handlers.NewUserHandler(userService)

	routeConfig := http.RouteConfig{
		App:          config.App,
		BasicHandler: basicHandler,
		UserHandler:  userHandler,
	}
	routeConfig.Setup()
}
