package group

import (
	"github.com/ZhuoYIZIA/money-liff-api/internal/entity"
	"github.com/ZhuoYIZIA/money-liff-api/internal/unity/response"
	"github.com/ZhuoYIZIA/money-liff-api/internal/user"
	"github.com/ZhuoYIZIA/money-liff-api/pkg/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var logger = log.TeeDefault()
var groupService = NewService()
var userService = user.NewService()

func Index(c *gin.Context) {
	lineId := c.GetHeader("Line-Id")
	queryPage := c.DefaultQuery("page", "1")
	queryPerPage := c.DefaultQuery("per_page", "10")
	sort := c.DefaultQuery("sort", "id")
	page, _ := strconv.ParseInt(queryPage, 10, 32)
	perPage, _ := strconv.ParseInt(queryPerPage, 10, 32)

	if lineId == "" {
		res := response.Unauthorized("")
		c.JSON(http.StatusUnauthorized, res)
		return
	}

	logger.Info("group create header", log.String("line id", lineId))

	userData := userService.GetUserByLineId(lineId)
	if userData == nil {
		res := response.Unauthorized("")
		c.JSON(http.StatusUnauthorized, res)
		return
	}

	pagination, err := groupService.GetListByUserWithPagination(userData, int(page), int(perPage), sort)
	if err != nil {
		res := response.InternalServerError(err.Error())
		c.JSON(http.StatusInternalServerError, res)
		return
	} else {
		res := response.List(*pagination, "")
		c.JSON(http.StatusOK, res)
		return
	}
}

func Create(c *gin.Context) {
	logger.Info("group create")

	lineId := c.GetHeader("Line-Id")
	if lineId == "" {
		res := response.Unauthorized("")
		c.JSON(http.StatusUnauthorized, res)
		return
	}

	logger.Info("group create header", log.String("line id", lineId))

	userData := userService.GetUserByLineId(lineId)
	if userData == nil {
		res := response.Unauthorized("")
		c.JSON(http.StatusUnauthorized, res)
		return
	}

	group := entity.Group{}
	if err := c.Bind(&group); err != nil {
		res := response.BadRequest("")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	logger.Info("group create request bind",
		log.String("Name", group.Name),
		log.Int("UserLimit", group.UserLimit),
		log.String("ImageUrl", group.ImageUrl))

	if err := group.Validate(); err != nil {
		res := response.BadRequest(err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := groupService.GenerateUUIDAndCreateByUser(&group, userData); err != nil {
		res := response.InternalServerError(err.Error())
		c.JSON(http.StatusInternalServerError, res)
		return
	} else {
		res := response.Created("")
		c.JSON(http.StatusCreated, res)
		return
	}

}
