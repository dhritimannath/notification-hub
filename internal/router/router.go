package router

import (
	"github.com/dhritimannath/notification-service/internal/config"
	"github.com/gin-gonic/gin"
)

func Init(init *config.Initialization) *gin.Engine  {
	router := gin.Default()
	api := router.Group("/api")
 	{
  		user := api.Group("/user")
		{
			user.POST("", init.UserCtrl.Create)
		  	user.GET("", init.UserCtrl.FindAll)
		  	user.GET(":id", init.UserCtrl.FindById)
			
			user.POST(":id/subscription", init.SubscriptionCtrl.Create)
			user.GET(":id/subscription", init.SubscriptionCtrl.FindAllByUserId)
			user.PATCH(":id/subscription/:subscriptionId", init.SubscriptionCtrl.UpdateById)
			user.DELETE(":id/subscription/:subscriptionId", init.SubscriptionCtrl.DeleteById)
		}
 	}
	return router
}