package group

import (
	"github.com/ZhuoYIZIA/money-liff-api/internal/user"
	"github.com/ZhuoYIZIA/money-liff-api/pkg/log"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.RouterGroup, groupService Service, userService user.Service, logger *log.Logger) {
	resource := Resource{service: groupService, logger: logger}

	route.GET("/group",
		user.AuthCheckMiddleware(userService),
		resource.Index)

	route.GET("/group/:group_uuid",
		user.AuthCheckMiddleware(userService),
		ParamsCheckMiddleware(groupService),
		resource.Show)

	route.POST("/group",
		user.AuthCheckMiddleware(userService),
		resource.Create)

	route.PUT("/group/:group_uuid",
		user.AuthCheckMiddleware(userService),
		ParamsCheckMiddleware(groupService),
		resource.Update)

	route.DELETE("/group/:group_uuid",
		user.AuthCheckMiddleware(userService),
		ParamsCheckMiddleware(groupService),
		resource.Delete)

	route.GET("/group/:group_uuid/user",
		user.AuthCheckMiddleware(userService),
		ParamsCheckMiddleware(groupService),
		resource.UserList)

	route.POST("/group/:group_uuid/user",
		user.AuthCheckMiddleware(userService),
		ParamsCheckMiddleware(groupService),
		resource.Join)

	route.DELETE("/group/:group_uuid/user/:user_uuid",
		user.AuthCheckMiddleware(userService),
		user.ParamsCheckMiddleware(userService),
		ParamsCheckMiddleware(groupService),
		resource.DeleteUserInGroup)
}
