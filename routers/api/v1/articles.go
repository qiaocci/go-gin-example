package v1

import (
	"fmt"
	"github.com/beego/beego/v2/core/validation"
	"github.com/gin-gonic/gin"
	"github.com/qiaocco/go-gin-example/models"
	"github.com/qiaocco/go-gin-example/pkg/e"
	"github.com/qiaocco/go-gin-example/pkg/settings"
	"github.com/qiaocco/go-gin-example/pkg/util"
	"github.com/unknwon/com"
	"net/http"
)

func GetArticles(c *gin.Context) {
	title := c.Query("title")
	desc := c.Query("desc")
	var state int = -1

	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	if title != "" {
		maps["title"] = title
	}
	if desc != "" {
		maps["desc"] = desc
	}
	fmt.Printf("maps=%v\n\n", maps)
	data["lists"] = models.GetArticles(util.GetPageOffset(c), settings.PageSize, maps)
	data["total"] = models.GetArticleTotal(maps)

	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Query("id")).MustInt()
	data := make(map[string]interface{})
	data["lists"] = models.GetArticle(id)

	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func AddArticle(c *gin.Context) {
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	createdBy := c.Query("createdBy")
	tagID := com.StrTo(c.Query("tagID")).MustInt()

	valid := validation.Validation{}
	valid.Required(title, "title").Message("标题必填")
	valid.Required(desc, "desc").Message("描述必填")
	valid.Required(content, "content").Message("内容必填")
	valid.Required(createdBy, "createdBy").Message("创建人必填")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		models.AddArticle(title, desc, content, createdBy, tagID)
		code = e.SUCCESS
	}
	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}

func EditArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	createdBy := c.Query("createdBy")
	tagID := com.StrTo(c.Query("tagID")).MustInt()

	maps := make(map[string]interface{})
	if title != "" {
		maps["title"] = title
	}
	if desc != "" {
		maps["desc"] = desc
	}
	if content != "" {
		maps["content"] = content
	}
	if createdBy != "" {
		maps["created_by"] = createdBy
	}
	maps["tag_id"] = tagID

	models.EditArticle(id, maps)
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})

}

func DeleteArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	models.DeleteArticle(id)
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}
