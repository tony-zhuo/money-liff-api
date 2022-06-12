package group

import (
	"github.com/ZhuoYIZIA/money-liff-api/internal/entity"
	"github.com/ZhuoYIZIA/money-liff-api/internal/unity/exception"
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
		res := exception.Unauthorized("")
		c.JSON(http.StatusUnauthorized, res)
		return
	}

	logger.Info("group create header", log.String("line id", lineId))

	userData := userService.GetUserByLineId(lineId)
	if userData == nil {
		res := exception.Unauthorized("")
		c.JSON(http.StatusUnauthorized, res)
		return
	}

	pagination, err := groupService.GetListByUserWithPagination(userData, int(page), int(perPage), sort)
	if err != nil {
		res := exception.InternalServerError(err.Error())
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
		res := exception.Unauthorized("")
		c.JSON(http.StatusUnauthorized, res)
		return
	}

	logger.Info("group create header", log.String("line id", lineId))

	userData := userService.GetUserByLineId(lineId)
	if userData == nil {
		res := exception.Unauthorized("")
		c.JSON(http.StatusUnauthorized, res)
		return
	}

	group := entity.Group{}
	if err := c.Bind(&group); err != nil {
		res := exception.BadRequest("")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	logger.Info("group create request bind",
		log.String("Name", group.Name),
		log.Int("UserLimit", group.UserLimit),
		log.String("ImageUrl", group.ImageUrl))

	if err := group.Validate(); err != nil {
		res := exception.BadRequest(err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := groupService.GenerateUUIDAndCreateByUser(&group, userData); err != nil {
		res := exception.InternalServerError(err.Error())
		c.JSON(http.StatusInternalServerError, res)
		return
	} else {
		res := response.Created("")
		c.JSON(http.StatusCreated, res)
		return
	}

}

func Update(c *gin.Context) {
	lineId := c.GetHeader("Line-Id")
	uuid := c.Param("uuid")
	logger.Info("Group update controller",
		log.String("line-id", lineId),
		log.String("uuid", uuid))

	if lineId == "" {
		res := exception.Unauthorized("")
		c.JSON(http.StatusUnauthorized, res)
		return
	}

	userData := userService.GetUserByLineId(lineId)
	if userData == nil {
		res := exception.Unauthorized("")
		c.JSON(http.StatusUnauthorized, res)
		return
	}

	request := entity.Group{}
	if err := c.Bind(&request); err != nil {
		res := exception.BadRequest("")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := request.Validate(); err != nil {
		res := exception.BadRequest(err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}

	group := groupService.GetGroupByUUID(uuid)
	if group == nil {
		res := exception.NotFound("Group not found.")
		c.JSON(http.StatusNotFound, res)
		return
	}

	if isAdmin := groupService.CheckUserIsAdmin(group, userData); !isAdmin {
		res := exception.Unauthorized("The user is not admin.")
		c.JSON(http.StatusUnauthorized, res)
		return
	}

	if err := groupService.UpdateGroupById(&request, group.Id); err != nil {
		res := exception.InternalServerError("")
		c.JSON(http.StatusInternalServerError, res)
		return
	} else {
		res := response.Ok("Update success.", entity.Group{
			UUID:      uuid,
			Name:      request.Name,
			UserLimit: request.UserLimit,
			ImageUrl:  request.ImageUrl,
		})
		c.JSON(http.StatusOK, res)
		return
	}
}

func Delete(c *gin.Context) {
	lineId := c.GetHeader("Line-Id")
	uuid := c.Param("uuid")
	logger.Info("Group update controller",
		log.String("line-id", lineId),
		log.String("uuid", uuid))

	if lineId == "" {
		res := exception.Unauthorized("")
		c.JSON(http.StatusUnauthorized, res)
		return
	}

	userData := userService.GetUserByLineId(lineId)
	if userData == nil {
		res := exception.Unauthorized("")
		c.JSON(http.StatusUnauthorized, res)
		return
	}

	group := groupService.GetGroupByUUID(uuid)
	if group == nil {
		res := exception.NotFound("Group not found.")
		c.JSON(http.StatusNotFound, res)
		return
	}

	if isAdmin := groupService.CheckUserIsAdmin(group, userData); !isAdmin {
		res := exception.Unauthorized("The user is not admin.")
		c.JSON(http.StatusUnauthorized, res)
		return
	}

	if err := groupService.DeleteGroupById(group.Id); err != nil {
		res := exception.InternalServerError("")
		c.JSON(http.StatusInternalServerError, res)
		return
	} else {
		res := response.Ok("Delete success.", nil)
		c.JSON(http.StatusOK, res)
		return
	}
}
