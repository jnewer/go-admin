package controller

import (
	e2 "go-admin/app/global/e"
	"go-admin/app/service"
	pkg "go-admin/app/util/file"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	user := service.GetProfile(c)
	if pkg.CheckNotExist(service.GetImgSavePath(user.Avatar)) {
		user.Avatar = e2.DefaultAvatar
	}
	site, _ := service.GetSiteConf()
	c.HTML(http.StatusOK, "index.html", gin.H{
		"site":      site,
		"user":      user,
		"copyright": template.HTML(site.Copyright), // 防止转义
	})
}

func FramePage(c *gin.Context) {
	c.HTML(http.StatusOK, "main.html", gin.H{})
}
