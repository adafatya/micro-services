package handlers

import (
	"net/http"

	"github.com/adafatya/micro-services/webapi/internal/dto"
	"github.com/adafatya/micro-services/webapi/internal/services"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	OrderService *services.OrderService
}

func NewOrderHandler(orderService *services.OrderService) *OrderHandler {
	return &OrderHandler{
		OrderService: orderService,
	}
}

func (o *OrderHandler) CreateOrder(c *gin.Context) {
	userID := c.GetInt("user_id")
	var data dto.CreateOrderRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, dto.MessageResponse{Message: "Format data tidak sesuai!"})
		return
	}
	data.UserID = userID

	resp, err := o.OrderService.CreateOrder(c, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	c.JSON(http.StatusOK, resp)
}
