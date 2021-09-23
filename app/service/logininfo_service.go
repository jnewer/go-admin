package service

import (
	dao2 "pear-admin-go/app/dao"
	"pear-admin-go/app/global"
	"pear-admin-go/app/global/api/request"
	"pear-admin-go/app/model"
)

func LoginInfoListJsonService(f request.LayerListForm) (count int, list []model.LoginInfo, err error) {
	if f.Page == 0 {
		f.Page = 1
	}
	if f.Limit == 0 {
		f.Limit = 10
	}
	list, count, err = dao2.NewLoginInfoImpl().FindByPage(f.Page, f.Limit)
	if err != nil {
		global.Log.Error("LoginInfoListJsonService.FindByPage:" + err.Error())
		return 0, nil, err
	}
	return count, list, nil
}
