package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	//查找、加载和反序列化JSON、TOML、YAML、HCL、INI、envfile和Java properties格式的配置文件。
	//提供一种机制为你的不同配置选项设置默认值。
	//提供一种机制来通过命令行参数覆盖指定选项的值。
	//提供别名系统，以便在不破坏现有代码的情况下轻松重命名参数。
	//当用户提供了与默认值相同的命令行或配置文件时，可以很容易地分辨出它们之间的区别。

	//Viper会按照下面的优先级。每个项目的优先级都高于它下面的项目:
	//显示调用Set设置值
	//命令行参数（flag）
	//环境变量
	//配置文件
	//key/value存储
	//默认值

	// 设置默认值
	viper.SetDefault("version", "0.0.1")

	// 读取配置文件
	viper.SetConfigFile("./config.yaml")  // 指定配置文件路径
	viper.SetConfigName("config")         // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")           // 如果配置文件的名称中没有扩展名，则需要配置此项
	viper.AddConfigPath("/etc/appname/")  // 查找配置文件所在的路径
	viper.AddConfigPath("$HOME/.appname") // 多次调用以添加多个搜索路径
	viper.AddConfigPath(".")              // 还可以在工作目录中查找配置
	err := viper.ReadInConfig()           // 查找并读取配置文件
	if err != nil {                       // 处理读取配置文件的错误
		fmt.Printf("Fatal error config file: %s \n", err)
		return
	}

	// 实时监听配置文件
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Println("Config file changed:", e.Name)
	})

	r := gin.Default()
	r.GET("/version", func(context *gin.Context) {
		context.String(200,viper.GetString("version"))
	})
	_ = r.Run(":8081")
}
