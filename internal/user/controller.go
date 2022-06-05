package user

import (
	"github.com/ZhuoYIZIA/money-liff-api/internal/entity"
	"github.com/gin-gonic/gin"
)

func GetUserOrRegister(c *gin.Context) {
	user := entity.User{}
	err := c.Bind(&user)
	if err != nil {
		//TODO: error handle
		c.JSON(400, err.Error())
		return
	}

	err = user.Validate()
	if err != nil {
		//TODO: error handle
		c.JSON(400, err.Error())
		return
	}

	service := NewService(&user)
	result := service.FirstOrCreate()
	c.JSON(200, result)
	return
}
