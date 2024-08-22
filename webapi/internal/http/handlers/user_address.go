package handlers

import (
	"net/http"

	"github.com/adafatya/micro-services/webapi/internal/dto"
	"github.com/adafatya/micro-services/webapi/internal/services"
	"github.com/gin-gonic/gin"
)

type UserAddressHandler struct {
	UserAddressService *services.UserAddressService
}

func NewUserAddressHandler(userAddressService *services.UserAddressService) *UserAddressHandler {
	return &UserAddressHandler{
		UserAddressService: userAddressService,
	}
}

func (u *UserAddressHandler) AddUserAddress(c *gin.Context) {
	userID := c.GetInt("user_id")

	var data dto.AddUserAddressRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, dto.AddUserAddressResponse{Message: "Format data tidak sesuai!"})
		return
	}

	data.UserID = userID
	resp, err := u.UserAddressService.AddUserAddress(c, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	c.JSON(http.StatusOK, resp)
}
