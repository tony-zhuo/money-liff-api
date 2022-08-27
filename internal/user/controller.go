package user

import (
	"github.com/ZhuoYIZIA/money-liff-api/internal/entity"
	"github.com/ZhuoYIZIA/money-liff-api/internal/unity/exception"
	"github.com/ZhuoYIZIA/money-liff-api/internal/unity/response"
	"github.com/ZhuoYIZIA/money-liff-api/internal/unity/validate_err_msg"
	"github.com/ZhuoYIZIA/money-liff-api/pkg/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller interface {
	GetUserOrRegister(c *gin.Context)
}

type controller struct {
	userService Service
	logger      *log.Logger
}

func NewController(userService Service, logger *log.Logger) Controller {
	return &controller{
		userService: userService,
		logger:      logger,
	}
}

func (ctr *controller) GetUserOrRegister(c *gin.Context) {
	user := entity.User{}
	if err := c.Bind(&user); err != nil {
		errMsg := validate_err_msg.Transfer(err).Error()
		res := exception.BadRequest(errMsg)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	ctr.logger.Info("user register API request data",
		log.String("line id", user.LineId),
		log.String("name", user.Name),
		log.String("avatar url", user.AvatarUrl))

	userResult, err := ctr.userService.RegisterOrFind(&user)
	if err != nil {
		res := exception.InternalServerError(err.Error())
		c.JSON(http.StatusInternalServerError, res)
		return
	} else {
		res := response.Ok("", userResult)
		c.JSON(http.StatusOK, res)
		return
	}
}
