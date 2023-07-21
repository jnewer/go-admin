package service

import (
	"go-admin/app/core/config"
	"go-admin/app/core/db"
	"testing"
)

func init() {
	config.InitConfig("../../config.toml")
	db.InitConn()
}

func TestRunTask(t *testing.T) {
	//task, _ := dao.NewTaskDaoImpl().FindOne(17)
	select {}
}
