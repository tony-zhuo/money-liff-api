package cost

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

func (r *Resource) Create(c *gin.Context) {
	groupData := c.MustGet("groupData").(*entity.Group)
	auth := c.MustGet("auth").(*entity.User)

	groupCostItem := entity.GroupCostItemRequestArg{}
	if err := c.Bind(&groupCostItem); err != nil {
		r.logger.Info("[Cost Create Controller] bind request error",
			log.Any("err", err))
		res := exception.BadRequest(err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if !r.service.CheckParticipantAmount(groupCostItem.TotalAmount, groupCostItem.Participants) {
		res := exception.BadRequest("Total amount not equal participants amount total.")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := r.service.CreateGroupCostAmdParticipants(&groupCostItem, groupData, auth); err != nil {
		res := exception.InternalServerError("")
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := response.Created("")
	c.JSON(http.StatusCreated, res)
	return

}
