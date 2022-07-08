package cost

import (
	"github.com/ZhuoYIZIA/money-liff-api/pkg/log"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.RouterGroup, costService Service, logger *log.Logger) {
	resource := Resource{service: costService, logger: logger}
	route.POST("/cost", resource.Create)
}
