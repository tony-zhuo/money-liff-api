package cost

import (
	"github.com/ZhuoYIZIA/money-liff-api/internal/entity"
	"github.com/ZhuoYIZIA/money-liff-api/internal/unity/exception"
	"github.com/ZhuoYIZIA/money-liff-api/internal/unity/response"
	"github.com/ZhuoYIZIA/money-liff-api/pkg/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller interface {
	Create(c *gin.Context)
}

type controller struct {
	costService Service
	logger      *log.Logger
}

func NewController(costService Service, logger *log.Logger) Controller {
	return &controller{
		costService: costService,
		logger:      logger,
	}
}

func (ctr *controller) Create(c *gin.Context) {
	groupData := c.MustGet("groupData").(*entity.Group)
	auth := c.MustGet("auth").(*entity.User)

	groupCostItem := entity.GroupCostItemRequestArg{}
	if err := c.Bind(&groupCostItem); err != nil {
		ctr.logger.Info("[Cost Create Controller] bind request error",
			log.Any("err", err))
		res := exception.BadRequest(err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if !ctr.costService.CheckParticipantAmount(groupCostItem.TotalAmount, groupCostItem.Participants) {
		res := exception.BadRequest("Total amount not equal participants amount total.")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := ctr.costService.CreateGroupCostAmdParticipants(&groupCostItem, groupData, auth); err != nil {
		res := exception.InternalServerError("")
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := response.Created("")
	c.JSON(http.StatusCreated, res)
	return

}
