package redis_demo

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client
var ctx = context.Background()

func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "20212021", // 密码
		DB:       0,          // 数据库
		PoolSize: 20,         // 连接池大小
	})

	_, err = rdb.Ping(ctx).Result()
	return err
}
func main() {
	if err := initClient(); err != nil {
		fmt.Printf("init redis client failed, errL:%v\n", err)
		return
	}

	defer rdb.Close()

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}
