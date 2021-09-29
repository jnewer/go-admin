package remote

import (
	"fmt"
	"github.com/cilidm/toolbox/OS"
	"github.com/pkg/sftp"
	"go.uber.org/zap"
	"os"
	"path"
	"pear-admin-go/app/dao"
	"pear-admin-go/app/global"
	"pear-admin-go/app/model"
	"runtime"
	"strings"
	"time"
)

type Remote struct {
	taskId       int
	serverId     int
	sourcePath   string       // 源路径
	dstPath      string       // 目标路径
	sourceClient *sftp.Client // 源服务器连接
	dstClient    *sftp.Client // 目标服务器连接
}

func NewRemote(taskId int, serverId int, sourcePath string, dstPath string, sourceClient *sftp.Client, dstClient *sftp.Client) *Remote {
	return &Remote{taskId: taskId, serverId: serverId, sourcePath: sourcePath, dstPath: dstPath, sourceClient: sourceClient, dstClient: dstClient}
}

// 校验目标地址权限
func (this *Remote) CheckAccess() {

}

func (this *Remote) LocalToRemote(fname string, fsize int64) error {
	if OS.IsWindows() {
		this.sourcePath = strings.ReplaceAll(this.sourcePath, "\\", "/")
		this.dstPath = strings.ReplaceAll(this.dstPath, "\\", "/")
		fname = strings.ReplaceAll(fname, "\\", "/")
		fname = strings.ReplaceAll(fname, this.sourcePath, "")
	}
	rf := path.Join(this.dstPath, fname) // 文件在服务器的路径及名称
	has, err := this.dstClient.Stat(rf)
	if err == nil && (has.Size() == fsize) {
		global.Log.Debug(fmt.Sprintf("文件%s已存在", rf))
		return nil
	}
	err = this.dstClient.MkdirAll(this.dstPath)
	if err != nil {
		return err
	}
	err = this.dstClient.Chmod(this.dstPath, os.ModePerm)
	if err != nil {
		return err
	}
	srcFile, err := os.Open(path.Join(this.sourcePath, fname))
	if err != nil {
		global.Log.Error("源文件无法读取", zap.Error(err))
		return err
	}
	defer srcFile.Close()
	dstFile, err := this.dstClient.Create(rf) // 如果文件存在，create会清空原文件 openfile会追加
	if err != nil {
		global.Log.Error("this.dstClient.Create", zap.Error(err))
		return err
	}
	defer dstFile.Close()

	buf := make([]byte, 10000)
	for {
		n, _ := srcFile.Read(buf)
		if n == 0 {
			break
		}
		dstFile.Write(buf[:n]) // 读多少 写多少
	}
	err = dao.NewTaskLogDaoImpl().Insert(model.TaskLog{
		TaskId:     this.taskId,
		ServerId:   this.serverId,
		SourcePath: path.Join(this.sourcePath, fname),
		DstPath:    rf,
		Size:       fsize,
		CreateTime: time.Now(),
	})
	if err != nil {
		return err
	}
	global.Log.Info(fmt.Sprintf("【%s】传输完毕", fname))
	return nil
}

func (this *Remote) Mkdir(p string) error {
	if runtime.GOOS == "windows" {
		p = strings.ReplaceAll(p, "\\", "/")
	}
	dst := path.Join(this.dstPath, p)
	fmt.Println(dst)
	err := this.dstClient.MkdirAll(dst)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) RemoteToRemote() {

}

func (this *Remote) RemoteToLocal() {

}
