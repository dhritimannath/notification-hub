package controllers

import (
	"github.com/dhritimannath/notification-service/internal/services"
	"github.com/gin-gonic/gin"
)

type SubscriptionController interface {
	Create(c *gin.Context)
	FindAllByUserId(c *gin.Context)
	UpdateById(c *gin.Context)
	DeleteById(c *gin.Context)
}

type SubscriptionControllerImpl struct {
	subscriptionService services.SubscriptionService
}


func SubscriptionControllerInit(subscriptionService services.SubscriptionService) *SubscriptionControllerImpl {
	return &SubscriptionControllerImpl{
		subscriptionService: subscriptionService,
	}
}

func (s *SubscriptionControllerImpl) Create(c *gin.Context)  {
	s.subscriptionService.Create(c)
}

func (s *SubscriptionControllerImpl) FindAllByUserId(c *gin.Context)  {
	s.subscriptionService.FindAllByUserId(c)
}

func (s *SubscriptionControllerImpl) UpdateById(c *gin.Context)  {
	s.subscriptionService.UpdateById(c)
}

func (s *SubscriptionControllerImpl) DeleteById(c *gin.Context)  {
	s.subscriptionService.DeleteById(c)
}
