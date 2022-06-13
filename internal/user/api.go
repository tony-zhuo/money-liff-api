package user

import (
	"github.com/ZhuoYIZIA/money-liff-api/pkg/log"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.RouterGroup, service Service, logger *log.Logger) {
	resource := Resource{service: service, logger: logger}

	route.POST("/user/register", resource.GetUserOrRegister)
}
