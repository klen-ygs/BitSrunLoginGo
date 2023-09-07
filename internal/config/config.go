package config

import (
	"github.com/Mmx233/BitSrunLoginGo/tools"
	"github.com/spf13/viper"
)

type User struct {
	Username string `yaml:"Username"`
	Password string `yaml:"Password"`
}

var UserInfo User

func LoadConfig(path string) {
	if path == "" {
		path = "./userinfo.yaml"
	}
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		tools.Notify("读取配置文件失败")
		panic(err)
	}
	err = viper.Unmarshal(&UserInfo)
	if err != nil {
		tools.Notify("读取配置文件失败")
		panic(err)
	}
}
