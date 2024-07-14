package redis

import (
	"context"
	"fmt"
	"github.com/easytech-international-sdn-bhd/esynx-auth/models"
	_redis "github.com/redis/go-redis/v9"
	"time"
)

type RedisInstance struct {
	ctx    context.Context
	Client *_redis.Client
}

func NewRedis(ctx context.Context, config models.RedisConfig) *RedisInstance {
	return &RedisInstance{
		ctx: ctx,
		Client: _redis.NewClient(&_redis.Options{
			Addr:         fmt.Sprintf("%s:%d", config.Host, config.Port),
			Password:     config.Pass,
			DB:           1,
			DialTimeout:  config.DialTimeout,
			ReadTimeout:  config.ReadTimeout,
			WriteTimeout: config.WriteTimeout,
			PoolSize:     config.PoolSize,
			PoolTimeout:  config.PoolTimeout,
		}),
	}
}

func (r *RedisInstance) GetToken(uid string) (string, error) {
	return r.Client.Get(r.ctx, uid).Result()
}

func (r *RedisInstance) SetToken(uid string, token string, expiration time.Duration) error {
	return r.Client.Set(r.ctx, uid, token, expiration).Err()
}

func (r *RedisInstance) DelToken(uid string) error {
	return r.Client.Del(r.ctx, uid).Err()
}
