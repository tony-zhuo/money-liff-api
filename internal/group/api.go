package group

import (
	"github.com/ZhuoYIZIA/money-liff-api/internal/unity/middleware"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.RouterGroup) {
	route.GET("/group", middleware.AuthMiddleware(), Index)
	//route.GET("/group/:uuid")
	route.POST("/group", middleware.AuthMiddleware(), Create)
	route.PUT("/group/:uuid", middleware.AuthMiddleware(), Update)
	route.DELETE("/group/:uuid", middleware.AuthMiddleware(), Delete)

	//route.GET("/group/:group_uuid/user", UserList)
}
