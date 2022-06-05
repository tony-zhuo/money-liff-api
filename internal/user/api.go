package user

import (
	"github.com/gin-gonic/gin"
)

func AuthRoutes(route *gin.RouterGroup) {
	route.POST("/user/register", GetUserOrRegister)
}
