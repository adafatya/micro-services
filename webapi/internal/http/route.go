package http

import (
	"github.com/adafatya/micro-services/webapi/internal/http/handlers"
	"github.com/gin-gonic/gin"
)

type RouteConfig struct {
	App          *gin.Engine
	BasicHandler *handlers.BasicHandler
	UserHandler  *handlers.UserHandler
}

func (config *RouteConfig) Setup() {
	config.App.GET("ping", config.BasicHandler.Ping)

	v1 := config.App.Group("api/v1")
	v1.POST("register", config.UserHandler.Register)
}
