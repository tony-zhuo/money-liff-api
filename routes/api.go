package routes

import (
	"github.com/ZhuoYIZIA/money-liff-api/internal/group"
	"github.com/ZhuoYIZIA/money-liff-api/internal/user"
	"github.com/ZhuoYIZIA/money-liff-api/pkg/database"
	"github.com/ZhuoYIZIA/money-liff-api/pkg/log"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}
	config.AllowHeaders = []string{"Line-Id"}
	router.Use(cors.New(config))

	logger := log.TeeDefault()
	db := database.Connection()
	userService := user.NewService(user.NewRepository(db, logger), logger)
	groupService := group.NewService(group.NewRepository(db, logger), logger)

	v1Router := router.Group("v1")
	{
		user.Routes(v1Router, userService, logger)
		group.Routes(v1Router, groupService, userService, logger)
	}

	return router
}
