package db

import (
	"fengjx/go-web-quickstart/internal/app/redis"
	"fmt"
	"time"
)

type getSupplier = func() (interface{}, error)

func GetByIdCache(id int64, obj interface{}, version string, supplier getSupplier) error {
	key := fmt.Sprintf("model:%T:%d:%s", obj, id, version)
	has, _ := redis.GetObj(key, &obj)
	if has {
		return nil
	}
	obj, err := supplier()
	if err != nil {
		return err
	}
	if obj != nil {
		err := redis.SetObj(key, obj, 2*time.Minute)
		return err
	}
	return nil
}
