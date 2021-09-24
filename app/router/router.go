package router

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"
	"pear-admin-go/app/core/config"

	"github.com/gin-gonic/gin"
	"pear-admin-go/app/middleware"
	"pear-admin-go/app/util/session"
)

func InitRouter(staticFs, templateFs embed.FS) *gin.Engine {
	gin.SetMode(config.Conf.App.RunMode)
	r := gin.New()

	t, _ := template.ParseFS(templateFs, "template/**/**/*.html")
	r.SetHTMLTemplate(t)

	r.Static(config.Conf.App.ImgUrlPath, config.Conf.App.ImgSavePath)
	r.Static("/runtime/file", "runtime/file")

	fads, _ := fs.Sub(staticFs, "static")
	r.StaticFS("/static", http.FS(fads))

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	r.Use(session.EnableCookieSession(config.Conf.App.JwtSecret))
	
	CommonRouter(r)
	SystemRouter(r)
	return r
}
