package controller

import (
	pkg "pear-admin-go/app/util/file"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	e2 "pear-admin-go/app/global/e"
	"pear-admin-go/app/service"
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
