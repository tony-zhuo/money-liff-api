package cost

import (
	"github.com/ZhuoYIZIA/money-liff-api/internal/entity"
	"github.com/ZhuoYIZIA/money-liff-api/pkg/log"
	"github.com/gin-gonic/gin"
)

type Resource struct {
	service Service
	logger  *log.Logger
}

func (r *Resource) Create(c *gin.Context) {
	groupData := c.MustGet("groupData").(entity.Group)
	auth := c.MustGet("auth").(entity.User)
}
