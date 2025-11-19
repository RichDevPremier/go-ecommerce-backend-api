package user

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	//public router

	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.GET("/search")
		userRouterPublic.GET("/detail/:id")
	}

	userRouterPrivate := Router.Group("/user")
	{
		userRouterPrivate.GET("/get_info")
	}
}