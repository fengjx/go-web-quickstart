package redis

import (
	"github.com/redis/go-redis/v9"
)

func GetDefaultClient() *redis.Client {
	return defaultCli
}

func GetClient(name string) *redis.Client {
	return cliMap[name]
}
