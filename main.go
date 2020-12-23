package main

import (
	"fmt"
	"net/http"
	"wechatNotify/pkg/setting"
	"wechatNotify/routers"
)

func main() {
	router := routers.InitRouter()
	// 启动一个http服务
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
