package user

import (
	"net/http"

	"github.com/ZhuoYIZIA/money-liff-api/internal/unity/exception"
	"github.com/gin-gonic/gin"
)

type Middleware interface {
	AuthCheckMiddleware() func(c *gin.Context)
	ParamsCheckMiddleware() func(c *gin.Context)
}

type middleware struct {
	userService Service
}

func NewMiddleware(userService Service) Middleware {
	return &middleware{
		userService: userService,
	}
}

func (m *middleware) AuthCheckMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		lineId := c.GetHeader("Line-Id")

		if lineId == "" {
			res := exception.Unauthorized("")
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
			return
		}

		userData, err := m.userService.GetUserByLineId(lineId)
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

func (m *middleware) ParamsCheckMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		lineId := c.Param("user_uuid")

		userData, err := m.userService.GetUserByLineId(lineId)
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
