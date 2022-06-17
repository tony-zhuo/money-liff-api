package user

import (
	"github.com/ZhuoYIZIA/money-liff-api/pkg/log"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.RouterGroup, userService Service, logger *log.Logger) {
	resource := Resource{service: userService, logger: logger}

	route.POST("/user/register", resource.GetUserOrRegister)
}
