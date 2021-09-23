package config

import (
	"github.com/spf13/viper"
	"log"
	"sync"
)

var (
	once sync.Once
	Conf = &conf{}
)

type conf struct {
	DB     DBConf
	Redis  RedisConf
	App    SettingConf
	Zaplog ZapLogConf
}

type DBConf struct {
	DBType string
	DBUser string
	DBPwd  string
	DBHost string
	DBName string
}

type RedisConf struct {
	RedisAddr string
	RedisPWD  string
	RedisDB   int
}

type SettingConf struct {
	HttpPort    int `json:"http-port"`
	RunMode     string
	PageSize    int
	JwtSecret   string
	ImgSavePath string
	ImgUrlPath  string
}

type ZapLogConf struct {
	Director string ` json:"director"  yaml:"director"`
}

func InitConfig(tomlPath string) {
	once.Do(func() {
		v := viper.New()
		v.SetConfigFile(tomlPath)
		err := v.ReadInConfig()
		if err != nil {
			log.Fatal("配置文件读取失败: ", err.Error())
		}
		err = v.Unmarshal(&Conf)
		if err != nil {
			log.Fatal("配置解析失败:", err.Error())
		}
	})
}
