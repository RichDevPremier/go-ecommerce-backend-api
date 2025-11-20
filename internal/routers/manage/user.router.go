package manage

import (
	"go-ecommerce-backend-api/internal/wire"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	//use wire
	userController, _ := wire.InitUserRouterHandler()
	//public router
	userRouterPublic := Router.Group("/admin/user")
	{
		userRouterPublic.POST("/register", userController.Register)
		// userRouterPublic.POST("/otp")
	}
	//private router
	userRouterPrivate := Router.Group("/admin/user")
	{
		userRouterPrivate.POST("/active_user")
	}
}