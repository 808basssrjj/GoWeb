package controllers

import (
	"bluebell/logic"
	"bluebell/models"
	"strconv"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// CreatePostHandler 发布帖子
func CreatePostHandler(c *gin.Context) {
	// 获取参数
	post := new(models.Post)
	if err := c.ShouldBindJSON(post); err != nil {
		zap.L().Debug("CreatePost params err", zap.Any("err", err))
		ResponseErrorWithMsg(c, CodeInvalidParma, err.Error())
		return
	}

	// 获取作者ID
	userID, err := getCurrentUser(c)
	if err != nil {
		zap.L().Error("getCurrentUser failed", zap.Error(err))
		ResponseError(c, CodeNeedLogin)
		return
	}
	post.AuthorId = userID

	// 创建
	err = logic.CreatePost(post)
	if err != nil {
		zap.L().Error("CreatePost failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, nil)
}

// PostDetailHandler 帖子详情
func PostDetailHandler(c *gin.Context) {
	// 帖子id
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ResponseError(c, CodeInvalidParma)
		return
	}

	// 查询数据库
	list, err := logic.PostDetail(id)
	if err != nil {
		zap.L().Error("logic.PostDetail failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, list)
}

// PostListHandler 帖子列表
func PostListHandler(c *gin.Context) {
	// 获取query参数
	//pageStr := c.Query("page")
	//sizeStr := c.Query("size")
	//page, err := strconv.ParseInt(pageStr, 10, 64)
	//if err != nil {
	//	page = 1
	//	return
	//}
	//size, err := strconv.ParseInt(sizeStr, 10, 64)
	//if err != nil {
	//	size = 10
	//	return
	//}
	param := &models.PostListParam{
		Page:  1,
		Size:  10,
		Order: models.OrderTime,
	}
	err := c.ShouldBindQuery(param)
	if err != nil {
		zap.L().Error("posts param err", zap.Error(err))
		ResponseError(c, CodeInvalidParma)
		return
	}

	list, err := logic.PostList(param)
	if err != nil {
		zap.L().Error("logic.PostList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, list)
}

// VoteHandler 帖子投票
func VoteHandler(c *gin.Context) {
	vote := &models.VoteParam{}
	if err := c.ShouldBindJSON(vote); err != nil {
		zap.L().Error("vote params error", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParma, err)
		return
	}

	userID, err := getCurrentUser(c)
	if err != nil {
		zap.L().Error("not login", zap.Error(err))
		ResponseError(c, CodeNeedLogin)
		return
	}

	// 具体逻辑
	err = logic.PostVote(userID, vote)
	if err != nil {
		zap.L().Error("logic.PostVote error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
