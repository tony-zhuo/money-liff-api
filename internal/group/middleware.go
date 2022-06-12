package group

import (
	"github.com/ZhuoYIZIA/money-liff-api/internal/unity/exception"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Middleware check group uuid exit
func Middleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		groupUuid := c.Param("group_uuid")
		groupService := NewService()

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
