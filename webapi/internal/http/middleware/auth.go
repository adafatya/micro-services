package middleware

import (
	"net/http"
	"strings"

	"github.com/adafatya/micro-services/webapi/internal/dto"
	"github.com/adafatya/micro-services/webapi/pkg/util"
	"github.com/gin-gonic/gin"
)

func LoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		bearer := strings.Split(authHeader, " ")
		if len(bearer) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.MessageResponse{Message: "Kesalahan format token"})

			return
		}

		userID, err := util.GetUserID(bearer[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.MessageResponse{Message: "Anda belum login"})
			return
		}

		c.Set("user_id", userID)

		c.Next()
	}
}

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		bearer := strings.Split(authHeader, " ")
		if len(bearer) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.MessageResponse{Message: "Kesalahan format token"})
			return
		}

		userID, err := util.GetUserID(bearer[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.MessageResponse{Message: "Anda belum login"})
			return
		}

		if userID != 1 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.MessageResponse{Message: "Anda bukan admin"})
			return
		}

		c.Set("user_id", userID)

		c.Next()
	}
}
