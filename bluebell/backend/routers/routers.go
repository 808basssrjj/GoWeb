package routers

import (
	"bluebell/controllers"
	"bluebell/logger"
	"bluebell/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUp() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := r.Group("/api/v1")
	// 用户注册
	v1.POST("/signup", controllers.SignUpHandler)
	// 用户登录
	v1.POST("/login", controllers.LoginHandler)
	// 刷新token
	v1.GET("/refresh_token", controllers.RefreshTokenHandler)

	v1.Use(middlewares.JWTAuthMiddleware())

	{
		v1.GET("/community", controllers.CommunityHandler)
		v1.GET("/community/:id", controllers.CommunityDetailHandler)

		v1.POST("/post", controllers.CreatePostHandler)
		v1.GET("/post/:id", controllers.PostDetailHandler)
		v1.GET("/posts", controllers.PostListHandler)

		v1.POST("/vote", controllers.VoteHandler)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
