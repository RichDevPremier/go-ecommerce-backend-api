package manage

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (pr *UserRouter) InitUSerRouter(Router *gin.RouterGroup) {
	//private router
	userRouterPrivate := Router.Group("/admin/user")
	{
		userRouterPrivate.POST("/active_user")
	}
}