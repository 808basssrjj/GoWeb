package redis

const (
	KeyPrefix         = "bluebell:"
	KeyPostTimeZSet   = "post:time"  // 发帖时间
	KeyPostScoreZSet  = "post:score" // 投票分数
	KeyPostVoteZSetPF = "post:vote:" // 用户和投票类型
)

func getRedisKey(key string) string {
	return KeyPrefix + key
}
