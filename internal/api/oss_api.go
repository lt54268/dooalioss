package api

import (
	"dooalioss/internal/model"
	"dooalioss/internal/service"
	"dooalioss/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var uploader model.Uploader = service.NewOssUploader() // 使用阿里云 OSS 上传器

// UploadFileHandler godoc
// @Summary 上传文件到阿里云 OSS
// @Description 接收文件并上传到阿里云 OSS
// @Tags 文件操作
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "文件"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /upload [post]
func UploadHandler(c *gin.Context) {
	// 使用工具函数解析文件
	file, fileName, err := utils.ParseFile(c, "file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "文件解析失败"})
		return
	}
	defer file.Close()

	// 使用通用上传接口上传文件
	info, err := uploader.Upload(file, fileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "上传失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "上传成功", "data": info})
}

// DownloadFileHandler godoc
// @Summary 下载文件
// @Description 从阿里云 OSS 下载指定文件
// @Tags 文件操作
// @Produce json
// @Param file_name query string true "文件名"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /download [get]
func DownloadFileHandler(c *gin.Context) {
	objectName := c.Query("objectName") // 从请求的查询参数中获取文件名

	if objectName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "object_Name 参数缺失",
		})
		return
	}

	// 调用服务函数下载文件
	data, err := service.DownloadFile(objectName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}

	// 设置文件下载响应头
	c.Header("Content-Disposition", "attachment; filename="+objectName)
	c.Data(http.StatusOK, "application/octet-stream", data)
}

// DeleteFileHandler godoc
// @Summary 删除文件
// @Description 从阿里云 OSS 删除指定文件
// @Tags 文件操作
// @Produce json
// @Param file_name query string true "文件名"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /delete [delete]
// DeleteFileHandler 处理文件删除请求
func DeleteFileHandler(c *gin.Context) {
	objectName := c.Query("objectName") // 从请求的查询参数中获取文件名

	if objectName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "object_Name 参数缺失",
		})
		return
	}

	// 调用服务函数删除文件
	err := service.DeleteFile(objectName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "文件删除成功",
	})
}

// ListFilesHandler godoc
// @Summary 获取文件列表
// @Description 从阿里云 OSS 获取文件列表
// @Tags 文件操作
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /list [get]
func ListFilesHandler(c *gin.Context) {
	// 调用服务函数获取文件列表
	files, err := service.ListFiles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "文件列表获取成功",
		"data": files,
	})
}

func CopyFileHandler(c *gin.Context) {
	srcBucket := c.Query("srcBucket")
	srcObject := c.Query("srcObject")
	destBucket := c.Query("destBucket")
	destObject := c.Query("destObject")

	// 验证必要的参数
	if srcBucket == "" || srcObject == "" || destObject == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "缺少必需的参数:srcBucket, srcObject, 或 destObject",
		})
		return
	}

	// 调用 CopyFile 函数执行文件拷贝
	err := service.CopyFile(srcBucket, srcObject, destBucket, destObject)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  fmt.Sprintf("文件拷贝失败: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "文件拷贝成功",
	})
}

func RenameFileHandler(c *gin.Context) {
	srcObject := c.Query("srcObject")
	destObject := c.Query("destObject")

	// 检查参数是否为空
	if srcObject == "" || destObject == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "缺少必需的参数:srcObject 和 destObject",
		})
		return
	}

	// 调用 RenameFile 函数执行重命名
	err := service.RenameFile(srcObject, destObject)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  fmt.Sprintf("文件重命名失败: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  fmt.Sprintf("文件'%s'重命名为'%s'成功", srcObject, destObject),
	})
}
