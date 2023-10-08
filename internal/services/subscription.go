package services

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/dhritimannath/notification-service/internal/constant"
	"github.com/dhritimannath/notification-service/internal/dtos"
	"github.com/dhritimannath/notification-service/internal/models"
	"github.com/dhritimannath/notification-service/internal/pkg"
	"github.com/dhritimannath/notification-service/internal/repositories"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)


type SubscriptionService interface {
	Create(*gin.Context)
	FindAllByUserId(*gin.Context)
	UpdateById(*gin.Context)
	DeleteById(*gin.Context)
}

type SubscriptionServiceImpl struct {
	subscriptionRepo repositories.SubscriptionRepository
}

func SubscriptionServiceInit(subscriptionRepo repositories.SubscriptionRepository) *SubscriptionServiceImpl {
	return &SubscriptionServiceImpl{
		subscriptionRepo: subscriptionRepo,
	}
}

func (s SubscriptionServiceImpl) Create(c *gin.Context)  {
	defer pkg.PanicHandler(c)
	var data dtos.CreateSubscription
	userId, _ := uuid.Parse(c.Param("id"))
   	if err := c.ShouldBindJSON(&data); err != nil {
      pkg.PanicException(constant.InvalidRequest)
   }
   var toCreate = models.Subscription{ID: uuid.New(),  Token: data.Token, Type: data.Type, UserId: userId}
   subscription , err := s.subscriptionRepo.Create(&toCreate)
   if err != nil {
      pkg.PanicException(constant.UnknownError)
   }

   c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, subscription))
}

func (s SubscriptionServiceImpl) FindAllByUserId(c *gin.Context)  {
	defer pkg.PanicHandler(c)
	userId, _ := uuid.Parse(c.Param("id"))
   data , err := s.subscriptionRepo.FindAllByUserId(userId)
   fmt.Println(data)
   if err != nil {
	fmt.Println(err)
      pkg.PanicException(constant.UnknownError)
   }

   c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (s SubscriptionServiceImpl) UpdateById(c *gin.Context)  {
	defer pkg.PanicHandler(c)
	var data dtos.UpdateSubscription
	subscriptionId, _ := uuid.Parse(c.Param("subscriptionId"))
   	if err := c.ShouldBindJSON(&data); err != nil {
      pkg.PanicException(constant.InvalidRequest)
   }
   subscription , err := s.subscriptionRepo.FindById(subscriptionId)
   if err != nil {
      pkg.PanicException(constant.DataNotFound)
   }

   subscription.Token = data.Token
   subscription.Type = data.Type

   s.subscriptionRepo.Save(&subscription)

   c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, subscription))
}

func (s SubscriptionServiceImpl) DeleteById(c *gin.Context)  {
	defer pkg.PanicHandler(c)
	subscriptionId, _ := uuid.Parse(c.Param("subscriptionId"))
   err := s.subscriptionRepo.DeleteById(subscriptionId)
   if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			pkg.PanicException(constant.DataNotFound)		
		} else {
			pkg.PanicException(constant.UnknownError)
		}
   }

   c.JSON(http.StatusNoContent, pkg.BuildResponse(constant.Success, pkg.Null()))
}