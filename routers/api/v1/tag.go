package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/qiaocci/go-gin-example/models"
	"github.com/qiaocci/go-gin-example/pkg/e"
	"github.com/qiaocci/go-gin-example/pkg/settings"
	"github.com/qiaocci/go-gin-example/pkg/util"
	"github.com/unknwon/com"
	"net/http"
)

func GetTags(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	state := -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}
	data["lists"] = models.GetTags(util.GetPageOffset(c), settings.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}
func AddTags(c *gin.Context) {

}
func EditTags(c *gin.Context) {

}
func DeleteTags(c *gin.Context) {

}
