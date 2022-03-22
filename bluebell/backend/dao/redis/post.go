package redis

import (
	"bluebell/models"
	"errors"
	"math"
	"time"

	"github.com/go-redis/redis"
)

const (
	oneWeekSeconds = 7 * 24 * 3600
	score          = 432 // 一票的分数
)

var (
	ErrTimeExpire = errors.New("投票时间已过")
)

/*
投票分为四种情况：1.投赞成票 2.投反对票 3.取消投票 4.反转投票
记录文章参与投票的人
更新文章分数：赞成票要加分；反对票减分

v=1时，有两种情况
	1.之前没投过票，现在要投赞成票
	2.之前投过反对票，现在要改为赞成票
v=0时，有两种情况
	1.之前投过赞成票，现在要取消
	2.之前投过反对票，现在要取消
v=-1时，有两种情况
	1.之前没投过票，现在要投反对票
	2.之前投过赞成票，现在要改为反对票
*/

// VoteForPost
// 帖子发布一个星期后 不再投票
func VoteForPost(userID, postID string, direction float64) error {
	// 1.判断投票类型
	// 获取帖子发布时间
	postTime := rdb.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekSeconds {
		return ErrTimeExpire
	}

	// 2.更新分数
	oldV := rdb.ZScore(getRedisKey(KeyPostVoteZSetPF+postID), postID).Val()
	var op float64
	if direction > oldV {
		op = 1 //加分数
	} else {
		op = -1 //减分数
	}
	diff := math.Abs(oldV - direction)
	pipeline := rdb.TxPipeline()
	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), diff*op*score, postID)

	// 3.记录用户投票信息
	if direction == 0 {
		// 取消投票
		pipeline.ZRem(getRedisKey(KeyPostVoteZSetPF+postID), userID)
	} else {
		pipeline.ZAdd(getRedisKey(KeyPostVoteZSetPF+postID), redis.Z{
			Score:  direction,
			Member: userID,
		})
	}
	_, err := pipeline.Exec()

	return err
}

// CreatePost 创建帖子
func CreatePost(pid int64) error {
	pipeline := rdb.TxPipeline()

	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: pid,
	})

	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  0,
		Member: pid,
	})

	_, err := pipeline.Exec()
	return err
}

// GetPostsInOrder 获取排序的id
func GetPostsInOrder(p *models.PostListParam) (ids []string, err error) {
	// 默认时间排序
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1
	return rdb.ZRevRange(key, start, end).Result()
}

// GetPostsScore 获取帖子的分数
func GetPostsScore(ids []string) (data []int64, err error) {

	key := getRedisKey(KeyPostScoreZSet)
	pipeline := rdb.TxPipeline()
	for _, v := range ids {
		pipeline.ZScore(key, v).Val()
	}

	cmders, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}

	data = make([]int64, 0, len(cmders))
	for _, cmder := range cmders {
		v := cmder.(*redis.FloatCmd).Val()
		data = append(data, int64(v))
	}
	return
}
