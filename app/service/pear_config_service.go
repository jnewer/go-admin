package service

import (
	"encoding/json"
	"pear-admin-go/app/core/cache"
	"pear-admin-go/app/dao"
	"pear-admin-go/app/model"
	"pear-admin-go/app/util/e"
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
	cache.Instance().Set(e.PearConfigCache, data, 0)
	return &data, nil
}
