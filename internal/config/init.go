package config

import (
	"github.com/dhritimannath/notification-service/internal/controllers"
	"github.com/dhritimannath/notification-service/internal/repositories"
	"github.com/dhritimannath/notification-service/internal/services"
)

type Initialization struct {
   userRepo repositories.UserRepository
   userSvc  services.UserService
   UserCtrl controllers.UserController

   subscriptionRepo repositories.SubscriptionRepository
   subscriptionSvc  services.SubscriptionService
   SubscriptionCtrl controllers.SubscriptionController
}

func NewInitialization(
   userRepo repositories.UserRepository,
   userService services.UserService,
   UserCtrl controllers.UserController, 
   subscriptionRepo repositories.SubscriptionRepository,
   subscriptionService  services.SubscriptionService,
   SubscriptionCtrl controllers.SubscriptionController,
   ) *Initialization {
   return &Initialization{
      userRepo: userRepo,
      userSvc:  userService,
      UserCtrl: UserCtrl,
      subscriptionRepo: subscriptionRepo,
      subscriptionSvc:  subscriptionService,
      SubscriptionCtrl: SubscriptionCtrl,
   }
}