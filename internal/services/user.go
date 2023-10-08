package services

import (
	"net/http"

	"github.com/dhritimannath/notification-service/internal/constant"
	"github.com/dhritimannath/notification-service/internal/dtos"
	"github.com/dhritimannath/notification-service/internal/models"
	"github.com/dhritimannath/notification-service/internal/pkg"
	"github.com/dhritimannath/notification-service/internal/repositories"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


type UserService interface {
	Create(*gin.Context)
	FindAll(*gin.Context)
	FindById(*gin.Context)
}

type UserServiceImpl struct {
	userRepo repositories.UserRepository
}

func UserServiceInit(userRepo repositories.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepo: userRepo,
	}
}

func (u UserServiceImpl) Create(c *gin.Context)  {
	defer pkg.PanicHandler(c)
	var data dtos.CreateUser
   	if err := c.ShouldBindJSON(&data); err != nil {
      pkg.PanicException(constant.InvalidRequest)
   }

   createdUser , err := u.userRepo.Create(&models.User{ID: uuid.New() ,ExternalId: &data.ExternalId})
   if err != nil {
      pkg.PanicException(constant.UnknownError)
   }

   c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, createdUser))
}

func (u UserServiceImpl) FindAll(c *gin.Context)  {
	defer pkg.PanicHandler(c)
   data , err := u.userRepo.FindAll()
   if err != nil {
      pkg.PanicException(constant.UnknownError)
   }

   c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u UserServiceImpl) FindById(c *gin.Context)  {
	defer pkg.PanicHandler(c)
	userId, _ := uuid.Parse(c.Param("id"))
   data , err := u.userRepo.FindById(userId)
   if data == (models.User{}) {
	pkg.PanicException(constant.DataNotFound)
 	}
   if err != nil {
      pkg.PanicException(constant.UnknownError)
   }

   c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}