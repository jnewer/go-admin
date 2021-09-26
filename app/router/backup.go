package router

import (
	"github.com/gin-gonic/gin"
	"pear-admin-go/app/controller"
)

func BackupRouter(r *gin.Engine) {
	br := r.Group("system")
	br.GET("server/list", controller.ServerList)
	br.GET("server/json", controller.ServerJson)
	br.GET("server/add", controller.ServerAdd)
	br.POST("server/add", controller.ServerAdd)
	br.GET("server/edit", controller.ServerEdit)
	br.POST("server/edit", controller.ServerEdit)
	br.POST("server/del", controller.ServerDel)

	br.GET("backup/list", controller.TaskList)
	br.GET("backup/json", controller.ServerJson)
	br.GET("backup/add", controller.TaskAdd)
	br.POST("backup/add", controller.TaskAdd)
	br.GET("backup/edit", controller.TaskEdit)
	br.POST("backup/edit", controller.TaskEdit)
	br.POST("backup/del", controller.TaskDel)
}
