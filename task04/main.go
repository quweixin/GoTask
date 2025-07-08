package main

import (
	"GoTask/task04/initialize"
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	initialize.InitDB()
	Router := initialize.Routers()

	srv := &http.Server{
		Addr:    ":9999",
		Handler: Router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	//创建一个通道 quit，用来接收系统信号
	quit := make(chan os.Signal)
	//使用 signal.Notify 注册要监听的信号：
	// SIGINT: Ctrl+C
	// SIGTERM: kill
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	//是一个阻塞操作，直到接收到上述信号才会继续执行
	<-quit
	log.Println("Shutdown Server ...")
	//创建一个带超时的上下文，最长等待 5 秒钟用于完成当前请求处理
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel() 确保在函数退出前调用 cancel，释放资源
	defer cancel()
	//调用 srv.Shutdown(ctx) 开始优雅关闭服务器：
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
