package router

import (
	"doocloud/internal/api"

	"github.com/gin-gonic/gin"
)

// SetupRoutes 设置 Gin 路由
func SetupRoutes(r *gin.Engine) {
	r.POST("/upload", api.UploadHandler)
}
