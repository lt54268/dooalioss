package model

import (
	"mime/multipart"
	"time"
)

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
	Upload(file multipart.File, objectName string) (*UploadResponse, error)
}

// FileInfo 包含文件基本信息
type FileInfo struct {
	Key           string    `json:"key"`
	ContentLength int64     `json:"content-length"`
	ETag          string    `json:"etag"`
	LastModified  time.Time `json:"last_modified"`
}

type UploadResponse struct {
	ContentLength int64     `json:"content-length"`
	ETag          string    `json:"etag"`
	LastModified  time.Time `json:"last-modified"`
}
