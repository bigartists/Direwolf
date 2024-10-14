package routes

import (
	"github.com/bigartists/Direwolf/internal/module/user"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {

	group := r.Group("/api/v1") // *gin.RouterGroup
	user.NewUserController().Build(group)

}
