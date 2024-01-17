package redis

import (
	"context"
	"fmt"

	"github.com/spf13/viper"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client
var ctx = context.Background()

func Init() (err error) {
	password := viper.GetString("redis.password")
	host := viper.GetString("redis.host")
	port := viper.GetInt("redis.port")
	db := viper.GetInt("redis.db")
	poolSize := viper.GetInt("redis.pool_size")
	MinIdleConns := viper.GetInt("redis.min_idle_conns")
	rdb = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", host, port),
		Password:     password, // 密码
		DB:           db,       // 数据库
		PoolSize:     poolSize, // 连接池大小
		MinIdleConns: MinIdleConns,
	})

	_, err = rdb.Ping(ctx).Result()

	return
}

func Close() {
	_ = rdb.Close()
}
