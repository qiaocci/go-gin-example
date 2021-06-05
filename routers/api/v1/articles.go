package v1

import (
	"github.com/beego/beego/v2/core/validation"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/qiaocco/go-gin-example/models"
	"github.com/qiaocco/go-gin-example/pkg/e"
	"github.com/qiaocco/go-gin-example/pkg/settings"
	"github.com/qiaocco/go-gin-example/pkg/util"
	"github.com/unknwon/com"
	"log"
	"net/http"
	"time"
)

// GetArticle godoc
// @Summary 查看文章
// @Description get article by ID
// @Param   id    query    int     true        "ID"
// @Success 200 {string} json  json "{"code":200,"data":{},"msg":"ok"}"
// @Router /articles/{id} [get]
func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)

	sentry.CaptureMessage("It works!")
	code := e.INVALID_PARAMS
	data := make(map[string]interface{})
	if !valid.HasErrors() {
		if models.ExistArticleByID(id) {
			data["lists"] = models.GetArticle(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
func GetArticles(c *gin.Context) {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	valid := validation.Validation{}

	if arg := c.Query("state"); arg != "" {
		state := com.StrTo(arg).MustInt()
		maps["state"] = state
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	if arg := c.Query("tagID"); arg != "" {
		tagID := com.StrTo(arg).MustInt()
		maps["tag_id"] = tagID
		valid.Min(tagID, 1, "tag_id").Message("标签ID必须大于0")
	}

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		data["lists"] = models.GetArticles(util.GetPageOffset(c), settings.PageSize, maps)
		data["total"] = models.GetArticleTotal(maps)
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

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
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	tagID := com.StrTo(c.Query("tagID")).MustInt()

	valid := validation.Validation{}
	valid.Required(title, "title").Message("标题必填")
	valid.Required(desc, "desc").Message("描述必填")
	valid.Required(content, "content").Message("内容必填")
	valid.Required(createdBy, "createdBy").Message("创建人必填")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	valid.Min(tagID, 1, "tag_id").Message("标签ID必须大于0")

	code := e.INVALID_PARAMS
	data := make(map[string]interface{})
	if !valid.HasErrors() {
		if models.ExistTagByID(tagID) {
			data["title"] = title
			data["desc"] = desc
			data["content"] = content
			data["CreatedBy"] = createdBy
			data["TagID"] = tagID
			data["State"] = state
			models.AddArticle(data)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
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
	modifiedBy := c.Query("modifiedBy")
	tagID := com.StrTo(c.Query("tagID")).MustInt()

	valid := validation.Validation{}
	if arg := c.Query("state"); arg != "" {
		state := com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
	valid.MaxSize(desc, 255, "desc").Message("简述最长为255字符")
	valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistArticleByID(id) {
			if models.ExistTagByID(tagID) {
				data := make(map[string]interface{})
				if tagID > 0 {
					data["tag_id"] = tagID
				}
				if title != "" {
					data["title"] = title
				}
				if desc != "" {
					data["desc"] = desc
				}
				if content != "" {
					data["content"] = content
				}

				data["modified_by"] = modifiedBy

				models.EditArticle(id, data)
				code = e.SUCCESS
			} else {
				code = e.ERROR_NOT_EXIST_TAG
			}
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})

}

func DeleteArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistArticleByID(id) {
			models.DeleteArticle(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}
