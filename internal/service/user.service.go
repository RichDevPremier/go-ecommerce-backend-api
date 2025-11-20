package service

import (
	"go-ecommerce-backend-api/internal/repo"
	"go-ecommerce-backend-api/response"
)

// import "go-ecommerce-backend-api/internal/repo"

// type UserService struct {
// 	userRepo *repo.UserRepo
// }

// func NewUserService() *UserService {
// 	return &UserService{
// 		userRepo: repo.NewUserRepo(),
// 	}
// }

// func (us *UserService) GetInfoUser() string {
// 	return us.userRepo.GetInfoUser()
// }

type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo repo.IUserRepository
}


func NewUserService(
	userRepo repo.IUserRepository,
) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}

// Register implements IUserService.
func (us *userService) Register(email string, purpose string) int {

	//1. Check email exists
	if us.userRepo.GetUserByEmail(email){
		return response.ErrCodeUserHasExists
	}
	return response.ErrCodeSuccess
}