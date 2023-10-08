package server

import (
	"github.com/dhritimannath/notification-service/internal/config"
	"github.com/dhritimannath/notification-service/internal/router"
)

func Init(init *config.Initialization)  {
	r := router.Init(init)
	r.Run("localhost:3000")
}