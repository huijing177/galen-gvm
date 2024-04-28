package inits

import (
	"context"

	"galen-gvm/global"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func Redis() *redis.Client {
	redisConfig := global.GVA_CONFIG.Redis

	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.GVA_LOG.Error("redis connect ping failed, err:", zap.Error(err))
		panic(err)
	}
	global.GVA_LOG.Info("redis connect ping response:", zap.String("pong", pong))
	return client
}
