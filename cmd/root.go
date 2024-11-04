package cmd

import (
	"doocloud/internal/config"
	"doocloud/router"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

// rootCmd 代表基本命令
var rootCmd = &cobra.Command{
	Use:   "doocloud",
	Short: "A brief description of your application",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 加载配置
		cfg := config.LoadOssConfig()

		// 创建 Gin 引擎
		r := gin.Default()

		// 设置路由
		router.SetupRoutes(r)

		// 启动服务器
		fmt.Println("Starting server on :" + cfg.Port)
		if err := r.Run(":" + cfg.Port); err != nil {
			fmt.Println("Error starting server:", err)
			return err
		}
		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// 添加全局标志
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
