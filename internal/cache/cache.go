package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"log"
)

var client *redis.Client
var ctxBg = context.Background()

func newRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	})
	return client
}
func InitCache() {
	client = newRedisClient()
	if err := client.Ping(ctxBg).Err(); err != nil {
		panic(err)
	}
	log.Println("redis connected")
}
func GetCacheClient() *redis.Client {
	return client
}
