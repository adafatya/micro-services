package config

import (
	"github.com/adafatya/micro-services/webapi/internal/http"
	"github.com/adafatya/micro-services/webapi/internal/http/controllers"
	"github.com/gin-gonic/gin"
)

type BootstrapConfig struct {
	App *gin.Engine
}

func Bootstrap(config *BootstrapConfig) {
	basicController := controllers.NewBasicController()

	routeConfig := http.RouteConfig{
		App:             config.App,
		BasicController: basicController,
	}
	routeConfig.Setup()
}
