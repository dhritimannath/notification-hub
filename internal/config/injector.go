// go:build wireinject
//go:build wireinject
// +build wireinject

package config

import (
	"github.com/dhritimannath/notification-service/internal/controllers"
	"github.com/dhritimannath/notification-service/internal/repositories"
	"github.com/dhritimannath/notification-service/internal/services"
	"github.com/google/wire"
)

var db = wire.NewSet(InitDb)

var userServiceSet = wire.NewSet(services.UserServiceInit,
 wire.Bind(new(services.UserService), new(*services.UserServiceImpl)),
)

var userRepoSet = wire.NewSet(repositories.UserRepositoryInit,
 wire.Bind(new(repositories.UserRepository), new(*repositories.UserRepositoryImpl)),
)

var userCtrlSet = wire.NewSet(controllers.UserControllerInit,
 wire.Bind(new(controllers.UserController), new(*controllers.UserControllerImpl)),
)

var subscriptionServiceSet = wire.NewSet(services.SubscriptionServiceInit,
	wire.Bind(new(services.SubscriptionService), new(*services.SubscriptionServiceImpl)),
)
   
var subscriptionRepoSet = wire.NewSet(repositories.SubscriptionRepositoryInit,
wire.Bind(new(repositories.SubscriptionRepository), new(*repositories.SubscriptionRepositoryImpl)),
)

var subscriptionCtrlSet = wire.NewSet(controllers.SubscriptionControllerInit,
wire.Bind(new(controllers.SubscriptionController), new(*controllers.SubscriptionControllerImpl)),
)

func Init() *Initialization {
 wire.Build(NewInitialization, db, userCtrlSet, userServiceSet, userRepoSet, subscriptionCtrlSet, subscriptionServiceSet, subscriptionRepoSet)
 return nil
}