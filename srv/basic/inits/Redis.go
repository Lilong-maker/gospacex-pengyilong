package inits

import (
	"context"
	"fmt"

	"github.com/Lilong-maker/consul"
	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()
var Rdb *redis.Client

func RedisInit() {
	RedisConfig := consul.Gen.Redis
	Addr := fmt.Sprintf("%s:%d",
		RedisConfig.Host,
		RedisConfig.Port,
	)
	Rdb = redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: RedisConfig.Password, // no password set
		DB:       RedisConfig.Database, // use default DB
	})
	err = Rdb.Ping(Ctx).Err()
	if err != nil {
		return
	}
	fmt.Println("redis连接成功")
}
