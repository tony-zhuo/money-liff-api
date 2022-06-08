package user

import (
	"github.com/ZhuoYIZIA/money-liff-api/internal/entity"
	"github.com/ZhuoYIZIA/money-liff-api/internal/unity/response"
	"github.com/ZhuoYIZIA/money-liff-api/pkg/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

var logger = log.TeeDefault()

func GetUserOrRegister(c *gin.Context) {
	user := entity.User{}
	if err := c.Bind(&user); err != nil {
		errMsg := err.Error()
		logger.Error("user register API bind error: ", log.String("err", errMsg))
		res := response.BadRequest(errMsg)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := user.Validate(); err != nil {
		errMsg := err.Error()
		logger.Error("user register API request validate error: ", log.String("err", errMsg))
		res := response.BadRequest(errMsg)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	logger.Info("user register API request data",
		log.String("line id", user.LineId),
		log.String("name", user.Name),
		log.String("avatar url", user.AvatarUrl))

	service := NewService()
	if err := service.CreateIfNotFound(&user); err != nil {
		res := response.InternalServerError("")
		c.JSON(http.StatusInternalServerError, res)
		return
	} else {
		res := response.Ok("", user)
		c.JSON(http.StatusOK, res)
		return
	}
}
