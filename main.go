package main

import (
	"fmt"
	"github.com/qiaocco/go-gin-example/models"
	"github.com/qiaocco/go-gin-example/pkg/logging"
	"github.com/qiaocco/go-gin-example/pkg/setting"
	"github.com/qiaocco/go-gin-example/routers"
	"log"
	"net/http"
)

func init()  {
	setting.Setup()
	models.SetUp()
	logging.SetUp()
}

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,  // 允许读取的最大时间
		WriteTimeout:   setting.ServerSetting.WriteTimeout, // 允许写入的最大时间
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("port=", setting.ServerSetting.HttpPort)
	err := s.ListenAndServe()
	if err != nil {
		log.Printf("server failed, err:%v\n", err)
	}
}
