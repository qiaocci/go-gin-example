package jwt

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qiaocco/go-gin-example/pkg/e"
	"github.com/qiaocco/go-gin-example/pkg/util"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int

		code = e.SUCCESS
		token := c.Query("token")
		if token == "" {
			fmt.Println("未传递jwt token")
			code = e.INVALID_PARAMS
		} else {
			claim, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claim.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": make(map[string]interface{}),
			})
			c.Abort()
			return
		}

		// Process request
		c.Next()
	}
}
