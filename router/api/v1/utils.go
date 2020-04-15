package v1

import "github.com/go-redis/redis"

var redisCli *redis.Client

func init() {
	redisCli = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
}

func GetRedisClient() *redis.Client {
	return redisCli
}
