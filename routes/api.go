package routes

import (
	"github.com/ZhuoYIZIA/money-liff-api/internal/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}
	router.Use(cors.New(config))

	v1Router := router.Group("v1")
	{
		user.Routes(v1Router)
	}

	return router
}
