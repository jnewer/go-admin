package main

import (
	"context"
	"embed"
	"fmt"
	"github.com/cilidm/toolbox/gconv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"pear-admin-go/app/core"
	"pear-admin-go/app/global"
	"pear-admin-go/app/router"
	"pear-admin-go/app/util/validate"
	"syscall"
	"time"
)

//go:embed template
var templateFs embed.FS

//go:embed static
var staticFs embed.FS

func main() {
	if err := validate.InitTrans("zh"); err != nil {
		fmt.Println("init trans failed, err:", err)
	}

	core.InitConfig("./config.toml")

	global.Log = core.InitLog()

	global.DBConn = core.InitConn()

	core.InitRedis()
	r := router.InitRouter(staticFs, templateFs)
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", core.Conf.App.HttpPort),
		Handler:        r,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Printf(`	欢迎使用 Pear Admin Go
	程序运行地址:http://127.0.0.1:%s
`, gconv.String(core.Conf.App.HttpPort))
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Log.Error(err.Error())
			os.Exit(0)
		}
	}()

	shutDown(s)
}

func shutDown(s *http.Server) {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("Shutdown Server ...")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := global.DBConn.Close()
	if err != nil {
		log.Fatal("Close DB error:", err.Error())
	}

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	fmt.Println("Server exiting")
}
