package upload

import (
	"github.com/ZhuoYIZIA/money-liff-api/pkg/log"
	"github.com/gin-gonic/gin"
)

type Route interface {
	Routes(route *gin.RouterGroup)
}

type route struct {
	uploadController Controller
	logger           *log.Logger
}

func NewRoute(uploadController Controller, logger *log.Logger) Route {
	return &route{
		uploadController: uploadController,
		logger:           logger,
	}
}

func (r *route) Routes(route *gin.RouterGroup) {
	route.POST("/upload/image", r.uploadController.UploadImage)
}
