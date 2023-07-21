package router

import (
	"embed"
	"go-admin/app/core/config"
	"html/template"
	"io/fs"
	"net/http"

	"go-admin/app/middleware"
	"go-admin/app/util/session"

	"github.com/gin-gonic/gin"
)

func InitRouter(staticFs, templateFs embed.FS) *gin.Engine {
	gin.SetMode(config.Instance().App.RunMode)
	r := gin.New()

	t, _ := template.ParseFS(templateFs, "template/**/**/*.html")
	r.SetHTMLTemplate(t)

	r.Static(config.Instance().App.ImgUrlPath, config.Instance().App.ImgSavePath)
	r.Static("/runtime/file", "runtime/file")

	fads, _ := fs.Sub(staticFs, "static")
	r.StaticFS("/static", http.FS(fads))

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	r.Use(session.EnableCookieSession(config.Instance().App.JwtSecret))
	CommonRouter(r)
	SystemRouter(r)
	TaskRouter(r)
	return r
}
