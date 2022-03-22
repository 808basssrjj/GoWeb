package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

const (
	CtxUserID = "userID"
)

var ErrorNotLogin = errors.New("用户未登录")

func getCurrentUser(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxUserID)
	if !ok {
		err = ErrorNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorNotLogin
		return
	}
	return
}
