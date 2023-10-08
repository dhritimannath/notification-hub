package controllers

import (
	"fmt"

	"github.com/dhritimannath/notification-service/internal/services"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	Create(c *gin.Context)
	FindAll(c *gin.Context)
	FindById(c *gin.Context)
}

type UserControllerImpl struct {
	userService services.UserService
}


func UserControllerInit(userService services.UserService) *UserControllerImpl {
	return &UserControllerImpl{
		userService: userService,
	}
}

func (u *UserControllerImpl) Create(c *gin.Context)  {
	u.userService.Create(c)
}

func (u *UserControllerImpl) FindAll(c *gin.Context)  {
	u.userService.FindAll(c)
}

func (u *UserControllerImpl) FindById(c *gin.Context)  {
	fmt.Print("FIND BY ID")
	u.userService.FindById(c)
}