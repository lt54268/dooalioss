package router

import (
	"dooalioss/internal/api"

	"github.com/gin-gonic/gin"
)

// 设置 Gin 路由
func SetupRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.POST("/upload", api.UploadHandler)
		v1.GET("/download", api.DownloadFileHandler)
		v1.DELETE("/delete", api.DeleteFileHandler)
		v1.GET("/list", api.ListFilesHandler)
		v1.POST("/copy", api.CopyFileHandler)
		v1.POST("/rename", api.RenameFileHandler)
		v1.GET("/logs", api.LogQueryHandler)
	}

	v2 := r.Group("/api/v2")
	{
		v2.GET("/list", api.ListFilesHandlerV2)
		v2.GET("/download", api.DownloadFileHandlerV2)
	}

	v3 := r.Group("/api/v3")
	{
		v3.GET("/download", api.DownloadFileHandlerV3)
	}
}
