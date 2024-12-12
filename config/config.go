package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name string
		Port string
	}
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}
}

var AppConfig *Config

func InitConfig() {
	// 设置配置文件名
	viper.SetConfigName("config")
	// 设置配置文件类型
	viper.SetConfigType("yml")
	// 设置配置文件路径
	viper.AddConfigPath("./config")
	
	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// 初始化配置结构体
	AppConfig = &Config{}

	// 将配置文件中的信息解析到结构体中
	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}

	initDB()
}
