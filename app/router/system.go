package router

import (
	controller "pear-admin-go/app/handler"
	"pear-admin-go/app/middleware"
)

func SystemInit() {
	common := New("system", "/")
	common.GET("/", middleware.CheckDefaultPage)
	common.GET("login", middleware.CheckLoginPage, controller.Login)
	common.POST("login", controller.LoginHandler)
	common.GET("logout", controller.Logout)
	common.POST("isLogin", nil)
	common.GET("not_found", controller.NotFound)
	common.GET("captcha", controller.GetCaptcha)
	common.POST("captcha_verify", controller.CaptchaVerify)

	r := New("system", "/system", middleware.AuthMiddleware)
	// index
	r.GET("/", controller.Index)
	r.GET("index", controller.Index)
	r.GET("main", controller.FrameIndex)

	// log
	r.GET("/log/list", controller.LogList)
	r.GET("/log/log_operate", controller.LogOperate)
	r.GET("/log/log_login", controller.LogLogin)

	// default
	r.POST("upload/def_upload", controller.DefaultUpload)
	r.GET("pear_config", controller.PearConfig)
	r.GET("menu_config", controller.GetMenu)
	r.GET("server_err", controller.ServerErr)
	r.GET("file", controller.ShowFile)

	// role
	r.GET("admin/list", controller.AdminList) // 管理员列表页
	r.GET("admin/list_json", controller.AdminListJson)
	r.GET("admin/add", controller.AdminAdd)
	r.POST("admin/add", controller.AdminAddHandler)
	r.GET("admin/edit", controller.AdminEdit)
	r.POST("admin/edit", controller.AdminEditHandler)
	r.POST("admin/edit_status", controller.AdminChangeStatus)
	r.DELETE("admin/delete", controller.AdminDelete)
	r.GET("role/list", controller.RoleList) // 角色管理列表页
	r.GET("role/list_json", controller.RoleListJson)
	r.GET("role/add", controller.RoleAdd)
	r.POST("role/add", controller.RoleAddHandler)
	r.GET("role/power", controller.RolePower)
	r.GET("role/getRolePower", controller.GetRolePower)
	r.POST("role/saveRolePower", controller.SaveRolePower)
	r.GET("role/edit", controller.RoleEdit)
	r.POST("role/edit", controller.RoleEditHandler)
	r.POST("role/delete", controller.RoleDeleteHandler)
	r.GET("auth/list", controller.AuthList)             // 权限因子列表页
	r.POST("auth/edit", controller.AuthNodeEdit)        // 新增、修改权限
	r.GET("auth/get_nodes", controller.GetNodes)        // 权限配置页面
	r.GET("auth/add", controller.AddNode)               // 新增权限
	r.GET("auth/edit", controller.EditNode)             // 修改权限
	r.POST("auth/get_node", controller.GetNode)         // 权限因子列表页
	r.POST("auth/delete", controller.AuthDelete)        // 权限因子列表页
	r.GET("auth/selectParent", controller.SelectParent) // 权限列表
	r.GET("ui/icon", controller.IconShow)

	// sysconf
	r.GET("site/list", controller.SiteList)
	r.POST("site/edit", controller.SiteEdit)
	r.GET("mail/list", controller.MailList)
	r.POST("mail/edit", controller.MailEdit)
	r.POST("mail/test", controller.MailTest)

	// user
	r.GET("user/edit", controller.UserShow)
	r.GET("user/uploadPage", controller.UploadPage)
	r.POST("user/edit", controller.ProfileEdit)
	r.POST("user/avatar", controller.AvatarEdit)
	r.GET("user/pwd", controller.PwdEdit)
	r.POST("user/pwd", controller.PwdEditHandler)
}
