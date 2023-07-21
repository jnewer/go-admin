package redis

import (
	"go-admin/app/core/config"
	"go-admin/app/core/log"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

var redisCli *redis.Client

func Instance() *redis.Client {
	if redisCli == nil {
		InitRedis()
	}
	return redisCli
}

func InitRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Instance().Redis.RedisAddr,
		Password: config.Instance().Redis.RedisPWD, // no password set
		DB:       config.Instance().Redis.RedisDB,  // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		log.Instance().Error("redis connect ping failed, err:", zap.Any("err", err))
	} else {
		log.Instance().Info("redis connect ping response:", zap.String("pong", pong))
		redisCli = client
	}
}
