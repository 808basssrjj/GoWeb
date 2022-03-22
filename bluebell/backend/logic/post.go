package logic

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/models"
	"bluebell/pkg/snowflake"
	"strconv"

	"go.uber.org/zap"
)

// CreatePost 帖子发布
func CreatePost(post *models.Post) error {
	// 生成帖子postId
	post.PostID = snowflake.GenID()

	err := mysql.CreatePost(post)
	if err != nil {
		return err
	}
	err = redis.CreatePost(post.PostID)
	return err
}

// PostDetail 帖子详情
func PostDetail(id int) (detail *models.ApiPostDetail, err error) {
	post, err := mysql.PostDetail(id)
	if err != nil {
		zap.L().Error("mysql.PostDetail failed", zap.Error(err))
		return
	}
	// 获取社区信息
	community, err := mysql.CommunityDetail(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.CommunityDetail failed", zap.Error(err))
		return
	}
	// 用户信息
	user, err := mysql.UserDetail(post.AuthorId)
	if err != nil {
		zap.L().Error("mysql.UserDetail failed", zap.Error(err))
		return
	}

	detail = &models.ApiPostDetail{
		AutoName:        user.Username,
		Post:            post,
		CommunityDetail: community,
	}
	return
}

// PostList 帖子列表
func PostList(param *models.PostListParam) (detail []models.PostList, err error) {
	//detail, err = mysql.PostList(offset, limit)

	// 1.从redis获取ids和score
	ids, err := redis.GetPostsInOrder(param)
	if err != nil {
		return nil, err
	}
	votes, err := redis.GetPostsScore(ids)
	if err != nil {
		return nil, err
	}

	// 2.查询数据库
	detail, err = mysql.PostListByIDs(ids)
	for i, vote := range votes {
		detail[i].VoteNum = vote
	}
	return
}

// PostVote 帖子投票
func PostVote(userID int64, vote *models.VoteParam) (err error) {
	return redis.VoteForPost(strconv.Itoa(int(userID)), vote.PostID, float64(vote.Direction))
}
