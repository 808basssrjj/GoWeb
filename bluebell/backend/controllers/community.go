package controllers

import (
	"bluebell/logic"
	"strconv"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// CommunityHandler 社区列表
func CommunityHandler(c *gin.Context) {
	list, err := logic.GetCommunity()
	if err != nil {
		zap.L().Error("logic.GetCommunity failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, list)
}

// CommunityDetailHandler 社区详情
func CommunityDetailHandler(c *gin.Context) {
	// 获取社区id
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParma)
		return
	}

	list, err := logic.CommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityDetail failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, list)
}
