package config

import (
	"doocloud/internal/model"
	"os"
)

// LoadConfig 从环境变量加载配置信息
func LoadConfig() *model.Config {
	return &model.Config{
		Port:     os.Getenv("PORT"),         // 从环境变量读取端口
		Region:   os.Getenv("OSS_REGION"),   // 从环境变量读取区域
		Endpoint: os.Getenv("OSS_ENDPOINT"), // 从环境变量读取 Endpoint
		Bucket:   os.Getenv("OSS_BUCKET"),   // 从环境变量读取 Bucket
	}
}
