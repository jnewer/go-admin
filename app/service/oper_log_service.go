package service

import (
	"encoding/json"
	"github.com/cilidm/toolbox/ip"
	"github.com/gin-gonic/gin"
	dao2 "pear-admin-go/app/dao"
	"pear-admin-go/app/global"
	"pear-admin-go/app/global/api/request"
	"pear-admin-go/app/model"
	"pear-admin-go/app/util/e"

	"time"
)

func CreateOperLog(c *gin.Context, f model.OperForm) error {
	var operName string
	user := GetProfile(c)
	if user == nil {
		operName = ""
	} else {
		operName = user.LoginName
	}
	outJson, _ := json.Marshal(f.OutContent)
	var oper model.OperLog
	oper.Title = f.Title
	oper.OperParam = f.InContent
	oper.JsonResult = string(outJson)
	oper.BusinessType = int(f.OutContent.Type)
	oper.OperatorType = 1 // 操作类别（0其它 1后台用户 2手机端用户）
	if f.OutContent.Code == 0 {
		oper.Status = 0
	} else {
		oper.Status = 1
	}
	oper.OperName = operName
	oper.RequestMethod = c.Request.Method
	oper.OperUrl = c.Request.URL.Path
	oper.Method = c.Request.Method
	oper.OperIp = c.ClientIP()
	oper.OperLocation = ip.GetCityByIp(oper.OperIp)
	oper.ErrorMsg = f.ErrorMsg
	oper.OperTime = time.Now().Format(e.TimeFormat)
	if err := dao2.NewOperLogDaoImpl().Insert(oper); err != nil {
		return err
	}
	return nil
}

func OperLogListJsonService(f request.LayerListForm) (count int, list []model.OperLog, err error) {
	if f.Page == 0 {
		f.Page = 1
	}
	if f.Limit == 0 {
		f.Limit = 10
	}
	list, count, err = dao2.NewOperLogDaoImpl().FindByPage(f.Page, f.Limit)
	if err != nil {
		global.Log.Error("LoginInfoListJsonService.FindByPage:" + err.Error())
		return 0, nil, err
	}
	return count, list, nil
}