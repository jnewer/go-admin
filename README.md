<div align="center">
<br/>
<br/>
  <h1 align="center">
    Pear Admin Go
  </h1>
</div>

#### 项目简介
>Pear Admin Go 基于 Gin框架  的后台管理系统
> 
>众人拾柴火焰高，欢迎参与项目~

>	go1.16	+	gin	+	mysql	+	权限验证	


####  项目结构

```
Pear Admin Golang
├─app  # 应用
├─database  # 数据库预设文件
├─static  # 前端css、js、img文件
├─template # 前端html文件
├─go.mod # go mod文件
├─config.toml.examole # 配置文件样例
└─main.go # 项目主入口执行文件

```



#### 项目安装

```bash
# 下 载
git clone https://gitee.com/pear-admin/pear-admin-golang

# 修 改 配 置
cp config.toml.example config.toml


```

#### 运行项目

```bash
go mod tidy
go run main.go
```

#### 未完成工作
- [ ] 修改路由结构
- [ ] 修改配置方式
- [ ] 去除多余函数
- [ ] 优化文件层级
- [ ] 增加其他功能