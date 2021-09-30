package remote

//type Remote struct {
//	taskId       int          // 添加log时使用
//	serverId     int          // 添加log时使用
//	sourcePath   string       // 源路径
//	dstPath      string       // 目标路径
//	sourceClient *sftp.Client // 源服务器连接
//	dstClient    *sftp.Client // 目标服务器连接
//}
//
//func NewRemote(taskId int, serverId int, sourcePath string, dstPath string, sourceClient *sftp.Client, dstClient *sftp.Client) *Remote {
//	return &Remote{taskId: taskId, serverId: serverId, sourcePath: sourcePath, dstPath: dstPath, sourceClient: sourceClient, dstClient: dstClient}
//}
//
//// 校验目标地址权限
//func (this *Remote) CheckAccess() {
//
//}
//
//func (this *Remote) LocalToRemote(fname string, fsize int64) error {
//	if OS.IsWindows() {
//		this.sourcePath = strings.ReplaceAll(this.sourcePath, "\\", "/")
//		this.dstPath = strings.ReplaceAll(this.dstPath, "\\", "/")
//		fname = strings.ReplaceAll(fname, "\\", "/")
//		fname = strings.ReplaceAll(fname, this.sourcePath, "")
//	}
//	rf := path.Join(this.dstPath, fname) // 文件在服务器的路径及名称
//	has, err := this.dstClient.Stat(rf)
//	if err == nil && (has.Size() == fsize) {
//		global.Log.Debug(fmt.Sprintf("文件%s已存在", rf))
//		return nil
//	}
//	err = this.dstClient.MkdirAll(this.dstPath)
//	if err != nil {
//		return err
//	}
//	err = this.dstClient.Chmod(this.dstPath, os.ModePerm)
//	if err != nil {
//		return err
//	}
//	srcFile, err := os.Open(path.Join(this.sourcePath, fname))
//	if err != nil {
//		global.Log.Error("源文件无法读取", zap.Error(err))
//		return err
//	}
//	defer srcFile.Close()
//	dstFile, err := this.dstClient.Create(rf) // 如果文件存在，create会清空原文件 openfile会追加
//	if err != nil {
//		global.Log.Error("this.dstClient.Create", zap.Error(err))
//		return err
//	}
//	defer dstFile.Close()
//
//	buf := make([]byte, 10000)
//	for {
//		n, _ := srcFile.Read(buf)
//		if n == 0 {
//			break
//		}
//		dstFile.Write(buf[:n]) // 读多少 写多少
//	}
//	err = dao.NewTaskLogDaoImpl().Insert(model.TaskLog{
//		TaskId:     this.taskId,
//		ServerId:   this.serverId,
//		SourcePath: path.Join(this.sourcePath, fname),
//		DstPath:    rf,
//		Size:       fsize,
//		CreateTime: time.Now(),
//	})
//	if err != nil {
//		return err
//	}
//	global.Log.Info(fmt.Sprintf("【%s】传输完毕", fname))
//	return nil
//}
//
//func (this *Remote) RemoteToRemote() {
//
//}
//
//// 远端->本地 使用 sourceClient
//func (this *Remote) RemoteToLocal(fname string, fsize int64) error { // 本地文件夹
//	dstFile := path.Join(this.dstPath, strings.ReplaceAll(fname, this.sourcePath, "")) // 需要保存的本地文件地址
//	has, err := check.CheckFile(dstFile)                                               // 是否已存在
//	if err != nil {
//		global.Log.Error("RemoteToLocal.CheckFile", zap.Error(err))
//		return err
//	}
//	if has != nil && has.Size() == fsize {
//		global.Log.Debug(fmt.Sprintf("文件%s已存在", dstFile))
//		return nil
//	}
//	dir, _ := path.Split(dstFile)
//	err = file.IsNotExistMkDir(dir)
//	if err != nil {
//		global.Log.Error("RemoteToLocal.IsNotExistMkDir", zap.Error(err))
//		return err
//	}
//
//	srcFile, err := this.sourceClient.Open(fname)
//	if err != nil {
//		global.Log.Error("RemoteToLocal.sourceClient.Open", zap.Error(err))
//		return err
//	}
//	defer srcFile.Close()
//	lf, err := os.Create(dstFile)
//	if err != nil {
//		return err
//	}
//	defer lf.Close()
//
//	if _, err = srcFile.WriteTo(lf); err != nil {
//		return err
//	}
//	global.Log.Info(fmt.Sprintf("copy %s finished!", srcFile.Name()))
//	return nil
//}
//
//func (this *Remote) WalkRemotePath(dirPath string, fp *pool.Pool) {
//	globPath := pathJoin(dirPath)
//	files, err := this.sourceClient.Glob(globPath)
//	if err != nil {
//		logging.Error(err)
//		return
//	}
//	for _, v := range files {
//		stat, err := this.sourceClient.Stat(v)
//		if err != nil {
//			global.Log.Error("WalkRemotePath.this.sourceClient.Stat", zap.Error(err))
//			continue
//		}
//		if stat.IsDir() {
//			this.WalkRemotePath(v, fp)
//		} else {
//			fp.Add(1)
//			go func(v string, size int64) {
//				defer fp.Done()
//				err = this.RemoteToLocal(v, size)
//				if err != nil {
//					global.Log.Error("WalkRemotePath.RemoteToLocal", zap.Error(err))
//				}
//			}(v, stat.Size())
//		}
//	}
//}
//
//func (this *Remote) Mkdir(p string) error {
//	if runtime.GOOS == "windows" {
//		p = strings.ReplaceAll(p, "\\", "/")
//	}
//	dst := path.Join(this.dstPath, p)
//	fmt.Println(dst)
//	err := this.dstClient.MkdirAll(dst)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func pathJoin(p string) (np string) {
//	if strings.HasSuffix(p, "/") == false {
//		p = p + "/"
//	}
//	if OS.IsWindows() {
//		np = strings.ReplaceAll(path.Join(p, "*"), "\\", "/")
//	} else {
//		np = path.Join(p, "*")
//	}
//	return
//}
