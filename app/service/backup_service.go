package service

import (
	"github.com/cilidm/toolbox/str"
	"pear-admin-go/app/dao"
	"pear-admin-go/app/global/request"
	"pear-admin-go/app/model"
	"time"
)

func TaskAdd(f request.TaskForm) error {
	var s model.Backup
	err := str.CopyFields(&s, f)
	if err != nil {
		return err
	}
	s.CreateTime = time.Now()
	err = dao.NewTaskDaoImpl().Insert(s)
	if err != nil {
		return err
	}
	return nil
}
