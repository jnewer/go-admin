package zapLog

import (
	"go.uber.org/zap"
	"pear-admin-go/app/core/log"
	"pear-admin-go/app/global"
)

func NewLog() *ZapLog {
	z := new(ZapLog)
	if global.Log == nil {
		log.InitLog()
	}
	z.log = global.Log
	return z
}

type ZapLog struct {
	log *zap.Logger
}

func (z *ZapLog) Info(msg, key string, infoMsg interface{}) {
	z.log.Info(msg, zap.Any(key, infoMsg))
}

func (z *ZapLog) Error(msg, key string, errMsg interface{}) {
	z.log.Error(msg, zap.Any(key, errMsg))
}

func (z *ZapLog) Warn(msg, key string, infoMsg interface{}) {
	z.log.Warn(msg, zap.Any(key, infoMsg))
}
