package main

import (
	"fmt"
	"github.com/qiaocco/go-gin-example/pkg/settings"
	"github.com/qiaocco/go-gin-example/routers"
	"log"
	"net/http"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", settings.HTTPPort),
		Handler:        router,
		ReadTimeout:    settings.ReadTimeout,  // 允许读取的最大时间
		WriteTimeout:   settings.WriteTimeout, // 允许写入的最大时间
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Printf("server failed, err:%v\n", err)
	}
}
