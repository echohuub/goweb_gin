package tool

import (
	"context"
	"github.com/go-redis/redis/v9"
	"github.com/mojocn/base64Captcha"
	"log"
	"time"
)

var RedisClient *RedisStore

type RedisStore struct {
	client *redis.Client
	ctx    context.Context
}

func InitRedisStore(ctx context.Context) *RedisStore {
	config := GetConfig().Redis
	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr + ":" + config.Port,
		Password: config.Password,
		DB:       config.Db,
	})
	RedisClient = &RedisStore{client: client, ctx: ctx}

	base64Captcha.SetCustomStore(RedisClient)

	return RedisClient
}

func (rs *RedisStore) Set(id string, value string) {
	err := rs.client.Set(rs.ctx, id, value, time.Minute*10).Err()
	if err != nil {
		log.Println(err)
	}
}

func (rs *RedisStore) Get(id string, clear bool) string {
	result, err := rs.client.Get(rs.ctx, id).Result()
	if err != nil {
		log.Println(err)
		return ""
	}
	if clear {
		err := rs.client.Del(rs.ctx, id).Err()
		if err != nil {
			log.Println(err)
			return result
		}
	}
	return result
}
