package config

import (
	"github.com/spf13/viper"
	"log"
)

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
