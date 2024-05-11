package initial

import (
	"caipiaotong/configs/connect"
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"log"
)

func newRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	})
	return client
}
func InitCache() {
	connect.Cache = newRedisClient()
	if err := connect.Cache.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
	log.Println("redis connected")
}
