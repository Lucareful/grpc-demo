package config

import (
	"github.com/spf13/viper"
)

var (
	cfg Config
)

type Config struct {
	Token Token
}

type Token struct {
	AppID     string
	AppSecret string
}

// InitConf 初始化加载配置
func InitConf() {
	viper.SetConfigFile("./token.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
}

// GetConf 获取配置信息
func GetConf() *Config {
	return &cfg
}
