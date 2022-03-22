package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func f1(c *gin.Context) {
	fmt.Println("f1")
}
func f2(c *gin.Context) {
	fmt.Println("f2 start")
	//c.Next()
	c.Set("name", "zhaoning")
	fmt.Println("f2 end")
}
func f3(c *gin.Context) {
	name, _ := c.Get("name")
	fmt.Println(name.(string))
	//c.Abort()
	fmt.Println("f3")
}
func f4(c *gin.Context) {
	fmt.Println("f4")
}

func main() {
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.JSON(200, "hello")
	})

	// 添加中间件 会把handler函数加到HandlersChain（切片)中
	// 在中间件函数中通过调用c.Next()实现嵌套调用（func1中调用func2；func2中调用func3），
	// 或者通过调用c.Abort()中断整个调用链条，从当前函数返回。
	userGroup := r.Group("/user", f1, f2)
	userGroup.Use(f3)
	{
		userGroup.GET("/index", f4)
	}
	_ = r.Run(":7071")
}
