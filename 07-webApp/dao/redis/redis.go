package redis

import (
	"07-webApp/setting"
	"fmt"
	"github.com/go-redis/redis"
)

var rdb *redis.Client

// Init 初始化连接
func Init(cfg *setting.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			//viper.GetString("redis.host"),
			//viper.GetInt("redis.port"),
			cfg.Host,
			cfg.Port,
		),
		//Password: viper.GetString("redis.password"), // no password set
		//DB:       viper.GetInt("redis.db"),          // use default DB
		//PoolSize: viper.GetInt("reids.pool_size"),
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})

	_, err = rdb.Ping().Result()
	return
}

func Close() {
	_ = rdb.Close()
}
