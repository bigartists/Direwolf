package server

import (
	"fmt"
	"github.com/bigartists/Direwolf/client"
	"github.com/bigartists/Direwolf/internal/routers"
	middlewares "github.com/bigartists/Direwolf/pkg/middleware"
	"github.com/bigartists/Direwolf/pkg/validators"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

func Run(port int) error {
	// 执行命令行
	client.InitDB()

	r := gin.Default()

	r.Use(middlewares.JwtAuthMiddleware())
	r.Use(middlewares.ErrorHandler())

	// 加载路由
	routes.SetupRouter(r)

	FrontendServer(r)
	// 加载 validator
	validators.Build()

	err := r.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	return nil
}

// NewApiServerCommand 初始化命令行参数
func NewApiServerCommand() (cmd *cobra.Command) {
	// 集成 cobra命令
	cmd = &cobra.Command{
		Use: "appServer",
		RunE: func(cmd *cobra.Command, args []string) error {
			port, err := cmd.Flags().GetInt("port")
			if err != nil {
				return err
			}
			return Run(port)
		},
	}
	// 添加 flag, name=port, 默认值是 9090
	cmd.Flags().Int("port", 9090, "appserver port")
	return
}
