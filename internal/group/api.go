package group

import (
	"github.com/ZhuoYIZIA/money-liff-api/internal/user"
	"github.com/ZhuoYIZIA/money-liff-api/pkg/log"
	"github.com/gin-gonic/gin"
)

type Route interface {
	Routes(route *gin.RouterGroup)
}

type route struct {
	groupController Controller
	userMiddleware  user.Middleware
	groupMiddleware Middleware
	logger          *log.Logger
}

func NewRoute(groupController Controller, userMiddleware user.Middleware, groupMiddleware Middleware, logger *log.Logger) Route {
	return &route{
		groupController: groupController,
		userMiddleware:  userMiddleware,
		groupMiddleware: groupMiddleware,
		logger:          logger,
	}
}

func (r *route) Routes(route *gin.RouterGroup) {

	route.GET("/group",
		r.userMiddleware.AuthCheckMiddleware(),
		r.groupController.Index)

	route.GET("/group/:group_uuid",
		r.userMiddleware.AuthCheckMiddleware(),
		r.groupMiddleware.ParamsCheckMiddleware(),
		r.groupController.Show)

	route.POST("/group",
		r.userMiddleware.AuthCheckMiddleware(),
		r.groupController.Create)

	route.PUT("/group/:group_uuid",
		r.userMiddleware.AuthCheckMiddleware(),
		r.groupMiddleware.ParamsCheckMiddleware(),
		r.groupController.Update)

	route.DELETE("/group/:group_uuid",
		r.userMiddleware.AuthCheckMiddleware(),
		r.groupMiddleware.ParamsCheckMiddleware(),
		r.groupController.Delete)

	route.GET("/group/:group_uuid/user",
		r.userMiddleware.AuthCheckMiddleware(),
		r.groupMiddleware.ParamsCheckMiddleware(),
		r.groupController.UserList)

	route.POST("/group/:group_uuid/user",
		r.userMiddleware.AuthCheckMiddleware(),
		r.groupMiddleware.ParamsCheckMiddleware(),
		r.groupController.Join)

	route.DELETE("/group/:group_uuid/user/:user_uuid",
		r.userMiddleware.AuthCheckMiddleware(),
		r.userMiddleware.ParamsCheckMiddleware(),
		r.groupMiddleware.ParamsCheckMiddleware(),
		r.groupController.DeleteUserInGroup)
}
