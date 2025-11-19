package initialize

import (
	"context"
	"fmt"
	"go-ecommerce-backend-api/global"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	r := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", r.Host, r.Port),
		Password: r.Password, // no password set
		DB:       r.Database,  // use default DB
		PoolSize: 10,
	})

	_, err := rdb.Ping(ctx).Result()
	if err !=nil {
		global.Logger.Error("Redis initialization Error:", zap.Error(err))
	}

	fmt.Println("Init Redis is running")
	global.Rdb = rdb
}