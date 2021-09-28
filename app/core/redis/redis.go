package redis

import (
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"pear-admin-go/app/core/config"
	"pear-admin-go/app/global"
)

func InitRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Instance().Redis.RedisAddr,
		Password: config.Instance().Redis.RedisPWD, // no password set
		DB:       config.Instance().Redis.RedisDB,  // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		global.Log.Error("redis connect ping failed, err:", zap.Any("err", err))
	} else {
		global.Log.Info("redis connect ping response:", zap.String("pong", pong))
		global.RedisConn = client
	}
}
