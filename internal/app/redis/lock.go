package redis

import (
	"time"

	"github.com/fengjx/go-halo/utils"
	"github.com/google/uuid"

	"github.com/fengjx/go-web-quickstart/internal/common/applog"
)

func TryLock(key string, timeout time.Duration) string {
	version := uuid.NewString()
	ok, err := GetDefaultClient().SetNX(Ctx, key, version, timeout).Result()
	if err != nil && ok {
		applog.Log.Errorf("set nx err - %s", err.Error())
		return ""
	}
	return version
}

func Unlock(key string, version string) bool {
	script := `
	if redis.call("get",KEYS[1]) == ARGV[1] then
		return redis.call("del",KEYS[1])
	else
		return 0
	end
	`
	keys := []string{key}
	result, err := GetDefaultClient().Eval(Ctx, script, keys, version).Result()
	if err != nil {
		applog.Log.Errorf("unlock err - %s", err.Error())
		return false
	}
	return utils.ToInt64(result) > 0
}

func DoInLock(key string, timeout time.Duration, fun func() interface{}) interface{} {
	version := TryLock(key, timeout)
	if version != "" {
		r := fun()
		Unlock(key, version)
		return r
	}
	return nil
}
