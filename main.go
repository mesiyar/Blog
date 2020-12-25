package main

import (
	"fmt"
	"net/http"
	"wechatNotify/pkg/logging"
	"wechatNotify/pkg/setting"
	"wechatNotify/pkg/util"
	"wechatNotify/routers"
)

func main() {
	router := routers.InitRouter()
	r := util.Redis{}
	if err := r.Init();err != nil  {
		logging.Warn(err.Error())
	}
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
