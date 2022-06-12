package group

import (
	"github.com/ZhuoYIZIA/money-liff-api/internal/entity"
	"github.com/ZhuoYIZIA/money-liff-api/internal/unity/exception"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ParamsCheckMiddleware check group uuid exit
func ParamsCheckMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		groupUuid := c.Param("group_uuid")
		groupService := NewService()

		params := entity.GroupParams{UUID: groupUuid}
		if err := params.Validate(); err != nil {
			res := exception.BadRequest(err.Error())
			c.JSON(http.StatusBadRequest, res)
			c.Abort()
			return
		}

		groupData := groupService.GetGroupByUUID(groupUuid)
		if groupData == nil {
			res := exception.NotFound("Group not found.")
			c.JSON(http.StatusNotFound, res)
			c.Abort()
			return
		}

		c.Set("groupData", groupData)
		c.Next()
	}
}
