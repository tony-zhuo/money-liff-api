package upload

import (
	"github.com/ZhuoYIZIA/money-liff-api/internal/unity/exception"
	"github.com/ZhuoYIZIA/money-liff-api/internal/unity/response"
	"github.com/ZhuoYIZIA/money-liff-api/pkg/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller interface {
	UploadImage(c *gin.Context)
}

type controller struct {
	uploadService Service
	logger        *log.Logger
}

func NewController(uploadService Service, logger *log.Logger) Controller {
	return &controller{
		uploadService: uploadService,
		logger:        logger,
	}
}

func (ctr *controller) UploadImage(c *gin.Context) {
	file, _ := c.FormFile("file")
	category := c.PostForm("category")

	filePath, err := ctr.uploadService.UploadImageAndGetPath(file, category)
	if err != nil {
		res := exception.InternalServerError(err.Error())
		c.JSON(http.StatusInternalServerError, res)
		return
	} else {
		res := response.Ok("", gin.H{"url": filePath})
		c.JSON(http.StatusOK, res)
		return
	}
}
