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

type Resource struct {
	service Service
	logger  *log.Logger
}

func (r *Resource) Index(c *gin.Context) {
	userData := c.MustGet("userData").(*entity.User)
	queryPage := c.DefaultQuery("page", "1")
	queryPerPage := c.DefaultQuery("per_page", "10")
	sort := c.DefaultQuery("sort", "id")
	page, _ := strconv.ParseInt(queryPage, 10, 32)
	perPage, _ := strconv.ParseInt(queryPerPage, 10, 32)

	pagination, err := r.service.GetListByUserWithPagination(userData, int(page), int(perPage), sort)
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

func (r *Resource) Create(c *gin.Context) {
	userData := c.MustGet("userData").(*entity.User)

	group := entity.Group{}
	if err := c.Bind(&group); err != nil {
		res := exception.BadRequest("")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	r.logger.Info("group create request bind",
		log.String("Name", group.Name),
		log.Int("UserLimit", group.UserLimit),
		log.String("ImageUrl", group.ImageUrl))

	if err := group.Validate(); err != nil {
		res := exception.BadRequest(err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := r.service.GenerateUUIDAndCreateByUser(&group, userData); err != nil {
		res := exception.InternalServerError(err.Error())
		c.JSON(http.StatusInternalServerError, res)
		return
	} else {
		res := response.Created("")
		c.JSON(http.StatusCreated, res)
		return
	}

}

func (r *Resource) Update(c *gin.Context) {
	userData := c.MustGet("userData").(*entity.User)
	groupData := c.MustGet("groupData").(*entity.Group)
	uuid := c.Param("uuid")
	r.logger.Info("Group update controller",
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

	if isAdmin := r.service.CheckUserIsAdmin(groupData, userData); !isAdmin {
		res := exception.Unauthorized("The user is not admin.")
		c.JSON(http.StatusUnauthorized, res)
		return
	}

	if err := r.service.UpdateGroupById(&request, groupData.Id); err != nil {
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

func (r *Resource) Delete(c *gin.Context) {
	userData := c.MustGet("userData").(*entity.User)
	groupData := c.MustGet("groupData").(*entity.Group)
	uuid := c.Param("uuid")
	r.logger.Info("Group update controller",
		log.String("line-id", userData.LineId),
		log.String("uuid", uuid))

	if isAdmin := r.service.CheckUserIsAdmin(groupData, userData); !isAdmin {
		res := exception.Unauthorized("The user is not admin.")
		c.JSON(http.StatusUnauthorized, res)
		return
	}

	if err := r.service.DeleteGroupById(groupData.Id); err != nil {
		res := exception.InternalServerError("")
		c.JSON(http.StatusInternalServerError, res)
		return
	} else {
		res := response.Ok("Delete success.", nil)
		c.JSON(http.StatusOK, res)
		return
	}
}

//func (r *Resource) UserList(c *gin.Context) {
//	lineId := c.GetHeader("Line-Id")
//	uuid := c.Param("group_uuid")
//	r.logger.Info("Group controller UserList",
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
