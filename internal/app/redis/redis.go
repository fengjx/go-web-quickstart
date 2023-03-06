package redis

import (
	"github.com/redis/go-redis/v9"
)

func Default() *redis.Client {
	return defaultCli
}

func GetCli(name string) *redis.Client {
	return cliMap[name]
}
