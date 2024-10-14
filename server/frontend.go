package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func FrontendServer(r *gin.Engine) {

	currentDir, _ := os.Getwd()
	distDir := filepath.Join(currentDir, "web", "dist")
	// 静态文件服务
	r.Static("/static", filepath.Join(distDir, "static"))

	r.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path

		// 检查请求的路径是否是一个文件
		if filepath.Ext(path) != "" {
			// 如果是文件请求，尝试从 dist 目录提供文件
			filePath := filepath.Join(distDir, path)
			if _, err := os.Stat(filePath); err == nil {
				c.File(filePath)
				return
			}
		}

		// 对于 API 请求，返回 404
		if strings.HasPrefix(path, "/api") {
			fmt.Println("API not found:", path)
			c.JSON(http.StatusNotFound, gin.H{"message": "404 API Not Found"})
			return
		}

		// 对于所有其他请求，返回 index.html
		fmt.Println("Serving index.html for:", path)
		c.File(filepath.Join(distDir, "index.html"))
	})

}
