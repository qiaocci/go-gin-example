package v1

import (
	"fmt"
	"github.com/beego/beego/v2/core/validation"
	"github.com/gin-gonic/gin"
	"github.com/qiaocco/go-gin-example/models"
	"github.com/qiaocco/go-gin-example/pkg/e"
	"github.com/qiaocco/go-gin-example/pkg/logging"
	"github.com/qiaocco/go-gin-example/pkg/util"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	code := e.INVALID_PARAMS
	data := make(map[string]interface{})
	if ok {
		isExist := models.CheckAuth(username, password)
		fmt.Printf("isExist: %v\n", isExist)
		if isExist {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		}
	} else {
		for _, err := range valid.Errors {
			fmt.Println("出错啦！！！")
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
