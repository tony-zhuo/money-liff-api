package group

import (
	"github.com/ZhuoYIZIA/money-liff-api/internal/entity"
	"github.com/ZhuoYIZIA/money-liff-api/internal/unity/exception"
	"github.com/ZhuoYIZIA/money-liff-api/internal/unity/response"
	"github.com/ZhuoYIZIA/money-liff-api/pkg/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var logger = log.TeeDefault()

func Index(c *gin.Context) {
	userData := c.MustGet("userData").(*entity.User)
	queryPage := c.DefaultQuery("page", "1")
	queryPerPage := c.DefaultQuery("per_page", "10")
	sort := c.DefaultQuery("sort", "id")
	page, _ := strconv.ParseInt(queryPage, 10, 32)
	perPage, _ := strconv.ParseInt(queryPerPage, 10, 32)
	groupService := NewService()

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
	userData := c.MustGet("userData").(*entity.User)
	groupService := NewService()

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
	userData := c.MustGet("userData").(*entity.User)
	groupData := c.MustGet("groupData").(*entity.Group)
	uuid := c.Param("uuid")
	groupService := NewService()
	logger.Info("Group update controller",
		log.String("line-id", userData.LineId),
		log.String("uuid", uuid))

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

	if isAdmin := groupService.CheckUserIsAdmin(groupData, userData); !isAdmin {
		res := exception.Unauthorized("The user is not admin.")
		c.JSON(http.StatusUnauthorized, res)
		return
	}

	if err := groupService.UpdateGroupById(&request, groupData.Id); err != nil {
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
	userData := c.MustGet("userData").(*entity.User)
	groupData := c.MustGet("groupData").(*entity.Group)
	uuid := c.Param("uuid")
	groupService := NewService()
	logger.Info("Group update controller",
		log.String("line-id", userData.LineId),
		log.String("uuid", uuid))

	if isAdmin := groupService.CheckUserIsAdmin(groupData, userData); !isAdmin {
		res := exception.Unauthorized("The user is not admin.")
		c.JSON(http.StatusUnauthorized, res)
		return
	}

	if err := groupService.DeleteGroupById(groupData.Id); err != nil {
		res := exception.InternalServerError("")
		c.JSON(http.StatusInternalServerError, res)
		return
	} else {
		res := response.Ok("Delete success.", nil)
		c.JSON(http.StatusOK, res)
		return
	}
}

//func UserList(c *gin.Context) {
//	lineId := c.GetHeader("Line-Id")
//	uuid := c.Param("group_uuid")
//	logger.Info("Group controller UserList",
//		log.String("line-id", lineId),
//		log.String("group_uuid", uuid))
//
//	if lineId == "" {
//		res := exception.Unauthorized("")
//		c.JSON(http.StatusUnauthorized, res)
//		return
//	}
//
//	userData := userService.GetUserByLineId(lineId)
//	if userData == nil {
//		res := exception.Unauthorized("")
//		c.JSON(http.StatusUnauthorized, res)
//		return
//	}
//
//	group := groupService.GetGroupByUUID(uuid)
//	if group == nil {
//		res := exception.NotFound("Group not found.")
//		c.JSON(http.StatusNotFound, res)
//		return
//	}
//}
