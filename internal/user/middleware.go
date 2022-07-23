package user

import (
	"github.com/ZhuoYIZIA/money-liff-api/internal/unity/exception"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthCheckMiddleware(userService Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		lineId := c.GetHeader("Line-Id")

		if lineId == "" {
			res := exception.Unauthorized("")
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
			return
		}

		userData, err := userService.GetUserByLineId(lineId)
		if err != nil {
			res := exception.InternalServerError("")
			c.JSON(http.StatusInternalServerError, res)
			c.Abort()
			return
		}
		if userData == nil {
			res := exception.Unauthorized("")
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
			return
		}

		c.Set("auth", userData)
		c.Next()
	}
}

func ParamsCheckMiddleware(userService Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		lineId := c.Param("user_uuid")

		userData, err := userService.GetUserByLineId(lineId)
		if err != nil {
			res := exception.InternalServerError("")
			c.JSON(http.StatusInternalServerError, res)
			c.Abort()
			return
		}
		if userData == nil {
			res := exception.Unauthorized("")
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
			return
		}

		c.Set("userData", userData)
		c.Next()
	}
}
