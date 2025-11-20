package controller

import (
	"go-ecommerce-backend-api/internal/service"
	"go-ecommerce-backend-api/response"

	"github.com/gin-gonic/gin"
)

// type UserController struct{
// 	userService *service.UserService
// }

// func NewUserController() *UserController{
// 	return &UserController{
// 		userService: service.NewUserService(),
// 	}
// }

// //uc user controller
// //us user service
// func (uc *UserController) GetUserById(c *gin.Context) {
// 	response.SuccessReponse(c, 20001, []string{"tipjs"})
// }

type UserController struct{
 	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController{
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Register(c *gin.Context){
	result := uc.userService.Register("", "")
	response.SuccessReponse(c, result, nil)
}