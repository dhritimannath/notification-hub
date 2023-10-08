package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDb() *gorm.DB {
	dbURL := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
	AppConfig.DbHost, AppConfig.DbPort, AppConfig.DbName,
	AppConfig.DbUsername, AppConfig.DbPassword)
    con, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

    if err != nil {
        log.Fatalln (err)
    }
	return con
}