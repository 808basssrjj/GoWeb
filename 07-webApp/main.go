package main

import (
	"07-webApp/dao/mysql"
	"07-webApp/dao/redis"
	"07-webApp/logger"
	"07-webApp/routers"
	"07-webApp/setting"
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 1.加载配置
	if err := setting.Init(); err != nil {
		fmt.Printf("init setting falied err:%v\n", err)
		return
	}

	// 2.初始化日志
	if err := logger.Init(setting.Conf.LogConfig); err != nil {
		fmt.Printf("init logger falied err:%v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success")

	// 3.初始化MySql
	if err := mysql.Init(setting.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql falied err:%v\n", err)
		return
	}
	defer mysql.Close()

	// 4.初始化Redis
	if err := redis.Init(setting.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis falied err:%v\n", err)
		return
	}
	defer redis.Close()

	// 5.注册路由
	r := routers.SetUp()

	// 6.启动服务(优雅关机)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler: r,
	}
	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	zap.L().Info("shutdown server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server shutdown:", zap.Error(err))
	}
	zap.L().Info("Server exiting")
}
