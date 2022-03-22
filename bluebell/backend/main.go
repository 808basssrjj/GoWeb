package main

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/logger"
	"bluebell/pkg/snowflake"
	"bluebell/routers"
	"bluebell/setting"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func main() {
	// 1.加载配置
	// 命令行参数 指定配置文件
	if len(os.Args) < 2 {
		fmt.Println("need config file.eg:bluebell config.yaml")
		return
	}
	if err := setting.Init(os.Args[1]); err != nil {
		fmt.Printf("init setting falied err:%v\n", err)
		return
	}

	// 2.初始化日志
	if err := logger.Init(setting.Conf.LogConfig, "dev"); err != nil {
		fmt.Printf("init logger falied err:%v\n", err)
		return
	}
	defer func(l *zap.Logger) {
		_ = l.Sync()
	}(zap.L())
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

	// 初始化ID生成器
	if err := snowflake.Init(setting.Conf.StartTime, setting.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake falied err:%v\n", err)
		return
	}

	// 5.注册路由
	r := routers.SetUp()

	// 6.启动服务(优雅关机)
	srv := &http.Server{
		//Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Addr:    fmt.Sprintf(":%d", setting.Conf.Port),
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
