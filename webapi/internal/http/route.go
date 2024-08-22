package http

import (
	"github.com/adafatya/micro-services/webapi/internal/http/handlers"
	"github.com/adafatya/micro-services/webapi/internal/http/middleware"
	"github.com/gin-gonic/gin"
)

type RouteConfig struct {
	App                *gin.Engine
	BasicHandler       *handlers.BasicHandler
	UserHandler        *handlers.UserHandler
	ProductHandler     *handlers.ProductHandler
	UserAddressHandler *handlers.UserAddressHandler
}

func (config *RouteConfig) Setup() {
	config.App.GET("ping", config.BasicHandler.Ping)

	v1 := config.App.Group("api/v1")
	v1.POST("register", config.UserHandler.Register)
	v1.POST("login", config.UserHandler.Login)

	login := v1
	login.Use(middleware.LoggedIn())
	login.POST("user/address", config.UserAddressHandler.AddUserAddress)
	login.GET("user/addresses", config.UserAddressHandler.GetUserAddresses)

	admin := v1.Group("admin")
	admin.Use(middleware.AdminOnly())
	admin.POST("product", config.ProductHandler.AddProduct)
}
