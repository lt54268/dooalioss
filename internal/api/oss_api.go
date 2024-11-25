package api

import (
	"dooalioss/internal/model"
	"dooalioss/internal/service"
	"dooalioss/utils"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var uploader model.Uploader = service.NewOssUploader()

// UploadHandler 文件上传接口
// @Summary 文件上传
// @Description 处理文件上传请求，可选择是否禁止覆盖已有文件
// @Tags 文件管理
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "待上传的文件"
// @Param forbidOverwrite query bool false "是否禁止覆盖已有文件（默认值为 false）"
// @Success 200 {object} model.UploadResponse "上传成功，返回文件信息"
// @Failure 400 {object} map[string]interface{} "文件解析失败"
// @Failure 500 {object} map[string]interface{} "上传失败"
// @Router /api/v1/upload [post]
func UploadHandler(c *gin.Context) {
	// 使用工具函数解析文件
	file, fileName, err := utils.ParseFile(c, "file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "文件解析失败"})
		return
	}
	defer file.Close()

	// 获取 forbidOverwrite 参数，默认值为 false，表示允许覆盖文件
	forbidOverwriteParam := c.DefaultQuery("forbidOverwrite", "false")
	forbidOverwrite := false
	if forbidOverwriteParam == "true" {
		forbidOverwrite = true
	}

	// 使用通用上传接口上传文件
	info, err := uploader.Upload(file, fileName, forbidOverwrite)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "上传失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "上传成功", "data": info})
}

// DownloadFileHandler 文件下载接口
// @Summary 下载文件
// @Description 根据文件名下载文件
// @Tags 文件管理
// @Accept json
// @Produce application/octet-stream
// @Param objectName query string true "文件名"
// @Success 200 {file} file "文件数据流"
// @Failure 400 {object} map[string]interface{} "object_Name 参数缺失"
// @Failure 500 {object} map[string]interface{} "文件下载失败"
// @Router /api/v1/download [get]
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

// DownloadFileHandlerV2 文件下载接口
// @Summary 获取文件下载链接
// @Description 根据文件名生成下载链接
// @Tags 文件管理
// @Accept json
// @Produce json
// @Param objectName query string true "文件名"
// @Success 200 {object} map[string]interface{} "成功返回下载链接"
// @Failure 400 {object} map[string]interface{} "objectName 参数缺失"
// @Failure 500 {object} map[string]interface{} "生成下载链接失败"
// @Router /api/v2/download [get]
func DownloadFileHandlerV2(c *gin.Context) {
	objectName := c.Query("objectName") // 从请求的查询参数中获取文件名

	if objectName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "objectName 参数缺失",
		})
		return
	}

	// 调用服务函数生成下载链接
	downloadURL, err := service.GenerateDownloadURL(objectName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}

	// 返回生成的下载链接
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "下载链接生成成功",
		"data": map[string]interface{}{
			"download_url": downloadURL,
		},
	})
}

// DownloadFileHandlerV3 文件下载接口
// @Summary 下载文件到本地
// @Description 根据文件名下载文件到本地目录
// @Tags 文件管理
// @Accept json
// @Produce json
// @Param objectName query string true "文件名"
// @Success 200 {object} map[string]interface{} "成功返回文件的本地路径"
// @Failure 400 {object} map[string]interface{} "objectName 参数缺失"
// @Failure 500 {object} map[string]interface{} "文件下载失败"
// @Router /api/v3/download [get]
func DownloadFileHandlerV3(c *gin.Context) {
	objectName := c.Query("objectName") // 从请求的查询参数中获取文件名

	if objectName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "objectName 参数缺失",
		})
		return
	}

	// 调用服务函数下载文件到本地
	localFilePath, err := service.DownloadFileToLocal(objectName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}

	// 返回本地文件路径
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "文件下载成功",
		"data": map[string]interface{}{
			"local_path": localFilePath,
		},
	})
}

// DeleteFileHandler 文件删除接口
// @Summary 删除文件
// @Description 根据文件名删除文件
// @Tags 文件管理
// @Accept json
// @Produce json
// @Param objectName query string true "文件名"
// @Success 200 {object} map[string]interface{} "文件删除成功"
// @Failure 400 {object} map[string]interface{} "object_Name 参数缺失"
// @Failure 500 {object} map[string]interface{} "文件删除失败"
// @Router /api/v1/delete [delete]
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

// ListFilesHandler 获取文件列表接口
// @Summary 获取文件列表
// @Description 返回存储中的所有文件列表
// @Tags 文件管理
// @Accept json
// @Produce json
// @Success 200 {object} model.FileInfo "文件列表获取成功"
// @Failure 500 {object} map[string]interface{} "文件列表获取失败"
// @Router /api/v1/list [get]
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

