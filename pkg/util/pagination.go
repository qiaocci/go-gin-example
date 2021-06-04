package util

import (
	"github.com/gin-gonic/gin"
	"github.com/qiaocco/go-gin-example/pkg/settings"
	"github.com/unknwon/com"
)

// GetPageOffset 获取分页偏移量, 也就是sql分页时offset的值
func GetPageOffset(c *gin.Context) int {
	result := 0
	page := com.StrTo(c.Query("page")).MustInt()
	if page > 0 {
		result = (page - 1) * settings.PageSize
	}
	return result
}
