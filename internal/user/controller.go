package user

import (
	"github.com/ZhuoYIZIA/money-liff-api/internal/entity"
	"github.com/ZhuoYIZIA/money-liff-api/internal/unity/exception"
	"github.com/ZhuoYIZIA/money-liff-api/internal/unity/response"
	"github.com/ZhuoYIZIA/money-liff-api/pkg/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Resource struct {
	service Service
	logger  *log.Logger
}

func (r *Resource) GetUserOrRegister(c *gin.Context) {
	user := entity.User{}
	if err := c.Bind(&user); err != nil {
		errMsg := err.Error()
		r.logger.Error("user register API bind error: ", log.String("err", errMsg))
		res := exception.BadRequest(errMsg)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := user.Validate(); err != nil {
		errMsg := err.Error()
		r.logger.Error("user register API request validate error: ", log.String("err", errMsg))
		res := exception.BadRequest(errMsg)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	r.logger.Info("user register API request data",
		log.String("line id", user.LineId),
		log.String("name", user.Name),
		log.String("avatar url", user.AvatarUrl))

	userResult, err := r.service.RegisterOrFind(&user)
	if err != nil {
		res := exception.InternalServerError("")
		c.JSON(http.StatusInternalServerError, res)
		return
	} else {
		res := response.Ok("", userResult)
		c.JSON(http.StatusOK, res)
		return
	}
}
