package redis

import (
	"context"
	"fengjx/go-web-quickstart/internal/app/appconfig"
	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

var cliMap = make(map[string]*redis.Client)
var Default *redis.Client

func Init() {
	for key, c := range appconfig.Conf.Redis {
		cli := redis.NewClient(&redis.Options{
			ClientName: key,
			Addr:       c.Addr,
			Password:   c.Password,
			DB:         c.DB,
		})
		cliMap[key] = cli
	}
	Default = cliMap["default"]
}
