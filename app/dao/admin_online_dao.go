package dao

import (
	"pear-admin-go/app/global"
	"pear-admin-go/app/model"
)

type AdminOnlineDao interface {
	Insert(online model.AdminOnline) error
	Delete(sessionID string) error
}

func NewAdminOnlineDaoImpl() AdminOnlineDao {
	online := new(AdminOnlineDaoImpl)
	return online
}

type AdminOnlineDaoImpl struct {
}

func (a *AdminOnlineDaoImpl) Insert(online model.AdminOnline) error {
	err := global.DBConn.Create(&online).Error
	return err
}

func (a *AdminOnlineDaoImpl) Delete(sessionID string) error {
	err := global.DBConn.Where("session_id = ?", sessionID).Delete(model.AdminOnline{}).Error
	return err
}
