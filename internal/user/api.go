package user

import (
	"github.com/ZhuoYIZIA/money-liff-api/pkg/log"
	"github.com/gin-gonic/gin"
)

type Route interface {
	Routes(route *gin.RouterGroup)
}

type route struct {
	userController Controller
	logger         *log.Logger
}

func NewRoute(userController Controller, logger *log.Logger) Route {
	return &route{
		userController: userController,
		logger:         logger,
	}
}

func (r *route) Routes(route *gin.RouterGroup) {
	route.POST("/user/register", r.userController.GetUserOrRegister)
}
