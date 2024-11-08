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
func (u *OssUploader) Upload(file multipart.File, objectName string) (*model.UploadResponse, error) {
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
		return nil, fmt.Errorf("failed to upload object: %v", err)
	}

	// 上传成功后，获取文件信息
	objectInfo, err := client.HeadObject(context.TODO(), &oss.HeadObjectRequest{
		Bucket: oss.Ptr(config.LoadOssConfig().OssBucket),
		Key:    oss.Ptr(objectName),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve object info: %v", err)
	}

	// 返回文件信息
	return &model.UploadResponse{
		ContentLength: objectInfo.ContentLength,
		ETag:          *objectInfo.ETag,
		LastModified:  *objectInfo.LastModified,
	}, nil
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
				Key:           oss.ToString(obj.Key),
				ContentLength: obj.Size,
				ETag:          oss.ToString(obj.ETag),
				LastModified:  oss.ToTime(obj.LastModified),
			})
		}
	}

	return fileInfos, nil
}

// CopyFile 拷贝文件到目标存储空间
func CopyFile(srcBucket, srcObject, destBucket, destObject string) error {
	region := os.Getenv("OSS_REGION")

	if srcBucket == "" || srcObject == "" || destObject == "" || region == "" {
		return errors.New("invalid parameters: source bucket, source object, destination object, and region are required")
	}

	// 如果目标存储空间未指定，默认为源存储空间
	if destBucket == "" {
		destBucket = srcBucket
	}

	cfg := oss.LoadDefaultConfig().
		WithCredentialsProvider(credentials.NewEnvironmentVariableCredentialsProvider()).
		WithRegion(region)

	client := oss.NewClient(cfg)

	request := &oss.CopyObjectRequest{
		Bucket:       oss.Ptr(destBucket),
		Key:          oss.Ptr(destObject),
		SourceBucket: oss.Ptr(srcBucket),
		SourceKey:    oss.Ptr(srcObject),
	}

	_, err := client.CopyObject(context.TODO(), request)
	if err != nil {
		return fmt.Errorf("failed to copy object: %v", err)
	}

	return nil
}

// RenameFile 将源对象重命名为目标对象
func RenameFile(srcObject, destObject string) error {
	bucketName := os.Getenv("OSS_BUCKET")
	region := os.Getenv("OSS_REGION")

	if bucketName == "" || region == "" || srcObject == "" || destObject == "" {
		return errors.New("invalid parameters: bucket name, region, source object, and destination object are required")
	}

	cfg := oss.LoadDefaultConfig().
		WithCredentialsProvider(credentials.NewEnvironmentVariableCredentialsProvider()).
		WithRegion(region)

	client := oss.NewClient(cfg)

	// 创建 CopyObject 请求，将源对象复制到目标位置
	copyRequest := &oss.CopyObjectRequest{
		Bucket:       oss.Ptr(bucketName),
		Key:          oss.Ptr(destObject),
		SourceKey:    oss.Ptr(srcObject),
		SourceBucket: oss.Ptr(bucketName),
	}

	// 执行 CopyObject 操作
	_, err := client.CopyObject(context.TODO(), copyRequest)
	if err != nil {
		return fmt.Errorf("failed to copy object '%s' to '%s': %v", srcObject, destObject, err)
	}

	// 创建 DeleteObject 请求，删除源对象
	deleteRequest := &oss.DeleteObjectRequest{
		Bucket: oss.Ptr(bucketName),
		Key:    oss.Ptr(srcObject),
	}

	// 执行 DeleteObject 操作
	_, err = client.DeleteObject(context.TODO(), deleteRequest)
	if err != nil {
		return fmt.Errorf("failed to delete source object '%s': %v", srcObject, err)
	}

	return nil
}
