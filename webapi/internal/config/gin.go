package config

import "github.com/gin-gonic/gin"

func NewGinApp() *gin.Engine {
	app := gin.Default()

	return app
}
