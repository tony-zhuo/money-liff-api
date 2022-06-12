package group

import (
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.RouterGroup) {
	route.GET("/group", Index)
	//route.GET("/group/:uuid")
	route.POST("/group", Create)
	route.PUT("/group/:uuid", Update)
	//route.DELETE("/group/:uuid")
}
