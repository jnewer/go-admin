package service

import (
	"pear-admin-go/app/core/config"
	"pear-admin-go/app/core/db"
	log2 "pear-admin-go/app/core/log"
	"pear-admin-go/app/dao"
	"pear-admin-go/app/global"
	"testing"
)

func init() {
	config.InitConfig("../../config.toml")
	global.Log = log2.InitLog()
	db.Instance() = db.InitConn()
}

func TestRunTask(t *testing.T) {
	task, _ := dao.NewTaskDaoImpl().FindOne(6)
	RunTask(*task)
	select {}
}
