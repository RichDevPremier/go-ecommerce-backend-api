package controller

import (
	"go-ecommerce-backend-api/internal/service"
	"go-ecommerce-backend-api/response"

	"github.com/gin-gonic/gin"
)

type UserController struct{
	userService *service.UserService
}

func NewUserController() *UserController{
	return &UserController{
		userService: service.NewUserService(),
	}
}


//uc user controller
//us user service
func (uc *UserController) GetUserById(c *gin.Context) {
	response.SuccessReponse(c, 20001, []string{"tipjs"})
}