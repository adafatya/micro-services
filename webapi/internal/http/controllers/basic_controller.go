package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BasicController struct {
}

func NewBasicController() *BasicController {
	return &BasicController{}
}

func (b *BasicController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong!",
	})
}
