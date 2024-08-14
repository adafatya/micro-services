package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BasicHandler struct {
}

func NewBasicHandler() *BasicHandler {
	return &BasicHandler{}
}

func (b *BasicHandler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong!",
	})
}
