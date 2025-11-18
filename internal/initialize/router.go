package initialize

import (
	c "go-ecommerce-backend-api/internal/controller"
	"go-ecommerce-backend-api/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.AuthenMiddleware())

	v1 := r.Group("/v1/2024")
	{
		v1.GET("/ping", c.NewPongController().GetPongById)
		v1.GET("/user/1", c.NewUserController().GetUserById)
	}

	v2 := r.Group("/v2/2024", c.NewUserController().GetUserById)
	{
		v2.GET("/ping")
	}
	return r
}
