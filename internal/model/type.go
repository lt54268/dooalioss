package model

import "mime/multipart"

// Config 用于存储配置信息
type Config struct {
	Port               string
	OssRegion          string
	OssEndpoint        string
	OssBucket          string
	OssAccessKeyId     string
	OssAccessKeySecret string
}

// Uploader 定义上传接口
type Uploader interface {
	Upload(file multipart.File, objectName string) (string, error)
}
