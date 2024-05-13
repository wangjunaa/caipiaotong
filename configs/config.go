package configs

import (
	"github.com/spf13/viper"
	"log"
)

func InitConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("configs")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	log.Println("config init success")
}
