package v1

import (
	"github.com/beego/beego/v2/core/validation"
	"github.com/gin-gonic/gin"
	"github.com/qiaocco/go-gin-example/models"
	"github.com/qiaocco/go-gin-example/pkg/e"
	"github.com/qiaocco/go-gin-example/pkg/settings"
	"github.com/qiaocco/go-gin-example/pkg/util"
	"github.com/unknwon/com"
	"log"
	"net/http"
)

func GetTags(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	if arg := c.Query("state"); arg != "" {
		state := com.StrTo(arg).MustInt()
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
	name := c.Query("name")
	state := com.StrTo(c.Query("state")).MustInt()
	createdBy := c.Query("createdBy")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最多100个字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最多100个字符")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistTagByName(name) {
			code = e.SUCCESS
			models.AddTag(name, state, createdBy)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func EditTags(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	modifiedBy := c.Query("modifiedBy")

	valid := validation.Validation{}

	var state = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(c.Query("state")).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}
	valid.Required(id, "id").Message("ID不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最多100个字符")
	valid.Required(modifiedBy, "created_by").Message("修改人不能为空")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS

		if models.ExistTagByID(id) {
			data := make(map[string]interface{})
			if name != "" {
				data["Name"] = name
			}
			if state != -1 {
				data["State"] = state
			}
			data["ModifiedBy"] = modifiedBy
			models.EditTag(id, data)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func DeleteTags(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.MinSize(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagByID(id) {
			models.DeleteTag(id)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
