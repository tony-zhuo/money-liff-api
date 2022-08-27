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

type Controller interface {
	Index(c *gin.Context)
	Show(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	UserList(c *gin.Context)
	Join(c *gin.Context)
	DeleteUserInGroup(c *gin.Context)
}

type controller struct {
	groupService Service
	logger       *log.Logger
}

func NewController(groupService Service, logger *log.Logger) Controller {
	return &controller{
		groupService: groupService,
		logger:       logger,
	}
}

func (ctr *controller) Index(c *gin.Context) {
	auth := c.MustGet("auth").(*entity.User)
	queryPage := c.DefaultQuery("page", "1")
	queryPerPage := c.DefaultQuery("per_page", "10")
	sort := c.DefaultQuery("sort", "id")
	page, _ := strconv.ParseInt(queryPage, 10, 32)
	perPage, _ := strconv.ParseInt(queryPerPage, 10, 32)

	pagination, err := ctr.groupService.GetListByUserWithPagination(auth, int(page), int(perPage), sort)
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

func (ctr *controller) Show(c *gin.Context) {
	groupData := c.MustGet("groupData").(*entity.Group)
	auth := c.MustGet("auth").(*entity.User)

	result, err := ctr.groupService.GetGroupWithCostItemInfo(auth, groupData)
	if err != nil {
		res := exception.InternalServerError("")
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	ctr.logger.Info("Show 4")

	res := response.Ok("", result)
	c.JSON(http.StatusOK, res)
	return
}

func (ctr *controller) Create(c *gin.Context) {
	auth := c.MustGet("auth").(*entity.User)

	group := entity.Group{}
	if err := c.Bind(&group); err != nil {
		res := exception.BadRequest("")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	ctr.logger.Info("group create request bind",
		log.String("Name", group.Name),
		log.Int("UserLimit", group.UserLimit),
		log.String("ImageUrl", group.ImageUrl))

	if err := group.Validate(); err != nil {
		res := exception.BadRequest(err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := ctr.groupService.GenerateUUIDAndCreateByUser(&group, auth); err != nil {
		res := exception.InternalServerError(err.Error())
		c.JSON(http.StatusInternalServerError, res)
		return
	} else {
		res := response.Created("")
		c.JSON(http.StatusCreated, res)
		return
	}

}

func (ctr *controller) Update(c *gin.Context) {
	auth := c.MustGet("auth").(*entity.User)
	groupData := c.MustGet("groupData").(*entity.Group)
	groupUuid := c.Param("group_uuid")
	ctr.logger.Info("Group update controller",
		log.String("line-id", auth.LineId),
		log.String("group_uuid", groupUuid))

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

	if isAdmin := ctr.groupService.CheckUserIsAdmin(groupData, auth); !isAdmin {
		res := exception.Unauthorized("The user is not admin.")
		c.JSON(http.StatusUnauthorized, res)
		return
	}

	if err := ctr.groupService.UpdateGroupById(&request, groupData.Id); err != nil {
		res := exception.InternalServerError("")
		c.JSON(http.StatusInternalServerError, res)
		return
	} else {
		res := response.Ok("Update success.", entity.Group{
			UUID:      groupUuid,
			Name:      request.Name,
			UserLimit: request.UserLimit,
			ImageUrl:  request.ImageUrl,
		})
		c.JSON(http.StatusOK, res)
		return
	}
}

func (ctr *controller) Delete(c *gin.Context) {
	auth := c.MustGet("auth").(*entity.User)
	groupData := c.MustGet("groupData").(*entity.Group)
	groupUuid := c.Param("group_uuid")
	ctr.logger.Info("Group update controller",
		log.String("line-id", auth.LineId),
		log.String("group_uuid", groupUuid))

	if isAdmin := ctr.groupService.CheckUserIsAdmin(groupData, auth); !isAdmin {
		res := exception.Unauthorized("The user is not admin.")
		c.JSON(http.StatusUnauthorized, res)
		return
	}

	if err := ctr.groupService.DeleteGroupById(groupData.Id); err != nil {
		res := exception.InternalServerError("")
		c.JSON(http.StatusInternalServerError, res)
		return
	} else {
		res := response.Ok("Delete success.", nil)
		c.JSON(http.StatusOK, res)
		return
	}
}

func (ctr *controller) UserList(c *gin.Context) {
	groupData := c.MustGet("groupData").(*entity.Group)
	queryPage := c.DefaultQuery("page", "1")
	queryPerPage := c.DefaultQuery("per_page", "10")
	sort := c.DefaultQuery("sort", "id")
	page, _ := strconv.ParseInt(queryPage, 10, 32)
	perPage, _ := strconv.ParseInt(queryPerPage, 10, 32)

	pagination, err := ctr.groupService.GetUserListWithPagination(groupData, int(page), int(perPage), sort)
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

func (ctr *controller) Join(c *gin.Context) {
	auth := c.MustGet("auth").(*entity.User)
	groupData := c.MustGet("groupData").(*entity.Group)

	if err := ctr.groupService.UserJoinGroup(auth, groupData); err != nil {
		res := exception.InternalServerError(err.Error())
		c.JSON(http.StatusInternalServerError, res)
		return
	} else {
		res := response.Created("")
		c.JSON(http.StatusOK, res)
		return
	}
}

func (ctr *controller) DeleteUserInGroup(c *gin.Context) {
	groupData := c.MustGet("groupData").(*entity.Group)
	userData := c.MustGet("userData").(*entity.User)

	if err := ctr.groupService.DeleteUserInGroup(groupData, userData); err != nil {
		res := exception.InternalServerError(err.Error())
		c.JSON(http.StatusInternalServerError, res)
		return
	} else {
		res := response.Ok("", nil)
		c.JSON(http.StatusOK, res)
		return
	}
}
