package service

import (
	"fmt"
	"github.com/cilidm/toolbox/str"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"pear-admin-go/app/core/scli"
	"pear-admin-go/app/dao"
	"pear-admin-go/app/global"
	"pear-admin-go/app/global/request"
	"pear-admin-go/app/model"
	"pear-admin-go/app/util/pool"
	"pear-admin-go/app/util/remote"
	"runtime"
	"strings"
	"time"
)

func TaskJson(f request.TaskPage) ([]model.Task, int, error) {
	//filters := make([]interface{}, 0)
	//if f.ServerName != "" {
	//	filters = append(filters, "server_name LIKE ?", "%"+f.ServerName+"%")
	//}
	//if f.ServerIp != "" {
	//	filters = append(filters, "server_ip LIKE ?", "%"+f.ServerIp+"%")
	//}
	//if f.Detail != "" {
	//	filters = append(filters, "detail LIKE ?", "%"+f.Detail+"%")
	//}
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
	err = dao.NewTaskDaoImpl().Insert(s)
	if err != nil {
		return err
	}
	go RunTask(s)
	return nil
}

func RunTask(task model.Task) {
	s, err := dao.NewTaskServerDaoImpl().FindOne(task.DstServer)
	if err != nil {
		global.Log.Error("runTask.FindOne", zap.Error(err))
		return
	}
	if s == nil {
		global.Log.Error(fmt.Sprintf("server id nod found:%d", task.DstServer))
		return
	}
	cli, err := scli.Instance(*s)
	if err != nil {
		global.Log.Error("runTask.scli.Instance", zap.Error(err))
		return
	}
	rm := remote.NewRemote(task.SourcePath, task.DstPath, nil, cli)
	fp = pool.NewPool(10)
	WalkPath(task.SourcePath, rm, task.SourcePath)
}

var fp *pool.Pool

func WalkPath(dir string, rm *remote.Remote, sourceDir string) {
	paths, err := filepath.Glob(strings.TrimRight(dir, "/") + "/*")
	if err != nil {
		global.Log.Error("WalkPath.Glob", zap.Error(err))
		return
	}
	for _, v := range paths {
		//if str.IsContain(conf.Conf.FileInfo.ExceptDir, v) || util.HasPathPrefix(conf.Conf.FileInfo.ExceptDir, v) {
		//	continue
		//}
		stat, err := os.Stat(v)
		if err != nil {
			global.Log.Error("WalkPath.os.Stat", zap.Error(err))
			continue
		}
		if stat.IsDir() {
			dname := string([]rune(strings.ReplaceAll(v, sourceDir, ""))[1:])
			err = rm.Mkdir(dname)
			if err != nil {
				global.Log.Error("WalkPath.rm.Mkdir", zap.Error(err))
				continue
			}
			WalkPath(v, rm, sourceDir)
		} else {
			global.Log.Info(fmt.Sprintf("发现文件 【%s】", v))
			fp.Add(1)
			go func(v string, size int64) {
				defer fp.Done()
				err = rm.LocalToRemote(v, size)
				if err != nil {
					global.Log.Error("WalkPath.LocalToRemote", zap.Error(err))
				}
			}(v, stat.Size())
		}
	}
}
