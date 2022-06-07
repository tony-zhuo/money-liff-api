package user

import (
	"github.com/ZhuoYIZIA/money-liff-api/internal/entity"
	"github.com/ZhuoYIZIA/money-liff-api/internal/unity"
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
		c.JSON(http.StatusBadRequest, unity.Exception{
			Status:  false,
			Code:    http.StatusBadRequest,
			Message: errMsg,
		})
		return
	}

	if err := user.Validate(); err != nil {
		errMsg := err.Error()
		logger.Error("user register API request validate error: ", log.String("err", errMsg))
		c.JSON(http.StatusBadRequest, unity.Exception{
			Status:  false,
			Code:    http.StatusBadRequest,
			Message: errMsg,
		})
		return
	}

	logger.Info("user register API request data",
		log.String("line id", user.LineId),
		log.String("name", user.Name),
		log.String("avatar url", user.AvatarUrl))

	service := NewService()
	if err := service.CreateIfNotFound(&user); err != nil {
		c.JSON(http.StatusInternalServerError, unity.Exception{
			Status:  false,
			Code:    http.StatusInternalServerError,
			Message: "server error",
		})
		return
	} else {
		c.JSON(http.StatusOK, unity.OkResponse{
			Status:  true,
			Code:    http.StatusOK,
			Message: "success",
			Data:    user,
		})
		return
	}
}
