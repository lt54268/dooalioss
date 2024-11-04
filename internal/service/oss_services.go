package service

import (
	"context"
	"doocloud/internal/config"
	"fmt"
	"mime/multipart"

	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
)

// OssUploader 实现了 Uploader 接口，支持阿里云 OSS 上传
type OssUploader struct{}

// NewOssUploader 返回 OssUploader 实例
func NewOssUploader() *OssUploader {
	return &OssUploader{}
}

// Upload 实现 Uploader 接口中的 Upload 方法
func (u *OssUploader) Upload(file multipart.File, objectName string) (string, error) {
	// 初始化 OSS 客户端
	cfg := oss.LoadDefaultConfig().
		WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			config.LoadOssConfig().OssAccessKeyId,
			config.LoadOssConfig().OssAccessKeySecret,
			"")).
		WithRegion(config.LoadOssConfig().OssRegion)

	client := oss.NewClient(cfg)

	// 创建上传请求
	request := &oss.PutObjectRequest{
		Bucket: oss.Ptr(config.LoadOssConfig().OssBucket),
		Key:    oss.Ptr(objectName),
		Body:   file,
	}

	// 上传文件
	_, err := client.PutObject(context.TODO(), request)
	if err != nil {
		return "", fmt.Errorf("failed to upload object: %v", err)
	}

	// 构建远程 URL
	remoteURL := fmt.Sprintf("https://%s.%s/%s", config.LoadOssConfig().OssBucket, config.LoadOssConfig().OssEndpoint, objectName)
	return remoteURL, nil
}
