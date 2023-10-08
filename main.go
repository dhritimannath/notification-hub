package main

import (
	"github.com/dhritimannath/notification-service/internal/config"
	"github.com/dhritimannath/notification-service/internal/server"
)

func main()  {
	config.InitConfig()
	init := config.Init()
	server.Init(init)
}