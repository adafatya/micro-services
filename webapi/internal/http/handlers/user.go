package handlers

import (
	"log"
	"net/http"

	"github.com/adafatya/micro-services/webapi/internal/dto"
	"github.com/adafatya/micro-services/webapi/internal/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (u *UserHandler) Register(c *gin.Context) {
	var data dto.UserRegisterRequest

	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, dto.UserRegisterResponse{
			Message: "Data tidak memenuhi format!",
		})
		return
	}

	msg, err := u.UserService.Register(c, data)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, dto.UserRegisterResponse{
			Message: msg,
		})
		return
	}

	c.JSON(http.StatusOK, dto.UserRegisterResponse{
		Message: msg,
	})
}
