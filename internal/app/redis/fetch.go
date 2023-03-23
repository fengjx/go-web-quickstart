package redis

import (
	"github.com/fengjx/go-web-quickstart/pkg/json"
	"time"
)

func SetObj(key string, obj interface{}, expiration time.Duration) error {
	jsonStr := json.ToJson(obj)
	return Default().Set(Ctx, key, jsonStr, expiration).Err()
}

// GetObj
// @desc 查询缓存，并返回对象
// @return bool 缓存是否有数据
// @return error 异常
func GetObj(key string, obj interface{}) (bool, error) {
	val, err := Default().Get(Ctx, key).Result()
	if err != nil {
		return false, err
	}
	if val == "" && err == nil {
		return false, err
	}
	err = json.FromJson(val, obj)
	if err != nil {
		return false, err
	}
	return true, nil
}
