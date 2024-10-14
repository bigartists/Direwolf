package middlewares

import (
	. "github.com/bigartists/Direwolf/pkg/utils"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var errMsg string
				switch errTyped := err.(type) {
				case string:
					errMsg = errTyped
				case error:
					errMsg = errTyped.Error()
				default:
					errMsg = "internal Server error occurred errorHandler"
				}
				//context.JSON(500, gin.H{"error": err})
				ret := ResultWrapper(context)(nil, errMsg)(Error)
				context.JSON(500, ret)
			}
		}()
		context.Next()
	}
}
