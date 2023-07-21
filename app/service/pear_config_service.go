package service

import (
	"encoding/json"
	"go-admin/app/core/cache"
	"go-admin/app/dao"
	e2 "go-admin/app/global/e"
	"go-admin/app/model"
)

func GetPearConfig() (*model.PearConfigForm, error) {
	pear, err := dao.NewSiteConfigDaoImpl().FindOne(model.PearSiteConfig)
	if err != nil {
		return nil, err
	}
	var data model.PearConfigForm
	err = json.Unmarshal([]byte(pear.ConfigData), &data)
	if err != nil {
		return nil, err
	}
	cache.Instance().Set(e2.PearConfigCache, data, 0)
	return &data, nil
}
