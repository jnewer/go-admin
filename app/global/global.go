package global

import (
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

var (
	Log       *zap.Logger
	RedisConn *redis.Client
)
