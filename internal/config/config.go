package config

import (
	"dooalioss/internal/model"
	"os"
)

// 从环境变量加载配置信息
func LoadOssConfig() *model.Config {
	return &model.Config{
		Port:               os.Getenv("PORT"),                  // 从环境变量读取端口
		OssRegion:          os.Getenv("OSS_REGION"),            // 从环境变量读取区域
		OssEndpoint:        os.Getenv("OSS_ENDPOINT"),          // 从环境变量读取 Endpoint
		OssBucket:          os.Getenv("OSS_BUCKET"),            // 从环境变量读取 Bucket
		OssAccessKeyId:     os.Getenv("OSS_ACCESS_KEY_ID"),     // 从环境变量读取 AccessKeyId
		OssAccessKeySecret: os.Getenv("OSS_ACCESS_KEY_SECRET"), // 从环境变量读取 AccessKeySecret
	}
}
