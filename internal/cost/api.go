package cost

import (
	"github.com/ZhuoYIZIA/money-liff-api/internal/group"
	"github.com/ZhuoYIZIA/money-liff-api/internal/user"
	"github.com/ZhuoYIZIA/money-liff-api/pkg/log"
	"github.com/gin-gonic/gin"
)

type Route interface {
	Routes(route *gin.RouterGroup)
}

type route struct {
	controller      Controller
	groupMiddleware group.Middleware
	userMiddleware  user.Middleware
	logger          *log.Logger
}

func NewRoutes(controller Controller, groupMiddleware group.Middleware, userMiddleware user.Middleware,
	logger *log.Logger) Route {
	return &route{
		controller:      controller,
		groupMiddleware: groupMiddleware,
		userMiddleware:  userMiddleware,
		logger:          logger,
	}
}

func (r *route) Routes(route *gin.RouterGroup) {
	// 新增群組付款資料
	route.POST(
		"/group/:group_uuid/cost",
		r.groupMiddleware.ParamsCheckMiddleware(),
		r.userMiddleware.AuthCheckMiddleware(),
		r.controller.Create)
}
