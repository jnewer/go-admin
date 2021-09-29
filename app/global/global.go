package global

import (
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

var (
	//DBConn    *gorm.DB
	Log       *zap.Logger
	RedisConn *redis.Client
)
