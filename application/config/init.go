package config

import (
	"github.com/spf13/viper"
	"log"
	"rest-api/application/config/database"
)

func Init() {
	InitViper()
	database.DbConnect()
}

// InitViper from file toml
func InitViper() {
	viper.SetConfigFile("toml")
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("viper not use")
	}
}
