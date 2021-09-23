package controller

import (
	pkg "github.com/cilidm/toolbox/file"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"pear-admin-go/app/service"
	"pear-admin-go/app/util/e"
)

func Index(c *gin.Context) {
	user := service.GetProfile(c)
	if pkg.CheckNotExist(service.GetImgSavePath(user.Avatar)) {
		user.Avatar = e.DefaultAvatar
	}
	site, _ := service.GetSiteConf()
	c.HTML(http.StatusOK, "index.html", gin.H{
		"site":      site,
		"user":      user,
		"copyright": template.HTML(site.Copyright), // 防止转义
	})
}

func FrameIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "main.html", gin.H{})
}
