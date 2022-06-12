package group

import (
	"github.com/ZhuoYIZIA/money-liff-api/internal/user"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.RouterGroup) {
	route.GET("/group", user.AuthCheckMiddleware(), Index)
	//route.GET("/group/:uuid")
	route.POST("/group", user.AuthCheckMiddleware(), Create)
	route.PUT("/group/:group_uuid", user.AuthCheckMiddleware(), ParamsCheckMiddleware(), Update)
	route.DELETE("/group/:group_uuid", user.AuthCheckMiddleware(), ParamsCheckMiddleware(), Delete)

	//route.GET("/group/:group_uuid/user", UserList)
}
