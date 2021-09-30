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
	go runtask.NewRunTask(s).SetSourceClient().SetDstClient().Run()	// todo 替换调RunTask 及remote
	return nil
}
//
//func RunTask(task model.Task) {
//	var sid int
//	if task.SourceType == 2 && task.DstType == 1 {
//		sid = task.SourceServer
//	} else if task.DstType == 2 && task.SourceType == 1 {
//		sid = task.DstServer
//	} else if task.SourceType == 2 && task.DstType == 2 {
//
//	}
//	s, err := dao.NewTaskServerDaoImpl().FindOne(sid)
//	if err != nil {
//		global.Log.Error("RunTask.FindOne", zap.Error(err))
//		return
//	}
//	if s == nil {
//		global.Log.Error(fmt.Sprintf("server id nod found:%d", task.DstServer))
//		return
//	}
//	cli, err := scli.Instance(*s)
//	if err != nil {
//		global.Log.Error("RunTask.scli.Instance", zap.Error(err))
//		return
//	}
//	fp = pool.NewPool(10)
//	var fcount int
//	if task.SourceType == 1 && task.DstType == 2 {
//		rm := remote.NewRemote(task.Id, task.DstServer, task.SourcePath, task.DstPath, nil, cli)
//		fcount = WalkLocalPathV2(task.SourcePath, rm, task.SourcePath)
//	} else if task.SourceType == 2 && task.DstType == 1 {
//		remote.NewRemote(task.Id, task.SourceServer, check.CheckWinPath(task.SourcePath), check.CheckWinPath(task.DstPath), cli, nil).
//			WalkRemotePath(task.SourcePath, fp)
//	}
//	err = dao.NewTaskDaoImpl().Update(task, map[string]interface{}{"task_file_num": fcount})
//	if err != nil {
//		global.Log.Error("RunTask.Update", zap.Error(err))
//		return
//	}
//}
//
//var fp *pool.Pool
//
//func WalkLocalPathV2(dir string, rm *remote.Remote, sourceDir string) int {
//	count := 0
//	_ = filepath.Walk(dir, func(v string, info fs.FileInfo, err error) error {
//		if err != nil {
//			global.Log.Error("WalkPathV2.Walk.err", zap.Error(err))
//			return nil
//		}
//		if info == nil {
//			global.Log.Error("WalkPathV2.Walk.info Is Nil")
//			return nil
//		}
//		stat, err := os.Stat(v)
//		if err != nil {
//			global.Log.Error("WalkPathV2.os.Stat", zap.Error(err))
//			return nil
//		}
//		if stat.IsDir() && v != sourceDir {
//			dname := string([]rune(strings.ReplaceAll(v, sourceDir, ""))[1:])
//			err = rm.Mkdir(dname)
//			if err != nil {
//				global.Log.Error("WalkPath.rm.Mkdir", zap.Error(err))
//				return nil
//			}
//		} else {
//			count++
//			fp.Add(1)
//			go func(v string, size int64) {
//				defer fp.Done()
//				err = rm.LocalToRemote(v, size)
//				if err != nil {
//					global.Log.Error("WalkPath.LocalToRemote", zap.Error(err))
//				}
//			}(v, stat.Size())
//		}
//		return nil
//	})
//	return count
//}
//
//// 递归方式无法准确获取文件查询的完成时间
//func WalkLocalPath(dir string, rm *remote.Remote, sourceDir string) {
//	paths, err := filepath.Glob(strings.TrimRight(dir, "/") + "/*")
//	if err != nil {
//		global.Log.Error("WalkPath.Glob", zap.Error(err))
//		return
//	}
//	for _, v := range paths {
//		stat, err := os.Stat(v)
//		if err != nil {
//			global.Log.Error("WalkPath.os.Stat", zap.Error(err))
//			continue
//		}
//		if stat.IsDir() {
//			dname := string([]rune(strings.ReplaceAll(v, sourceDir, ""))[1:])
//			err = rm.Mkdir(dname)
//			if err != nil {
//				global.Log.Error("WalkPath.rm.Mkdir", zap.Error(err))
//				continue
//			}
//			WalkLocalPath(v, rm, sourceDir)
//		} else {
//			fp.Add(1)
//			go func(v string, size int64) {
//				defer fp.Done()
//				err = rm.LocalToRemote(v, size)
//				if err != nil {
//					global.Log.Error("WalkPath.LocalToRemote", zap.Error(err))
//				}
//			}(v, stat.Size())
//		}
//	}
//}
