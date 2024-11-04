package api

import (
	"doocloud/internal/model"
	"doocloud/internal/service"
	"doocloud/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

var uploader model.Uploader = service.NewOssUploader() // 使用阿里云 OSS 上传器

func UploadHandler(c *gin.Context) {
	// 使用工具函数解析文件
	file, fileName, err := utils.ParseFile(c, "file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "文件解析失败"})
		return
	}
	defer file.Close()

	// 使用通用上传接口上传文件
	url, err := uploader.Upload(file, fileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "上传失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "上传成功", "url": url})
}
