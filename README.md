<div align="center">
<br/>
<br/>
  <h1 align="center">
    Go Admin
  </h1>
</div>

#### 项目结构

```
go-admin
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
git clone https://github.com/jnewer/go-admin.git

# 修 改 配 置
cp config.toml.example config.toml


```

#### 运行项目

```bash
go mod tidy
go run main.go
```
