package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetInt64Param(c *gin.Context, param string) (int64, error) {
	paramStr := c.Param(param)
	paramInt64, err := strconv.ParseInt(paramStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return paramInt64, nil
}
