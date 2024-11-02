package routers

import (
	"github.com/bigartists/Direwolf/internal/module/model"
	"github.com/bigartists/Direwolf/internal/module/user"
	"github.com/bigartists/Direwolf/pkg/middlewares"
	"github.com/bigartists/Direwolf/pkg/validators"
	"github.com/bigartists/Direwolf/server"
	"github.com/gin-gonic/gin"
)

func ProvideRouter(userController *user.UserController,
	modelController *model.ModelController,
	authMiddleware *middlewares.AuthMiddleware) *gin.Engine {
	r := gin.Default()

	r.Use(authMiddleware.JwtAuthMiddleware())
	r.Use(middlewares.ErrorHandler())

	api := r.Group("/api/v1")
	userController.Build(api)
	modelController.Build(api)

	server.FrontendServer(r)

	// 加载 validator
	validators.Build()

	return r
}
