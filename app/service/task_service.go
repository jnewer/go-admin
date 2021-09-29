package service

import (
	"fmt"
	"github.com/cilidm/toolbox/str"
	"go.uber.org/zap"
	"io/fs"
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
	tid, err := dao.NewTaskDaoImpl().Insert(s)
	if err != nil {
		return err
	}
	s.Id = tid
	go RunTask(s)
	return nil
}

func RunTask(task model.Task) {
	s, err := dao.NewTaskServerDaoImpl().FindOne(task.DstServer)
	if err != nil {
		global.Log.Error("RunTask.FindOne", zap.Error(err))
		return
	}
	if s == nil {
		global.Log.Error(fmt.Sprintf("server id nod found:%d", task.DstServer))
		return
	}
	cli, err := scli.Instance(*s)
	if err != nil {
		global.Log.Error("RunTask.scli.Instance", zap.Error(err))
		return
	}
	rm := remote.NewRemote(task.Id, task.DstServer, task.SourcePath, task.DstPath, nil, cli)
	fp = pool.NewPool(10)
	fcount := WalkPathV2(task.SourcePath, rm, task.SourcePath)
	//fcount, err := dao.NewTaskLogDaoImpl().CountByTaskId(task.Id)
	//if err != nil {
	//	global.Log.Error("RunTask.CountByTaskId", zap.Error(err))
	//	return
	//}
	err = dao.NewTaskDaoImpl().Update(task, map[string]interface{}{"task_file_num": fcount})
	if err != nil {
		global.Log.Error("RunTask.Update", zap.Error(err))
		return
	}
}

var fp *pool.Pool

func WalkPathV2(dir string, rm *remote.Remote, sourceDir string) int {
	count := 0
	_ = filepath.Walk(dir, func(v string, info fs.FileInfo, err error) error {
		if err != nil {
			global.Log.Error("WalkPathV2.Walk.err", zap.Error(err))
			return nil
		}
		if info == nil {
			global.Log.Error("WalkPathV2.Walk.info Is Nil")
			return nil
		}
		stat, err := os.Stat(v)
		if err != nil {
			global.Log.Error("WalkPathV2.os.Stat", zap.Error(err))
			return nil
		}
		if stat.IsDir() && v != sourceDir {
			dname := string([]rune(strings.ReplaceAll(v, sourceDir, ""))[1:])
			err = rm.Mkdir(dname)
			if err != nil {
				global.Log.Error("WalkPath.rm.Mkdir", zap.Error(err))
				return nil
			}
		} else {
			count++
			fp.Add(1)
			go func(v string, size int64) {
				defer fp.Done()
				err = rm.LocalToRemote(v, size)
				if err != nil {
					global.Log.Error("WalkPath.LocalToRemote", zap.Error(err))
				}
			}(v, stat.Size())
		}
		return nil
	})
	return count
}

// 递归方式无法准确获取文件查询的完成时间
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
			//global.Log.Info(fmt.Sprintf("发现文件 【%s】", v))
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
