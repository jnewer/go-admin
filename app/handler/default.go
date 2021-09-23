package controller

import (
	"github.com/cilidm/toolbox/OS"
	f "github.com/cilidm/toolbox/file"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"pear-admin-go/app/core"
	"pear-admin-go/app/global/api/response"
	"pear-admin-go/app/model"
	"pear-admin-go/app/service"
	"pear-admin-go/app/util/e"
	"pear-admin-go/app/util/gocache"
	"strings"
	"time"
)

func DefaultUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).SetType(model.OperAdd).WriteJsonExit()
		return
	}
	if file.Size > e.DefUploadSize {
		response.ErrorResp(c).SetMsg("文件大小超限").SetType(model.OperAdd).Log(e.DefaultUpload, file).WriteJsonExit()
		return
	}
	day := time.Now().Format(e.TimeFormatDay)
	savePath := filepath.Join(core.Conf.App.ImgSavePath, day) // 按年月日归档保存
	err = f.IsNotExistMkDir(savePath)
	if err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).SetType(model.OperAdd).Log(e.DefaultUpload, file).WriteJsonExit()
		return
	}
	if err := c.SaveUploadedFile(file, filepath.Join(savePath, file.Filename)); err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).SetType(model.OperAdd).Log(e.DefaultUpload, file).WriteJsonExit()
		return
	}
	backFilePath := filepath.Join(filepath.Join(core.Conf.App.ImgUrlPath, day), file.Filename)
	if OS.IsWindows() {
		backFilePath = strings.ReplaceAll(backFilePath, "\\", "/")
	}
	response.SuccessResp(c).SetData(backFilePath).SetType(model.OperAdd).Log(e.DefaultUpload, file).WriteJsonExit()
}

func PearConfig(c *gin.Context) {
	var (
		data *model.PearConfigForm
		err  error
	)
	conf, found := gocache.Instance().Get(e.PearConfigCache)
	if found && conf != nil {
		d, ok := conf.(model.PearConfigForm)
		if ok {
			data = &d
		} else {
			response.ErrorResp(c).WriteJsonExit()
			return
		}
	} else {
		data, err = service.GetPearConfig()
		if err != nil {
			response.ErrorResp(c).SetMsg(err.Error()).WriteJsonExit()
			return
		}
	}
	c.JSON(http.StatusOK, data)
}

func GetMenu(c *gin.Context) {
	menuResp := service.MenuServiceV2(c)
	c.JSON(http.StatusOK, menuResp.MenuResp)
}

func ServerErr(c *gin.Context) {
	c.HTML(http.StatusOK, "server_err.html", nil)
}

func ShowFile(c *gin.Context) {
	fp := c.Query("filePath")
	f, err := os.Open(fp)
	if err != nil {
		c.Redirect(http.StatusFound, "/not_found")
		return
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		c.Redirect(http.StatusFound, "/not_found")
		return
	}
	c.String(http.StatusOK, string(b))
}