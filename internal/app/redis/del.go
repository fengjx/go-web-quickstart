package redis

import (
	"fengjx/go-web-quickstart/internal/app/applog"
	"fmt"
	"time"
)

func DelWithLock(key string) bool {
	lockKey := fmt.Sprintf("lock:%s", key)
	res := DoInLock(lockKey, time.Second*1, func() interface{} {
		result, err := Default().Del(Ctx, key).Result()
		if err != nil {
			applog.Log.Errorf("del with lock err - %s", err.Error())
		}
		return result > 0
	})
	return res.(bool)
}
