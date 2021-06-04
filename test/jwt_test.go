package test

import (
	"fmt"
	"github.com/qiaocco/go-gin-example/pkg/util"
	"testing"
)

func TestJwtGenerateToken(t *testing.T) {
	token, err := util.GenerateToken("qiaocc", "123")
	if err != nil {
		return
	}
	fmt.Println(token)
}
