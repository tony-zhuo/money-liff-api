package upload

import (
	"github.com/ZhuoYIZIA/money-liff-api/pkg/log"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.RouterGroup, uploadService Service, logger *log.Logger) {
	resource := Resource{service: uploadService, logger: logger}

	route.POST("/upload/image", resource.UploadImage)
}
