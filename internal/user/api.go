package user

import (
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.RouterGroup) {
	route.POST("/user/register", GetUserOrRegister)
}
