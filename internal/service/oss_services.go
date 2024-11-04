package service

import (
	"context"
	"dooalioss/internal/config"
	"dooalioss/internal/model"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"

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

// DownloadFile 从阿里云OSS下载文件
func DownloadFile(objectName string) ([]byte, error) {
	bucketName := os.Getenv("OSS_BUCKET")
	region := os.Getenv("OSS_REGION")

	if bucketName == "" || region == "" || objectName == "" {
		return nil, errors.New("invalid parameters: bucket name, region, and object name are required")
	}

	cfg := oss.LoadDefaultConfig().
		WithCredentialsProvider(credentials.NewEnvironmentVariableCredentialsProvider()).
		WithRegion(region)

	client := oss.NewClient(cfg)

	request := &oss.GetObjectRequest{
		Bucket: oss.Ptr(bucketName),
		Key:    oss.Ptr(objectName),
	}

	// 发起下载请求
	result, err := client.GetObject(context.TODO(), request)
	if err != nil {
		return nil, fmt.Errorf("failed to get object: %v", err)
	}
	defer result.Body.Close()

	// 读取文件内容
	data, err := io.ReadAll(result.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read object data: %v", err)
	}

	return data, nil
}

// DeleteFile 从阿里云OSS删除文件
func DeleteFile(objectName string) error {
	bucketName := os.Getenv("OSS_BUCKET")
	region := os.Getenv("OSS_REGION")

	if bucketName == "" || region == "" || objectName == "" {
		return errors.New("invalid parameters: bucket name, region, and object name are required")
	}

	cfg := oss.LoadDefaultConfig().
		WithCredentialsProvider(credentials.NewEnvironmentVariableCredentialsProvider()).
		WithRegion(region)

	client := oss.NewClient(cfg)

	request := &oss.DeleteObjectRequest{
		Bucket: oss.Ptr(bucketName),
		Key:    oss.Ptr(objectName),
	}

	_, err := client.DeleteObject(context.TODO(), request)
	if err != nil {
		return fmt.Errorf("failed to delete object: %v", err)
	}

	return nil
}

// ListFiles 从阿里云OSS获取文件列表
func ListFiles() ([]model.FileInfo, error) {
	bucketName := os.Getenv("OSS_BUCKET")
	region := os.Getenv("OSS_REGION")

	if bucketName == "" || region == "" {
		return nil, errors.New("invalid parameters: bucket name and region are required")
	}

	cfg := oss.LoadDefaultConfig().
		WithCredentialsProvider(credentials.NewEnvironmentVariableCredentialsProvider()).
		WithRegion(region)

	client := oss.NewClient(cfg)

	request := &oss.ListObjectsV2Request{
		Bucket: oss.Ptr(bucketName),
	}

	p := client.NewListObjectsV2Paginator(request)
	var fileInfos []model.FileInfo

	for p.HasNext() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return nil, fmt.Errorf("failed to get objects list: %v", err)
		}

		// 收集每个对象的信息
		for _, obj := range page.Contents {
			fileInfos = append(fileInfos, model.FileInfo{
				Key:          oss.ToString(obj.Key),
				Size:         obj.Size,
				LastModified: oss.ToTime(obj.LastModified).Format("2006-01-02 15:04:05"),
			})
		}
	}

	return fileInfos, nil
}
