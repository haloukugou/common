package bootstrap

import (
	"dj/config"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var Redis *redis.Client

func RedisConnect() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Config.Redis.Host + ":" + fmt.Sprintf("%d", config.Config.Redis.Port),
		Password: config.Config.Redis.Auth,
		DB:       config.Config.Redis.Db,
	})
	return client
}
