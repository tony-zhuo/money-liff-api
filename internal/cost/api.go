package cost

import (
	"github.com/ZhuoYIZIA/money-liff-api/internal/group"
	"github.com/ZhuoYIZIA/money-liff-api/internal/user"
	"github.com/ZhuoYIZIA/money-liff-api/pkg/log"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.RouterGroup, costService Service, groupService group.Service, userService user.Service, logger *log.Logger) {
	resource := Resource{service: costService, logger: logger}

	// 新增群組付款資料
	route.POST(
		"/group/:group_uuid/cost",
		group.ParamsCheckMiddleware(groupService),
		user.AuthCheckMiddleware(userService),
		resource.Create)
}
