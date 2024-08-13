package http

import (
	"github.com/adafatya/micro-services/webapi/internal/http/controllers"
	"github.com/gin-gonic/gin"
)

type RouteConfig struct {
	App             *gin.Engine
	BasicController *controllers.BasicController
}

func NewRoute(app *gin.Engine, basicController *controllers.BasicController) *RouteConfig {
	return &RouteConfig{
		App:             app,
		BasicController: basicController,
	}
}

func (config *RouteConfig) Setup() {
	config.App.GET("ping", config.BasicController.Ping)
}
