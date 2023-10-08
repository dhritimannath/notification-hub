package config

import (
	"log"

	"github.com/spf13/viper"
)

var AppConfig Config

type Config struct {
	AppPort uint
	DbHost string
	DbPort uint
	DbUsername string
	DbPassword string
	DbName string
}

func InitConfig() {
	var err error
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file")
	}

	AppConfig.AppPort = viper.GetUint("APP_PORT")
	AppConfig.DbHost = viper.GetString("DB_HOST")
	AppConfig.DbPort = viper.GetUint("DB_PORT")
	AppConfig.DbUsername = viper.GetString("DB_USERNAME")
	AppConfig.DbPassword = viper.GetString("DB_PASSWORD")
	AppConfig.DbName = viper.GetString("DB_NAME")
}