package service

import (
	"github.com/cilidm/toolbox/str"
	"pear-admin-go/app/dao"
	"pear-admin-go/app/global/request"
	"pear-admin-go/app/model"
	"pear-admin-go/app/util/runtask"
	"runtime"
	"strings"
	"time"
)

func TaskJson(f request.TaskPage) ([]model.Task, int, error) {
	ts, count, err := dao.NewTaskDaoImpl().FindByPage(f.Page, f.Limit)
	if err != nil {
		return nil, 0, err
	}
	return ts, count, nil
}

func TaskAdd(f request.TaskForm) error {
	var s model.Task
	err := str.CopyFields(&s, f)
	if err != nil {
		return err
	}
	if runtime.GOOS == "windows" {
		f.SourcePath = strings.ReplaceAll(f.SourcePath, "\\", "/")
		f.DstPath = strings.ReplaceAll(f.DstPath, "\\", "/")
	}
	s.CreateTime = time.Now()
	tid, err := dao.NewTaskDaoImpl().Insert(s)
	if err != nil {
		return err
	}
	s.Id = tid
	go runtask.NewRunTask(s).SetSourceClient().SetDstClient().Run()
	return nil
}
