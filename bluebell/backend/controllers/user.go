package controllers

import (
	"bluebell/dao/mysql"
	"bluebell/logic"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"strings"
)

//SignUpHandler 用户注册
func SignUpHandler(c *gin.Context) {
	// 1.参数校验
	p := new(models.SignUpParam)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("invalid param", zap.Error(err))
		//c.JSON(http.StatusOK, gin.H{
		//	"msg": err.Error(),
		//})
		ResponseError(c, CodeInvalidParma)
		return
	}

	// 2.业务处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		//c.JSON(http.StatusOK, gin.H{
		//	"msg": "注册失败！",
		//})
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
		} else {
			ResponseError(c, CodeServerBusy)
		}
		return
	}
	// 3.返回响应
	ResponseSuccess(c, nil)
}

//LoginHandler 用户登录
func LoginHandler(c *gin.Context) {
	// 1.参数校验
	p := new(models.LoginParam)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParma)
		return
	}

	// 2.业务处理
	aToken, rToken, err := logic.Login(p)
	//aToken, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.Login failed", zap.String("username", p.Username), zap.Error(err))
		//c.JSON(http.StatusOK, gin.H{
		//	"msg": "用户不存在或密码错误！",
		//})
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
		} else {
			ResponseError(c, CodeInvalidPassword)
		}
		return
	}
	// 3.返回响应
	ResponseSuccess(c, gin.H{
		"accessToken":  aToken,
		"refreshToken": rToken,
		"userID":       strconv.Itoa(p.UserID),
		"username":     p.Username,
	})
}

func RefreshTokenHandler(c *gin.Context) {
	rt := c.Query("refresh_token")

	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		ResponseErrorWithMsg(c, CodeInvalidAuth, "请求头缺少Auth Token")
		c.Abort()
		return
	}
	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		ResponseErrorWithMsg(c, CodeInvalidAuth, "Token格式不对")
		c.Abort()
		return
	}
	aToken, rToken, err := jwt.RefreshToken(parts[1], rt)
	fmt.Println(err)
	c.JSON(http.StatusOK, gin.H{
		"access_token":  aToken,
		"refresh_token": rToken,
	})
}
