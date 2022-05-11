package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"wechatNotify/pkg/setting"
	"wechatNotify/pkg/util"
	"wechatNotify/routers"
)

func main() {
	router := routers.InitRouter()
	util.InitRedis()
	// 启动一个http服务
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()
	pid := os.Getpid()
	if setting.PidFile != "" {
		pidFile, err := os.OpenFile(setting.PidFile, os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatalf("Fail to create pid file: %s", err)
		}
		defer pidFile.Close()
		if _, err := pidFile.WriteString(fmt.Sprintf("%d", pid)); err != nil {
			log.Fatalf("Fail to write pid to file: %s", err)
		}
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// 删除创建的 pid 文件
	if setting.PidFile != "" {
		os.Remove(setting.PidFile)
	}

	log.Println("Server exiting")
}