// ListFilesHandlerV2 获取文件列表接口（V2版本）
// @Summary 获取文件列表（V2版本）
// @Description 获取指定目录下的文件列表，并支持分页查询，返回文件列表及下一页的分页标记（Continuation Token）
// @Tags 文件管理
// @Accept json
// @Produce json
// @Param prefix query string false "文件前缀"
// @Param continuationToken query string false "分页标记，继续上次查询的位置"
// @Param maxKeys query int false "每页返回的文件数，最大值为1000，默认为1000"
// @Success 200 {object} map[string]interface{} "文件列表获取成功"
// @Failure 400 {object} map[string]interface{} "Invalid maxKeys parameter"
// @Failure 500 {object} map[string]interface{} "获取文件列表失败"
// @Router /api/v2/list [get]
func ListFilesHandlerV2(c *gin.Context) {
	// 获取请求参数
	prefix := c.DefaultQuery("prefix", "")                       // 默认为 ""
	continuationToken := c.DefaultQuery("continuationToken", "") // 默认为 ""
	maxKeysParam := c.DefaultQuery("maxKeys", "1000")            // 默认为 1000
	maxKeys, err := strconv.Atoi(maxKeysParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"msg":   "Invalid maxKeys parameter",
			"error": err.Error(),
		})
		return
	}

	// 调用服务函数获取文件列表
	files, nextContinuationToken, err := service.ListFilesV2(prefix, continuationToken, maxKeys)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":                  http.StatusOK,
		"msg":                   "文件列表获取成功",
		"data":                  files,
		"NextContinuationToken": nextContinuationToken,
	})
}

// CopyFileHandler 文件拷贝接口
// @Summary 拷贝文件
// @Description 根据源存储桶和对象，将文件拷贝到目标存储桶和对象
// @Tags 文件管理
// @Accept json
// @Produce json
// @Param srcBucket query string true "源存储桶名称"
// @Param srcObject query string true "源对象名称（文件名）"
// @Param destBucket query string false "目标存储桶名称（可选）"
// @Param destObject query string true "目标对象名称（文件名）"
// @Success 200 {object} map[string]interface{} "文件拷贝成功"
// @Failure 400 {object} map[string]interface{} "缺少必需的参数"
// @Failure 500 {object} map[string]interface{} "文件拷贝失败"
// @Router /api/v1/copy [post]
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

// RenameFileHandler 文件重命名接口
// @Summary 重命名文件
// @Description 根据源对象名称将文件重命名为目标对象名称
// @Tags 文件管理
// @Accept json
// @Produce json
// @Param srcObject query string true "源对象名称（当前文件名）"
// @Param destObject query string true "目标对象名称（新的文件名）"
// @Success 200 {object} map[string]interface{} "文件重命名成功"
// @Failure 400 {object} map[string]interface{} "缺少必需的参数"
// @Failure 500 {object} map[string]interface{} "文件重命名失败"
// @Router /api/v1/rename [post]
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

// LogQueryHandler 处理日志查询请求
// @Summary 查询日志
// @Description 根据条件查询日志
// @Tags 日志管理
// @Accept json
// @Produce json
// @Param query query string false "查询条件"
// @Param startTime query int64 false "开始时间，Unix 时间戳"
// @Param endTime query int64 false "结束时间，Unix 时间戳"
// @Param limit query int false "返回日志条数上限"
// @Param offset query int false "分页起始位置"
// @Success 200 {object} map[string]interface{} "日志查询成功"
// @Failure 500 {object} map[string]interface{} "日志查询失败"
// @Router /api/v1/logs [get]
func LogQueryHandler(c *gin.Context) {
	projectName := os.Getenv("OSS_PROJECT_NAME")
	logStoreName := os.Getenv("OSS_LOG_STORE_NAME")
	if projectName == "" || logStoreName == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "项目名称或日志库名称未配置"})
		return
	}

	query := c.DefaultQuery("query", "*")
	startTimeStr := c.Query("startTime")
	endTimeStr := c.Query("endTime")
	limitStr := c.DefaultQuery("limit", "1000")
	offsetStr := c.DefaultQuery("offset", "0")

	startTime, err := strconv.ParseInt(startTimeStr, 10, 64)
	if err != nil {
		startTime = time.Now().Add(-30 * 24 * time.Hour).Unix() // 默认过去一个月
	}
	endTime, err := strconv.ParseInt(endTimeStr, 10, 64)
	if err != nil {
		endTime = time.Now().Unix() // 默认当前时间
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 1000 // 默认返回最多 1000 条日志
	}
	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0 // 默认从第 0 条日志开始
	}

	service := service.NewLogQueryService()
	logs, err := service.QueryLogs(projectName, logStoreName, query, startTime, endTime, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "日志查询失败", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "日志查询成功", "data": logs})
}
