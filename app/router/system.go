package router

import (
	"github.com/gin-gonic/gin"
	controller "pear-admin-go/app/handler"
	"pear-admin-go/app/middleware"
)

func SystemRouter(r *gin.Engine) {
	sr := r.Group("system", middleware.AuthMiddleware)

	// index
	sr.GET("/", controller.Index)
	sr.GET("index", controller.Index)
	sr.GET("main", controller.FrameIndex)

	// log
	sr.GET("/log/list", controller.LogList)
	sr.GET("/log/log_operate", controller.LogOperate)
	sr.GET("/log/log_login", controller.LogLogin)

	// default
	sr.POST("upload/def_upload", controller.DefaultUpload)
	sr.GET("pear_config", controller.PearConfig)
	sr.GET("menu_config", controller.GetMenu)
	sr.GET("server_err", controller.ServerErr)
	sr.GET("file", controller.ShowFile)

	// role
	sr.GET("admin/list", controller.AdminList) // 管理员列表页
	sr.GET("admin/list_json", controller.AdminListJson)
	sr.GET("admin/add", controller.AdminAdd)
	sr.POST("admin/add", controller.AdminAddHandler)
	sr.GET("admin/edit", controller.AdminEdit)
	sr.POST("admin/edit", controller.AdminEditHandler)
	sr.POST("admin/edit_status", controller.AdminChangeStatus)
	r.DELETE("admin/delete", controller.AdminDelete)
	sr.GET("role/list", controller.RoleList) // 角色管理列表页
	sr.GET("role/list_json", controller.RoleListJson)
	sr.GET("role/add", controller.RoleAdd)
	sr.POST("role/add", controller.RoleAddHandler)
	sr.GET("role/power", controller.RolePower)
	sr.GET("role/getRolePower", controller.GetRolePower)
	sr.POST("role/saveRolePower", controller.SaveRolePower)
	sr.GET("role/edit", controller.RoleEdit)
	sr.POST("role/edit", controller.RoleEditHandler)
	sr.POST("role/delete", controller.RoleDeleteHandler)
	sr.GET("auth/list", controller.AuthList)             // 权限因子列表页
	sr.POST("auth/edit", controller.AuthNodeEdit)        // 新增、修改权限
	sr.GET("auth/get_nodes", controller.GetNodes)        // 权限配置页面
	sr.GET("auth/add", controller.AddNode)               // 新增权限
	sr.GET("auth/edit", controller.EditNode)             // 修改权限
	sr.POST("auth/get_node", controller.GetNode)         // 权限因子列表页
	sr.POST("auth/delete", controller.AuthDelete)        // 权限因子列表页
	sr.GET("auth/selectParent", controller.SelectParent) // 权限列表
	sr.GET("ui/icon", controller.IconShow)

	// sysconf
	sr.GET("site/list", controller.SiteList)
	sr.POST("site/edit", controller.SiteEdit)
	sr.GET("mail/list", controller.MailList)
	sr.POST("mail/edit", controller.MailEdit)
	sr.POST("mail/test", controller.MailTest)

	// user
	sr.GET("user/edit", controller.UserShow)
	sr.GET("user/uploadPage", controller.UploadPage)
	sr.POST("user/edit", controller.ProfileEdit)
	sr.POST("user/avatar", controller.AvatarEdit)
	sr.GET("user/pwd", controller.PwdEdit)
	sr.POST("user/pwd", controller.PwdEditHandler)
}
