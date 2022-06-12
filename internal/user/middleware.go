package user

import (
	"github.com/ZhuoYIZIA/money-liff-api/internal/unity/exception"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthCheckMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		lineId := c.GetHeader("Line-Id")
		userService := NewService()

		if lineId == "" {
			res := exception.Unauthorized("")
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
			return
		}

		userData := userService.GetUserByLineId(lineId)
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